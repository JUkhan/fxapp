package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"go.uber.org/fx"
)

// ApplicationConfig ...
type ApplicationConfig struct {
	Address string `yaml:"address"`
}
type dbConfig struct {
	URL string `yaml:"url"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
	DB                dbConfig `yaml:"dbConfig"`
}

func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	// handle error
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	// handle error
	if err != nil {
		panic(err)
	}

	return &conf
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
