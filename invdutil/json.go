package invdutil

import (
	"encoding/json"
	"reflect"
)

func JsonEqual(jsonString1, jsonString2 string) (bool, error) {
	var jsonObject1 interface{}
	var jsonObject2 interface{}

	err := json.Unmarshal([]byte(jsonString1), &jsonObject1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(jsonString2), &jsonObject2)

	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(jsonObject1, jsonObject2), nil
}
