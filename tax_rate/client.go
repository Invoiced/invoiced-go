package tax_rate

import "github.com/Invoiced/invoiced-go"

type Client struct {
	*invoiced.Api
}

type TaxRates []*Client

func (c *Client) Create(request *invoiced.TaxRateRequest) (*Client, error) {
	endpoint := invoiced.RateEndpoint
	resp := new(Client)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Retrieve(id string) (*Client, error) {
	endpoint := invoiced.RateEndpoint + "/" + id
	taxRate := &Client{c.Client, new(invoiced.TaxRate)}

	_, err := c.Api.Get(endpoint, taxRate)
	if err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (c *Client) Update(request *invoiced.TaxRateRequest) error {
	endpoint := invoiced.RateEndpoint + "/" + c.Id
	resp := new(Client)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.TaxRate = resp.TaxRate

	return nil
}

func (c *Client) Delete() error {
	endpoint := invoiced.RateEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (TaxRates, error) {
	endpoint := invoiced.RateEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

	taxRates := make(TaxRates, 0)

NEXT:
	tmpTaxRates := make(TaxRates, 0)

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
