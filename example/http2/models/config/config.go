package config

import (
	"os"

	"github.com/kataras/iris"

	"github.com/BurntSushi/toml"
)

func init() {
	cfg = &Configuration{
		DB:    newDBConf(),
		App:   newAppConf(),
		Redis: newRedisConf(),
	}
}

var cfg *Configuration

// Configuration .
type Configuration struct {
	DB    *DBConf
	App   *iris.Configuration
	Redis *RedisConf
}

// Get .
func Get() *Configuration {
	return cfg
}

func configure(obj interface{}, fileName string, def bool) {
	path := os.Getenv("FREEDOM_PROJECT_CONFIG")
	if path == "" {
		path = "./conf"
	}
	_, err := toml.DecodeFile(path+"/"+fileName, obj)
	if err != nil && !def {
		panic(err)
	}
}