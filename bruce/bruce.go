package bruce

import (
	"flag"
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
var configFileName *string

func Init() {

	configFileName = flag.String("config", "./config.json", "configuration file")

	flag.Parse()

	if readFromFile(*configFileName) != nil {
		log.Fatalf("error: failed to read from config file")
		indexer = nil
	}

	go watcher()
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

//
// execute this as a goroutine to keep a watch of any changes to the current config file
// no need for a terminate signal for the goroutine -> cleanup later
//
func watcher() {

	ticker := time.NewTicker(30 * time.Second)

	for {

		<-ticker.C

		fileInfo, err := os.Stat(indexer.Filename)
		if err != nil {
			log.Printf("error: failed to read fileinfo of configfile, reason=%s", err.Error())
			break
		}

		if !fileInfo.ModTime().Equal(indexer.FileModTime) {
			log.Print("change in config file detected")
			if err = readFromFile(indexer.Filename); err != nil {
				log.Printf("error: configfile update detected; failed to read in new config, reason=%s", err.Error())
			}
		}
	}
}
