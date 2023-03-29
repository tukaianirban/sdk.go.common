package bruce

//
// Each configuration data store (local, gcp, aws) must implement this interface
//
type ConfigStore interface {
	GetBoolean(string) (bool, error)
	GetString(string) (string, error)
	GetInt(string) (int, error)
	GetFloat64(string) (float64, error)
	GetArray(string) ([]interface{}, error)
	GetMap(string) (map[string]interface{}, error)
}

func GetBoolean(key string) (bool, error) {

	return instance.GetBoolean(key)
}

func GetString(key string) (string, error) {

	return instance.GetString(key)
}

func GetInt(key string) (int, error) {

	return instance.GetInt(key)
}

func GetFloat64(key string) (float64, error) {

	return instance.GetFloat64(key)
}

func GetArray(key string) ([]interface{}, error) {

	return instance.GetArray(key)
}

func GetMap(key string) (map[string]interface{}, error) {

	return instance.GetMap(key)
}
