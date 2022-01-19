package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type TaxRate struct {
	*Connection
	*invdendpoint.TaxRate
}

type TaxRates []*TaxRate

func (c *Connection) NewTaxRate() *TaxRate {
	taxRate := new(invdendpoint.TaxRate)
	return &TaxRate{c, taxRate}
}

func (c *TaxRate) Create(request *invdendpoint.TaxRateRequest) (*TaxRate, error) {
	endpoint := invdendpoint.RateEndpoint
	resp := new(TaxRate)

	err := c.create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *TaxRate) Retrieve(id string) (*TaxRate, error) {
	endpoint := invdendpoint.RateEndpoint + "/" + id
	taxRate := &TaxRate{c.Connection, new(invdendpoint.TaxRate)}

	_, err := c.retrieveDataFromAPI(endpoint, taxRate)
	if err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (c *TaxRate) Update(request *invdendpoint.TaxRateRequest) error {
	endpoint := invdendpoint.RateEndpoint + "/" + c.Id
	resp := new(TaxRate)

	err := c.update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.TaxRate = resp.TaxRate

	return nil
}

func (c *TaxRate) Delete() error {
	endpoint := invdendpoint.RateEndpoint + "/" + c.Id

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *TaxRate) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (TaxRates, error) {
	endpoint := invdendpoint.RateEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	taxRates := make(TaxRates, 0)

NEXT:
	tmpTaxRates := make(TaxRates, 0)

	endpointTmp, err := c.retrieveDataFromAPI(endpoint, &tmpTaxRates)

	if err != nil {
		return nil, err
	}

	taxRates = append(taxRates, tmpTaxRates...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, taxRate := range taxRates {
		taxRate.Connection = c.Connection
	}

	return taxRates, nil
}
