package webhook_attempt

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

const endpoint = "/webhook_attempts"

type Client struct {
	*invoiced.Api
}

type WebhookAttempts []*invoiced.WebhookAttempt

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (WebhookAttempts, error) {
	endpoint2 := invoiced.AddFilterAndSort(endpoint, filter, sort)

	webhookAttempts := make(WebhookAttempts, 0)

NEXT:
	tmpWebhookAttempts := make(WebhookAttempts, 0)

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
	endpoint2 := endpoint + "/" + strconv.FormatInt(webhookId, 10) + "/retries"
	return c.Api.Create(endpoint2, nil, nil)
}
