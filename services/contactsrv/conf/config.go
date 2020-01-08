package conf

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var Conf *Config

type Config struct {
	DBNetWork string `yaml:"db_network"`
	SecretConf
}

type SecretConf struct {
	CorPID string // 企业ID

	Secret string // 密钥
	Token  string //签名计算，自定义生成
	AesKey string //消息加密，自定义生成
}

func InitConfig() {
	Conf = &Config{}

	corPID := os.Getenv("CorPID")

	secret := os.Getenv("ContactSecret")
	token := os.Getenv("ContactToken")
	aesKey := os.Getenv("ContactAesKey")

	if corPID == "" || secret == "" || token == "" || aesKey == "" {
		panic("service environment is not set")
	}

	Conf.CorPID = corPID
	Conf.Secret = secret
	Conf.Token = token
	Conf.AesKey = aesKey

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
