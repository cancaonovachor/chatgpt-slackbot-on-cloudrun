package slackhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"chatgptbot/gpt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type Client struct {
	gptClient   *gpt.Client
	slackClient *slack.Client
}

type RequestBody struct {
	Message      PubSubMessage `json:"message"`
	Subscription string        `json:"subscription"`
}

type PubSubMessage struct {
	Attributes map[string]interface{} `json:"attributes"`
	Data       []byte                 `json:"data"`
	MessageId  string                 `json:"messageId"`
}

func NewClient(gptClient *gpt.Client, slackClient *slack.Client) *Client {
	return &Client{
		gptClient:   gptClient,
		slackClient: slackClient,
	}
}

func (c *Client) SlackEvent(ctx context.Context, reqBody []byte) error {
	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(slackBotToken)
	botUserId, err := getBotUserId(api)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(reqBody), slackevents.OptionNoVerifyToken())
	if err != nil {
		return err
	}

	switch eventsAPIEvent.Type {
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			channelId := ev.Channel
			threadTs := ev.ThreadTimeStamp
			if threadTs == "" {
				threadTs = ev.TimeStamp
			}

			// Fetch replies
			replies, _, _, err := api.GetConversationReplies(&slack.GetConversationRepliesParameters{
				ChannelID: channelId,
				Timestamp: threadTs,
			})
			if err != nil {
				log.Printf("Error: %s", err)

			}

			// Create threadMessages
			threadMessages := createThreadMessages(replies, botUserId)

			// Get GPT-3 answer
			gptAnswerText, err := c.gptClient.GenerateText(threadMessages)
			if err != nil {
				log.Printf("Error: %s", err)

			}

			// Reply to the thread
			_, _, err = api.PostMessage(channelId, slack.MsgOptionText(gptAnswerText, false), slack.MsgOptionTS(threadTs))
			if err != nil {
				log.Printf("Error: %s", err)
			}
		}
	}
	return nil
}

func createThreadMessages(replies []slack.Message, botUserId string) []gpt.Message {
	threadMessages := []gpt.Message{}
	for _, message := range replies {
		role := "user"
		if message.User == botUserId {
			role = "assistant"
		}

		content := strings.ReplaceAll(message.Text, fmt.Sprintf("<@%s>", botUserId), "")
		threadMessages = append(threadMessages, gpt.Message{
			Role:    role,
			Content: content,
		})
	}

	return threadMessages
}

func getBotUserId(api *slack.Client) (string, error) {
	resp, err := api.AuthTest()
	if err != nil {
		return "", err
	}
	return resp.UserID, nil
}
