package require

import (
	"fmt"
)

type RequiredConfig struct {
	paths map[string]string
}

type PathInterfaceGetter interface {
	Get(path string) interface{}
}

func New() *RequiredConfig {
	return &RequiredConfig{
		paths: make(map[string]string),
	}
}

func (r *RequiredConfig) RequireBool(path string) {
	r.paths[path] = "bool"
}

func (r *RequiredConfig) RequireInt(path string) {
	r.paths[path] = "int"
}

func (r *RequiredConfig) RequireInt64(path string) {
	r.paths[path] = "int"
}

func (r *RequiredConfig) RequireFloat32(path string) {
	r.paths[path] = "float32"
}

func (r *RequiredConfig) RequireFloat64(path string) {
	r.paths[path] = "float64"
}

func (r *RequiredConfig) RequireString(path string) {
	r.paths[path] = "string"
}

func (r *RequiredConfig) RequireBoolSlice(path string) {
	r.paths[path] = "[]bool"
}

func (r *RequiredConfig) RequireIntSlice(path string) {
	r.paths[path] = "[]int"
}

func (r *RequiredConfig) RequireFloatSlice(path string) {
	r.paths[path] = "[]float"
}

func (r *RequiredConfig) RequireFloat64Slice(path string) {
	r.paths[path] = "[]float64"
}

func (r *RequiredConfig) RequireStringSlice(path string) {
	r.paths[path] = "[]string"
}

func (r *RequiredConfig) Verify(config PathInterfaceGetter) []error {
	var es []error
	for path, stringType := range r.paths {
		value := config.Get(path)
		switch value.(type) {
		case bool:
			if stringType != "bool" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case int:
			if stringType != "int" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case int64:
			if stringType != "int64" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case float32:
			if stringType != "float32" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case float64:
			if stringType != "float64" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case string:
			if stringType != "string" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case []bool:
			if stringType != "[]bool" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case []int:
			if stringType != "[]int" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case []float32:
			if stringType != "[]float32" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case []float64:
			if stringType != "[]float64" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case []string:
			if stringType != "[]string" {
				es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			}
			break
		case nil:
			es = append(es, fmt.Errorf("config value %s did not match expected type %s", path, stringType))
			break
		default:
			es = append(es, fmt.Errorf("unexpected type found at %s", path))
		}
	}

	return es
}