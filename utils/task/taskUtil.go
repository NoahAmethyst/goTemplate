package task

import (
	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

var c *cron.Cron

func PutTask(corn string, f func()) error {
	if c == nil {
		c = cron.New()
	}
	spec := corn
	err := c.AddFunc(spec, f)
	if err != nil {
		log.Error().Fields(map[string]interface{}{
			"action": "put timing task",
			"error":  err.Error(),
		}).Send()
		return err
	}
	c.Start()
	select {}
}
