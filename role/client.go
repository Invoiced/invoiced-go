package role

import (
	"strconv"

	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Retrieve(id int64) (*invoiced.Role, error) {
	resp := new(invoiced.Role)
	_, err := c.Api.Get("/roles/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Roles, error) {
	endpoint := invoiced.AddFilterAndSort("/roles", filter, sort)

	roles := make(invoiced.Roles, 0)

NEXT:
	tmpRoles := make(invoiced.Roles, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpRoles)

	if err != nil {
		return nil, err
	}

	roles = append(roles, tmpRoles...)

	if endpointTmp != "" {
		goto NEXT
	}

	return roles, nil
}
