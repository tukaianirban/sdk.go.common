package local

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tukaianirban/sdk.go.common/bruce/readers"
	"gopkg.in/yaml.v3"
)

type FileReader struct {
	Filename    string
	Reader      *readers.Reader
	FileModTime time.Time
}

var indexerInstance *FileReader

/**
create a new indexer by loading config from the local file
**/
func Init(configFile string, refresh bool) (*FileReader, error) {

	reader, err := ReadFromFile(configFile)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}

	indexerInstance = &FileReader{
		Filename:    configFile,
		Reader:      reader,
		FileModTime: fileInfo.ModTime(),
	}

	if refresh {
		go indexerInstance.watcher()
	}

	return indexerInstance, nil
}

func ReadFromFile(filename string) (*readers.Reader, error) {

	parts := strings.Split(filename, ".")
	suffix := parts[len(parts)-1]

	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	switch suffix {
	case "json":
		err = json.Unmarshal(filebytes, &data)
		if err != nil {
			return nil, err
		}
		return readers.NewReaderFromMap(data), nil

	case "yaml", "yml":
		err = yaml.Unmarshal(filebytes, &data)
		if err != nil {
			return nil, err
		}
		return readers.NewReaderFromMap(data), nil

	default:
		// unrecognised file type
		return nil, errors.New("unrecognised config file type")
	}

}

//
// execute this as a goroutine to keep a watch of any changes to the current config file
// no need for a terminate signal for the goroutine -> cleanup later
//
func (indexer *FileReader) watcher() {

	ticker := time.NewTicker(5 * time.Minute)

	for {

		<-ticker.C

		fileInfo, err := os.Stat(indexer.Filename)
		if err != nil {
			log.Printf("error: failed to read fileinfo of configfile, reason=%s", err.Error())
			break
		}

		if fileInfo.ModTime().Equal(indexer.FileModTime) {
			continue
		}

		log.Print("change in local config file detected")

		//
		// if loading new changed configfile causes error,
		// ignore it and continue using the old object as config object
		// as previously
		newReader, err := ReadFromFile(indexer.Filename)
		if err != nil {
			log.Printf("Error: failed to reload config from file, reason=%s", err.Error())
			continue
		}

		indexer.Reader = newReader
		indexer.FileModTime = fileInfo.ModTime()
	}
}
