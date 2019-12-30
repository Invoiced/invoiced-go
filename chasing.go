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
	endPoint := c.MakeEndPointURL(invdendpoint.ChasingCadencesEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	chasing := make(ChasingCadences, 0)

NEXT:
	tmpChasing := make(ChasingCadences, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpChasing)

	if apiErr != nil {
		return nil, apiErr
	}

	chasing = append(chasing, tmpChasing...)

	if endPointTmp != "" {
		goto NEXT
	}

	for _, chase := range chasing {
		chase.Connection = c.Connection

	}

	return chasing, nil

}
