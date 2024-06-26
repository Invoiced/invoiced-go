package invoiced

import (
	"encoding/json"
	"errors"
	"strings"
)

type Event struct {
	Id        int64           `json:"id"`
	Object    string          `json:"object"`
	Type      string          `json:"type"`
	Timestamp int64           `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
	User      *User           `json:"user"`
}

type EventObject struct {
	Object         *json.RawMessage `json:"object"`
	PreviousObject *json.RawMessage `json:"previous"`
}

type Events []*Event

func (e *Event) ParseEventObject() (*json.RawMessage, error) {
	data := e.Data

	eo := new(EventObject)

	b, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, eo)

	if err != nil {
		return nil, err
	}

	if eo.Object == nil {
		return nil, errors.New("Could not parse event object")
	}

	return eo.Object, nil
}

func (e *Event) ParseEventPreviousObject() (*json.RawMessage, error) {
	data := e.Data

	eo := new(EventObject)

	b, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, eo)

	if err != nil {
		return nil, err
	}

	if eo.Object == nil {
		return nil, errors.New("Could not parse event object")
	}

	return eo.PreviousObject, nil
}

func (e *Event) ParseInvoiceEvent() (*Invoice, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Invoice)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseInvoicePreviousEvent() (*Invoice, error) {
	eoData, err := e.ParseEventPreviousObject()
	if err != nil {
		return nil, err
	}

	if eoData == nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Invoice)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseCreditNoteEvent() (*CreditNote, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(CreditNote)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseCreditNotePreviousEvent() (*CreditNote, error) {
	eoData, err := e.ParseEventPreviousObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(CreditNote)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseCustomerPreviousEvent() (*Customer, error) {
	eoData, err := e.ParseEventPreviousObject()
	if err != nil {
		return nil, err
	}

	if eoData == nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Customer)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseSubscriptionEvent() (*Subscription, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Subscription)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParsePaymentEvent() (*Payment, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Payment)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParsePaymentPreviousEvent() (*Payment, error) {
	eoData, err := e.ParseEventPreviousObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Payment)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func (e *Event) ParseCustomerEvent() (*Customer, error) {
	eoData, err := e.ParseEventObject()
	if err != nil {
		return nil, err
	}

	b, err := eoData.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bClean := CleanMetaDataArray(b)

	ie := new(Customer)

	err = json.Unmarshal(bClean, ie)

	if err != nil {
		return nil, err
	}

	return ie, nil
}

func CleanMetaDataArray(b []byte) []byte {
	s := string(b)
	s1 := strings.Replace(s, `"metadata": []`, ` "metadata": null`, -1)
	s1 = strings.Replace(s1, `"metadata":[]`, ` "metadata": null`, -1)
	return []byte(s1)
}
