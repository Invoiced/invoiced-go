package invoiced

import (
	"strconv"
)

type RoleClient struct {
	*Client
	*Role
}

func (c *Client) NewRole() *RoleClient {
	role := new(Role)
	return &RoleClient{c, role}
}

func (c *RoleClient) Retrieve(id int64) (*RoleClient, error) {
	endpoint := RoleEndpoint + "/" + strconv.FormatInt(id, 10)

	role := new(RoleClient)

	_, err := c.Api.Get(endpoint, role)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *RoleClient) ListAll(filter *Filter, sort *Sort) ([]*RoleClient, error) {
	endpoint := RoleEndpoint

	endpoint = AddFilterAndSort(endpoint, filter, sort)

	roles := make(Roles, 0)

NEXT:
	tmpRoles := make(Roles, 0)

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
