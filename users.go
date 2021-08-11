package main

import (
	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uuid := c.GetString("uuid")
	u, err := wilhelmiina.GetUser(uuid, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    u.ToData(),
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("uuid")
	u, err := wilhelmiina.GetUser(id, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    u.ToData(),
	})
}

func GetUserByUn(c *gin.Context) {
	name := c.Param("username")
	u, err := wilhelmiina.GetUserByUn(name, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    u.ToData(),
	})
}

func CreateNewUser(c *gin.Context) {
	type createUserReq struct {
		Username  string
		Password  string
		FirstName string
		Surname   string
		Role      wilhelmiina.Role
	}
	CheckPerms(c, CanCreateUser)
	if c.IsAborted() {
		return
	}
	var req createUserReq
	c.BindJSON(&req)
	_, err := wilhelmiina.CreateUser(req.Username, req.FirstName, req.Surname, req.Password, req.Role, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func DeleteUser(c *gin.Context) {
	CheckPerms(c, CanDeleteUser)
	if c.IsAborted() {
		return
	}
	id := c.Param("uuid")
	err := wilhelmiina.DeleteUser(id, db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

type updateReq struct {
	NewUsername  string
	NewFirstName string
	NewSurname   string
	NewPassword  string
}

func UpdateUser(c *gin.Context) {
	CheckPerms(c, CanUpdateUser)
	if c.IsAborted() {
		return
	}
	req := updateReq{}
	c.BindJSON(&req)
	id := c.Param("uuid")
	if req.NewPassword != "" {
		wilhelmiina.ChangePassword(req.NewPassword, id, db)
	}
	if req.NewUsername != "" {
		wilhelmiina.ChangeUsername(req.NewUsername, id, db)
	}
	if req.NewFirstName != "" {
		wilhelmiina.ChangeFirstName(req.NewFirstName, id, db)
	}
	if req.NewSurname != "" {
		wilhelmiina.ChangeLastName(req.NewSurname, id, db)
	}
}

func UpdateCurrentUser(c *gin.Context) {
	req := updateReq{}
	c.BindJSON(&req)
	id := c.GetString("uuid")
	var err error
	if req.NewPassword != "" {
		err = wilhelmiina.ChangePassword(req.NewPassword, id, db)
	}
	if req.NewUsername != "" {
		err = wilhelmiina.ChangeUsername(req.NewUsername, id, db)
	}
	if req.NewFirstName != "" {
		err = wilhelmiina.ChangeFirstName(req.NewFirstName, id, db)
	}
	if req.NewSurname != "" {
		err = wilhelmiina.ChangeLastName(req.NewSurname, id, db)
	}
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
	})
}

func GetUserList(c *gin.Context) {
	CheckPerms(c, CanListUsers)
	if c.IsAborted() {
		return
	}
	ul, err := wilhelmiina.GetUserList(db)
	if err != nil {
		AbortWithErr(c, err)
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    ul,
	})
}

func GetTeacherList(c *gin.Context) {
	list, err := wilhelmiina.GetTeacherList(db)
	if err != nil {
		AbortWithErr(c, err)
		return
	}
	c.JSON(200, DefaultResp{
		Success: true,
		Data:    list,
	})
}
