package main

import (
	"context"
	"fmt"

	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client
var rdbCtx = context.Background()

func main() {
	var err error
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	dbstr := "host=postgres port=5432 user=admin password=secret dbname=test sslmode=disable"
	db, err = gorm.Open(postgres.Open(dbstr))
	if err != nil {
		fmt.Println(dbstr)
		panic(err)
	}
	fmt.Println("Connected to database")
	err = wilhelmiina.CreateTables(db)
	if err != nil {
		panic(err)
	}

	ul, err := wilhelmiina.GetUserList(db)
	if err != nil {
		panic(err)
	}

	if len(ul) == 0 {
		fmt.Println("There are no users in database")
		fmt.Println("Creating default admin user with credentials:")
		fmt.Println("Username: admin")
		fmt.Println("Password: admin")
		wilhelmiina.CreateUser("admin", "Admin", "Admin", "admin", wilhelmiina.Admin, db)
	}
	StartRouter(RouterSettings{
		HTTPS:   false,
		Address: ":8080",
		Debug:   true,
	})
}
