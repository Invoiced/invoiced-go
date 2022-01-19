package invoiced

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Roles []*Role
