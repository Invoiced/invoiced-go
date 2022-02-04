package invoice

import (
	"strconv"

	"github.com/strongdm/invoiced-go/v2"
)

type Client struct {
	*invoiced.Api
}

func (c *Client) Create(request *invoiced.InvoiceRequest) (*invoiced.Invoice, error) {
	resp := new(invoiced.Invoice)
	err := c.Api.Create("/invoices", request, resp)
	return resp, err
}

func (c *Client) Retrieve(id int64) (*invoiced.Invoice, error) {
	resp := new(invoiced.Invoice)
	_, err := c.Api.Get("/invoices/"+strconv.FormatInt(id, 10), resp)
	return resp, err
}

func (c *Client) Update(id int64, request *invoiced.InvoiceRequest) (*invoiced.Invoice, error) {
	resp := new(invoiced.Invoice)
	err := c.Api.Update("/invoices/"+strconv.FormatInt(id, 10), request, resp)
	return resp, err
}

func (c *Client) Void(id int64) (*invoiced.Invoice, error) {
	resp := new(invoiced.Invoice)
	err := c.Api.PostWithoutData("/invoices/"+strconv.FormatInt(id, 10)+"/void", resp)
	return resp, err
}

func (c *Client) Delete(id int64) error {
	return c.Api.Delete("/invoices/" + strconv.FormatInt(id, 10))
}

func (c *Client) Count() (int64, error) {
	return c.Api.Count("/invoices")
}

func (c *Client) ListAll(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Invoices, error) {
	return c.ListAllHelper(invoiced.AddFilterAndSort("/invoices", filter, sort), filter, sort)
}

func (c *Client) ListAllHelper(endpoint string, filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Invoices, error) {
	invoices := make(invoiced.Invoices, 0)
NEXT:

	tmpInvoices, endpoint, err := c.ListHelper(endpoint, filter, sort)

	if err != nil {
		return nil, err
	}

	invoices = append(invoices, tmpInvoices...)

	if endpoint != "" {
		goto NEXT
	}

	return invoices, nil
}

func (c *Client) ListHelper(url string, filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Invoices, string, error) {
	if len(url) == 0 {
		url = invoiced.AddFilterAndSort("/invoices", filter, sort)
	}

	invoices := make(invoiced.Invoices, 0)

	nextEndpoint, err := c.Api.Get(url, &invoices)
	if err != nil {
		return nil, "", err
	}

	return invoices, nextEndpoint, nil
}

func (c *Client) ListAllInvoicesStartDate(filter *invoiced.Filter, sort *invoiced.Sort, invoiceDate int64) (invoiced.Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, invoiceDate, 0)
}

func (c *Client) ListAllInvoicesEndDate(filter *invoiced.Filter, sort *invoiced.Sort, invoiceDate int64) (invoiced.Invoices, error) {
	return c.ListAllInvoicesStartEndDate(filter, sort, 0, invoiceDate)
}

func (c *Client) ListAllInvoicesStartEndDate(filter *invoiced.Filter, sort *invoiced.Sort, startDate, endDate int64) (invoiced.Invoices, error) {
	url := "/invoices"
	url = invoiced.AddFilterAndSort(url, filter, sort)

	if startDate > 0 {
		startDateString := strconv.FormatInt(startDate, 10)
		url = invoiced.AddQueryParameter(url, "start_date", startDateString)
	}

	if endDate > 0 {
		endDateString := strconv.FormatInt(endDate, 10)
		url = invoiced.AddQueryParameter(url, "end_date", endDateString)
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *Client) ListAllInvoicesUpdatedDate(filter *invoiced.Filter, sort *invoiced.Sort, invoiceDate int64) (invoiced.Invoices, error) {
	url := "/invoices"
	url = invoiced.AddFilterAndSort(url, filter, sort)

	if invoiceDate > 0 {
		updatedAfterString := strconv.FormatInt(invoiceDate, 10)
		url = invoiced.AddQueryParameter(url, "updated_after", updatedAfterString)
	}

	return c.ListAllHelper(url, filter, sort)
}

func (c *Client) List(filter *invoiced.Filter, sort *invoiced.Sort) (invoiced.Invoices, string, error) {
	return c.ListHelper("", filter, sort)
}

func (c *Client) ListInvoiceByNumber(invoiceNumber string) (*invoiced.Invoice, error) {
	filter := invoiced.NewFilter()
	err := filter.Set("number", invoiceNumber)
	if err != nil {
		return nil, err
	}

	invoices, apiError := c.ListAll(filter, nil)

	if apiError != nil {
		return nil, apiError
	}

	if len(invoices) == 0 {
		return nil, nil
	}

	return invoices[0], nil
}

func (c *Client) SendEmail(id int64, request *invoiced.SendEmailRequest) error {
	return c.Api.Create("/invoices/"+strconv.FormatInt(id, 10)+"/emails", request, nil)
}

func (c *Client) SendText(id int64, request *invoiced.SendTextMessageRequest) (invoiced.TextMessages, error) {
	resp := new(invoiced.TextMessages)
	err := c.Api.Create("/invoices/"+strconv.FormatInt(id, 10)+"/text_messages", request, resp)
	return *resp, err
}

func (c *Client) SendLetter(id int64) (*invoiced.Letter, error) {
	resp := new(invoiced.Letter)
	err := c.Api.Create("/invoices/"+strconv.FormatInt(id, 10)+"/letters", nil, resp)
	return resp, err
}

func (c *Client) Pay(id int64) (*invoiced.Invoice, error) {
	resp := new(invoiced.Invoice)
	err := c.Api.Create("/invoices/"+strconv.FormatInt(id, 10)+"/pay", nil, resp)
	return resp, err
}

func (c *Client) ListAttachments(id int64) (invoiced.Files, error) {
	endpoint := "/invoices/" + strconv.FormatInt(id, 10) + "/attachments"

	files := make(invoiced.Files, 0)

NEXT:
	tempFiles := make(invoiced.Files, 0)

	endpoint, err := c.Api.Get(endpoint, &tempFiles)
	if err != nil {
		return nil, err
	}

	files = append(files, tempFiles...)

	if endpoint != "" {
		goto NEXT
	}

	return files, nil
}

func (c *Client) RetrieveNotes(id int64) (invoiced.Notes, error) {
	endpoint := "/invoices/" + strconv.FormatInt(id, 10) + "/notes"

	notes := make(invoiced.Notes, 0)

NEXT:
	tmpNotes := make(invoiced.Notes, 0)

	endpoint, err := c.Api.Get(endpoint, &tmpNotes)

	if err != nil {
		return nil, err
	}

	notes = append(notes, tmpNotes...)

	if endpoint != "" {
		goto NEXT
	}

	return notes, nil
}

func (c *Client) CreatePaymentPlan(id int64, request *invoiced.PaymentPlanRequest) (*invoiced.PaymentPlan, error) {
	resp := new(invoiced.PaymentPlan)
	err := c.Api.Create("/invoices/"+strconv.FormatInt(id, 10)+"/payment_plan", request, resp)
	return resp, err
}

func (c *Client) RetrievePaymentPlan(id int64) (*invoiced.PaymentPlan, error) {
	resp := new(invoiced.PaymentPlan)
	_, err := c.Api.Get("/invoices/"+strconv.FormatInt(id, 10)+"/payment_plan", resp)
	return resp, err
}

func (c *Client) CancelPaymentPlan(id int64) error {
	return c.Api.Delete("/invoices/" + strconv.FormatInt(id, 10) + "/payment_plan")
}
