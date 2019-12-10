package invdendpoint

import (
	"encoding/json"
	"strings"
)

const CustomersEndPoint = "/customers"

type Customers []Customer

type Customer struct {
	Id                 int64                  `json:"id,omitempty"`                //The customer’s unique ID
	Object 			   string                 `json:"object,omitempty"`
	Name               string                 `json:"name,omitempty"`              //Customer name
	Number             string                 `json:"number,omitempty"`            //A unique ID to help tie your customer to your external systems
	Email              string                 `json:"email,omitempty"`             //Email address
	AutoPay            bool                   `json:"autopay,omitempty"`   		   //Autopay
	AutoPayDelays	   int  				  `json:"autopay_delay_days,omitempty"`//Number of days to delay AutoPay
	PaymentTerms       string                 `json:"payment_terms,omitempty"`     //Payment terms used for manual collection mode, i.e. “NET 30”
	StripeToken        string    			  `json:"stripe_token,omitempty"`      //When provided sets the customer’s payment source to the tokenized Stripe card
	PaymentSourceRAW   json.RawMessage        `json:"payment_source,omitempty"`    //Holds the raw payment json information for later parsing
	PaymentSource      *PaymentSource         `json:"-"`                           //Customer’s payment source, if attached
	PaymentSourceReady bool                   `json:"-"`                           //If true than run UnbundlePaymentSource()
	AttentionTo        string                 `json:"attention_to,omitempty"`      //Used for ATTN: address line if company
	Address1           string                 `json:"address1,omitempty"`          //First address line
	Address2           string                 `json:"address2,omitempty"`          //Optional second address line
	City               string                 `json:"city,omitempty"`              //City
	State              string                 `json:"state,omitempty"`             //State or province
	PostalCode         string                 `json:"postal_code,omitempty"`       //Zip or postal code
	Country            string                 `json:"country,omitempty"`           //Two-letter ISO code
	Language           string                 `json:"language,omitempty"`          //Two-letter ISO code
	Chase              bool                   `json:"boolean,omitempty"`   		   //Chasing enabled? - defaults to true
	ChasingCadence     int64				  `json:"chasing_cadence,omitempty"`   //Cadence ID
	Phone              string                 `json:"phone,omitempty"`             //Phone #
	CreditHold         bool                   `json:"credit_hold,omitempty"`        //When true, customer is on credit hold
	CreditLimit        float64                `json:"credit_limit,omitempty"`       //Customer credit limit
	Owner              int64                  `json:"owner,omitempty"`             //Customer credit limit
	Taxable            bool                   `json:"taxable,omitempty"`           //Customer taxable?
	Taxes              []Rate                 `json:"taxes,omitempty"`             //Collection of Tax Rate IDs
	TaxId			   string                 `json:"taxid,omitempty"`             //Tax ID to be displayed on documents
	AvalaraEntityUseCode string               `json:"avalara_entity_use_code,omitempty"` //Avalara-specific entity use code
	AvalaraExemptionNumber string             `json:"avalara_exemption_number,omitempty"`// Tax-exempt number to pass to Avalara
	Type               string                 `json:"type,omitempty"`              //Organization type, company or person
	ParentCustomer      int64                 `json:"parent_customer,omitempty"`   //Parent customer ID
	Notes              string                 `json:"notes,omitempty"`             //Private customer notes
	SignUpPage         string                 `json:"sign_up_page,omitempty"`
	SignUpUrl          string                 `json:"sign_up_url,omitempty"` //URL to download the latest account statement
	StatementPdfUrl    string                 `json:"statement_pdf_url,omitempty"` //URL to download the latest account statement
	CreatedAt          int64                  `json:"created_at,omitempty"`        //Timestamp when created
	MetaData           map[string]interface{} `json:"metadata,omitempty"`          //A hash of key/value pairs that can store additional information about this object.
	DisabledPaymentMethods []string           `json:"disabled_payment_methods,omitempty"`
}

type PaymentSource struct {
	CardObject        *CardObject
	BankAccountObject *BankAccountObject
	Type              string
}

type CardObject struct {
	Id       int64  `json:"id,omitempty"`        //The card’s unique ID
	Object   string `json:"object,omitempty"`    //card
	Brand    string `json:"brand,omitempty"`     //Card brand
	Last4    string `json:"last4,omitempty"`     //Last 4 digits of card
	ExpMonth int    `json:"exp_month,omitempty"` //Expiry month
	ExpYear  int    `json:"exp_year,omitempty"`  //Expiry year
	Funding  string `json:"funding,omitempty"`   //Funding instrument, can be credit, debit, prepaid, or unknown
}

type BankAccountObject struct {
	Id            int64  `json:"id,omitempty"`             //The bank account’s unique ID
	Object        string `json:"object,omitempty"`         //bank_account
	BankName      string `json:"bank_name,omitempty"`      //Bank name
	Last4         string `json:"last4,omitempty"`          //Last 4 digits of bank account
	RoutingNumber string `json:"routing_number,omitempty"` //Bank routing number
	Verified      bool   `json:"verified,omitempty"`       //Whether the bank account has been verified with instant verification or micro-deposits
	Currency      string `json:"currency,omitempty"`       //3-letter ISO code
}

func (c *Customer) addSource() error {
	s := string(c.PaymentSourceRAW)

	if c.PaymentSourceReady {
		return nil
	}

	if strings.Contains(s, "card") {
		card := new(CardObject)
		err := json.Unmarshal(c.PaymentSourceRAW, card)

		if err != nil {
			return err
		}
		paymentSource := new(PaymentSource)
		c.PaymentSource = paymentSource
		c.PaymentSource.CardObject = card
		c.PaymentSource.Type = "card"
		c.PaymentSourceReady = true

	} else if strings.Contains(s, "bank_account") {
		bankAcct := new(BankAccountObject)
		err := json.Unmarshal(c.PaymentSourceRAW, bankAcct)

		if err != nil {
			return err
		}
		paymentSource := new(PaymentSource)
		c.PaymentSource = paymentSource
		c.PaymentSource.BankAccountObject = bankAcct
		c.PaymentSource.Type = "bank_account"
		c.PaymentSourceReady = true
	}

	return nil

}

//We only need to run this if c.PaymentSourceReady is false
func (c *Customer) UnbundlePaymentSource() (*PaymentSource, error) {
	if c.PaymentSourceReady {
		return c.PaymentSource, nil
	}

	err := c.addSource()

	if err != nil {
		return nil, err
	}

	return c.PaymentSource, nil

}

func (c *Customer) BundlePaymentSource() error {

	if len(c.PaymentSourceRAW) != 0 || c.PaymentSource == nil {
		return nil
	}

	if c.PaymentSource.Type == "bank_account" {

		b, err := json.Marshal(c.PaymentSource.BankAccountObject)

		if err != nil {
			return err
		}

		c.PaymentSourceRAW = b
		c.PaymentSourceReady = true

	} else if c.PaymentSource.Type == "card" {

		b, err := json.Marshal(c.PaymentSource.CardObject)

		if err != nil {
			return err
		}

		c.PaymentSourceRAW = b
		c.PaymentSourceReady = true

	}

	return nil

}

func (c *Customer) String() string {

	b, _ := json.MarshalIndent(c, "", "    ")

	return string(b)
}
