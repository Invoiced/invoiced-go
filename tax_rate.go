package invdapi

import (
	"errors"
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

func (c *TaxRate) Create(taxRate *TaxRate) (*TaxRate, error) {
	endpoint := invdendpoint.RateEndpoint

	taxRateResp := new(TaxRate)

	if taxRate == nil {
		return nil, errors.New("taxRate is nil")
	}

	// safe prune file data for creation
	invdTaxRateDataToCreate, err := SafeTaxRateForCreation(taxRate.TaxRate)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdTaxRateDataToCreate, taxRateResp)

	if apiErr != nil {
		return nil, apiErr
	}

	taxRateResp.Connection = c.Connection

	return taxRateResp, nil
}

func (c *TaxRate) Save() error {
	endpoint := invdendpoint.RateEndpoint + "/" + c.Id

	taxRateResp := new(TaxRate)

	invdTaxRatDataToUpdate, err := SafeTaxRateForUpdating(c.TaxRate)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, invdTaxRatDataToUpdate, taxRateResp)

	if apiErr != nil {
		return apiErr
	}

	c.TaxRate = taxRateResp.TaxRate

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

func (c *TaxRate) Retrieve(id string) (*TaxRate, error) {
	endpoint := invdendpoint.RateEndpoint + "/" + id

	taxRateEndpoint := new(invdendpoint.TaxRate)

	taxRate := &TaxRate{c.Connection, taxRateEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, taxRate)
	if err != nil {
		return nil, err
	}

	return taxRate, nil
}

func (c *TaxRate) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (TaxRates, error) {
	endpoint := invdendpoint.RateEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	taxRates := make(TaxRates, 0)

NEXT:
	tmpTaxRates := make(TaxRates, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpTaxRates)

	if apiErr != nil {
		return nil, apiErr
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

// SafetaxRateForCreation prunes tax TaxRate data for just fields that can be used for creation of a tax TaxRate
func SafeTaxRateForCreation(taxRate *invdendpoint.TaxRate) (*invdendpoint.TaxRate, error) {
	if taxRate == nil {
		return nil, errors.New("taxRate is nil")
	}

	taxRateData := new(invdendpoint.TaxRate)
	taxRateData.Id = taxRate.Id
	taxRateData.Name = taxRate.Name
	taxRateData.Currency = taxRate.Currency
	taxRateData.Value = taxRate.Value
	taxRateData.Inclusive = taxRate.Inclusive
	taxRateData.IsPercent = taxRate.IsPercent
	taxRateData.Metadata = taxRate.Metadata

	return taxRateData, nil
}

// SafeTaxRateForUpdating prunes plan data for just fields that can be used for updating of a plan
func SafeTaxRateForUpdating(taxRate *invdendpoint.TaxRate) (*invdendpoint.TaxRate, error) {
	if taxRate == nil {
		return nil, errors.New("taxRate is nil")
	}

	taxRateData := new(invdendpoint.TaxRate)
	taxRateData.Name = taxRate.Name
	taxRateData.Metadata = taxRate.Metadata

	return taxRateData, nil
}
