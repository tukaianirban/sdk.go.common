package bruce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

var Bruce map[string]interface{}

func ReadFromFile(filename string) error {

	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(filebytes, &Bruce)
	return err
}

func get(key string) (interface{}, error) {

	if Bruce == nil {
		return nil, errors.New("Bruce not init")
	}

	keyparts := strings.Split(key, ".")
	if len(keyparts) == 0 {
		return nil, errors.New("invalid key format")
	}

	var localBruce interface{} = Bruce
	for _, part := range keyparts {

		v, ok := localBruce.(map[string]interface{})
		if !ok {
			return nil, errors.New("key depth too much")
		}

		val, ok := v[part]
		if !ok {
			return nil, errors.New("key not found")
		}

		localBruce = val
	}

	return localBruce, nil
}

func GetString(key string) (string, error) {

	intVal, err := get(key)
	if err != nil {
		return "", err
	}

	v, ok := intVal.(string)
	if !ok {
		return "", errors.New("not a string value")
	}

	return v, nil
}

func GetInt(key string) (int, error) {

	intVal, err := get(key)
	if err != nil {
		return 0, err
	}

	vfloat, ok := intVal.(float64)
	if !ok {
		return 0, errors.New("not a number value")
	}

	vint := int(vfloat)

	if float64(vint) != vfloat {
		return 0, errors.New("not a int value")
	} else {
		return vint, nil
	}
}

func GetFloat64(key string) (float64, error) {

	intVal, err := get(key)
	if err != nil {
		return 0, err
	}

	v, ok := intVal.(float64)
	if !ok {
		return 0, errors.New("not a float64 value")
	}

	return v, nil
}

func GetArray(key string) ([]interface{}, error) {

	intVal, err := get(key)
	if err != nil {
		return nil, err
	}

	v, ok := intVal.([]interface{})
	if !ok {
		return nil, errors.New("not a array value")
	}

	return v, nil
}

func GetMap(key string) (map[string]interface{}, error) {

	intVal, err := get(key)
	if err != nil {
		return nil, err
	}

	v, ok := intVal.(map[string]interface{})
	if !ok {
		return nil, errors.New("not a map value")
	}

	return v, nil
}
