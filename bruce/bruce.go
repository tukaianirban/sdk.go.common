package bruce

import (
	"github.com/tukaianirban/sdk.go.common/jsonreader"
)

var bruceObject *jsonreader.JsonReader

func ReadFromFile(filename string) error {

	var err error

	bruceObject, err = jsonreader.ReadFromFile(filename)
	return err
}

// func get(key string) (interface{}, error) {

// 	if bruceObject == nil {
// 		return nil, errors.New("bruce not init")
// 	}

// 	keyparts := strings.Split(key, ".")
// 	if len(keyparts) == 0 {
// 		return nil, errors.New("invalid key format")
// 	}

// 	var localBruce interface{} = bruceObject
// 	for _, part := range keyparts {

// 		v, ok := localBruce.(map[string]interface{})
// 		if !ok {
// 			return nil, errors.New("key depth too much")
// 		}

// 		val, ok := v[part]
// 		if !ok {
// 			return nil, errors.New("key not found")
// 		}

// 		localBruce = val
// 	}

// 	return localBruce, nil
// }

func GetString(key string) (string, error) {

	return bruceObject.GetString(key)
}

func GetInt(key string) (int, error) {

	return bruceObject.GetInt(key)
}

func GetFloat64(key string) (float64, error) {

	return bruceObject.GetFloat64(key)
}

func GetArray(key string) ([]interface{}, error) {

	return bruceObject.GetArray(key)
}

func GetMap(key string) (map[string]interface{}, error) {

	return bruceObject.GetMap(key)
}
