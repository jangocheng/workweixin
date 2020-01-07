package conf

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Conf *Config

type Config struct {
	DBNetWork string `yaml:"db_network"`
}

func InitConfig() {
	Conf = &Config{}

	var (
		productFile = "product.yaml"
		defaultFile *string
	)

	defaultFile = &productFile

	configFile := flag.String("c", "product.yaml", "config file")
	flag.Parse()

	if configFile == nil || *configFile == "" {
		configFile = defaultFile
	}

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("read config file error %#v", err)
	}
	if err := yaml.Unmarshal(data, Conf); err != nil {
		log.Fatalf("unmarshal config file error %#v", err)
	}
	if Conf.DBNetWork == "" {
		log.Fatal("config data is invalid")
	}
}
