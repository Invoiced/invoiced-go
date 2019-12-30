package invdendpoint

const CouponsEndPoint = "/coupons"

type Coupon struct {
	Id             string                  `json:"id,omitempty"`              //The discount’s unique ID
	Object         string                 `json:"object,omitempty"`          //Object type, coupon
	Name           string                 `json:"name,omitempty"`            //Coupon name
	Currency       string                 `json:"currency,omitempty"`        //3-letter ISO code
	Value          int64                 `json:"value,omitempty"`           //Amount
	IsPercent      bool                   `json:"is_percent,omitempty"`      //When true the value is a %
	Exclusive      bool                   `json:"exclusive,omitempty"`       //exclusive
	ExpirationDate int64                  `json:"expiration_date,omitempty"` //Date coupon expires
	MaxRedemptions int64                  `json:"max_redemptions,omitempty"` //Max number of times coupon can be used
	CreatedAt      int64                  `json:"created_at,omitempty"`      //Timestamp when created
	Metadata       map[string]interface{} `json:"metadata,omitempty"`        //A hash of key/value pairs that can store additional information about this object

}
