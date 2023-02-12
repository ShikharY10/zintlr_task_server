package main

import (
	"fmt"

	"github.com/ShikharY10/zintlr_internship_task_app_server/config"
	"github.com/ShikharY10/zintlr_internship_task_app_server/controllers"
	"github.com/ShikharY10/zintlr_internship_task_app_server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDataBase("mongodb+srv://db-mongodb-blr1-59698-480f7686.mongo.ondigitalocean.com", "doadmin", "3Uz59w1m02V76oyk")
	controller := controllers.Controller{
		DB: db,
	}
	router := gin.New()
	routes.Routes(router, controller)
	fmt.Println("[SETUP DONE!] [STARTING SERVER AT: 8000]")
	router.Run(":8000")
}
