package msteams

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Webhook struct {
	URL string
}

func NewWebhook(webhookUrl string) *Webhook {
	return &Webhook{URL: webhookUrl}
}

type Response struct {
	Body       string
	Status     string
	StatusCode int
	Header     http.Header
}

func (c *Webhook) SendMessage(message *Message) (*Response, error) {
	return c.SendMessageWithContext(context.Background(), message)
}

func (c *Webhook) SendMessageWithContext(ctx context.Context, message *Message) (*Response, error) {
	data, err := jsoniter.Marshal(*message)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       string(responseByte),
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
	}, nil
}
