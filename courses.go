package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	id := c.Param("courseid")
	course, err := wilhelmiina.GetCourse(id, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    course,
	})
}

func GetCoursesForSubj(c *gin.Context) {
	subjid := c.Param("subjectid")
	courses, err := wilhelmiina.GetCoursesForSubject(subjid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    courses,
	})
}

func NewCourse(c *gin.Context) {
	CheckPerms(c, CanAddCourse)
	if c.IsAborted() {
		return
	}
	type newCourseReq struct {
		SubjID          string
		CourseName      string
		CourseDesc      string
		CourseShortName string
	}
	var req newCourseReq
	c.BindJSON(&req)

	_, err := wilhelmiina.NewCourse(req.CourseName, req.CourseShortName, req.CourseDesc, req.SubjID, db)

	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func EditCourse(c *gin.Context) {
	CheckPerms(c, CanEditCourse)
	if c.IsAborted() {
		return
	}

	type editCourseReq struct {
		CourseName      string
		CourseDesc      string
		CourseNameShort string
	}
	var req editCourseReq
	c.BindJSON(&req)

	cid := c.Param("courseid")
	course, err := wilhelmiina.GetCourse(cid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}

	if req.CourseName != "" {
		err := course.SetName(req.CourseName, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}
	if req.CourseNameShort != "" {
		err := course.SetShortName(req.CourseNameShort, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}
	if req.CourseDesc != "" {
		err := course.SetDescription(req.CourseDesc, db)
		if err != nil {
			AbortWithErr(c, err)
			return
		}
	}

	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func DeleteCourse(c *gin.Context) {
	CheckPerms(c, CanRemoveCourse)
	if c.IsAborted() {
		return
	}
	cid := c.Param("courseid")
	err := wilhelmiina.DeleteCourse(cid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}
