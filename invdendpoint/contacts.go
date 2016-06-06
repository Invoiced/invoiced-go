package invdendpoint

//Contacts can be attached to customers. A contact could represent an additional email recipient for a customer, or perhaps an address in addition to the billing address, like a shipping address.
type Contact struct {
	Id         int64  `json:"id,omitempty"`          //The customer’s unique ID
	Name       string `json:"name,omitempty"`        //Contact name
	Email      string `json:"email,omitempty"`       //Email address
	Primary    bool   `json:"primary,omitempty"`     //When true the contact will be copied on any account communications
	Address1   string `json:"address1,omitempty"`    //First address line
	Address2   string `json:"address2,omitempty"`    //Optional second address line
	City       string `json:"city,omitempty"`        //City
	State      string `json:"state,omitempty"`       //State or province
	PostalCode string `json:"postal_code,omitempty"` //Zip or postal code
	Country    string `json:"country,omitempty"`     //Two-letter ISO code
	CreatedAt  int64  `json:"created_at,omitempty"`  //Timestamp when created
}

type Contacts []Contact
