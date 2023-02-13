package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShikharY10/zintlr_internship_task_app_server/config"
	"github.com/ShikharY10/zintlr_internship_task_app_server/controllers"
	"github.com/ShikharY10/zintlr_internship_task_app_server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// LOADING ENVIRONMENT VARIABLES
	godotenv.Load()

	mongodbConnectionString, found := os.LookupEnv("MONGO_CONN_STRING")
	if !found {
		log.Fatal("environment variable, MONGO_CONN_STRING not found.")
	}

	db := config.ConnectToDataBase(mongodbConnectionString)
	controller := controllers.Controller{
		DB: db,
	}
	router := gin.New()
	routes.Routes(router, controller)
	fmt.Println("[SETUP DONE!] [STARTING SERVER AT: 8000]")
	router.Run(":8000")
}
