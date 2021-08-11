package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

func GetMessagesForCurrentUser(c *gin.Context) {
	id := c.GetString("uuid")
	msglist, err := wilhelmiina.GetMessagesForId(id, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    msglist,
	})
}

func GetReplies(c *gin.Context) {
	msgid := c.Param("messageid")
	CheckAccessToMessage(c, msgid)
	if c.IsAborted() {
		return
	}

	replies, err := wilhelmiina.GetReplies(msgid, db)
	if err != nil {
		AbortWithErr(c, err)
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    replies,
	})
}

func GetMessage(c *gin.Context) {
	msgid := c.Param("messageid")
	CheckAccessToMessage(c, msgid)
	if c.IsAborted() {
		return
	}
	msg, err := wilhelmiina.GetMessage(msgid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    msg,
	})
}

func SendMessage(c *gin.Context) {
	type sendReq struct {
		Title     string
		Contents  string
		Recievers []string
	}
	uid := c.GetString("uuid")
	var req sendReq
	c.BindJSON(&req)
	recievers := append(req.Recievers, uid)
	_, err := wilhelmiina.SendMessage(uid, recievers, req.Title, req.Contents, "", db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func ReplyToMessage(c *gin.Context) {
	type sendReq struct {
		Title     string
		Contents  string
		Recievers []string
	}
	uid := c.GetString("uuid")
	replyto := c.Param("messageid")
	var req sendReq
	c.BindJSON(&req)
	recievers := append(req.Recievers, uid)
	_, err := wilhelmiina.SendMessage(uid, recievers, req.Title, req.Contents, replyto, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func DeleteMessage(c *gin.Context) {
	msgid := c.Param("messageid")
	CheckAccessToMessage(c, msgid)
	if c.IsAborted() {
		return
	}

	err := wilhelmiina.DeleteMessage(msgid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
}
