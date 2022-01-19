package invoiced

import "net/url"

type MemberRequests []MemberRequest

type MemberRequest struct {
	Email           *string              `json:"email,omitempty"`
	FirstName       *string              `json:"first_name,omitempty"`
	LastName        *string              `json:"last_name,omitempty"`
	RestrictionMode *string              `json:"restriction_mode,omitempty"`
	Restrictions    *map[string][]string `json:"restrictions,omitempty"`
	Role            *string              `json:"role,omitempty"`
}

type Member struct {
	CreatedAt            int64               `json:"created_at"`
	UpdatedAt            int64               `json:"updated_at"`
	EmailUpdateFrequency string              `json:"email_update_frequency"`
	Id                   int64               `json:"id"`
	LastSignedIn         int64               `json:"last_accessed"`
	RestrictionMode      string              `json:"restriction_mode"`
	Restrictions         map[string][]string `json:"restrictions"`
	Role                 string              `json:"role"`
	User                 *User               `json:"user"`
}

type Members []*Member

type UserEmailUpdateRequest struct {
	EmailUpdateFrequency *string `json:"email_update_frequency"`
}

type UserInvite struct {
	Id int64 `json:"id"`
}

type User struct {
	Email            string `json:"email,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	Id               int64  `json:"id,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Registered       bool   `json:"registered,omitempty"`
	TwoFactorEnabled bool   `json:"two_factor_enabled,omitempty"`
}

func (u *User) GenerateRegistrationURL(sandbox bool) string {
	regURl := ""
	if sandbox {
		regURl = "https://app.sandbox.invoiced.com/register"
	} else {
		regURl = "https://app.invoiced.com/register"
	}

	regUrlParts, _ := url.Parse(regURl)
	query := regUrlParts.Query()
	query.Add("email", u.Email)
	query.Add("first_name", u.FirstName)
	query.Add("last_name", u.LastName)
	regUrlParts.RawQuery = query.Encode()

	return regUrlParts.String()
}
