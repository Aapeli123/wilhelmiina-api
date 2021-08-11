package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type RouterSettings struct {
	HTTPS    bool
	KeyPath  string
	CertPath string
	Address  string
	Debug    bool
}

type DefaultResp struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func AbortWithErr(c *gin.Context, err error) {
	c.AbortWithStatusJSON(500, DefaultResp{
		Success: false,
		Error:   err.Error(),
	})
}

func StartRouter(options RouterSettings) {
	if !options.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/login", HandleLogin)  // Handle login
	auth.GET("/logout", HandleLogout) // Logout the current user

	users := api.Group("/user")
	users.Use(getUserFromSessionID)

	users.GET("/", GetUser)                   // Get current user
	users.GET("/id/:uuid", GetUserByID)       // Get user by id
	users.GET("/name/:username", GetUserByUn) // Get user by username
	users.GET("/list", GetUserList)           // Get list of users
	users.GET("/teachers", GetTeacherList)

	users.POST("/new", CreateNewUser) // Create new user

	users.PATCH("/update/:uuid", UpdateUser)  // Update some user
	users.PATCH("/update", UpdateCurrentUser) // Update current user

	users.DELETE("/:uuid", DeleteUser) // Delete user

	messages := api.Group("/messages")
	messages.Use(getUserFromSessionID)
	messages.GET("/", GetMessagesForCurrentUser)    // Get messages for user
	messages.GET("/message/:messageid", GetMessage) // Get message data
	messages.GET("/replies/:messageid", GetReplies) // Get replies for message

	messages.POST("/send", SendMessage)                // Send message
	messages.POST("/reply/:messageid", ReplyToMessage) // Reply to message

	messages.DELETE("/:messageid", DeleteMessage) // Delete message

	subjects := api.Group("/subjects")
	subjects.Use(getUserFromSessionID)
	subjects.GET("/", GetSubjects)          // Get all subjects
	subjects.GET("/:subjectid", GetSubject) // Get subject info

	subjects.POST("/new", NewSubject) // Create new subject

	subjects.PATCH("/:subjectid", EditSubj) // Edit some subject

	subjects.DELETE("/:subjectid", DeleteSubject) // Delete some subject

	courses := api.Group("/courses")
	courses.Use(getUserFromSessionID)
	courses.GET("/:courseid", GetCourse)                     // Get specific course
	courses.GET("/forsubject/:subjectid", GetCoursesForSubj) // Get courses for subject

	courses.POST("/new", NewCourse) // Create new course

	courses.PATCH("/:courseid", EditCourse) // Edit some course

	courses.DELETE("/:courseid", DeleteCourse) // Delete some course

	groups := api.Group("/groups")
	groups.Use(getUserFromSessionID)
	groups.GET("/:groupid", GetGroup)                      // Get some group
	groups.GET("/forcourse/:courseid", GetGroupsForCourse) // Get groups for some course
	groups.GET("/join/:groupid", JoinGroup)

	groups.POST("/new", NewGroup) // Create new group

	groups.PATCH("/add/:groupid", AddUserToGroup)            // Add some user to some group
	groups.PATCH("/:groupid", EditGroup)                     // Edit some group
	groups.PATCH("/:groupid/times", EditGroupTimeData)       // Edit some group times
	groups.PATCH("/setTeacher/:groupid", AssingGroupTeacher) // Set some user as the teacher of a group

	groups.DELETE("/:groupid", DeleteGroup) // Delete group
	if options.HTTPS {
		router.RunTLS(options.Address, options.CertPath, options.KeyPath)
		return
	}
	router.Run(options.Address)
}
