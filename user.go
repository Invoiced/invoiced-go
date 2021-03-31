package invdapi

import (
	"fmt"
	"strconv"
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type User struct {
	*Connection
	*invdendpoint.UserResponse
}

type Users []*User

func (c *Connection) NewUser() *User {
	user := new(invdendpoint.UserResponse)
	return &User{c, user}
}

func (c *User) Create(userRequest *invdendpoint.UserRequest) (*User, error) {
	endpoint := invdendpoint.UsersEndpoint

	userResp := new(User)

	apiErr := c.create(endpoint, userRequest, userResp)

	if apiErr != nil {
		return nil, apiErr
	}

	userResp.Connection = c.Connection

	return userResp, nil
}

func (c *User) Save(userRequest *invdendpoint.UserRequest, id int64) error {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	userResp := new(User)

	apiErr := c.update(endpoint, userRequest, userResp)

	if apiErr != nil {
		return apiErr
	}

	userResp.Connection = c.Connection

	return nil
}

func (c *User) Delete(id int64) error {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) Retrieve(id int64) (*User, error) {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10)

	userResp := new(User)

	_, err := c.retrieveDataFromAPI(endpoint, userResp)
	if err != nil {
		return nil, err
	}

	userResp.Connection = c.Connection

	return userResp, nil
}

func (c *User) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Users, error) {
	endpoint := invdendpoint.UsersEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	users := make(Users, 0)

NEXT:
	tmpUsers := make(Users, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpUsers)

	if apiErr != nil {
		return nil, apiErr
	}

	users = append(users, tmpUsers...)

	fmt.Println("endpointTmp => ",endpoint)

	if endpoint != "" {
		goto NEXT
	}

	for _, user := range users {
		user.Connection = c.Connection
	}

	return users, nil
}

func (c *User) SetUserEmailFrequency(userEmailFrequency string, id int64) (*User, error) {
	endpoint := invdendpoint.UsersEndpoint + "/" + strconv.FormatInt(id, 10) + "/frequency"

	userResp := new(User)

	userRequest := new(invdendpoint.UserEmailUpdateRequest)
	userRequest.Id = id
	userRequest.EmailUpdateFrequency = userEmailFrequency

	apiErr := c.update(endpoint, userRequest, userResp)

	if apiErr != nil {
		return nil,apiErr
	}


	userResp.Connection = c.Connection

	return userResp, nil
}


