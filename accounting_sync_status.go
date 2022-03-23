package invoiced


type AccountingSyncStatus struct {
	AccountingId     string `json:"accounting_id"`
	AccountingSystem string `json:"accounting_system"`
	Error            string `json:"error"`
	FirstSynced      int64  `json:"first_synced"`
	LastSynced       int64  `json:"last_synced"`
	Source           string `json:"source"`
	Synced           bool   `json:"synced"`
}
