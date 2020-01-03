package conf

import (
	"os"
)

var Conf *Config

type Config struct {
	CorPID string // 企业ID

	Secret string // 密钥
	Token  string //签名计算，自定义生成
	AesKey string //消息加密，自定义生成
}

func init() {
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
}
