package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/slack-go/slack/slackevents"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func main() {
	http.HandleFunc("/", EventHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Received a new request.")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		log.Print("Received a new event: ", eventsAPIEvent.Type)
		ctx := context.Background()
		client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		topic := client.Topic(os.Getenv("PUBSUB_TOPIC"))
		res := topic.Publish(ctx, &pubsub.Message{
			Data: buf.Bytes(),
		})

		_, err = res.Get(ctx)
		if err != nil {
			log.Fatalf("Failed to publish: %v", err)
		}
		log.Print(res)
	}

	w.WriteHeader(http.StatusOK)
}
