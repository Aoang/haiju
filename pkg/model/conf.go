package model

import (
	"encoding/json"
	"github.com/Aoang/haiju/pkg/logger"
	"io/ioutil"
	"os"
)

var log = logger.NewLogger(os.Stdout)

// Conf of server
var Conf *Configuration

// Models represents all models..
var Models = []interface{}{
	&User{},&Location{},&Community{},&Property{},
	&Building{},&Room{},&File{},
}

// Configuration (config.json).
type Configuration struct {
	Server                string // server scheme, host and port
	RuntimeMode           string // dev
	LogLevel              string // logging level: trace/debug/info/warn/error/fatal
	ShowSQL               bool   // whether print sql in log
	Redis                 string // Redis connection URL
	RedisPassword         string // Redis Password
	MySQL                 string // MySQL connection URL
	StaticRoot            string // static resources file root path
}

// LoadConf loads the configurations
func LoadConf() {
	data, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(Conf.LogLevel)
}
