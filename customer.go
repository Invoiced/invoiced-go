package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"log"
)

type Customer struct {
	*Connection
	*invdendpoint.Customer
}

type Customers []*Customer

func (c *Connection) NewCustomer() *Customer {
	customer := new(invdendpoint.Customer)
	return &Customer{c, customer}

}

func (c *Customer) Count() (int64, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Customer) Create(customer *Customer) (*Customer, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	custResp := new(Customer)

	apiErr := c.create(endPoint, customer, custResp)

	if apiErr != nil {
		return nil, apiErr
	}

	custResp.Connection = c.Connection

	return custResp, nil

}

func (c *Customer) Delete() *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Customer) Save() *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), c.Id)
	custResp := new(Customer)
	apiErr := c.update(endPoint, c, custResp)

	if apiErr != nil {
		return apiErr
	}

	c.Customer = custResp.Customer

	return nil

}

func (c *Customer) Retrieve(id int64) (*Customer, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.CustomersEndPoint), id)

	custEndPoint := new(invdendpoint.Customer)

	customer := &Customer{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, customer)

	if apiErr != nil {
		return nil, apiErr
	}

	return customer, nil

}

func (c *Customer) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := make(Customers, 0)

	log.Println("customer connection => ", c.Connection)

NEXT:
	tmpCustomers := make(Customers, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpCustomers)

	if apiErr != nil {
		return nil, apiErr
	}

	customers = append(customers, tmpCustomers...)

	if endPoint != "" {
		goto NEXT
	}

	for _, customer := range customers {
		customer.Connection = c.Connection

	}

	return customers, nil

}

func (c *Customer) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Customers, string, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.CustomersEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	customers := make(Customers, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &customers)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, customer := range customers {
		customer.Connection = c.Connection

	}

	return customers, nextEndPoint, nil

}

func (c *Customer) ListCustomersByName(customerName string) (Customers, *APIError) {

	filter := invdendpoint.NewFilter()
	filter.Set("name", customerName)

	customers, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	return customers, nil

}

func (c *Customer) ListCustomerByNumber(customerNumber string) (*Customer, *APIError) {

	filter := invdendpoint.NewFilter()
	filter.Set("number", customerNumber)

	customers, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return customers[0], nil

}
