package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/licenseclassifier/v2/classifier"
)

var defaultThreshold = 0.8
var baseLicenses = "./classifier/licenses"

//export New
func New() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

func main() {
	// license := "/home/avishrant/GitRepo/scancode.io/LICENSE"
	license := "/home/avishrant/GitRepo/licenseclassifier/forbidden.go"
	// license := "/home/avishrant/GitRepo/Classifier/Licensev2/classifier/scenarios/114431182"
	b, err := ioutil.ReadFile(license)
	if err != nil {
		log.Fatalf("Couldn't read scenario %s: %v", license, err)
	}
	// lines := strings.SplitN(string(b), "EXPECTED:", 2)
	// lines = strings.SplitN(lines[1], "\n", 2)
	data := []byte(string(b))

	// New Classifier
	c, err := New()
	if err != nil {
		fmt.Printf("couldn't instantiate standard test classifier: %v", err)
	}
	m := c.Match(data)
	for i := 0; i < m.Len(); i++ {
		fmt.Printf("Name : %s\nConfidence : %f\nMatchType : %s\nStartLine : %d\nEndLine : %d\n\n", m[i].Name, m[i].Confidence, m[i].MatchType, m[i].StartLine, m[i].EndLine)
	}
}
