package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"chatgptbot/gpt"
	"chatgptbot/slackhandler"
)

type APIHandler struct {
	gptClient   *gpt.Client
	slackClient *slackhandler.Client
}

func NewAPIHandler(gptClient *gpt.Client, slackClient *slackhandler.Client) *APIHandler {
	return &APIHandler{
		gptClient:   gptClient,
		slackClient: slackClient,
	}
}

func (h *APIHandler) HandleChatGptEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var requestBody slackhandler.RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	if err := h.slackClient.SlackEvent(ctx, requestBody.Message.Data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("done")
	w.WriteHeader(http.StatusOK)
}
