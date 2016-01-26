package invdapi

import "errors"

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Param   string `json:"param"`
}

func (c *Connection) APIErrorToError(apiError *APIError) error {
	e := apiError.Type + " " + apiError.Message + " " + apiError.Param

	err := errors.New(e)

	return err
}
