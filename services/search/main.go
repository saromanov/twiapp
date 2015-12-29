package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"

	"github.com/saromanov/notesapp/configs"
	"github.com/saromanov/notesapp/logging"
	"github.com/saromanov/twiapp/service"
)

// Getting home timeline

type Response struct {
	Info    string
	API     string
	Time    string
	Error   string
	Request string
	Data    string
}

func main() {
	logger := logging.NewLogger(nil)
	key := os.Getenv("CONSUMER_KEY")
	secret := os.Getenv("CONSUMER_SECRET")
	if key == "" {
		logger.Error("Consumer key is not found")
		return
	}

	if secret == "" {
		logger.Error("Consumer Secret is not found")
		return
	}

	anaconda.SetConsumerKey(key)
	anaconda.SetConsumerSecret(secret)
	api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")

	serv, err := service.CreateService(cfg)

	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}

	serv.HandleFunc("/api/twitter/timeline", func(w http.ResponseWriter, r *http.Request) {
		var Error string
		tweets, err := api.GetHomeTimeline()
		if err != nil {
			Error = fmt.Sprintf("%v", err)
			logger.Error(Error)
		}
		vars := mux.Vars(r)
		title := vars["title"]

		tweets_string, err := json.Marshal(tweets)
		if err != nil {
			Error = fmt.Sprintf("%v", err)
			logger.Error(Error)
		}

		resp := Response{Request: "GET", API: "timeline", Info: "Getting timeline by user", Time: time.Now().String(),
			Data: tweets_string, Error: Error}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}

	})

	logger.Info("Service twitter is started")
	serv.Start()
}
