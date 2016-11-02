package invdutil

import "testing"

func TestReadAPIKeyFromYaml(t *testing.T) {

	path := "test.yaml"
	apiKey, _ := ReadAPIKeyFromYaml(path)

	if apiKey != "777ab44$$$@!2131211" {
		t.Fatal("Error Reading API Key")
	}

}
