package schedules

import (
	"github.com/robfig/cron/v3"
	"github.com/vnotes/workweixin_app/cores/notifications"
)

func RegisterCronJob() {
	var c = cron.New()
	_, _ = c.AddFunc("0 1 * * *", func() {
		notifications.AppMsgPush()
	})
	c.Start()
}
