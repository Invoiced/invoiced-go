package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

func (c *Connection) ListAllCustomersAuto(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (*invdendpoint.Customers, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := new(invdendpoint.Customers)

NEXT:
	tmpCustomers := new(invdendpoint.Customers)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, tmpCustomers)

	if apiErr != nil {
		return nil, apiErr
	}

	*customers = append(*customers, *tmpCustomers...)

	if endPoint != "" {
		goto NEXT
	}

	return customers, apiErr

}

func (c *Connection) ListAllCustomers(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (*invdendpoint.Customers, string, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := new(invdendpoint.Customers)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, customers)

	if apiErr != nil {
		return nil, "", apiErr
	}

	return customers, nextEndPoint, apiErr

}

func (c *Connection) ListCustomer(id int64) (*invdendpoint.Customer, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), id)

	customer := new(invdendpoint.Customer)

	_, apiErr := c.retrieveDataFromAPI(endPoint, customer)

	if apiErr != nil {
		return nil, apiErr
	}

	return customer, apiErr

}

func (c *Connection) ListCustomersByName(customerName string) (*invdendpoint.Customers, *APIError) {

	filter := invdendpoint.NewFilter()
	filter.Set("name", customerName)

	invdCustomers, apiError := c.ListAllCustomersAuto(filter, nil)

	return invdCustomers, apiError

}

func (c *Connection) ListCustomerByNumber(customerNumber string) (*invdendpoint.Customer, *APIError) {

	filter := invdendpoint.NewFilter()
	filter.Set("number", customerNumber)

	invdCustomers, apiError := c.ListAllCustomersAuto(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(*invdCustomers) == 0 {
		return nil, nil
	}

	return &((*invdCustomers)[0]), nil

}

func (c *Connection) CountCustomer() (int64, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)

	count, apiErr := c.count(endPoint)

	return count, apiErr

}

func (c *Connection) CreateCustomer(customer *invdendpoint.Customer) (*invdendpoint.Customer, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	customerResponse := new(invdendpoint.Customer)

	apiErr := c.create(endPoint, customer, customerResponse)

	if apiErr != nil {
		return nil, apiErr
	}

	return customerResponse, apiErr

}

func (c *Connection) UpdateCustomer(id int64, customer *invdendpoint.Customer) (*invdendpoint.Customer, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), id)
	customerResponse := new(invdendpoint.Customer)
	apiErr := c.update(endPoint, customer, customerResponse)

	if apiErr != nil {
		return nil, apiErr
	}

	return customerResponse, apiErr

}

func (c *Connection) DeleteCustomer(id int64) *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return apiErr

}
