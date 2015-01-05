package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/darkhelmet/twitterstream"
	"github.com/gorilla/mux"
	"github.com/wdalmut/rover/proxy"
	"github.com/wdalmut/rover/proxy/wrap"
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

	router := mux.NewRouter()
	server := proxy.Server{
		Router: router,
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:8082",
			Handler: router,
		},
		Wrapper: &wrap.HtmlWrapper{},
	}

	server.ListenAndServe()

	classifier := sentiment.NewRoverClassifierFromFiles("files/rt-polarity.pos", "files/rt-polarity.neg")

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)

	client := async.NewClient(
		config.ConsumerKey,
		config.ConsumerSecret,
		config.AccessToken,
		config.AccessSecret,
	)

	client.TrackAndServe("sun", func(tweet *twitterstream.Tweet) {
		class := classifier.Classify(tweet.Text)

		fmt.Printf("Tweet: %s is %v\n\n", tweet.Text, class)
	})
}
