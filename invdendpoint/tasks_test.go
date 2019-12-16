package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalTaskObject(t *testing.T) {
	s := `{
  "action": "phone",
  "chase_step_id": null,
  "complete": false,
  "completed_by_user_id": null,
  "completed_date": null,
  "created_at": 1571347283,
  "customer_id": 481594,
  "due_date": 1571288400,
  "id": 788,
  "name": "Call customer",
  "user_id": 1976
}`

	so := new(Task)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
