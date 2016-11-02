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
	apiErr := &APIError{typeE, message, param}
	return apiErr
}

func (apiErr *APIError) Error() string {
	b, err := json.Marshal(apiErr)

	if err != nil {
		return ""
	}

	return string(b)

}
