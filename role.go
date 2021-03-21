package invdapi

import (
"strconv"
"github.com/Invoiced/invoiced-go/invdendpoint"
)

type Role struct {
	*Connection
	*invdendpoint.Role
}

type Roles []*Role

func (c *Connection) NewRole() *Role {
	role := new(invdendpoint.Role)
	return &Role{c, role}
}


func (c *Role) Retrieve(id int64) (*Role, error) {
	endpoint := invdendpoint.RoleEndpoint+ "/" + strconv.FormatInt(id, 10)

	role := new(Role)

	_, err := c.retrieveDataFromAPI(endpoint, role)

	if err != nil {
		return nil, err
	}

	role.Connection = c.Connection

	return role, nil
}

func (c *Role) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) ([]*Role, error) {
	endpoint := invdendpoint.RoleEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	roles := make(Roles, 0)

NEXT:
	tmpRoles := make(Roles, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpRoles)

	if apiErr != nil {
		return nil, apiErr
	}

	roles = append(roles, tmpRoles...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, role := range roles {
		role.Connection = c.Connection
	}

	return roles, nil
}


