package chasing

import (
	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.ChasingCadences, error) {
	endpoint := invoiced.AddFilterAndSort("/chasing_cadences", filter, sort)

	chasing := make(invoiced.ChasingCadences, 0)

NEXT:
	tmpChasing := make(invoiced.ChasingCadences, 0)

	endpointTmp, err := c.Api.Get(endpoint, &tmpChasing)

	if err != nil {
		return nil, err
	}

	chasing = append(chasing, tmpChasing...)

	if endpointTmp != "" {
		goto NEXT
	}

	return chasing, nil
}
