package cores

import (
	"fmt"
	"log"
)

func GetAccessToken(corPID, secret string) (*AccessToken, error) {
	appTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corPID, secret)
	rsp := &AccessToken{}
	if err := InitClient("GET", appTokenUrl, nil).HttpDo(rsp); err != nil {
		log.Printf("call app token error %#v", err)
		return nil, err
	}
	return rsp, nil
}
