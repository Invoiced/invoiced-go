package invdutil

import "gopkg.in/yaml.v2"
import "path/filepath"
import "io/ioutil"

type YamlConfigFile struct {
	Apikey string
}

func ReadAPIKeyFromYaml(path string) string {
	filename, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var yamlConfig YamlConfigFile

	err = yaml.Unmarshal(yamlFile, &yamlConfig)

	if err != nil {
		panic(err)
	}

	return yamlConfig.Apikey

}
