package api

import (
	"github.com/strongdm/invoiced-go/v2"
	"github.com/strongdm/invoiced-go/v2/charge"
	"github.com/strongdm/invoiced-go/v2/chasing"
	"github.com/strongdm/invoiced-go/v2/coupon"
	"github.com/strongdm/invoiced-go/v2/credit_balance_adjustment"
	"github.com/strongdm/invoiced-go/v2/credit_note"
	"github.com/strongdm/invoiced-go/v2/customer"
	"github.com/strongdm/invoiced-go/v2/estimate"
	"github.com/strongdm/invoiced-go/v2/event"
	"github.com/strongdm/invoiced-go/v2/file"
	"github.com/strongdm/invoiced-go/v2/invoice"
	"github.com/strongdm/invoiced-go/v2/item"
	"github.com/strongdm/invoiced-go/v2/member"
	"github.com/strongdm/invoiced-go/v2/note"
	"github.com/strongdm/invoiced-go/v2/notification"
	"github.com/strongdm/invoiced-go/v2/payment"
	"github.com/strongdm/invoiced-go/v2/plan"
	"github.com/strongdm/invoiced-go/v2/role"
	"github.com/strongdm/invoiced-go/v2/subscription"
	"github.com/strongdm/invoiced-go/v2/task"
	"github.com/strongdm/invoiced-go/v2/tax_rate"
	"github.com/strongdm/invoiced-go/v2/webhook_attempt"
)

type Client struct {
	Api                     *invoiced.Api
	Charge                  charge.Client
	ChasingCadence          chasing.Client
	Coupon                  coupon.Client
	CreditBalanceAdjustment credit_balance_adjustment.Client
	CreditNote              credit_note.Client
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
	Task                    task.Client
	TaxRate                 tax_rate.Client
	WebhookAttempt          webhook_attempt.Client
}

func New(key string, sandbox bool) *Client {
	apiClient := invoiced.New(key, sandbox)

	return &Client{
		Api:                     apiClient,
		Charge:                  charge.Client{Api: apiClient},
		ChasingCadence:          chasing.Client{Api: apiClient},
		Coupon:                  coupon.Client{Api: apiClient},
		CreditBalanceAdjustment: credit_balance_adjustment.Client{Api: apiClient},
		CreditNote:              credit_note.Client{Api: apiClient},
		Customer:                customer.Client{Api: apiClient},
		Estimate:                estimate.Client{Api: apiClient},
		Event:                   event.Client{Api: apiClient},
		File:                    file.Client{Api: apiClient},
		Invoice:                 invoice.Client{Api: apiClient},
		Item:                    item.Client{Api: apiClient},
		Member:                  member.Client{Api: apiClient},
		Note:                    note.Client{Api: apiClient},
		Notification:            notification.Client{Api: apiClient},
		Payment:                 payment.Client{Api: apiClient},
		Plan:                    plan.Client{Api: apiClient},
		Role:                    role.Client{Api: apiClient},
		Subscription:            subscription.Client{Api: apiClient},
		Task:                    task.Client{Api: apiClient},
		TaxRate:                 tax_rate.Client{Api: apiClient},
		WebhookAttempt:          webhook_attempt.Client{Api: apiClient},
	}
}
