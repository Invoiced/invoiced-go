package invdendpoint

const RoleEndpoint = "/roles"

type Roles []Role

type Role struct {
	Id               string   `json:"id,omitempty"`
	Name 			 string  `json:"name,omitempty"`
}
