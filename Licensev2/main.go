package main

import "C"
import (
	"fmt"
	"io/ioutil"

	"github.com/google/licenseclassifier/v2/classifier"
)

var defaultThreshold = 0.8
var baseLicenses = "./classifier/licenses"

func New() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

func FindMatch(path string) *C.char {
	var status string
	b, err := ioutil.ReadFile(path)
	if err != nil {
		status = "Couldn't read file at : " + path
	}

	data := []byte(string(b))

	// New Classifier
	c, err := New()
	if err != nil {
		status = "Couldn't instantiate standard test classifier: " + err.Error()
	}
	m := c.Match(data)
	for i := 0; i < m.Len(); i++ {
		status = fmt.Sprintf("Name : %s\nConfidence : %f\nMatchType : %s\nStartLine : %d\nEndLine : %d\n\n", m[i].Name, m[i].Confidence, m[i].MatchType, m[i].StartLine, m[i].EndLine)
	}
	return C.CString(status)
}
func main() {
	license := "/home/avishrant/GitRepo/scancode.io/LICENSE"
	fmt.Println(C.GoString(FindMatch(license)))
}
