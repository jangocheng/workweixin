package cores

import (
	"fmt"
	"log"
)

type AppAccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAppAccessToken() (*AppAccessToken, error) {
	corpID := GetConfig().CorPID
	corpSecret := GetConfig().CorSecret
	appTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpID, corpSecret)
	rsp := &AppAccessToken{}
	if err := InitClient("GET", appTokenUrl, nil).HttpResult(rsp); err != nil {
		log.Printf("call app token error %#v", err)
		return nil, err
	}
	return rsp, nil
}
