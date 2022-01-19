package invoiced

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalChasingCadenceObject(t *testing.T) {
	s := `{
	"assignment_conditions": "",
	"assignment_mode": "none",
	"created_at": 1571950909,
	"frequency": "day_of_month",
	"id": 210,
	"last_run": 1575723606,
	"min_balance": 100,
	"name": "Standard Cadence",
	"next_run": 1578402000,
	"num_customers": 1,
	"object": "chasing_cadence",
	"paused": false,
	"run_date": 7,
	"steps": [{
		"action": "email",
		"assigned_user_id": null,
		"created_at": 1571950909,
		"email_template_id": "5d3605831a1f8",
		"id": 801,
		"name": "1st Email",
		"schedule": "past_due_age:0",
		"sms_template_id": null
	}],
	"time_of_day": 7
}`

	so := new(ChasingCadence)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
}
