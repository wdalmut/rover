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

func NewRoverClassifierFromFiles(positiveSentences, negativeSentencies string) *RoverClassifier {
	goodStuff := extractFromFile(positiveSentences)
	badStuff := extractFromFile(negativeSentencies)

	return NewRoverClassifierFromPieces(goodStuff, badStuff)
}

func NewRoverClassifierFromPieces(goodStuff, badStuff []string) *RoverClassifier {
	classifier := &RoverClassifier{
		bayesian.NewClassifier(Good, Bad),
	}

	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff, Bad)

	return classifier
}

func (c *RoverClassifier) Classify(sentence string) bayesian.Class {
	scores, _, _ := c.LogScores(tokenize(sentence))

	if scores[0] > scores[1] {
		return Good
	} else {
		return Bad
	}
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
