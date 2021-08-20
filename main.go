package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Aapeli123/wilhelmiina-student-manager"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client
var rdbCtx = context.Background()

func main() {

	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DATABASE")

	apiPort := os.Getenv("API_PORT")
	environment := os.Getenv("ENVIRONMENT")

	var err error
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:" + redisPort,
		Password: redisPassword,
	})
	dbstr := fmt.Sprintf("host=postgres port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresPort, postgresUser, postgresPassword, postgresDatabase,
	)
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
	debug := environment == "test"
	StartRouter(RouterSettings{
		HTTPS:   false,
		Address: ":" + apiPort,
		Debug:   debug,
	})
}
