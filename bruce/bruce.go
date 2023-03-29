package bruce

import (
	"flag"
	"strings"

	"github.com/tukaianirban/sdk.go.common/bruce/local"
	"github.com/tukaianirban/sdk.go.common/log"
)

const (
	PROVIDER_LOCAL = "local"
	PROVIDER_GCP   = "gcp"
	PROVIDER_AWS   = "aws"
)

var instance ConfigStore

var configFile = flag.String("configfile", "./sample_config.json", "configuration file")
var provider = flag.String("configprovider", "local", "store to load configs from")

/**
Bruce Init() cannot be replaced by init() as flag parsing in init() is not recommended
Mode: determines if config is fetched from local file, gcp, aws, etc
args: optional param depending on the mode (ex: for local mode, the config file)
**/
func Init() {

	if instance != nil {
		log.Print("Bruce is already init'ed. Cannot re-init it")
		return
	}

	flag.Parse()

	var err error

	switch strings.ToLower(*provider) {
	case PROVIDER_LOCAL:
		instance, err = local.Init(*configFile)
		if err != nil {
			log.Fatal("Failed to load config from local file=%s, reason=%s", *configFile, err.Error())
			return
		}

	case PROVIDER_GCP:
		log.Fatal("Loading config from GCP is not implemented yet")
		return

	case PROVIDER_AWS:
		log.Fatal("Loading config from AWS is not implemented yet")

	default:
		log.Fatal("Invalid mode for config loading=%s", *provider)
	}
}
