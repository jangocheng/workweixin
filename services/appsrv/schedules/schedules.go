package schedules

import (
	"github.com/robfig/cron/v3"
)

func init() {
	var c = cron.New()
	_, _ = c.AddFunc("0 1 * * *", func() {
		AppMsgPush()
	})
	c.Start()
}
