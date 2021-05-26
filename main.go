/*
	Error Codes:
	E1 - Failure in loading licenses from given directory.
	E2 - File Not Found at specified path.
	E3 - No matching license found.
*/

package main

import "C"
import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/avishrantssh/GoLicenseClassifier/classifier"
)

// Default Threshold for Filtering the results
var defaultThreshold = 0.8

// Base Licenses Root Directory
var baseLicenses = "./classifier/licenses"

// Create a classifier instance and load base licenses
func CreateClassifier() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

//export FindMatch
func FindMatch(filepath *C.char) *C.char {
	var status []string
	patharr := GetPaths(C.GoString(filepath))

	// A simple channel implementation to lock function until execution is complete
	sem := make(chan struct{})
	c, err := CreateClassifier()

	// fmt.Println("Finished Reading licenses")
	if err != nil {
		return C.CString("E1")
	}

	for _, path := range patharr {

		// Go Routine far faster Processing
		go func(path string) {
			b, err := ioutil.ReadFile(path)
			// File Not Found
			if err != nil {
				status = append(status, "E2,"+path)
			}

			data := []byte(string(b))
			// Internal Error in Initializing Classifier

			m := c.Match(data)
			var tmp string
			for i := 0; i < m.Len(); i++ {
				tmp += fmt.Sprintf("(%s,%f,%s,%d,%d),", m[i].Name, m[i].Confidence, m[i].MatchType, m[i].StartLine, m[i].EndLine)
			}

			// If No valid license is found
			if tmp == "" {
				status = append(status, "E3,"+path)
			} else {
				status = append(status, path+":"+tmp)
			}

			sem <- struct{}{}
		}(path)
	}
	for range patharr {
		<-sem
	}

	return C.CString(strings.Join(status, "\n"))
}

// GetPaths function is used to convert new-line seperated filepaths to a string array.
func GetPaths(filepath string) []string {
	return strings.SplitN(filepath, "\n", -1)
}

//export LoadCustomLicenses
func LoadCustomLicenses(path *C.char) {
	baseLicenses = C.GoString(path)
}

func main() {}
