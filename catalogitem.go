package invdapi

import (
"errors"
"github.com/Invoiced/invoiced-go/invdendpoint"
)

type CatalogItem struct {
	*Connection
	*invdendpoint.CatalogItem
}

type CatalogItems []*CatalogItem

func (c *Connection) NewCatalogItem() *CatalogItem {
	catalogItem := new(invdendpoint.CatalogItem)
	return &CatalogItem{c, catalogItem}

}

func (c *CatalogItem) Create(catalogItem *CatalogItem) (*CatalogItem, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CatalogItemEndPoint)

	catalogItemResp := new(CatalogItem)


	if catalogItem == nil {
		return nil, errors.New("catalogItem is nil")
	}


	//safe prune file data for creation
	invdCatalogItemDataToCreate, err := SafeCatalogItemForCreation(catalogItem.CatalogItem)

	if err != nil {
		return nil, err
	}

	apiErr := c.create(endPoint, invdCatalogItemDataToCreate, catalogItemResp)

	if apiErr != nil {
		return nil, apiErr
	}

	catalogItemResp.Connection = c.Connection

	return catalogItemResp, nil

}

func (c *CatalogItem) Save() error {
	endPoint := c.MakeEndPointURL(invdendpoint.CatalogItemEndPoint) + "/" + c.Id

	catalogItemResp := new(CatalogItem)

	catalogItemDataToUpdate, err := SafeCatalogItemForUpdating(c.CatalogItem)

	if err != nil {
		return err
	}

	apiErr := c.update(endPoint, catalogItemDataToUpdate, catalogItemResp)

	if apiErr != nil {
		return apiErr
	}

	c.CatalogItem = catalogItemResp.CatalogItem

	return nil

}

func (c *CatalogItem) Delete() error {
	endPoint := c.MakeEndPointURL(invdendpoint.CatalogItemEndPoint) + "/" + c.Id

	err := c.delete(endPoint)

	if err != nil {
		return err
	}

	return nil

}

func (c *CatalogItem) Retrieve(id string) (*CatalogItem, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CatalogItemEndPoint) + "/" + id

	catalogItemEndPoint := new(invdendpoint.CatalogItem)

	catalogItem := &CatalogItem{c.Connection, catalogItemEndPoint}

	_, err := c.retrieveDataFromAPI(endPoint, catalogItem)

	if err != nil {
		return nil, err
	}

	return catalogItem, nil

}

func (c *CatalogItem) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (CatalogItems, error) {
	endPoint := c.MakeEndPointURL(invdendpoint.CatalogItemEndPoint)

	endPoint = addFilterSortToEndPoint(endPoint, filter, sort)

	expandedValues := invdendpoint.NewExpand()

	expandedValues.Set(defaultExpandInvoice)

	endPoint = addExpandToEndPoint(endPoint, expandedValues)

	catalogItems := make(CatalogItems, 0)

NEXT:
	tmpCatalogItems := make(CatalogItems, 0)

	endPointTmp, apiErr := c.retrieveDataFromAPI(endPoint, &tmpCatalogItems)

	if apiErr != nil {
		return nil, apiErr
	}

	catalogItems = append(catalogItems, tmpCatalogItems...)

	if endPointTmp != "" {
		goto NEXT
	}

	for _, catalogItem := range catalogItems {
		catalogItem.Connection = c.Connection

	}

	return catalogItems, nil

}

//SafeCustomerForCreation prunes customer data for just fields that can be used for creation of a customer
func SafeCatalogItemForCreation(catalogItem *invdendpoint.CatalogItem) (*invdendpoint.CatalogItem, error) {

	if catalogItem == nil {
		return nil, errors.New("task is nil")
	}

	catalogItemData := new(invdendpoint.CatalogItem)
	catalogItemData.Id = catalogItem.Id
	catalogItemData.Name = catalogItem.Name
	catalogItemData.Currency = catalogItem.Currency
	catalogItemData.UnitCost = catalogItem.UnitCost
	catalogItemData.Description = catalogItem.Description
	catalogItemData.Type = catalogItem.Type
	catalogItemData.Taxable = catalogItem.Taxable
	catalogItemData.AvalaraTaxCode = catalogItem.AvalaraTaxCode
	catalogItemData.GlAccount = catalogItem.GlAccount
	catalogItemData.Discountable = catalogItem.Discountable
	catalogItemData.MetaData = catalogItem.MetaData

	return catalogItemData, nil
}

//SafeTaskForUpdating prunes customer data for just fields that can be used for updating of a customer
func SafeCatalogItemForUpdating(catalogItem *invdendpoint.CatalogItem) (*invdendpoint.CatalogItem, error) {

	if catalogItem == nil {
		return nil, errors.New("task is nil")
	}

	catalogItemData := new(invdendpoint.CatalogItem)

	catalogItemData.Name = catalogItem.Name
	catalogItemData.Description = catalogItem.Description
	catalogItemData.Type = catalogItem.Type
	catalogItemData.MetaData = catalogItem.MetaData

	return catalogItemData, nil
}
