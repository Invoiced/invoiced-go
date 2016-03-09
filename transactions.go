package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"log"
	"strconv"
)

type Transaction struct {
	*Connection
	*invdendpoint.Transaction
}

type Transactions []*Transaction

func (c *Connection) NewTransaction() *Transaction {
	transaction := new(invdendpoint.Transaction)
	return &Transaction{c, transaction}

}

func (c *Transaction) Count() (int64, error) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Transaction) Create(transaction *Transaction) (*Transaction, error) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)
	txnResp := new(Transaction)

	apiErr := c.create(endPoint, transaction, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil

}

func (c *Transaction) Delete() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Transaction) Save() error {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), c.Id)
	txnResp := new(Transaction)
	apiErr := c.update(endPoint, c, txnResp)

	if apiErr != nil {
		return apiErr
	}

	c.Transaction = txnResp.Transaction

	return nil

}

func (c *Transaction) Retrieve(id int64) (*Transaction, error) {
	endPoint := makeEndPointSingular(c.makeEndPointURL(invdendpoint.TransactionsEndPoint), id)

	custEndPoint := new(invdendpoint.Transaction)

	transaction := &Transaction{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, transaction)

	if apiErr != nil {
		return nil, apiErr
	}

	return transaction, nil

}

func (c *Transaction) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Transactions, error) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	transactions := make(Transactions, 0)

NEXT:
	tmpTransactions := make(Transactions, 0)

	endPoint, apiErr := c.retrieveDataFromAPI(endPoint, &tmpTransactions)

	if apiErr != nil {
		return nil, apiErr
	}

	transactions = append(transactions, tmpTransactions...)

	if endPoint != "" {
		goto NEXT
	}

	for _, transaction := range transactions {
		transaction.Connection = c.Connection

	}

	return transactions, nil

}

func (c *Transaction) List(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Transactions, string, error) {
	endPoint := c.makeEndPointURL(invdendpoint.TransactionsEndPoint)
	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	transactions := make(Transactions, 0)

	nextEndPoint, apiErr := c.retrieveDataFromAPI(endPoint, &transactions)

	if apiErr != nil {
		return nil, "", apiErr
	}

	for _, transaction := range transactions {
		transaction.Connection = c.Connection

	}

	return transactions, nextEndPoint, nil

}

func (c *Transaction) ListByNumber(transactionNumber string) (*Transaction, error) {

	filter := invdendpoint.NewFilter()
	filter.Set("number", transactionNumber)

	transactions, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions[0], nil

}

func (c *Transaction) ListSuccessfulByInvoiceID(invoiceID int64) (Transactions, error) {

	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	log.Println("invoiceIDString", invoiceIDString)
	filter := invdendpoint.NewFilter()
	filter.Set("invoice", invoiceIDString)
	filter.Set("status", "succeeded")

	transactions, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions, nil

}

func (c *Transaction) ListSuccessfulChargesByInvoiceID(invoiceID int64) (Transactions, error) {

	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	log.Println("invoiceIDString", invoiceIDString)
	filter := invdendpoint.NewFilter()
	filter.Set("invoice", invoiceIDString)
	filter.Set("status", "succeeded")
	filter.Set("type", "charge")

	transactions, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions, nil

}

func (c *Transaction) ListSuccessfulPaymentsByInvoiceID(invoiceID int64) (Transactions, error) {

	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	log.Println("invoiceIDString", invoiceIDString)
	filter := invdendpoint.NewFilter()
	filter.Set("invoice", invoiceIDString)
	filter.Set("status", "succeeded")
	filter.Set("type", "payment")

	transactions, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions, nil

}

func (c *Transaction) ListSuccessfulChargesAndPaymentsByInvoiceID(invoiceID int64) (Transactions, error) {

	charges, err := c.ListSuccessfulChargesByInvoiceID(invoiceID)

	if err != nil {
		return nil, err
	}

	payments, err := c.ListSuccessfulPaymentsByInvoiceID(invoiceID)

	if err != nil {
		return nil, err
	}

	chargesPayments := append(charges, payments...)

	return chargesPayments, nil

}
