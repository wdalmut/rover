package main

import (
	"fmt"

	"github.com/wdalmut/rover/sentiment"
)

func main() {
	classifier := sentiment.NewRoverClassifierFromFiles("files/rt-polarity.pos", "files/rt-polarity.neg")

	class := classifier.Classify("I am so happy")
	fmt.Printf("Sentence is %v\n", class)

	class = classifier.Classify("I am so unhappy")
	fmt.Printf("Sentence is %v\n", class)
}
