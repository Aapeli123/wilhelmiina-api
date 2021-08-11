package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

var MaxGroupMembers = 30

var CanCreateUser = wilhelmiina.Moderator
var CanUpdateUser = wilhelmiina.Moderator
var CanDeleteUser = wilhelmiina.Moderator
var CanListUsers = wilhelmiina.Moderator

var CanAddSubject = wilhelmiina.Moderator
var CanRemoveSubject = wilhelmiina.Moderator
var CanEditSubject = wilhelmiina.Moderator

var CanAddCourse = wilhelmiina.Moderator
var CanRemoveCourse = wilhelmiina.Moderator
var CanEditCourse = wilhelmiina.Moderator

var CanAddGroup = wilhelmiina.Moderator
var CanRemoveGroup = wilhelmiina.Moderator
var CanEditGroup = wilhelmiina.Moderator
var CanListGroupUsers = wilhelmiina.Teacher

var CanAddToGroup = wilhelmiina.Teacher
var CanJoinGroup = wilhelmiina.Student

func CheckPerms(c *gin.Context, minimumPerm wilhelmiina.Role) {
	id := c.GetString("uuid")
	u, err := wilhelmiina.GetUser(id, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	if u.Role < minimumPerm {
		c.AbortWithStatusJSON(403, DefaultResp{
			Success: false,
			Error:   "You don't have the permission to do that...",
		})
		return
	}
}

func CheckAccessToMessage(c *gin.Context, messageid string) {
	id := c.GetString("uuid")
	ul, err := wilhelmiina.GetRecieversForMessage(messageid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	for _, u := range ul {
		if u.UUID == id {
			return
		}
	}
	c.AbortWithStatusJSON(403, DefaultResp{
		Success: false,
		Error:   "You don't have access to this message",
	})
}
