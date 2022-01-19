invoiced-go
========

This repository contains the Go client library for the [Invoiced](https://invoiced.com) API.

[![Build Status](https://travis-ci.com/Invoiced/invoiced-go.svg?branch=master)](https://travis-ci.com/Invoiced/invoiced-go)
[![Coverage Status](https://coveralls.io/repos/github/Invoiced/invoiced-go/badge.svg?branch=master)](https://coveralls.io/github/Invoiced/invoiced-go?branch=master)

## Installing

The Invoiced package can be installed like this:

```
go get -u https://github.com/Invoiced/invoiced-go
```

## Requirements

- Go 1.11+

## Usage

First, you must instantiate a new client

```go
import "github.com/Invoiced/invoiced-go"

conn := invdapi.NewConnection("API_KEY", false)
```

Then, API calls can be made like this:

```go
import "github.com/Invoiced/invoiced-go/invdendpoint"
import "fmt"

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
customerResponse, err := customerConn.Create(&invdendpoint.CustomerRequest{
	Name: invdapi.String("Test Customer")
})

if err != nil {
    panic(err)
}

fmt.Println("Customer Response => ", customerResponse.Customer)

//Let's create a new invoice
invoiceToCreate := &invdendpoint.InvoiceRequest{}
invoiceToCreate.Customer = invdapi.Int64(customerResponse.Id)

//Create a Line Item
lineItem := &invdendpoint.LineItem{}
lineItem.Description = invdapi.String("Retina MacBook Pro")
lineItem.Quantity = invdapi.Float64(5)
lineItem.UnitCost = invdapi.Float64(1999.22)

lineItems := append([]invdendpoint.LineItem{}, lineItem)

invoiceToCreate.Items = lineItems

//Add a Payment Term
invoiceToCreate.PaymentTerms = invdapi.String("NET15")

invoiceResponse, err := invoiceConn.Create(invoiceToCreate)

if err != nil {
    panic(err)
}

fmt.Println("Invoice Response => ", invoiceResponse.Invoice)
```

If you want to use the sandbox API instead then you must set the second argument on the client to `true` like this:

```go
conn := invdapi.NewConnection("SANDBOX_API_KEY", false)
```

## Developing

The test suite can be run with:

```
go test ./...
```
