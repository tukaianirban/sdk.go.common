package bruce

import "github.com/tukaianirban/sdk.go.common/errs"

func GetString(key string) (string, error) {

	if indexer == nil {
		return "", errs.ErrBruceNotInit
	}

	return indexer.Reader.GetString(key)
}

func GetInt(key string) (int, error) {

	return indexer.Reader.GetInt(key)
}

func GetFloat64(key string) (float64, error) {

	return indexer.Reader.GetFloat64(key)
}

func GetArray(key string) ([]interface{}, error) {

	return indexer.Reader.GetArray(key)
}

func GetMap(key string) (map[string]interface{}, error) {

	return indexer.Reader.GetMap(key)
}