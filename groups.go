package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

func GetGroup(c *gin.Context) {
	groupid := c.Param("groupid")
	g, err := wilhelmiina.GetGroup(groupid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    g,
	})
}

func GetGroupsForCourse(c *gin.Context) {
	cid := c.Param("courseid")
	groups, err := wilhelmiina.GetGroupsForCourse(cid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    groups,
	})
}

func JoinGroup(c *gin.Context) {
	gid := c.Param("groupid")
	uid := c.GetString("uuid")
	users, err := wilhelmiina.GetGroupUsers(gid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	if len(users) > MaxGroupMembers {
		c.JSON(400, DefaultResp{
			Success: false,
			Error:   "Group is full.",
		})
		return
	}
	_, err = wilhelmiina.CreateReservation(uid, gid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func NewGroup(c *gin.Context) {
	CheckPerms(c, CanAddGroup)
	if !c.IsAborted() {
		return
	}
	type groupCreateReq struct {
		GroupName string
		CourseID  string
		StartDate int64
		EndData   int64
		TimeData  []wilhelmiina.GroupTimeData
	}
	var req groupCreateReq
	c.BindJSON(&req)
	_, err := wilhelmiina.NewGroup(req.GroupName, req.CourseID, req.StartDate, req.EndData, req.TimeData, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func AddUserToGroup(c *gin.Context) {
	type addToGroupReq struct {
		UUID string
	}
	CheckPerms(c, CanAddToGroup)
	if c.IsAborted() {
		return
	}
	var req addToGroupReq
	gid := c.Param("groupid")
	c.BindJSON(&req)
	_, err := wilhelmiina.CreateReservation(req.UUID, gid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func EditGroup(c *gin.Context) {
	CheckPerms(c, CanEditGroup)
	if c.IsAborted() {
		return
	}

	type editGroupReq struct {
		Name string
	}
	var req editGroupReq
	c.BindJSON(&req)
	err := wilhelmiina.ChangeGroupName(req.Name, c.Param("groupid"), db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func EditGroupTimeData(c *gin.Context) {
	CheckPerms(c, CanEditGroup)
	if c.IsAborted() {
		return
	}

	type editGroupReq struct {
		TimeData []wilhelmiina.GroupTimeData
	}
	var req editGroupReq
	c.BindJSON(&req)
	err := wilhelmiina.UpdateGroupTimes(c.Param("groupid"), req.TimeData, db)

	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func AssingGroupTeacher(c *gin.Context) {
	CheckPerms(c, CanEditGroup)
	if c.IsAborted() {
		return
	}
	type editGroupReq struct {
		TeacherID string
	}
	var req editGroupReq
	c.BindJSON(&req)
	gd, err := wilhelmiina.GetGroup(c.Param("groupid"), db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	err = gd.GroupInfo.AssingTeacher(req.TeacherID, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func DeleteGroup(c *gin.Context) {
	CheckPerms(c, CanRemoveGroup)
	if c.IsAborted() {
		return
	}
	err := wilhelmiina.DeleteGroup(c.Param("groupid"), db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}
