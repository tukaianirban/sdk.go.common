package bruce

import (
	"log"
)

//
// different data stores must return config clients adhering to this interface
//
type DataStoreInterface interface {
	GetBoolean(string) (bool, error)
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

/**
Mode: determines if config is fetched from local file, gcp, aws, etc
args: optional param depending on the mode (ex: for local mode, the config file)
**/
func Init(mode int, args ...string) {

	if configIndexer != nil {
		log.Fatal("bruce has already been init")
	}

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
		// if len(args) < 1 {
		// 	log.Fatalf("error: GCP config init params not provided")
		// }

		// pass in the GCP project-id
		// configIndexer = initConfigModeGCP(args[0])
		fallthrough

	default:
		log.Fatalf("fatal: config mode not implemented")
	}
}

func GetBoolean(key string) (bool, error) {

	return configIndexer.GetBoolean(key)
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
