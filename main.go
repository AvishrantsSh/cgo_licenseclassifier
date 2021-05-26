package main

import "C"
import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/avishrantssh/GoLicenseClassifier/classifier"
)

var defaultThreshold = 0.8
var baseLicenses = "./classifier/licenses"

// CreateClassifier creates a classifier instance
func CreateClassifier() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

//export FindMatch
func FindMatch(filepath *C.char) *C.char {
	var status []string
	patharr := GetPaths(C.GoString(filepath))
	sem := make(chan struct{})
	for _, path := range patharr {
		// Go Routine far faster Processing
		go func(path string) {
			b, err := ioutil.ReadFile(path)
			// File Not Found
			if err != nil {
				status = append(status, "E1,"+path)
			}

			data := []byte(string(b))

			c, err := CreateClassifier()
			// Internal Error in Initializing Classifier
			if err != nil {
				status = append(status, "E2,"+err.Error())
			}

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
