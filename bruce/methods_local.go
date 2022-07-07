package bruce

import (
	"log"
	"os"
	"time"

	"github.com/tukaianirban/sdk.go.common/jsonreader"
)

type fileIndexer struct {
	Filename    string
	Reader      *jsonreader.JsonReader
	FileModTime time.Time
}

var indexer *fileIndexer

func initConfigModeLocal(configFile string) (*fileIndexer, error) {

	if err := readFromFile(configFile); err != nil {
		return nil, err
	}

	go indexer.watcher()

	return indexer, nil
}

//
// execute this as a goroutine to keep a watch of any changes to the current config file
// no need for a terminate signal for the goroutine -> cleanup later
//
func (self fileIndexer) watcher() {

	ticker := time.NewTicker(30 * time.Second)

	for {

		<-ticker.C

		fileInfo, err := os.Stat(indexer.Filename)
		if err != nil {
			log.Printf("error: failed to read fileinfo of configfile, reason=%s", err.Error())
			break
		}

		if !fileInfo.ModTime().Equal(indexer.FileModTime) {
			log.Print("change in local config file detected")
			if err = readFromFile(indexer.Filename); err != nil {
				log.Printf("error: configfile update detected; failed to read in new config, reason=%s", err.Error())
			}
		}
	}
}

func readFromFile(filename string) error {

	reader, err := jsonreader.ReadFromFile(filename)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	indexer = &fileIndexer{
		Filename:    filename,
		Reader:      reader,
		FileModTime: fileInfo.ModTime(),
	}

	return nil
}

func (self fileIndexer) GetString(key string) (string, error) {

	return self.Reader.GetString(key)
}

func (self fileIndexer) GetInt(key string) (int, error) {

	return self.Reader.GetInt(key)
}

func (self fileIndexer) GetFloat64(key string) (float64, error) {

	return self.Reader.GetFloat64(key)
}

func (self fileIndexer) GetArray(key string) ([]interface{}, error) {

	return self.Reader.GetArray(key)
}

func (self fileIndexer) GetMap(key string) (map[string]interface{}, error) {

	return self.Reader.GetMap(key)
}
