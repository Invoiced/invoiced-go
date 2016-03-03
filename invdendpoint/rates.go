package invdendpoint

const RatesEndPoint = "/rates/"

type Rate struct {
	CreatedAt int64   `json:"created_at,omitempty"`
	UpdatedAt int64   `json:"updated_at,omitempty"`
	Id        int64   `json:"id,omitempty"`
	Type      string  `json:"type,omitempty"`
	Name      string  `json:"name,omitempty"`
	Number    string  `json:"number,omitempty"`
	IsPercent bool    `json:"is_percent,omitempty"`
	Value     float64 `json:"value,omitempty"`
	Archived  bool    `json:"archived,omitempty"`
}
