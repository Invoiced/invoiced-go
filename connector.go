package invdapi

import (
	"bytes"
	"encoding/json"
	"github.com/Invoiced/invoiced-go/invdendpoint"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const requestURL = "https://api.invoiced.com"
const requestType = "application/json"

const version = "0.1"

func Version() string {
	return version
}

type Connection struct {
	key             string
	itemsPerRequest int
	client          *http.Client
	url             string
}

type InvoicedToken struct {
	Key string `json:"invoicedApiKey"`
}

func NewConnection(key string) *Connection {
	c := new(Connection)
	c.key = key
	c.client = new(http.Client)
	c.url = requestURL

	return c
}

func checkStatusForError(status int, r io.Reader) *APIError {
	if status < 400 {
		return nil
	}

	body, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	apiError := new(APIError)

	err = json.Unmarshal(body, apiError)

	if err != nil {
		apiError.Type = string(body)

	}

	return apiError

}

func pushDataIntoStruct(endPointData interface{}, respBody io.Reader) {

	body, err := ioutil.ReadAll(respBody)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, endPointData)

	if err != nil {
		panic(err)
	}

}

func parseURLLinksFromString(s string) map[string]string {

	urlAndLinkMap := make(map[string]string)

	rawURLLinksAndRelations := strings.Split(s, ",")

	for _, rawURLLinkRelation := range rawURLLinksAndRelations {
		parsedRawURLAndRelation := strings.Split(rawURLLinkRelation, ";")
		url := parseRawURL(parsedRawURLAndRelation[0])
		relation := parseRawRelation(parsedRawURLAndRelation[1])

		urlAndLinkMap[relation] = url
	}

	return urlAndLinkMap
}

func parseRawRelation(s string) string {
	//parse rel="last" => last

	first := strings.Index(s, "\"")
	last := strings.LastIndex(s, "\"")

	trimmed := s[first+1 : last]

	trimmed = strings.TrimSpace(trimmed)

	return trimmed

}

func parseRawURL(s string) string {
	//<https://api.invoiced.com/invoices?page=1>

	trimmed := strings.Trim(s, " < ")

	trimmed = strings.Trim(trimmed, " > ")

	trimmed = strings.TrimSpace(trimmed)

	return trimmed

}

func addFilterSortToEndPoint(endpoint string, filter *invdendpoint.Filter, sort *invdendpoint.Sort) string {

	emptyFilter := true
	emptySort := true

	if filter != nil && filter.String() != "" {
		emptyFilter = false
	}

	if sort != nil && sort.String() != "" {
		emptySort = false
	}

	if !emptyFilter && !emptySort {
		return endpoint + "?" + filter.String() + "&" + sort.String()
	} else if !emptyFilter && emptySort {
		return endpoint + "?" + filter.String()
	} else if emptyFilter && !emptySort {
		return endpoint + "?" + sort.String()
	}

	return endpoint

}

func makeEndPointSingular(endpoint string, id int64) string {
	return endpoint + "/" + strconv.FormatInt(id, 10)
}

func (c *Connection) makeEndPointURL(endPoint string) string {

	return c.url + endPoint
}

func (c *Connection) setItemsPerRequest(items int) {
	c.itemsPerRequest = items
}

func (c *Connection) get(endPoint string) *http.Response {

	req, err := http.NewRequest("GET", endPoint, nil)
	req.SetBasicAuth(c.key, "")

	resp, err := c.client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp

}

func (c *Connection) post(endPoint string, body io.Reader) *http.Response {

	req, err := http.NewRequest("POST", endPoint, body)
	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp

}

func (c *Connection) patch(endPoint string, body io.Reader) *http.Response {

	req, err := http.NewRequest("PATCH", endPoint, body)
	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp

}

func (c *Connection) deleteRequest(endPoint string) *http.Response {

	req, err := http.NewRequest("DELETE", endPoint, nil)
	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp

}

func (c *Connection) create(endPoint string, requestData interface{}, responseData interface{}) *APIError {

	b, err := json.Marshal(requestData)

	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(b)

	resp := c.post(endPoint, body)

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	pushDataIntoStruct(responseData, resp.Body)

	return nil

}

func (c *Connection) delete(endPoint string) *APIError {

	resp := c.deleteRequest(endPoint)

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	return nil

}

func (c *Connection) update(endPoint string, requestData interface{}, responseData interface{}) *APIError {

	b, err := json.Marshal(requestData)

	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(b)

	resp := c.patch(endPoint, body)

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	pushDataIntoStruct(responseData, resp.Body)

	return nil

}

func (c *Connection) count(endPoint string) (int64, *APIError) {
	resp := c.get(endPoint)

	defer resp.Body.Close()

	apiErr := checkStatusForError(resp.StatusCode, resp.Body)

	if apiErr != nil {
		return 0, apiErr
	}

	s := resp.Header.Get("X-Total-Count")

	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return i, nil

}

func (c *Connection) retrieveDataFromAPI(endPoint string, endPointData interface{}) (string, *APIError) {

	nextURL := ""

	resp := c.get(endPoint)

	defer resp.Body.Close()

	link := resp.Header.Get("Link")

	if link != "" {
		nextMap := parseURLLinksFromString(link)

		if nextMap["self"] != nextMap["next"] {
			nextURL = nextMap["next"]
		}
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return "", apiError
	}

	pushDataIntoStruct(endPointData, resp.Body)

	return nextURL, nil

}
