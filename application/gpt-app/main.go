package main

import (
	"log"
	"net/http"
	"os"

	"chatgptbot/api"
	"chatgptbot/gpt"
	"chatgptbot/logger"
	"chatgptbot/slackhandler"

	"github.com/go-chi/chi"
	"github.com/slack-go/slack"
)

func main() {
	gptClient := gpt.NewClient()
	slackClient := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	slackHandler := slackhandler.NewClient(gptClient, slackClient)
	apiHandler := api.NewAPIHandler(gptClient, slackHandler)

	r := chi.NewRouter()
	r.Use(logger.Logger)
	r.Post("/", apiHandler.HandleChatGptEvent)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
