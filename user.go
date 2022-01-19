package invoiced

import (
	"net/url"
	"strconv"
)

type MemberClient struct {
	*Client
	*Member
}

type Users []*MemberClient

func (c *Client) NewMember() *MemberClient {
	user := new(Member)
	return &MemberClient{c, user}
}

func (c *MemberClient) Create(request *MemberRequest) (*MemberClient, error) {
	endpoint := UsersEndpoint

	resp := new(MemberClient)

	err := c.Api.Create(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *MemberClient) Retrieve(id int64) (*MemberClient, error) {
	endpoint := UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	resp := new(MemberClient)

	_, err := c.Api.Get(endpoint, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *MemberClient) Update(request *MemberRequest, id int64) error {
	endpoint := UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	resp := new(MemberClient)

	err := c.Api.Update(endpoint, request, resp)

	if err != nil {
		return err
	}

	return nil
}

func (c *MemberClient) Delete(id int64) error {
	endpoint := UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.Api.Delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *MemberClient) ListAll(filter *Filter, sort *Sort) (Users, error) {
	endpoint := UsersEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

	users := make(Users, 0)

NEXT:
	tmpUsers := make(Users, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpUsers)

	if err != nil {
		return nil, err
	}

	users = append(users, tmpUsers...)

	if endpoint != "" {
		goto NEXT
	}

	return users, nil
}

func (c *MemberClient) SetUserEmailFrequency(id int64, request *UserEmailUpdateRequest) (*MemberClient, error) {
	endpoint := UsersEndpoint + "/" + strconv.FormatInt(id, 10) + "/frequency"

	resp := new(MemberClient)
	err := c.Api.Update(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *MemberClient) SendInvite(id int64) error {
	endpoint := UsersEndpoint + "/" + strconv.FormatInt(id, 10) + "/invites"

	request := new(UserInvite)
	request.Id = id

	err := c.Api.Create(endpoint, request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *MemberClient) GenerateRegistrationURL() string {
	regURl := ""

	if c.Client.Api.Sandbox {
		regURl = "https://app.sandbox.invoiced.com/register"
	} else {
		regURl = "https://app.invoiced.com/register"
	}

	u, _ := url.Parse(regURl)
	q := u.Query()
	q.Add("email", c.User.Email)
	q.Add("first_name", c.User.FirstName)
	q.Add("last_name", c.User.LastName)
	u.RawQuery = q.Encode()

	return u.String()
}
