package config

import (
	"fmt"
	conf "github.com/ONSDigital/lfs-imports/config"
	"github.com/caarlos0/env/v6"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type configuration struct {
	Database conf.DatabaseConfiguration
}

var DatabaseConfiguration configuration

func init() {
	configFile, err := ioutil.ReadFile(fileName())

	if err != nil {
		log.Fatal(fmt.Errorf("cannot read configuration %+v", err))
	}

	DatabaseConfiguration = configuration{}

	err = toml.Unmarshal(configFile, &DatabaseConfiguration)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot unmarshall configuration file %+v", err))
	}

	// Parse environment variables
	if err := env.Parse(&DatabaseConfiguration); err != nil {
		log.Fatal(fmt.Errorf("cannot parse environment variables %+v", err))
	}

}

func fileName() string {
	runEnv := os.Getenv("CONFIG")

	if len(runEnv) == 0 {
		runEnv = "development"
	}

	filename := []string{"config.", runEnv, ".toml"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
