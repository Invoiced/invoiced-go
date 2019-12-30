//Transactions can represent a charge, payment, refund, or adjustment.
// We record charge and refund transactions for you that happen through Invoiced. The payment transaction type is designated for recording offline payments like checks. Finally, an adjustment transaction represents any additional credit or debits to a customer’s balance.
// Most transactions will be associated with an invoice, however, not all. For example, if you wanted to credit your customer for $20 you would create an adjustment transaction for -$20 using the customer ID only instead of the invoice ID.
// We currently support the following payment methods on transactions:
// credit_card
// ach
// paypal
// wire_transfer
// check
// cash
// other
package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"strconv"
	"errors"
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
	endPoint := c.MakeEndPointURL(invdendpoint.TransactionsEndPoint)

	count, apiErr := c.count(endPoint)

	if apiErr != nil {
		return -1, apiErr
	}

	return count, nil

}

func (c *Transaction) Create(transaction *Transaction) (*Transaction, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.TransactionsEndPoint)
	txnResp := new(Transaction)

	if transaction == nil {
		return nil, errors.New("transaction cannot be nil")
	}

	//safe prune invoice data for creation
	invdTransDataToCreate,err := SafeTransactionForCreation(transaction.Transaction)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdTransDataToCreate, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil

}

func (c *Transaction) Delete() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TransactionsEndPoint), c.Id)

	apiErr := c.delete(endPoint)

	if apiErr != nil {
		return apiErr
	}

	return nil

}

func (c *Transaction) Save() error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TransactionsEndPoint), c.Id)
	txnResp := new(Transaction)

	//safe prune invoice data for updating
	invdTransDataToUpdate,err := SafeTransactionForUpdate(c.Transaction)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, invdTransDataToUpdate, txnResp)

	if apiErr != nil {
		return apiErr
	}

	c.Transaction = txnResp.Transaction

	return nil

}

func (c *Transaction) Retrieve(id int64) (*Transaction, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TransactionsEndPoint), id)

	custEndPoint := new(invdendpoint.Transaction)

	transaction := &Transaction{c.Connection, custEndPoint}

	_, apiErr := c.retrieveDataFromAPI(endPoint, transaction)

	if apiErr != nil {
		return nil, apiErr
	}

	return transaction, nil

}

func (c *Transaction) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Transactions, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.TransactionsEndPoint)
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
	endPoint := c.MakeEndPointURL(invdendpoint.TransactionsEndPoint)
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

func (c *Transaction) ListSuccessfulByInvoiceID(invoiceID int64) (Transactions, error) {

	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)

	if err != nil {
		return nil,err
	}

	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil,err
	}

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

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)
	if err != nil {
		return nil, err
	}
	err = filter.Set("status", "succeeded")
	if err != nil {
		return nil, err
	}
	err = filter.Set("type", "charge")
	if err != nil {
		return nil, err
	}

	transactions, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions, nil

}

func (c *Transaction) ListSuccessfulRefundsByInvoiceID(invoiceID int64) (Transactions, error) {

	invoiceIDString := strconv.FormatInt(invoiceID, 10)

	filter := invdendpoint.NewFilter()
	err  := filter.Set("invoice", invoiceIDString)

	if err != nil {
		return nil, err
	}
	
	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil, err
	}

	err = filter.Set("type", "refund")

	if err != nil {
		return nil, err
	}

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

	filter := invdendpoint.NewFilter()
	err := filter.Set("invoice", invoiceIDString)

	if err != nil {
		return nil, err
	}
	err = filter.Set("status", "succeeded")

	if err != nil {
		return nil, err
	}

	err = filter.Set("type", "payment")

	if err != nil {
		return nil, err
	}

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

func (c *Transaction) SendReceipt(emailReq *invdendpoint.EmailRequest) (invdendpoint.EmailResponses, error) {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.InvoicesEndPoint), c.Id) + "/emails"

	emailResp := new(invdendpoint.EmailResponses)

	err := c.create(endPoint, emailReq, emailResp)

	if err != nil {
		return nil, err
	}

	return *emailResp, nil

}

func (c *Transaction) Refund(refund float64) error {
	endPoint := makeEndPointSingular(c.MakeEndPointURL(invdendpoint.TransactionsEndPoint), c.Id) + "/refunds"
	transaction := new(invdendpoint.Transaction)
	err := c.create(endPoint, nil, transaction)

	if err != nil {
		return nil
	}

	c.Transaction = transaction

	return nil

}

func (c *Transaction) InitiateCharge(chargeRequest *invdendpoint.ChargeRequest) (*Transaction, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.ChargesEndPoint)
	txnResp := new(Transaction)

	if chargeRequest == nil {
		return nil, errors.New("chargeRequest cannot be nil")
	}

	apiErr := c.create(endPoint, chargeRequest, txnResp)

	if apiErr != nil {
		return nil, apiErr
	}

	txnResp.Connection = c.Connection

	return txnResp, nil

}


//SafeTransactionForCreation prunes transaction data for just fields that can be used for creation of a transaction
func SafeTransactionForCreation(transaction *invdendpoint.Transaction) (*invdendpoint.Transaction, error) {

	if transaction == nil  {
		return nil, errors.New("Transaction is nil")
	}

	transData :=new(invdendpoint.Transaction)
	transData.Customer = transaction.Customer
	transData.Invoice = transaction.Invoice
	transData.CreditNote = transaction.CreditNote
	transData.Type = transaction.Type
	transData.Method = transaction.Method
	transData.Status = transaction.Status
	transData.Gateway = transaction.Gateway
	transData.GatewayId = transaction.GatewayId
	transData.Currency = transaction.Currency
	transData.Amount = transaction.Amount
	transData.Notes = transaction.Notes
	transData.Metadata = transaction.Metadata


	return transData,nil
}

//SafeTransactionForUpdate prunes transaction data for just fields that can be used for creation of a transactiobn
func SafeTransactionForUpdate(transaction *invdendpoint.Transaction) (*invdendpoint.Transaction, error) {

	if transaction == nil {
		return nil, errors.New("Transaction is nil")
	}

	transData :=new(invdendpoint.Transaction)

	transData.Date = transaction.Date
	transData.Method = transaction.Method
	transData.Status = transaction.Status
	transData.Gateway = transaction.Gateway
	transData.GatewayId = transaction.GatewayId
	transData.Currency = transaction.Currency
	transData.Amount = transaction.Amount
	transData.Notes = transaction.Notes
	transData.Metadata = transaction.Metadata


	return transData, nil

}