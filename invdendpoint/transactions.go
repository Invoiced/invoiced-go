package invdendpoint

const TransactionsEndPoint = "/transactions"

type Transactions []Transaction

type Transaction struct {
	Id                int64   `json:"id,omitempty"`
	PdfUrl            string  `json:"pdf_url,omitempty"`
	Customer          int64   `json:"customer,omitempty"`
	Invoice           int64   `json:"invoice,omitempty"`
	Theme             string  `json:"theme,omitempty"`
	Date              int64   `json:"date,omitempty"`
	Type              string  `json:"type,omitempty"`
	Method            string  `json:"method,omitempty"`
	Status            string  `json:"status,omitempty"`
	Gateway           string  `json:"gateway,omitempty"`
	GatewayId         string  `json:"gateway_id,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	Fee               float64 `json:"fee,omitempty"`
	Notes             string  `json:"notes,omitempty"`
	Sent              bool    `json:"sent,omitempty"`
	FailureReason     string  `json:"failure_reason,omitempty"`
	ParentTransaction string  `json:"parent_transaction,omitempty"`

	CreatedAt int64       `json:"created_at,omitempty"`
	UpdatedAt int64       `json:"updated_at,omitempty"`
	MetaData  interface{} `json:"metadata,omitempty"`
}
