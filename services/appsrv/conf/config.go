package conf

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Conf *Config

type Config struct {
	SecretConf
	DBNetWork   string `yaml:"db_network"`
	ToDoNetWork string `yaml:"todo_network"`
}

type SecretConf struct {
	CorPID string // 企业ID

	AgentID string //应用ID
	Secret  string // 应用密钥
	Token   string //应用签名计算，自定义生成
	AesKey  string //应用消息加密，自定义生成
}

func InitConfig() {
	Conf = &Config{}

	corPID := os.Getenv("CorPID")

	agentID := os.Getenv("AgentID")
	secret := os.Getenv("AppSecret")
	token := os.Getenv("AppToken")
	aesKey := os.Getenv("AppAesKey")

	if corPID == "" || secret == "" || token == "" || aesKey == "" || agentID == "" {
		panic("service environment is not set")
	}

	Conf.AgentID = agentID
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
	if Conf.DBNetWork == "" || Conf.ToDoNetWork == "" {
		log.Fatal("config data is invalid")
	}
}