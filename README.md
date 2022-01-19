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

client := invoiced.New("API_KEY", false)
```

Then, API calls can be made like this:

```go
import "github.com/Invoiced/invoiced-go"
import "fmt"

//Get All The Invoices With Auto Pagination
invoices, err := client.Invoice.ListAll(nil, nil)

if err != nil {
    panic(err)
}

//Let's print all the invoices
for _, invoice := range invoices {
    fmt.Println(invoice)
}

//Let's create a new customer
customer, err := client.Customer.Create(&invoiced.CustomerRequest{
	Name: invoiced.String("Test Customer")
})

if err != nil {
    panic(err)
}

fmt.Println("Customer Response => ", customer)

//Let's create a new invoice
invoiceToCreate := &invoiced.InvoiceRequest{}
invoiceToCreate.Customer = invoiced.Int64(customerResponse.Id)

//Create a Line Item
lineItem := &invoiced.LineItem{}
lineItem.Description = invoiced.String("Retina MacBook Pro")
lineItem.Quantity = invoiced.Float64(5)
lineItem.UnitCost = invoiced.Float64(1999.22)

lineItems := append([]invoiced.LineItem{}, lineItem)

invoiceToCreate.Items = lineItems

//Add a Payment Term
invoiceToCreate.PaymentTerms = invoiced.String("NET15")

invoiceResponse, err := client.Invoice.Create(invoiceToCreate)

if err != nil {
    panic(err)
}

fmt.Println("Invoice Response => ", invoiceResponse.Invoice)
```

If you want to use the sandbox API instead then you must set the second argument on the client to `true` like this:

```go
client := invoiced.New("SANDBOX_API_KEY", false)
```

## Developing

The test suite can be run with:

```
go test ./...
```
