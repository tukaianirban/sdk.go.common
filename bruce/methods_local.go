package bruce

import (
	"os"
	"time"

	"github.com/tukaianirban/sdk.go.common/log"

	"github.com/tukaianirban/sdk.go.common/jsonreader"
)

type fileIndexer struct {
	Filename    string
	Reader      *jsonreader.JsonReader
	FileModTime time.Time
}

var indexerInstance *fileIndexer

/**
create a new indexer by loading config from the local file
**/
func initConfigModeLocal(configFile string) (*fileIndexer, error) {

	reader, err := jsonreader.ReadFromFile(configFile)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}

	indexerInstance = &fileIndexer{
		Filename:    configFile,
		Reader:      reader,
		FileModTime: fileInfo.ModTime(),
	}

	// if err := readFromFile(configFile); err != nil {
	// 	return nil, err
	// }

	go indexerInstance.watcher()

	return indexerInstance, nil
}

//
// execute this as a goroutine to keep a watch of any changes to the current config file
// no need for a terminate signal for the goroutine -> cleanup later
//
func (indexer fileIndexer) watcher() {

	ticker := time.NewTicker(30 * time.Second)

	for {

		<-ticker.C

		fileInfo, err := os.Stat(indexer.Filename)
		if err != nil {
			log.Print("error: failed to read fileinfo of configfile, reason=%s", err.Error())
			break
		}

		if !fileInfo.ModTime().Equal(indexer.FileModTime) {
			log.Print("change in local config file detected")

			// if loading new changed configfile causes error, ignore it and continue working
			// as previously
			if err = indexer.reloadConfig(); err != nil {
				log.Print("error: failed to read in changed configfile, reason=%s", err.Error())
			}
		}
	}
}

// func readFromFile(filename string) error {

// 	reader, err := jsonreader.ReadFromFile(filename)
// 	if err != nil {
// 		return err
// 	}

// 	fileInfo, err := os.Stat(filename)
// 	if err != nil {
// 		return err
// 	}

// 	indexer = &fileIndexer{
// 		Filename:    filename,
// 		Reader:      reader,
// 		FileModTime: fileInfo.ModTime(),
// 	}

// 	return nil
// }

/**
reload config from the same file
**/
func (indexer *fileIndexer) reloadConfig() error {

	reader, err := jsonreader.ReadFromFile(indexer.Filename)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(indexer.Filename)
	if err != nil {
		return err
	}

	indexer.Reader = reader
	indexer.FileModTime = fileInfo.ModTime()

	return nil
}

func (indexer fileIndexer) GetBoolean(key string) (bool, error) {

	return indexer.Reader.GetBoolean(key)
}

func (indexer fileIndexer) GetString(key string) (string, error) {

	return indexer.Reader.GetString(key)
}

func (indexer fileIndexer) GetInt(key string) (int, error) {

	return indexer.Reader.GetInt(key)
}

func (indexer fileIndexer) GetFloat64(key string) (float64, error) {

	return indexer.Reader.GetFloat64(key)
}

func (indexer fileIndexer) GetArray(key string) ([]interface{}, error) {

	return indexer.Reader.GetArray(key)
}

func (indexer fileIndexer) GetMap(key string) (map[string]interface{}, error) {

	return indexer.Reader.GetMap(key)
}
