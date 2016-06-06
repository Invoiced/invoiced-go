package invdendpoint

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnMarshalCustomerObject(t *testing.T) {
	s := `{
  "id": 15444,
  "name": "Acme",
  "number": "CUST-0001",
  "email": "billing@acmecorp.com",
  "collection_mode": "auto",
  "payment_terms": null,
  "payment_source": {
    "id": 850,
    "object": "card",
    "brand": "Visa",
    "last4": 4242,
    "exp_month": 2,
    "exp_year": 20,
    "funding": "credit"
  },
  "taxes": [],
  "type": "company",
  "attention_to": null,
  "address1": null,
  "address2": null,
  "city": null,
  "state": null,
  "postal_code": null,
  "country": "US",
  "tax_id": null,
  "phone": null,
  "notes": null,
  "statement_pdf_url": "https://dundermifflin.invoiced.com/statements/t3NmhUomra3g3ueSNnbtUgrr/pdf",
  "created_at": 1415222128,
  "metadata": {}
}`

	so := new(Customer)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}

func TestCustomerUnbundlePaymentSourceCreditCard(t *testing.T) {
	sampleCustomer := `{
  "id": 15444,
  "name": "Acme",
  "number": "CUST-0001",
  "email": "billing@acmecorp.com",
  "collection_mode": "auto",
  "payment_terms": null,
  "payment_source": {
    "id": 850,
    "object": "card",
    "brand": "Visa",
    "last4": 4242,
    "exp_month": 2,
    "exp_year": 20,
    "funding": "credit"
  },
  "taxes": [],
  "type": "company",
  "attention_to": null,
  "address1": null,
  "address2": null,
  "city": null,
  "state": null,
  "postal_code": null,
  "country": "US",
  "tax_id": null,
  "phone": null,
  "notes": null,
  "statement_pdf_url": "https://dundermifflin.invoiced.com/statements/t3NmhUomra3g3ueSNnbtUgrr/pdf",
  "created_at": 1415222128,
  "metadata": {}
}`

	customer := new(Customer)
	err := json.Unmarshal([]byte(sampleCustomer), customer)

	if err != nil {
		t.Fatal(err)
	}

	paymentSource, err := customer.UnbundlePaymentSource()

	if err != nil {
		t.Fatal(err)
	}

	if paymentSource == nil {
		t.Fatal("PaymentSource should not be nil")
	}

	if paymentSource.Type != "card" {
		t.Fatal("Payment source should be a card")
	}

	correctCard := new(CardObject)
	err = json.Unmarshal(customer.PaymentSourceRAW, correctCard)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(paymentSource.CardObject, correctCard) {
		t.Fatal("Card do not match up")
	}

}

func TestCustomerUnbundlePaymentSourceBankAccount(t *testing.T) {
	sampleCustomer := `{
  "id": 15444,
  "name": "Acme",
  "number": "CUST-0001",
  "email": "billing@acmecorp.com",
  "collection_mode": "auto",
  "payment_terms": null,
  "payment_source": {
  "id": 4321,
  "object": "bank_account",
  "bank_name": "Wells Fargo",
  "last4": 7890,
  "routing_number": 110000000,
  "verified": true,
  "currency": "usd"
},
  "taxes": [],
  "type": "company",
  "attention_to": null,
  "address1": null,
  "address2": null,
  "city": null,
  "state": null,
  "postal_code": null,
  "country": "US",
  "tax_id": null,
  "phone": null,
  "notes": null,
  "statement_pdf_url": "https://dundermifflin.invoiced.com/statements/t3NmhUomra3g3ueSNnbtUgrr/pdf",
  "created_at": 1415222128,
  "metadata": {}
}`

	customer := new(Customer)
	err := json.Unmarshal([]byte(sampleCustomer), customer)

	if err != nil {
		t.Fatal(err)
	}

	paymentSource, err := customer.UnbundlePaymentSource()

	if err != nil {
		t.Fatal(err)
	}

	if paymentSource == nil {
		t.Fatal("PaymentSource should not be nil")
	}

	if paymentSource.Type != "bank_account" {
		t.Fatal("Payment source should be bank_account")
	}

	correctBankAcct := new(BankAccountObject)
	err = json.Unmarshal(customer.PaymentSourceRAW, correctBankAcct)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(paymentSource.BankAccountObject, correctBankAcct) {
		t.Fatal("Card do not match up")
	}

}
