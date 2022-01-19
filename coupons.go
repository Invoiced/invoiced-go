package invoiced

type CouponRequest struct {
	Currency       *string                 `json:"currency,omitempty"`
	Duration       *int64                  `json:"durationo,omitempty"`
	Exclusive      *bool                   `json:"exclusive,omitempty"`
	ExpirationDate *int64                  `json:"expiration_date,omitempty"`
	Id             *string                 `json:"id,omitempty"`
	IsPercent      *bool                   `json:"is_percent,omitempty"`
	MaxRedemptions *int64                  `json:"max_redemptions,omitempty"`
	Metadata       *map[string]interface{} `json:"metadata,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	Value          *int64                  `json:"value,omitempty"`
}

type Coupon struct {
	CreatedAt      int64                  `json:"created_at"`
	Currency       *string                `json:"currency"`
	Duration       *int64                 `json:"duration"`
	Exclusive      bool                   `json:"exclusive"`
	ExpirationDate *int64                 `json:"expiration_date"`
	Id             string                 `json:"id"`
	IsPercent      bool                   `json:"is_percent"`
	MaxRedemptions *int64                 `json:"max_redemptions"`
	Metadata       map[string]interface{} `json:"metadata"`
	Name           string                 `json:"name"`
	Object         string                 `json:"object"`
	UpdatedAt      int64                  `json:"updated_at"`
	Value          int64                  `json:"value"`
}

type Coupons []*Coupon
