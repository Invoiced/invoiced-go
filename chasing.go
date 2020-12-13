package invdapi

import (
	"github.com/Invoiced/invoiced-go/invdendpoint"
)

type ChasingCadence struct {
	*Connection
	*invdendpoint.ChasingCadence
}

type ChasingCadences []*ChasingCadence

func (c *Connection) NewChasingCadence() *ChasingCadence {
	chasing := new(invdendpoint.ChasingCadence)
	return &ChasingCadence{c, chasing}
}

func (c *ChasingCadence) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (ChasingCadences, error) {
	endpoint := addFilterAndSort(invdendpoint.ChasingCadenceEndpoint, filter, sort)

	chasing := make(ChasingCadences, 0)

NEXT:
	tmpChasing := make(ChasingCadences, 0)

	endpointTmp, apiErr := c.retrieveDataFromAPI(endpoint, &tmpChasing)

	if apiErr != nil {
		return nil, apiErr
	}

	chasing = append(chasing, tmpChasing...)

	if endpointTmp != "" {
		goto NEXT
	}

	for _, chase := range chasing {
		chase.Connection = c.Connection
	}

	return chasing, nil
}
