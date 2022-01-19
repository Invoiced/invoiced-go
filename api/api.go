package api

import (
	"github.com/Invoiced/invoiced-go"
	"github.com/Invoiced/invoiced-go/charge"
	"github.com/Invoiced/invoiced-go/chasing"
	"github.com/Invoiced/invoiced-go/coupon"
	"github.com/Invoiced/invoiced-go/credit_note"
	"github.com/Invoiced/invoiced-go/customer"
	"github.com/Invoiced/invoiced-go/estimate"
	"github.com/Invoiced/invoiced-go/event"
	"github.com/Invoiced/invoiced-go/file"
	"github.com/Invoiced/invoiced-go/invoice"
	"github.com/Invoiced/invoiced-go/item"
	"github.com/Invoiced/invoiced-go/note"
	"github.com/Invoiced/invoiced-go/notification"
	"github.com/Invoiced/invoiced-go/payment"
	"github.com/Invoiced/invoiced-go/plan"
	"github.com/Invoiced/invoiced-go/role"
	"github.com/Invoiced/invoiced-go/subscription"
	"github.com/Invoiced/invoiced-go/task"
	"github.com/Invoiced/invoiced-go/tax_rate"
	"github.com/Invoiced/invoiced-go/webhook_attempt"
)

type Client struct {
	Api            *invoiced.Api
	Charge         charge.Client
	ChasingCadence chasing.Client
	Coupon coupon.Client
	CreditNote credit_note.Client
	Customer       customer.Client
	Estimate estimate.Client
	Event event.Client
	File file.Client
	Invoice        invoice.Client
	Item item.Client
	Note note.Client
	Notification notification.Client
	Payment payment.Client
	Plan plan.Client
	Role role.Client
	Subscription subscription.Client
	Task task.Client
	TaxRate tax_rate.Client
	WebhookAttempt webhook_attempt.Client
}

func New(key string, sandbox bool) *Client {
	apiClient := invoiced.New(key, sandbox)

	return &Client{
		Api:            apiClient,
		Charge:         charge.Client{apiClient},
		ChasingCadence: chasing.Client{apiClient},
		Coupon: coupon.Client{apiClient},
		CreditNote: credit_note.Client{apiClient},
		Customer:       customer.Client{apiClient},
		Estimate: estimate.Client{apiClient},
		Event: event.Client{apiClient},
		File: file.Client{apiClient},
		Invoice:        invoice.Client{apiClient},
		Item: item.Client{apiClient},
		Note: note.Client{apiClient},
		Notification: notification.Client{apiClient},
		Payment: payment.Client{apiClient},
		Plan: plan.Client{apiClient},
		Role: role.Client{apiClient},
		Subscription: subscription.Client{apiClient},
		Task: task.Client{apiClient},
		TaxRate: tax_rate.Client{apiClient},
		WebhookAttempt: webhook_attempt.Client{apiClient},
	}
}
