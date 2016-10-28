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

	if so.Id != 15444 {
		t.Fatal("Id is incorrect")
	}

	if so.Name != "Acme" {
		t.Fatal("Name is incorrect")
	}

	if so.Number != "CUST-0001" {
		t.Fatal("Number is incorrects")
	}

	if so.CollectionMode != "auto" {
		t.Fatal("Collection Mode is incorrect")
	}

	if so.PaymentSourceRAW == nil {
		t.Fatal("Payment Source RAW is nil")
	}

	if so.Type != "company" {

		t.Fatal("Type is incorrect")
	}

	if so.StatementPdfUrl != "https://dundermifflin.invoiced.com/statements/t3NmhUomra3g3ueSNnbtUgrr/pdf" {
		t.Fatal("Statement PDF is incorrect")
	}

	if so.CreatedAt != 1415222128 {
		t.Fatal("Created At is incorrect")
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
    "last4": "4242",
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
  "last4": "7890",
  "routing_number": "110000000",
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

func TestCustomerUnmarshalCardObject(t *testing.T) {

	s := `{
  "id": 850,
  "object": "card",
  "brand": "Visa",
  "last4": "4242",
  "exp_month": 2,
  "exp_year": 20,
  "funding": "credit"
}`

	so := new(CardObject)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 850 {
		t.Fatal("Id is incorrect")
	}

	if so.Object != "card" {
		t.Fatal("Object is incorrect")
	}

	if so.Brand != "Visa" {
		t.Fatal("Brand is incorrect")
	}

	if so.Last4 != "4242" {
		t.Fatal("Last 4 is incorrect")
	}

	if so.ExpMonth != 2 {
		t.Fatal("ExpMonth is incorrect")
	}

	if so.ExpYear != 20 {
		t.Fatal("ExpYear in incorrect")
	}

	if so.Funding != "credit" {
		t.Fatal("Funding is incorrect")
	}

}

func TestCustomerUnmarshalBankObject(t *testing.T) {

	s := `{
  "id": 4321,
  "object": "card",
  "bank_name": "Wells Fargo",
  "last4": "7890",
  "routing_number": "110000000",
  "verified": true,
  "currency": "usd"
}`

	so := new(BankAccountObject)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 4321 {
		t.Fatal("Id is incorrect")
	}

	if so.Object != "card" {
		t.Fatal("Object is incorrect")
	}

	if so.BankName != "Wells Fargo" {
		t.Fatal("Bank Name is incorrect")
	}

	if so.Last4 != "7890" {
		t.Fatal("Last 4 is incorrect")
	}

	if so.RoutingNumber != "110000000" {
		t.Fatal("Routing Number is incorrect")
	}

	if !so.Verified {
		t.Fatal("Verified in incorrect")
	}

	if so.Currency != "usd" {
		t.Fatal("Currency is incorrect")
	}

}
