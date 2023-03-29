package local

/**
The fileIndexer object must implement the interface methods
of Bruce to be able to comply to being a provider for bruce configs
*/

func (indexer *FileReader) GetBoolean(key string) (bool, error) {

	return indexer.Reader.GetBoolean(key)
}

func (indexer *FileReader) GetString(key string) (string, error) {

	return indexer.Reader.GetString(key)
}

func (indexer *FileReader) GetInt(key string) (int, error) {

	return indexer.Reader.GetInt(key)
}

func (indexer *FileReader) GetFloat64(key string) (float64, error) {

	return indexer.Reader.GetFloat64(key)
}

func (indexer *FileReader) GetArray(key string) ([]interface{}, error) {

	return indexer.Reader.GetArray(key)
}

func (indexer *FileReader) GetMap(key string) (map[string]interface{}, error) {

	return indexer.Reader.GetMap(key)
}
