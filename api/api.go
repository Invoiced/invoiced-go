package api

import (
	"github.com/Invoiced/invoiced-go"
	"github.com/Invoiced/invoiced-go/charge"
	"github.com/Invoiced/invoiced-go/chasing"
	"github.com/Invoiced/invoiced-go/customer"
	"github.com/Invoiced/invoiced-go/webhook_attempt"
)

type Client struct {
	Api            *invoiced.Api
	Charge         charge.Client
	ChasingCadence chasing.Client
	Customer       customer.Client
	WebhookAttempt webhook_attempt.Client
}

func New(key string, sandbox bool) *Client {
	apiClient := invoiced.New(key, sandbox)

	return &Client{
		Api:            apiClient,
		Charge:         charge.Client{apiClient},
		ChasingCadence: chasing.Client{apiClient},
		Customer:       customer.Client{apiClient},
		WebhookAttempt: webhook_attempt.Client{apiClient},
	}
}
