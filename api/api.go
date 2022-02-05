package api

import (
	"github.com/Invoiced/invoiced-go/v2"
	"github.com/Invoiced/invoiced-go/v2/charge"
	"github.com/Invoiced/invoiced-go/v2/chasing"
	"github.com/Invoiced/invoiced-go/v2/coupon"
	"github.com/Invoiced/invoiced-go/v2/creditbalanceadjustment"
	"github.com/Invoiced/invoiced-go/v2/creditnote"
	"github.com/Invoiced/invoiced-go/v2/customer"
	"github.com/Invoiced/invoiced-go/v2/estimate"
	"github.com/Invoiced/invoiced-go/v2/event"
	"github.com/Invoiced/invoiced-go/v2/file"
	"github.com/Invoiced/invoiced-go/v2/invoice"
	"github.com/Invoiced/invoiced-go/v2/item"
	"github.com/Invoiced/invoiced-go/v2/member"
	"github.com/Invoiced/invoiced-go/v2/note"
	"github.com/Invoiced/invoiced-go/v2/notification"
	"github.com/Invoiced/invoiced-go/v2/payment"
	"github.com/Invoiced/invoiced-go/v2/plan"
	"github.com/Invoiced/invoiced-go/v2/role"
	"github.com/Invoiced/invoiced-go/v2/subscription"
	"github.com/Invoiced/invoiced-go/v2/task"
	"github.com/Invoiced/invoiced-go/v2/taxrate"
	"github.com/Invoiced/invoiced-go/v2/webhookattempt"
)

type Client struct {
	Api                     *invoiced.Api
	Charge                  charge.Client
	ChasingCadence          chasing.Client
	Coupon                  coupon.Client
	CreditBalanceAdjustment creditbalanceadjustment.Client
	CreditNote              creditnote.Client
	Customer                customer.Client
	Estimate                estimate.Client
	Event                   event.Client
	File                    file.Client
	Invoice                 invoice.Client
	Item                    item.Client
	Member                  member.Client
	Note                    note.Client
	Notification            notification.Client
	Payment                 payment.Client
	Plan                    plan.Client
	Role                    role.Client
	Subscription            subscription.Client
	Task           task.Client
	TaxRate        taxrate.Client
	WebhookAttempt webhookattempt.Client
}

func New(key string, sandbox bool) *Client {
	apiClient := invoiced.New(key, sandbox)

	return &Client{
		Api:                     apiClient,
		Charge:                  charge.Client{Api: apiClient},
		ChasingCadence:          chasing.Client{Api: apiClient},
		Coupon:                  coupon.Client{Api: apiClient},
		CreditBalanceAdjustment: creditbalanceadjustment.Client{Api: apiClient},
		CreditNote:              creditnote.Client{Api: apiClient},
		Customer:                customer.Client{Api: apiClient},
		Estimate:                estimate.Client{Api: apiClient},
		Event:                   event.Client{Api: apiClient},
		File:                    file.Client{Api: apiClient},
		Invoice:        invoice.Client{Api: apiClient},
		Item:           item.Client{Api: apiClient},
		Member:         member.Client{Api: apiClient},
		Note:           note.Client{Api: apiClient},
		Notification:   notification.Client{Api: apiClient},
		Payment:        payment.Client{Api: apiClient},
		Plan:           plan.Client{Api: apiClient},
		Role:           role.Client{Api: apiClient},
		Subscription:   subscription.Client{Api: apiClient},
		Task:           task.Client{Api: apiClient},
		TaxRate:        taxrate.Client{Api: apiClient},
		WebhookAttempt: webhookattempt.Client{Api: apiClient},
	}
}
