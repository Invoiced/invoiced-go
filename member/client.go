package member

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.MemberRequest) (*invoiced.Member, error) {
	resp := new(invoiced.Member)
	err := c.Api.Create("/members", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Member, error) {
	resp := new(invoiced.Member)
	_, err := c.Api.Get("/members/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.MemberRequest) (*invoiced.Member, error) {
	resp := new(invoiced.Member)
	err := c.Api.Update("/members/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/members/" + strconv.FormatInt(id, 10))
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Members, error) {
	endpoint := invoiced.AddFilterAndSort("/members", filter, sort)

	users := make(invoiced.Members, 0)

NEXT:
	tmpUsers := make(invoiced.Members, 0)

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

func (c *Client) SetUserEmailFrequency(id int64, request *invoiced.UserEmailUpdateRequest) (*Client, error) {
	endpoint := "/members/" + strconv.FormatInt(id, 10) + "/frequency"

	resp := new(Client)
	err := c.Api.Update(endpoint, request, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) SendInvite(id int64) error {
	endpoint := "/members/" + strconv.FormatInt(id, 10) + "/invites"

	request := new(invoiced.UserInvite)
	request.Id = id

	err := c.Api.Create(endpoint, request, nil)

	if err != nil {
		return err
	}

	return nil
}
