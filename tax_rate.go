package invoiced

type TaxRateClient struct {
	*Client
	*TaxRate
}

type TaxRates []*TaxRateClient

func (c *Client) NewTaxRate() *TaxRateClient {
	taxRate := new(TaxRate)
	return &TaxRateClient{c, taxRate}
}

func (c *TaxRateClient) Create(request *TaxRateRequest) (*TaxRateClient, error) {
	endpoint := RateEndpoint
	resp := new(TaxRateClient)

	err := c.Api.Create(endpoint, request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TaxRateClient) Retrieve(id string) (*TaxRateClient, error) {
	endpoint := RateEndpoint + "/" + id
	taxRate := &TaxRateClient{c.Client, new(TaxRate)}

	_, err := c.Api.Get(endpoint, taxRate)
	if err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (c *TaxRateClient) Update(request *TaxRateRequest) error {
	endpoint := RateEndpoint + "/" + c.Id
	resp := new(TaxRateClient)

	err := c.Api.Update(endpoint, request, resp)
	if err != nil {
		return err
	}

	c.TaxRate = resp.TaxRate

	return nil
}

func (c *TaxRateClient) Delete() error {
	endpoint := RateEndpoint + "/" + c.Id

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *TaxRateClient) ListAll(filter *Filter, sort *Sort) (TaxRates, error) {
	endpoint := RateEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

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
