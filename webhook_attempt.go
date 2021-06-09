package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"strconv"
)

type WebhookAttempt struct {
	*Connection
	*invdendpoint.WebhookAttempt
}

type WebhookAttempts []*WebhookAttempt

func (c *Connection) NewWebhookAttempt() *WebhookAttempt {
	webhookAttempt := new(invdendpoint.WebhookAttempt)
	return &WebhookAttempt{c, webhookAttempt}
}

func (c *WebhookAttempt) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (WebhookAttempts, error) {
	endpoint := invdendpoint.WebhookEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	webhookAttempts := make(WebhookAttempts, 0)

NEXT:
	tmpWebhookAttempts := make(WebhookAttempts, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpWebhookAttempts)

	if apiErr != nil {
		return nil, apiErr
	}

	webhookAttempts = append(webhookAttempts, tmpWebhookAttempts...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, webhookAttempt := range tmpWebhookAttempts {
		webhookAttempt.Connection = c.Connection
	}

	return webhookAttempts, nil
}

func (c *WebhookAttempt) ReAttempt(webhookId int64) error {

		endpoint := invdendpoint.WebhookEndpoint + "/" + strconv.FormatInt(webhookId, 10) + "/retries"

		err := c.create(endpoint, nil, nil)
		if err != nil {
			return err
		}

		return  nil
	}
