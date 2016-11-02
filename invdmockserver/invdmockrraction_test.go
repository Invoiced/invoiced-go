package invdmockserver

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRRActionObject(t *testing.T) {

	jsonString := ` {
    "request": {
        "method": "POST",
        "url": "/customers"

    },
    "response": {
        "status": 400,
        "body": "{\n    \"type\": \"invalid_request\",\n    \"message\": \"Name missing\",\n    \"param\": \"name\"\n}",
        "headers": {
            "Content-Type": "application/json"
        }
    }
}`

	rrAction := new(RRActionObject)

	err := json.Unmarshal([]byte(jsonString), rrAction)

	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(rrAction.Response.Body)

}

func TestRRActionMap(t *testing.T) {

	jsonString := ` {
    "request": {
        "method": "POST",
        "url": "/customers",
         "bodyPatterns": [
          {
            "equalToJson": "{ \"email\":\"billing@acmecorp.com\",\n  \"collection_mode\":\"manual\",\n  \"payment_terms\":\"NET 30\",\n  \"type\":\"company\" \n }"
          }]

    },
    "response": {
        "status": 400,
        "body": "{\n    \"type\": \"invalid_request\",\n    \"message\": \"Name missing\",\n    \"param\": \"name\"\n}",
        "headers": {
            "Content-Type": "application/json"
        }
    }
}`

	rrActionObject1 := new(RRActionObject)

	err := json.Unmarshal([]byte(jsonString), rrActionObject1)

	if err != nil {
		t.Fatal(err)
	}

	rrActionMap := NewRRActionMap()

	rrActionMap.Put(rrActionObject1)

	rrActionObject2, found, err := rrActionMap.Get("POST", "/customers", "{ \n  \"collection_mode\":\"manual\",\n  \"payment_terms\":\"NET 30\",\n  \"type\":\"company\" \n,\"email\":\"billing@acmecorp.com\" }")

	if err != nil {
		t.Fatal(err)
	}

	if !found {
		t.Fatal("rrActionObject should have been found")
	}

	if !reflect.DeepEqual(rrActionObject1, rrActionObject2) {
		t.Fatal("Objects should be identical")
	}

	jsonString2 := ` {
    "request": {
        "method": "POST",
        "url": "/customers"
    },
    "response": {
        "status": 401,
        "body": "{\n    \"type\": \"Unauthorized\",\n    \"message\": \"Incorrect or missing API key\"\n}",
        "headers": {
            "Content-Type": "application/json"
        }
    }
}`

	rrActionObject3 := new(RRActionObject)

	err = json.Unmarshal([]byte(jsonString2), rrActionObject3)

	if err != nil {
		t.Fatal(err)
	}

	rrActionMap.Put(rrActionObject3)

	rrActionObject4, found, err := rrActionMap.Get("POST", "/customers", "")

	if err != nil {
		t.Fatal(err)
	}

	if !found {
		t.Fatal("rrActionObject should have been found")
	}

	if !reflect.DeepEqual(rrActionObject3, rrActionObject4) {
		t.Fatal("Objects should be identical")
	}
}
