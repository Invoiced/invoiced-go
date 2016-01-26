invdapi
========

[![Build Status](https://travis-ci.org/Invoiced/invdapi.svg?branch=master)](https://travis-ci.org/Invoiced/invdapi)
[![Coverage Status](https://coveralls.io/repos/Invoiced/invdapi/badge.svg?branch=master&service=github)](https://coveralls.io/github/Invoiced/invdapi?branch=master)

This repository contains the Go client library for the [Invoiced](https://invoiced.com) API.

##API Documentation
[The API Documenation](http://invoiced.com/docs/api/) is good to look at, as it list fields which can used for creating versus updating objects.  The Go Client Library uses all of the endpoint object's fields.  The server will  simply return an error if you try to use a field that should not be used during a create or update call.

## Installing

The Invoiced Go Client can be installed liked this:

```
go get -u https://github.com/Invoiced/invdapi
```


## Requirements

- >= Go 1.4

## Version
0.1
```go
//Will print out the version.
invd.Version()
```

## Usage

```go
package main

import "github.com/Invoiced/invdapi"
import "github.com/Invoiced/invdapi/invdendpoint"
import "fmt"

func main() {

    key := "YOUR DEVELOPER KEY"

    conn := invdapi.NewConnection(key)

    //Get All The Invoices With Auto Pagination
    invoices, apiErr := conn.GetAllInvoicesAuto(nil, nil)

    if apiErr != nil {
        panic(apiErr)
    }

    //Let's print all the invoices
    for _, invoice := range *invoices {
        fmt.Println(invoice)
    }

    //Let's create a new customer

    customerToCreate := new(invdendpoint.Customer)
    customerToCreate.Name = "Test Customer"

    customerResponse, apiErr := conn.CreateCustomer(customerToCreate)

    if apiErr != nil {
        panic(apiErr)
    }

    fmt.Println("Customer Response => ", customerResponse)

    //Let's create a new invoice

    invoiceToCreate := new(invdendpoint.Invoice)
    invoiceToCreate.Customer = customerResponse.Id

    //Create a Line Item
    lineItem := invdendpoint.LineItem{}
    lineItem.Description = "Retina MacBook Pro"
    lineItem.Quantity = 5
    lineItem.UnitCost = 1999.22

    lineItems := append([]invdendpoint.LineItem{}, lineItem)

    invoiceToCreate.Items = lineItems

    //Add a Payment Term
    invoiceToCreate.PaymentTerms = "NET15"

    invoiceResponse, apiErr := conn.CreateInvoice(invoiceToCreate)

    fmt.Println("Invoice Response => ", invoiceResponse)

}
```

##Endpoints Implemented
This Library has implemented all the endpoint objects. However only customer,invoice has CRUD methods implemented.  Transactions,Subscriptions,and Plans are TODO.

## Testing


For Testing you can set the api key in connector_test.go by changing the location of the api key yaml file.

```go
func init() {
   apikey = invoicedutil.ReadAPIKeyFromYaml("/usr/local/keys/invoicedapikey.yaml")
}
```

The format of the yaml file is
```
apikey: YOUR_DEVELOPER_KEY
```