package unauthenticated

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	if c.Query("sleep") == "" {
		c.Redirect(301, "/app/")
	} else {
		sleep, err := strconv.Atoi(c.Query("sleep"))

		if err != nil {
			return
		}
		if sleep > 0 && sleep < 11 {
			time.Sleep(time.Duration(sleep) * time.Second)
		}

	}
}
