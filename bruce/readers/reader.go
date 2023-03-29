package readers

import (
	"errors"
	"strings"
)

type Reader struct {
	data map[string]interface{}
}

func NewReaderFromMap(data map[string]interface{}) *Reader {
	return &Reader{
		data: data,
	}
}

func (self *Reader) get(key string) (interface{}, error) {

	keyparts := strings.Split(key, ".")
	if len(keyparts) == 0 {
		return nil, errors.New("invalid key format")
	}

	var localReader interface{} = self.data

	for _, part := range keyparts {

		v, ok := localReader.(map[string]interface{})
		if !ok {
			return nil, errors.New("key depth too much")
		}

		val, ok := v[part]
		if !ok {
			return nil, errors.New("key not found")
		}

		localReader = val
	}

	return localReader, nil
}

func (self *Reader) GetBoolean(key string) (bool, error) {

	boolVal, err := self.get(key)
	if err != nil {
		return false, err
	}

	v, ok := boolVal.(bool)
	if !ok {
		return false, errors.New("not a boolean value")
	} else {
		return v, nil
	}
}

func (self *Reader) GetString(key string) (string, error) {

	intVal, err := self.get(key)
	if err != nil {
		return "", err
	}

	v, ok := intVal.(string)
	if !ok {
		return "", errors.New("not a string value")
	} else {
		return v, nil
	}
}

func (self *Reader) GetInt(key string) (int, error) {

	intVal, err := self.get(key)
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

func (self *Reader) GetFloat64(key string) (float64, error) {

	intVal, err := self.get(key)
	if err != nil {
		return 0, err
	}

	v, ok := intVal.(float64)
	if !ok {
		return 0, errors.New("not a float64 value")
	}

	return v, nil
}

func (self *Reader) GetArray(key string) ([]interface{}, error) {

	intVal, err := self.get(key)
	if err != nil {
		return nil, err
	}

	v, ok := intVal.([]interface{})
	if !ok {
		return nil, errors.New("not a array value")
	}

	return v, nil
}

func (self Reader) GetMap(key string) (map[string]interface{}, error) {

	intVal, err := self.get(key)
	if err != nil {
		return nil, err
	}

	v, ok := intVal.(map[string]interface{})
	if !ok {
		return nil, errors.New("not a map value")
	}

	return v, nil
}
