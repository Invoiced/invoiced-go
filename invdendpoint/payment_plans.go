package invdendpoint

type PaymentPlanRequest struct {
	Installments []*PaymentPlanInstallment `json:"installments,omitempty"`
}

type PaymentPlan struct {
	Approval     *PaymentPlanApproval     `json:"approval,omitempty"`
	CreatedAt    int64                    `json:"created_at,omitempty"`
	Id           int64                    `json:"id,omitempty"`
	Installments []PaymentPlanInstallment `json:"installments,omitempty"`
	Object       string                   `json:"object,omitempty"`
	Status       string                   `json:"status,omitempty"`
	UpdatedAt    int64                    `json:"updated_at,omitempty"`
}

type PaymentPlanApproval struct {
	Id        int64  `json:"id,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}

type PaymentPlanInstallment struct {
	Amount  float64 `json:"amount,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	Date    int64   `json:"date,omitempty"`
	Id      int64   `json:"id,omitempty"`
}

type PaymentPlanInstallments []PaymentPlanInstallment
