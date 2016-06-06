package invdendpoint

const TransactionsEndPoint = "/transactions"

type Transactions []Transaction

type Transaction struct {
	Id                int64                  `json:"id,omitempty"`
	Customer          int64                  `json:"customer,omitempty"`           //Customer ID, required if invoice ID is not supplied
	Invoice           int64                  `json:"invoice,omitempty"`            //Invoice ID, if any
	Type              string                 `json:"type,omitempty"`               //Transaction type, charge, payment, refund, or adjustment - required
	Date              int64                  `json:"date,omitempty"`               //Transaction date, defaults to current timestamp
	Method            string                 `json:"method,omitempty"`             //Payment instrument used to facilitate transaction, defaults to other
	Status            string                 `json:"status,omitempty"`             //Transaction status, one of succeeded, pending, or failed, defaults to succeeded
	Gateway           string                 `json:"gateway,omitempty"`            //Payment gateway that processed the transaction, if any
	GatewayId         string                 `json:"gateway_id,omitempty"`         //Transaction ID from the payment gateway, or check # if method is check
	Currency          string                 `json:"currency,omitempty"`           //3-letter ISO code
	Amount            float64                `json:"amount,omitempty"`             //Transaction amount
	Fee               float64                `json:"fee,omitempty"`                //Processing fees
	Notes             string                 `json:"notes,omitempty"`              //Internal notes
	FailureReason     string                 `json:"failure_reason,omitempty"`     //Failure message from the payment gateway (only available when status = failed)
	ParentTransaction int64                  `json:"parent_transaction,omitempty"` //ID of the original transaction for refunds
	PdfUrl            string                 `json:"pdf_url,omitempty"`            //URL to download the invoice as a PDF
	CreatedAt         int64                  `json:"created_at,omitempty"`         //Timestamp when created
	MetaData          map[string]interface{} `json:"metadata,omitempty"`           //A hash of key/value pairs that can store additional information about this object.
}

type Refund struct {
	Amount float64 `json:"amount,omitempty"` //Amount to refund - required
}
