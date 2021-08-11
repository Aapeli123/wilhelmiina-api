package main

import (
	"time"

	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type loginParams struct {
	Username string
	Password string
}

func HandleLogin(c *gin.Context) {
	var lp loginParams
	c.BindJSON(&lp)
	user, err := wilhelmiina.GetUserByUn(lp.Username, db)
	if err != nil {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	res, err := user.CheckPassword(lp.Password)
	if err != nil {
		c.AbortWithStatusJSON(500, DefaultResp{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if !res {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   "Wrong password",
		})
		return
	}
	sessionID := uuid.New().String()
	rdb.Set(rdbCtx, sessionID, user.UUID, time.Minute*20)
	c.SetCookie("session", sessionID, 0, "/", "", false, true)
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func HandleLogout(c *gin.Context) {
	sid, err := c.Cookie("session")
	if err != nil {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	n, err := rdb.Del(rdbCtx, sid).Result()
	if err != nil {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	if n == 0 {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   "Session not found",
		})
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}
