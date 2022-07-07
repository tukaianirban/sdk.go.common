package bruce

import (
	"log"
)

//
// different data stores must return config clients adhering to this interface
//
type DataStoreInterface interface {
	GetString(string) (string, error)
	GetInt(string) (int, error)
	GetFloat64(string) (float64, error)
	GetArray(string) ([]interface{}, error)
	GetMap(string) (map[string]interface{}, error)
}

const (
	MODE_LOCAL = iota
	MODE_GCP
	MODE_AWS
)

var configIndexer DataStoreInterface

func Init(mode int, args ...string) {

	var err error

	switch mode {
	case MODE_LOCAL:
		if len(args) < 1 {
			log.Fatalf("error: Local config init params not provided")
		}

		// pass in the local config file name
		if configIndexer, err = initConfigModeLocal(args[0]); err != nil {
			log.Fatalf("error initializing local mode config, reason=%s", err.Error())
		}

	case MODE_GCP:
		if len(args) < 1 {
			log.Fatalf("error: GCP config init params not provided")
		}

		// pass in the GCP project-id
		configIndexer = initConfigModeGCP(args[0])

	default:
		log.Fatalf("fatal: config mode not implemented")
	}
}

func GetString(key string) (string, error) {

	return configIndexer.GetString(key)
}

func GetInt(key string) (int, error) {

	return configIndexer.GetInt(key)
}

func GetFloat64(key string) (float64, error) {

	return configIndexer.GetFloat64(key)
}

func GetArray(key string) ([]interface{}, error) {

	return configIndexer.GetArray(key)
}

func GetMap(key string) (map[string]interface{}, error) {

	return configIndexer.GetMap(key)
}
