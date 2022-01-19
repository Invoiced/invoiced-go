package role

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Retrieve(id int64) (*Client, error) {
	endpoint := invoiced.RoleEndpoint + "/" + strconv.FormatInt(id, 10)

	role := new(Client)

	_, err := c.Api.Get(endpoint, role)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) ([]*Client, error) {
	endpoint := invoiced.RoleEndpoint

	endpoint = invoiced.AddFilterAndSort(endpoint, filter, sort)

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
