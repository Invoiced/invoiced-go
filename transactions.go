package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

func (c *Connection) ListTransaction(id int64) (*invdendpoint.Transaction, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), id)

	transaction := new(invdendpoint.Transaction)

	_, apiErr := c.retrieveDataFromAPI(endPoint, transaction)

	if apiErr != nil {
		return nil, apiErr
	}

	return transaction, apiErr

}

func (c *Connection) ListAllTransactionAuto(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (*invdendpoint.Transactions, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	transactions := new(invdendpoint.Transactions)

NEXT:
	tmpTransactions := new(invdendpoint.Transactions)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, tmpTransactions)

	if apiErr != nil {
		return nil, apiErr
	}

	*transactions = append(*transactions, *tmpTransactions...)

	if endPoint != "" {
		goto NEXT
	}

	return transactions, apiErr

}

func (c *Connection) DeleteTransaction(id int64) *APIError {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), id)

	apiErr := c.delete(endPoint)

	return apiErr

}

func (c *Connection) UpdateTransaction(id int64, transactionToUpdate *invdendpoint.Transaction) (*invdendpoint.Transaction, *APIError) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), id)
	transactionCreated := new(invdendpoint.Transaction)

	apiErr := c.update(endPoint, transactionToUpdate, transactionCreated)

	return transactionCreated, apiErr

}

func (c *Connection) CreateTransaction(transactionToCreate *invdendpoint.Transaction) (*invdendpoint.Transaction, *APIError) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)

	transactionCreated := new(invdendpoint.Transaction)

	apiErr := c.create(endPoint, transactionToCreate, transactionCreated)

	return transactionCreated, apiErr

}
