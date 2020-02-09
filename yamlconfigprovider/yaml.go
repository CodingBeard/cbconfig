package yamlconfigprovider

import (
	"github.com/codingbeard/cbconfig/flatten"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type Yaml struct{}

func (y Yaml) Unmarshall(reader io.Reader) (map[string]interface{}, error) {
	defaultConfBytes, e := ioutil.ReadAll(reader)

	if e != nil {
		return make(map[string]interface{}), e
	}

	var yamlConfig map[string]interface{}

	e = yaml.Unmarshal(defaultConfBytes, &yamlConfig)
	if e != nil {
		return make(map[string]interface{}), e
	}

	out, e := flatten.Flatten(yamlConfig, "", flatten.DotStyle)
	if e != nil {
		return make(map[string]interface{}), e
	}

	return out, e
}
