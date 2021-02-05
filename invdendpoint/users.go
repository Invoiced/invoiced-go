package invdendpoint

const UsersEndpoint = "/members"

type UserRequests []UserRequest

type UserRequest struct {
	Email           string              `json:"email,omitempty"`
	FirstName       string              `json:"first_name,omitempty"`
	LastName        string              `json:"last_name,omitempty"`
	Role            string              `json:"role,omitempty"`
	RestrictionMode string              `json:"restriction_mode,omitempty"`
	Restrictions    map[string][]string `json:"restrictions,omitempty"`
}

type UserResponse struct {
	CreatedAt            int64               `json:"created_at,omitempty"`
	EmailUpdateFrequency string              `json:"email_update_frequency,omitempty"`
	Id                   int64               `json:"id,omitempty"`
	RestrictionMode      string              `json:"restriction_mode,omitempty"`
	Restrictions         map[string][]string `json:"restrictions,omitempty"`
	Role                 string              `json:"role,omitempty"`
	User                 *User               `json:"user,omitempty"`
}

type User struct {
	Email            string `json:"email,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	Id               int64  `json:"id,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Registered       bool   `json:"registered,omitempty"`
	TwoFactorEnabled bool   `json:"two_factor_enabled,omitempty"`
}
