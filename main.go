package main

import "C"
import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/avishrantssh/GoLicenseClassifier/classifier"
	"github.com/avishrantssh/GoLicenseClassifier/result"
)

var ROOT = ""

// Default Threshold for Filtering the results
var defaultThreshold = 0.8

// Default Licenses Root Directory
var default_path = "./classifier/default"
var licensePath string

// Normalize Copyright Literals
var copyliteralRE = regexp.MustCompile(`&copy;|&copy|&#169;|&#xa9;|&#XA9;|u00A9|u00a9|\\xa9|\\XA9|\\251|©|/( C/)|(?i:/(c/))`)

// Regexp for Detecting Copyrights
var copyrightRE = regexp.MustCompile(`(?m)(?i:Copyright)\s+(?i:\(c\)\s+)?(?:\d{2,4})(?:[-,]\s*\d{2,4})*,?\s*(?i:by)?\s*(.*?(?i:\s+Inc\.)?)[.,-]?\s*(?i:All rights reserved\.?)?\s*$`)

// Removing in-text special code literals
var endliteralRE = regexp.MustCompile(`\\n|\\f|\\r|\\0`)

// Maximum Parallel Running Goroutines
var MAX_ROUTINES = 10000

// Create a classifier instance and load base licenses
func CreateClassifier() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(licensePath)
}

//export FindMatch
func FindMatch(root *C.char, fpaths *C.char, outputPath *C.char) *C.char {
	ROOT = C.GoString(root)
	if licensePath == "" {
		licensePath = filepath.Join(ROOT, default_path)
	}
	patharr := GetPaths(C.GoString(fpaths))
	res := new(result.JSON_struct)
	res.Init(ROOT, len(patharr))
	c, err := CreateClassifier()
	if err != nil {
		return C.CString("ERROR:" + err.Error())
	}

	// Guard channel for ensuring thar no more than 'MAX_ROUTINES' run at any given time.
	guard := make(chan struct{}, MAX_ROUTINES)

	// A simple channel implementation to lock function until execution is complete
	var wg sync.WaitGroup
	wg.Add(len(patharr))

	for index, path := range patharr {
		// Spawn a thread for each iteration in the loop.
		guard <- struct{}{}
		go func(index int, path string) {
			defer wg.Done()
			finfo := new(result.FileInfo)
			finfo.Path = path
			b, err := ioutil.ReadFile(path)
			// File Not Found
			if err != nil {
				finfo.Errors = err.Error()
				res.AddFile(index, finfo)
				finfo = nil
				<-guard
				return
			}

			data := []byte(string(b))
			m := c.Match(data)
			for i := 0; i < m.Len(); i++ {
				finfo.Licenses = append(finfo.Licenses, result.License{
					Expression: m[i].Name,
					Confidence: m[i].Confidence,
					StartLine:  m[i].StartLine,
					EndLine:    m[i].EndLine,
					StartToken: m[i].StartTokenIndex,
					EndToken:   m[i].EndTokenIndex})
			}

			cpInfo, holder, tokens := CopyrightInfo(string(b))
			for i := 0; i < len(cpInfo); i++ {
				finfo.Copyrights = append(finfo.Copyrights, result.CpInfo{
					Expression: cpInfo[i],
					StartIndex: tokens[i][0],
					EndIndex:   tokens[i][1],
					Holder:     holder[i],
				})
			}
			res.AddFile(index, finfo)
			finfo = nil
			data = nil
			<-guard
		}(index, path)
	}

	// Wait for `wg.Done()` to be exectued the number of times specified in the `wg.Add()` call.
	wg.Wait()
	f_error := res.Finish(C.GoString(outputPath))
	res = nil
	if f_error != nil {
		return C.CString(f_error.Error())
	}
	return C.CString("Done")
}

// GetPaths function is used to convert new-line seperated filepaths to a string array.
func GetPaths(filepath string) []string {
	return strings.SplitN(filepath, "\n", -1)
}

//export LoadCustomLicenses
func LoadCustomLicenses(path *C.char) int {
	licensePath = C.GoString(path)
	return 1
}

//export SetThreshold
func SetThreshold(thresh int) int {
	if thresh < 0 || thresh > 100 {
		return 1
	}
	defaultThreshold = float64(thresh) / 100.0
	return 1
}

// CopyrightHolder finds a copyright notification, if it exists, and returns
// the copyright holder.
func CopyrightInfo(contents string) ([]string, []string, [][]int) {
	str := endliteralRE.ReplaceAllString(contents, "\n")
	normalized_str := copyliteralRE.ReplaceAllString(str, "(c)")

	matches := copyrightRE.FindAllStringSubmatch(normalized_str, -1)
	tokens := copyrightRE.FindAllStringSubmatchIndex(normalized_str, -1)

	var cpInfo, holder []string
	for _, match := range matches {
		if len(match) == 2 {
			cpInfo = append(cpInfo, strings.TrimSpace(match[0]))
			holder = append(holder, strings.TrimSpace(match[1]))
		}
	}
	return cpInfo, holder, tokens
}

func main() {}
