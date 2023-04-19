package slackhandler

import (
	"context"
	"encoding/json"
	"log"

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
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(reqBody), slackevents.OptionNoVerifyToken())
	if err != nil {
		return err
	}

	switch eventsAPIEvent.Type {
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch event := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			message, err := c.buildMessage(ctx, event)
			if err != nil {
				return err
			}

			response, err := c.gptClient.GenerateText(message, 500)
			if err != nil {
				log.Println("error: ", err.Error())
				_, _, _ = c.slackClient.PostMessageContext(ctx, event.Channel, slack.MsgOptionText("エラーが起きました\n"+err.Error(), false))
			} else if response != "" {
				_, _, _ = c.slackClient.PostMessageContext(ctx,
					event.Channel,
					slack.MsgOptionText(response, false),
					slack.MsgOptionTS(event.TimeStamp),
				)
			}
		}
	}
	return nil
}

func (c *Client) buildMessage(ctx context.Context, event *slackevents.AppMentionEvent) (string, error) {
	if event.ThreadTimeStamp == "" {
		return event.Text, nil
	}

	threadTimestamp := event.ThreadTimeStamp
	channelID := event.Channel
	history, _, _, err := c.slackClient.GetConversationReplies(&slack.GetConversationRepliesParameters{
		ChannelID: channelID,
		Timestamp: threadTimestamp,
	})
	if err != nil {
		log.Fatalf("Error fetching conversation replies: %v", err)
		return "", err
	}

	conversationHistory := ""
	for _, message := range history {
		if message.ThreadTimestamp == threadTimestamp {
			author := "User"
			if message.BotID != "" {
				author = "ChatGPT"
			}
			conversationHistory += author + ": " + message.Text + "\n"
		}
	}

	return conversationHistory + "ChatGPT:", nil
}
