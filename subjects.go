package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

func GetSubjects(c *gin.Context) {
	list, err := wilhelmiina.GetSubjects(db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    list,
	})
}

func GetSubject(c *gin.Context) {
	subjid := c.Param("subjectid")
	s, err := wilhelmiina.GetSubject(subjid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    s,
	})
}

func NewSubject(c *gin.Context) {
	CheckPerms(c, CanAddSubject)
	if c.IsAborted() {
		return
	}
	type newSubjReq struct {
		name      string
		shortname string
		desc      string
	}
	var req newSubjReq
	c.BindJSON(&req)

	_, err := wilhelmiina.CreateSubject(req.name, req.shortname, req.desc, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}

	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func EditSubj(c *gin.Context) {
	CheckPerms(c, CanEditSubject)
	if c.IsAborted() {
		return
	}
	type editSubjReq struct {
		name      string
		shortname string
		desc      string
	}
	var req editSubjReq
	subjid := c.Param("subjectid")

	if req.name != "" {
		err := wilhelmiina.ChangeSubjectName(req.name, subjid, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}
	if req.desc != "" {
		err := wilhelmiina.ChangeSubjectShortName(req.shortname, subjid, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}
	if req.shortname != "" {
		err := wilhelmiina.ChangeSubjectDesc(req.desc, subjid, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func DeleteSubject(c *gin.Context) {
	CheckPerms(c, CanRemoveSubject)
	if c.IsAborted() {
		return
	}
	subjid := c.Param("subjectid")
	err := wilhelmiina.DeleteSubject(subjid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}
