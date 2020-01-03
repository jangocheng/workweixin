package cores

import (
	"fmt"
	"log"
)

type AccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func InitAccessToken(corPID, secret string) (*AccessToken, error) {
	appTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corPID, secret)
	rsp := &AccessToken{}
	if err := InitClient("GET", appTokenUrl, nil).HttpDo(rsp); err != nil {
		log.Printf("call app token error %#v", err)
		return nil, err
	}
	return rsp, nil
}

func GetAccessToken(corPID, secret string) string {
	token, err := InitAccessToken(corPID, secret)
	if err != nil {
		return ""
	}
	return token.AccessToken
}
