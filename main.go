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

// Regexp for Detecting Copyrights
var copyrightRE = regexp.MustCompile(`(?m)(?i:Copyright)\s+(?i:©\s+|\(c\)\s+)?(?:\d{2,4})(?:[-,]\s*\d{2,4})*,?\s*(?i:by)?\s*(.*?(?i:\s+Inc\.)?)[.,]?\s*(?i:All rights reserved\.?)?\s*$`)

// Removing in-text special code literals
var endliteralRE = regexp.MustCompile(`\\n|\\f|\\r`)

// Create a classifier instance and load base licenses
func CreateClassifier() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(licensePath)
}

//export FindMatch
func FindMatch(root *C.char, fpaths *C.char, getjson bool) *C.char {
	ROOT = C.GoString(root)
	if licensePath == "" {
		licensePath = filepath.Join(ROOT, default_path)
	}
	patharr := GetPaths(C.GoString(fpaths))
	res := new(result.JSON_struct)
	res.Init(len(patharr))
	c, err := CreateClassifier()

	if err != nil {
		return C.CString("ERROR:" + err.Error())
	}

	// A simple channel implementation to lock function until execution is complete
	var wg sync.WaitGroup
	wg.Add(len(patharr))

	for index, path := range patharr {
		// Spawn a thread for each iteration in the loop.
		go func(index int, path string) {
			defer wg.Done()
			finfo := result.FileInfo{}
			finfo.Path = path
			b, err := ioutil.ReadFile(path)
			// File Not Found
			if err != nil {
				finfo.Errors = err.Error()
				return
			}

			data := []byte(string(b))
			m := c.Match(data)
			for i := 0; i < m.Len(); i++ {
				finfo.Licenses = append(finfo.Licenses, result.License{
					Expression: m[i].Name,
					Confidence: m[i].Confidence,
					Startline:  m[i].StartLine,
					Endline:    m[i].EndLine,
					Starttoken: m[i].StartTokenIndex,
					Endttoken:  m[i].EndTokenIndex})
			}

			cpInfo, holder := CopyrightInfo(string(b))
			if len(cpInfo) > 0 {
				finfo.Copyrights = append(finfo.Copyrights, result.CpInfo{
					Expression: cpInfo,
					Holders:    holder})
			}
			res.AddFile(index, &finfo)
		}(index, path)
	}

	// Wait for `wg.Done()` to be exectued the number of times specified in the `wg.Add()` call.
	wg.Wait()
	f_error := res.Finish("./test.json")
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
func CopyrightInfo(contents string) ([]string, []string) {
	str := endliteralRE.ReplaceAllString(contents, "\n")
	matches := copyrightRE.FindAllStringSubmatch(str, -1)
	var cpInfo, holder []string
	for _, match := range matches {
		if len(match) == 2 {
			cpInfo = append(cpInfo, match[0])
			holder = append(holder, match[1])
		}
	}
	return cpInfo, holder
}

func main() {}
