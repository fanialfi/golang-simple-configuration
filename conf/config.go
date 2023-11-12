package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

var shared *_Configuration

type _Configuration struct {
	Log struct {
		Verbose bool `json:"Verbose"`
	} `json:"log"`

	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `json:"server"`
}

func init() {
	if shared != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	bts, err := os.ReadFile(filepath.Join(basePath, "conf", "config.json"))
	if err != nil {
		panic(err.Error())
	}

	shared = new(_Configuration)
	err = json.Unmarshal(bts, &shared)
	if err != nil {
		panic(err.Error())
	}
}

func Configuration() _Configuration {
	return *shared
}
