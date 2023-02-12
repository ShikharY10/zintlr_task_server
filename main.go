package main

import (
	"fmt"

	"github.com/ShikharY10/zintlr_internship_task_app_server/config"
	"github.com/ShikharY10/zintlr_internship_task_app_server/controllers"
	"github.com/ShikharY10/zintlr_internship_task_app_server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDataBase("127.0.0.1", "rootuser", "rootpass")
	controller := controllers.Controller{
		DB: db,
	}
	router := gin.New()
	routes.Routes(router, controller)
	fmt.Println("[SETUP DONE!] [STARTING SERVER AT: 8000]")
	router.Run(":8000")
}
