package cores

var c *Config

type Config struct {
	CorPID string // 企业ID

	AgentID   string //应用ID
	CorSecret string // 应用密钥
	Token     string //应用签名计算，自定义生成
	AesKey    string //应用消息加密，自定义生成
}

func SetConfig(corPID, agentID, corSecret, token, aesKey string) {
	c = &Config{
		CorPID:    corPID,
		AgentID:   agentID,
		CorSecret: corSecret,
		Token:     token,
		AesKey:    aesKey,
	}
}

func GetConfig() *Config {
	return c
}
