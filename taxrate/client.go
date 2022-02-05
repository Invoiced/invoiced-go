package taxrate

import (
	"github.com/Invoiced/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.TaxRateRequest) (*invoiced.TaxRate, error) {
	resp := new(invoiced.TaxRate)
	err := c.Api.Create("/tax_rates", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id string) (*invoiced.TaxRate, error) {
	resp := new(invoiced.TaxRate)
	_, err := c.Api.Get("/tax_rates/"+id, resp)
	return resp, err
}

func (c *Client) Update(id string, request *invoiced.TaxRateRequest) (*invoiced.TaxRate, error) {
	resp := new(invoiced.TaxRate)
	err := c.Api.Update("/tax_rates/"+id, request, resp)
	return resp, err
}

func (c *Client) Delete(id string) error {
	return c.Api.Delete("/tax_rates/" + id)
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.TaxRates, error) {
	endpoint := invoiced.AddFilterAndSort("/tax_rates", filter, sort)

	taxRates := make(invoiced.TaxRates, 0)

NEXT:
	tmpTaxRates := make(invoiced.TaxRates, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpTaxRates)
	if err != nil {
		return nil, err
	}

	taxRates = append(taxRates, tmpTaxRates...)

	if endpointTmp != "" {
		goto NEXT
	}

	return taxRates, nil
}
