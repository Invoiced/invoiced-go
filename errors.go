package invdapi

import (
	"encoding/json"
)

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Param   string `json:"param"`
}

func NewAPIError(typeE, message, param string) *APIError {
	err := &APIError{typeE, message, param}
	return err
}

func (a *APIError) Error() string {
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}

	return string(b)
}
