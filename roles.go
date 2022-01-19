package invoiced

const RoleEndpoint = "/roles"

type Roles []Role

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
