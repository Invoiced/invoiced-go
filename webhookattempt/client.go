package webhookattempt

import (
	"github.com/Invoiced/invoiced-go/v2"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.WebhookAttempts, error) {
	endpoint2 := invoiced.AddFilterAndSort("/webhook_attempts", filter, sort)

	webhookAttempts := make(invoiced.WebhookAttempts, 0)

NEXT:
	tmpWebhookAttempts := make(invoiced.WebhookAttempts, 0)

	endpoint, err := c.Api.Get(endpoint2, &tmpWebhookAttempts)

	if err != nil {
		return nil, err
	}

	webhookAttempts = append(webhookAttempts, tmpWebhookAttempts...)

	if endpoint != "" {
		goto NEXT
	}

	return webhookAttempts, nil
}

func (c *Client) ReAttempt(webhookId int64) error {
	return c.Api.PostWithoutData("/webhook_attempts/"+strconv.FormatInt(webhookId, 10)+"/retries", nil)
}
