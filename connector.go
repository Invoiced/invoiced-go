package invdapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Invoiced/invoiced-go/invdendpoint"
)

const (
	productionUrl = "https://api.invoiced.com"
	sandboxUrl    = "https://api.sandbox.invoiced.com"
	requestType   = "application/json"
)

const version = "6.0.0"

func Version() string {
	return version
}

type Connection struct {
	key     string
	client  *http.Client
	baseUrl string
}

func NewConnection(key string, sandbox bool) *Connection {
	url := productionUrl
	if sandbox {
		url = sandboxUrl
	}

	return &Connection{
		key:     key,
		client:  new(http.Client),
		baseUrl: url,
	}
}

func checkStatusForError(status int, r io.Reader) error {
	if status < 400 {
		return nil
	}

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	apiError := new(APIError)

	err = json.Unmarshal(body, apiError)

	if err != nil {
		apiError.Type = string(body)
	}

	return errors.New(apiError.Error())
}

func pushDataIntoStruct(endpointData interface{}, respBody io.Reader) error {
	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, endpointData)

	if err != nil {
		return err
	}

	return nil
}

func parseLinkHeader(s string) map[string]string {
	urlAndLinkMap := make(map[string]string)

	rawURLLinksAndRelations := strings.Split(s, ",")

	for _, rawURLLinkRelation := range rawURLLinksAndRelations {
		parsedRawURLAndRelation := strings.Split(rawURLLinkRelation, ";")
		url := parseLinkUrl(parsedRawURLAndRelation[0])
		relation := parseRelValue(parsedRawURLAndRelation[1])

		urlAndLinkMap[relation] = url
	}

	return urlAndLinkMap
}

func parseRelValue(s string) string {
	// parse rel="last" => last

	first := strings.Index(s, "\"")
	last := strings.LastIndex(s, "\"")

	trimmed := s[first+1 : last]

	trimmed = strings.TrimSpace(trimmed)

	return trimmed
}

func parseLinkUrl(s string) string {
	//<https://api.invoiced.com/invoices?page=1>
	trimmed := strings.TrimSpace(s)

	trimmed = strings.Trim(trimmed, "<")

	trimmed = strings.Trim(trimmed, ">")

	trimmed = strings.TrimSpace(trimmed)

	return trimmed
}

func addFilterAndSort(url string, filter *invdendpoint.Filter, sort *invdendpoint.Sort) string {
	emptyFilter := true
	emptySort := true

	if filter != nil && filter.String() != "" {
		emptyFilter = false
	}

	if sort != nil && sort.String() != "" {
		emptySort = false
	}

	if !emptyFilter && !emptySort {
		return url + "?" + filter.String() + "&" + sort.String()
	} else if !emptyFilter && emptySort {
		return url + "?" + filter.String()
	} else if emptyFilter && !emptySort {
		return url + "?" + sort.String()
	}

	return url
}

func addQueryParameter(url string, name string, value string) string {
	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}

	return url + name + "=" + value
}

func (c *Connection) get(endpoint string) (*http.Response, error) {
	url := c.baseUrl + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.key, "")

	resp, err := c.client.Do(req)

	return resp, err
}

func (c *Connection) post(endpoint string, body io.Reader) (*http.Response, error) {
	url := c.baseUrl + endpoint
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	return resp, err
}

func (c *Connection) patch(endpoint string, body io.Reader) (*http.Response, error) {
	url := c.baseUrl + endpoint
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	return resp, err
}

func (c *Connection) deleteRequest(endpoint string) (*http.Response, error) {
	url := c.baseUrl + endpoint
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.key, "")
	req.Header.Set("Content-Type", requestType)

	resp, err := c.client.Do(req)

	return resp, err
}

func (c *Connection) create(endpoint string, requestData interface{}, responseData interface{}) error {
	b, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(b)

	resp, err := c.post(endpoint, body)
	if err != nil {
		return err
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	err = pushDataIntoStruct(responseData, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) delete(endpoint string) error {
	resp, err := c.deleteRequest(endpoint)
	if err != nil {
		return err
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	return nil
}

func (c *Connection) update(endpoint string, requestData interface{}, responseData interface{}) error {
	b, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(b)

	resp, err := c.patch(endpoint, body)
	if err != nil {
		return err
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	err = pushDataIntoStruct(responseData, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) postWithoutData(endpoint string, responseData interface{}) error {
	resp, err := c.post(endpoint, nil)
	if err != nil {
		return err
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return apiError
	}

	err = pushDataIntoStruct(responseData, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) count(endpoint string) (int64, error) {
	resp, err := c.get(endpoint)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	apiErr := checkStatusForError(resp.StatusCode, resp.Body)

	if apiErr != nil {
		return 0, apiErr
	}

	s := resp.Header.Get("X-Total-Count")

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func (c *Connection) retrieveDataFromAPI(endpoint string, endpointData interface{}) (string, error) {
	nextURL := ""

	resp, err := c.get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	link := resp.Header.Get("Link")

	if link != "" {
		nextMap := parseLinkHeader(link)

		if nextMap["self"] != nextMap["next"] {
			nextURL = nextMap["next"]
		}
	}

	apiError := checkStatusForError(resp.StatusCode, resp.Body)

	if apiError != nil {
		return "", apiError
	}

	err = pushDataIntoStruct(endpointData, resp.Body)

	if err != nil {
		return "", err
	}

	return strings.Replace(nextURL,c.baseUrl,"",-1), nil
}
