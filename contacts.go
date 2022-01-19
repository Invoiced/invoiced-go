package invoiced

type ContactRequest struct {
	Address1   *string `json:"address1,omitempty"`
	Address2   *string `json:"address2,omitempty"`
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	Department *string `json:"department,omitempty"`
	Email      *string `json:"email,omitempty"`
	Name       *string `json:"name,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Primary    *bool   `json:"primary"`
	SmsEnabled *bool   `json:"sms_enabled"`
	State      *string `json:"state,omitempty"`
	Title      *string `json:"title,omitempty"`
}

type Contact struct {
	Address1   *string `json:"address1"`
	Address2   *string `json:"address2"`
	City       *string `json:"city"`
	Country    *string `json:"country"`
	CreatedAt  int64   `json:"created_at"`
	Department *string `json:"department"`
	Email      *string `json:"email"`
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Object     string  `json:"object"`
	Phone      *string `json:"phone"`
	PostalCode *string `json:"postal_code"`
	Primary    bool    `json:"primary"`
	SmsEnabled bool    `json:"sms_enabled"`
	State      *string `json:"state"`
	Title      *string `json:"title"`
	UpdatedAt  int64   `json:"updated_at"`
}

type Contacts []*Contact
