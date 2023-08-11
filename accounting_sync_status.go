package invoiced

import "encoding/json"

type AccountingSyncStatus struct {
	Synced           bool            `json:"synced"`
	Error            json.RawMessage `json:"error"`
	AccountingSystem string          `json:"accounting_system"`
	AccountingId     string          `json:"accounting_id"`
	Source           string          `json:"source"`
	FirstSynced      int64           `json:"first_synced"`
	LastSynced       int64           `json:"last_synced"`
}
