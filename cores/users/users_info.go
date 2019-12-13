package users

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
