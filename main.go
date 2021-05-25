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

func New() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

//export FindMatch
func FindMatch(filepath *C.char) *C.char {
	var status []string
	patharr := GetPaths(C.GoString(filepath))

	for _, path := range patharr {
		b, err := ioutil.ReadFile(path)
		// File Not Found
		if err != nil {
			status = append(status, "E1,"+path)
		}

		data := []byte(string(b))

		c, err := New()
		// Internal Error in Initializing Classifier
		if err != nil {
			status = append(status, "E2,"+err.Error())
		}

		m := c.Match(data)
		for i := 0; i < m.Len(); i++ {
			status = append(status, fmt.Sprintf("(%s,%f,%s,%d,%d),", m[i].Name, m[i].Confidence, m[i].MatchType, m[i].StartLine, m[i].EndLine))
		}

		// If No valid license is found
		if len(m) == 0 {
			status = append(status, "E3,"+path)
		}
	}
	return C.CString(strings.Join(status, "\n"))
}

func GetPaths(filepath string) []string {
	return strings.SplitN(filepath, "\n", -1)
}

func main() {}
