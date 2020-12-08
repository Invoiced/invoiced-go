package invdendpoint

type User struct {
	Email            string `json:"email,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	Id               int64  `json:"id,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Registered       bool   `json:"registered,omitempty"`
	TwoFactorEnabled bool   `json:"two_factor_enabled,omitempty"`
}

// unmarshaling for this struct is tested in notes_test.go and other places
