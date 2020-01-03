package conf

import (
	"os"
)

var Conf *Config

type Config struct {
	CorPID string // 企业ID

	AgentID string //应用ID
	Secret  string // 应用密钥
	Token   string //应用签名计算，自定义生成
	AesKey  string //应用消息加密，自定义生成
}

func init() {
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
}
