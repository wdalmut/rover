package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/darkhelmet/twitterstream"
	"github.com/wdalmut/rover/sentiment"
	"github.com/wdalmut/twitterstream/async"
)

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func main() {
	classifier := sentiment.NewRoverClassifierFromFiles("files/rt-polarity.pos", "files/rt-polarity.neg")

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)

	fmt.Printf("%v", config)

	client := async.NewClient(
		config.ConsumerKey,
		config.ConsumerSecret,
		config.AccessToken,
		config.AccessSecret,
	)

	client.TrackAndServe("bieber", func(tweet *twitterstream.Tweet) {
		class := classifier.Classify(tweet.Text)

		fmt.Printf("Tweet: %s is %v\n\n", tweet.Text, class)
	})
}
