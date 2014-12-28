package sentiment

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jbrukh/bayesian"
)

const (
	Good bayesian.Class = "Good"
	Bad  bayesian.Class = "Bad"
)

type RoverClassifier struct {
	*bayesian.Classifier
}

func NewRoverClassifierFromFiles(positiveSentences string, negativeSentencies string) *RoverClassifier {
	classifier := &RoverClassifier{
		bayesian.NewClassifier(Good, Bad),
	}

	goodStuff := extractFromFile(positiveSentences)
	badStuff := extractFromFile(negativeSentencies)

	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff, Bad)

	return classifier
}

func extractFromFile(filename string) []string {
	var tokens []string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Errorf("Unable to open that file %v", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tokens = append(tokens, tokenize(scanner.Text())...)
	}

	return tokens
}
