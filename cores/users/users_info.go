package users

import (
	"fmt"
	"log"

	"github.com/vnotes/workweixin_app/cores"
)

type AppDetail struct {
	AgentID        int   `json:"agentid"`
	AllowUserInfos Users `json:"allow_userinfos"`
}

type Users struct {
	User []*UserInfo `json:"user"`
}

type UserInfo struct {
	UserID string `json:"userid"`
}

func InitAppUsers() {
	tokens, err := cores.GetAppAccessToken()
	if err != nil {
		return
	}
	agentID := cores.GetConfig().AgentID
	uri := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=%s&agentid=%s", tokens.AccessToken, agentID)
	rsp := &AppDetail{}
	if err := cores.InitClient("GET", uri, nil).HttpResult(rsp); err != nil {
		return
	}
	userIDs := make([]string, 0)
	for _, v := range rsp.AllowUserInfos.User {
		userIDs = append(userIDs, v.UserID)
	}
	log.Print(userIDs)
}
