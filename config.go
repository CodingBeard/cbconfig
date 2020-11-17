package cbconfig

import (
	"fmt"
	"io"
)

type Config struct {
	data map[string]interface{}
}

type Provider interface {
	Unmarshall(reader io.Reader) (map[string]interface{}, error)
}

func New() *Config {
	return &Config{
		data: make(map[string]interface{}),
	}
}

func (c *Config) SetFromReader(provider Provider, reader io.Reader) error {
	c.data = make(map[string]interface{})
	return c.OverrideFromReader(provider, reader)
}

func (c *Config) OverrideFromReader(provider Provider, reader io.Reader) error {
	out, e := provider.Unmarshall(reader)
	if e != nil {
		return e
	}

	for key, value := range out {
		c.data[key] = value
	}

	return nil
}

func (c *Config) Get(path string) interface{} {
	return c.data[path]
}

func (c *Config) GetRequired(path string) (interface{}, error) {
	value, ok := c.data[path]
	if !ok {
		return value, fmt.Errorf("%s not found in the config", path)
	}
	return value, nil
}

func (c *Config) GetBool(path string) bool {
	value := c.data[path]
	boolValue, _ := value.(bool)
	return boolValue
}

func (c *Config) GetRequiredBool(path string) (bool, error) {
	var returnValue bool
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	returnValue, ok := value.(bool)
	if !ok {
		return returnValue, fmt.Errorf("%s was not a bool", path)
	}
	return returnValue, nil
}

func (c *Config) GetInt(path string) int {
	value := c.data[path]
	intValue, _ := value.(int)
	return intValue
}

func (c *Config) GetRequiredInt(path string) (int, error) {
	var returnValue int
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	returnValue, ok := value.(int)
	if !ok {
		return returnValue, fmt.Errorf("%s was not a int", path)
	}
	return returnValue, nil
}

func (c *Config) GetFloat32(path string) float32 {
	value := c.data[path]
	float32Value, _ := value.(float32)
	return float32Value
}

func (c *Config) GetRequiredFloat32(path string) (float32, error) {
	var returnValue float32
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	returnValue, ok := value.(float32)
	if !ok {
		return returnValue, fmt.Errorf("%s was not a float32", path)
	}
	return returnValue, nil
}

func (c *Config) GetFloat64(path string) float64 {
	value := c.data[path]
	float64Value, _ := value.(float64)
	return float64Value
}

func (c *Config) GetRequiredFloat64(path string) (float64, error) {
	var returnValue float64
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	returnValue, ok := value.(float64)
	if !ok {
		return returnValue, fmt.Errorf("%s was not a float64", path)
	}
	return returnValue, nil
}

func (c *Config) GetString(path string) string {
	value := c.data[path]
	stringValue, _ := value.(string)
	return stringValue
}

func (c *Config) GetRequiredString(path string) (string, error) {
	var returnValue string
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	returnValue, ok := value.(string)
	if !ok {
		return returnValue, fmt.Errorf("%s was not a string", path)
	}
	return returnValue, nil
}

func (c *Config) GetIntSlice(path string) []int {
	value := c.data[path]
	var returnValue []int
	interfaceSlice, ok := value.([]interface{})
	if ok {
		for _, intInterface := range interfaceSlice {
			intValue, ok := intInterface.(int)
			if ok {
				returnValue = append(returnValue, intValue)
			}
		}
	}
	return returnValue
}

func (c *Config) GetRequiredIntSlice(path string) ([]int, error) {
	var returnValue []int
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	intSliceInterfaces, ok := value.([]interface{})
	if ok {
		for _, intInterface := range intSliceInterfaces {
			intValue, ok := intInterface.(int)
			if ok {
				returnValue = append(returnValue, intValue)
			} else {
				return returnValue, fmt.Errorf("%s was not a int slice", path)
			}
		}
	} else {
		return returnValue, fmt.Errorf("%s was not a int slice", path)
	}
	return returnValue, nil
}

func (c *Config) GetStringSlice(path string) []string {
	value := c.data[path]
	var returnValue []string
	interfaceSlice, ok := value.([]interface{})
	if ok {
		for _, stringInterface := range interfaceSlice {
			stringValue, ok := stringInterface.(string)
			if ok {
				returnValue = append(returnValue, stringValue)
			}
		}
	}
	return returnValue
}

func (c *Config) GetRequiredStringSlice(path string) ([]string, error) {
	var returnValue []string
	value, e := c.GetRequired(path)
	if e != nil {
		return returnValue, e
	}
	stringSliceInterfaces, ok := value.([]interface{})
	if ok {
		for _, stringInterface := range stringSliceInterfaces {
			stringValue, ok := stringInterface.(string)
			if ok {
				returnValue = append(returnValue, stringValue)
			} else {
				return returnValue, fmt.Errorf("%s was not a string slice", path)
			}
		}
	} else {
		return returnValue, fmt.Errorf("%s was not a string slice", path)
	}
	return returnValue, nil
}
