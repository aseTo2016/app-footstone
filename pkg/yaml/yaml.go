package yaml

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

// LoadYamlData loads data from yaml file.
func LoadYamlData(fileName string, result interface{}) error {
	rawData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(rawData, result)
}
