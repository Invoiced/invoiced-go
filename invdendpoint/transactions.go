package invdendpoint

const TransactionsEndPoint = "/transactions/"

type Transaction struct {
	Id                int64   `json:"id,omitempty"`
	Customer          int64   `json:"customer,omitempty"`
	Invoice           int64   `json:"invoice,omitempty"`
	Type              string  `json:"type,omitempty"`
	Date              int64   `json:"date,omitempty"`
	Theme             string  `json:"theme,omitempty"`
	Method            string  `json:"method,omitempty"`
	Status            int64   `json:"status,omitempty"`
	Gateway           string  `json:"gateway,omitempty"`
	GatewayId         string  `json:"gateway_id,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	Fee               float64 `json:"fee,omitempty"`
	Notes             string  `json:"notes,omitempty"`
	Sent              bool    `json:"sent,omitempty"`
	ParentTransaction int64   `json:"parent_transaction,omitempty"`
	PdfUrl            int64   `json:"pdf_url,omitempty"`
	Net               int64   `json:"net,omitempty"`
	CreatedAt         int64   `json:"created_at,omitempty"`
	UpdatedAt         int64   `json:"updated_at,omitempty"`
	FailureReason     string  `json:"failure_reason,omitempty"`
}
