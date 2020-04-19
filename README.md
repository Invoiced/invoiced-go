invdapi
========

[![Build Status](https://travis-ci.org/Invoiced/invoiced-go.svg?branch=master)](https://travis-ci.org/Invoiced/invoiced-go)
[![Coverage Status](https://coveralls.io/repos/github/Invoiced/invoiced-go/badge.svg?branch=master)](https://coveralls.io/github/Invoiced/invoiced-go?branch=master)

This repository contains the Go client library for the [Invoiced](https://invoiced.com) API.

## API Documentation

[The API Documenation](http://invoiced.com/docs/api/) is good to look at, as it list fields which can used for creating versus updating objects.  The Go Client Library uses all of the endpoint object's fields.  The server will  simply return an error if you try to use a field that should not be used during a create or update call.

## Installing

The Invoiced Go Client can be installed liked this:

```
go get -u https://github.com/Invoiced/invoiced-go
```


## Requirements

- >= Go 1.11

## Version
5.4.2
```go
//Will print out the version.
invd.Version()
```

## Usage

```go
package main

import "github.com/Invoiced/invoiced-go"
import "github.com/Invoiced/invoiced-go/invdendpoint"
import "fmt"

func main() {

    key := "YOUR DEVELOPER KEY"

    conn := invdapi.NewConnection(key, true)

    //Get All The Invoices With Auto Pagination
    invoiceConn := conn.NewInvoice()
    invoices, err := invoiceConn.ListAll(nil, nil)

    if err != nil {
        panic(err)
    }

    //Let's print all the invoices
    for _, invoice := range invoices {
        fmt.Println(invoice)
    }

    //Let's create a new customer
    customerConn := conn.NewCustomer()

    customerToCreate := conn.NewCustomer()
    customerToCreate.Name = "Test Customer"

    customerResponse, err := customerConn.Create(customerToCreate)

    if err != nil {
        panic(err)
    }

    fmt.Println("Customer Response => ", customerResponse.Customer)

    //Let's create a new invoice
    invoiceToCreate := conn.NewInvoice()
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

    invoiceResponse, err := invoiceConn.Create(invoiceToCreate)

    if err != nil {
        panic(err)
    }

    fmt.Println("Invoice Response => ", invoiceResponse.Invoice)
}
```
