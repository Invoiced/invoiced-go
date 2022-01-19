package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"net/url"
	"strconv"
	"strings"
)

type Member struct {
	*Connection
	*invdendpoint.Member
}

type Users []*Member

func (c *Connection) NewMember() *Member {
	user := new(invdendpoint.Member)
	return &Member{c, user}
}

func (c *Member) Create(request *invdendpoint.MemberRequest) (*Member, error) {
	endpoint := invdendpoint.UsersEndpoint

	resp := new(Member)

	err := c.create(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Member) Retrieve(id int64) (*Member, error) {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	resp := new(Member)

	_, err := c.retrieveDataFromAPI(endpoint, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Member) Update(request *invdendpoint.MemberRequest, id int64) error {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	resp := new(Member)

	err := c.update(endpoint, request, resp)

	if err != nil {
		return err
	}

	resp.Connection = c.Connection

	return nil
}

func (c *Member) Delete(id int64) error {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Member) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Users, error) {
	endpoint := invdendpoint.UsersEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	users := make(Users, 0)

NEXT:
	tmpUsers := make(Users, 0)

	endpoint, err := c.retrieveDataFromAPI(endpoint, &tmpUsers)

	if err != nil {
		return nil, err
	}

	users = append(users, tmpUsers...)

	if endpoint != "" {
		goto NEXT
	}

	for _, user := range users {
		user.Connection = c.Connection
	}

	return users, nil
}

func (c *Member) SetUserEmailFrequency(id int64, request *invdendpoint.UserEmailUpdateRequest) (*Member, error) {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10) + "/frequency"

	resp := new(Member)
	err := c.update(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	resp.Connection = c.Connection

	return resp, nil
}

func (c *Member) SendInvite(id int64) error {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10) + "/invites"

	request := new(invdendpoint.UserInvite)
	request.Id = id

	err := c.create(endpoint, request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Member) GenerateRegistrationURL() string {
	regURl := ""

	if strings.Contains(c.Connection.baseUrl, "sandbox") {
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
