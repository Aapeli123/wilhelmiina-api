package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type sidRequired struct {
	Success bool
	Err     string
}

func getUserFromSessionID(c *gin.Context) {
	sid, err := c.Cookie("session")
	if err != nil {
		c.AbortWithStatusJSON(403, sidRequired{
			Success: false,
			Err:     "SessionID not found. You should probably login again...",
		})
		return
	}

	uuid, err := rdb.Get(rdbCtx, sid).Result()
	if err != nil {
		c.AbortWithStatusJSON(403, sidRequired{
			Success: false,
			Err:     "SessionID not found. You should probably login again...",
		})
		return
	}

	rdb.Expire(rdbCtx, sid, 20*time.Minute)

	c.Set("uuid", uuid)
}
