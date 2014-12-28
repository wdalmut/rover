package tokenizer

import (
	"regexp"
)

func tokenize(sentence string) []string {
	re := regexp.MustCompile("\\w+")
	return re.FindAllString(sentence, -1)
}
