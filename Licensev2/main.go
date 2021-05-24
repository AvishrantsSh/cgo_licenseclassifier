package main

import "github.com/google/licenseclassifier/v2/classifier"

var defaultThreshold = 0.8
var baseLicenses = "./classifier/licenses"

//export New
func New() (*classifier.Classifier, error) {
	c := classifier.NewClassifier(defaultThreshold)
	return c, c.LoadLicenses(baseLicenses)
}

func main() {}
