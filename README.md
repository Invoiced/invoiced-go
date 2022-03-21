invoiced-go
========

**SDM changes**: 
- Add Active flag to Customer
- Add 'customer' tag to InvoiceRequest


========

This repository contains the Go client library for the [Invoiced](https://invoiced.com) API.

[![Build Status](https://travis-ci.com/Invoiced/invoiced-go.svg?branch=master)](https://travis-ci.com/Invoiced/invoiced-go)
[![Coverage Status](https://coveralls.io/repos/github/Invoiced/invoiced-go/badge.svg?branch=master)](https://coveralls.io/github/Invoiced/invoiced-go?branch=master)

## Installing

The Invoiced package can be installed like this:

```
go get -u https://github.com/strongdm/invoiced-go/v2
```

## Requirements

- Go 1.11+

## Usage

First, you must instantiate a new client

```go
import "github.com/strongdm/invoiced-go/v2/api"

client := api.New("API_KEY", false)
```

Then, API calls can be made like this:

```go
import "github.com/strongdm/invoiced-go/v2"
import "fmt"

// Get all invoices with auto pagination
invoices, err := client.Invoice.ListAll(nil, nil)
if err != nil {
    panic(err)
}

// Let's print all the invoices
for _, invoice := range invoices {
    fmt.Println(invoice)
}

// Let's create a new customer
customer, err := client.Customer.Create(&invoiced.CustomerRequest{
	Name: invoiced.String("Test Customer")
})

if err != nil {
    panic(err)
}

fmt.Println("Customer Response => ", customer)

// Let's create a new invoice
invoice, err := client.Invoice.Create(&invoiced.InvoiceRequest{
    Customer: invoiced.Int64(customerResponse.Id),
    PaymentTerms: invoiced.String("NET 30"),
    Items: []*invoiced.LineItemRequest{
        {
            Description: invoiced.String("Retina MacBook Pro"),
            Quantity: invoiced.Float64(5),
            UnitCost: invoiced.Float64(1999.22),
        },
    },
})

if err != nil {
    panic(err)
}

fmt.Println("Invoice Response => ", invoice)
```

If you want to use the sandbox API instead then you must set the second argument on the client to `true` like this:

```go
client := api.New("SANDBOX_API_KEY", false)
```

## Developing

The test suite can be run with:

```
go test ./...
```
