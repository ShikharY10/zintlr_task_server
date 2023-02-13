package routes

import (
	"github.com/ShikharY10/zintlr_internship_task_app_server/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, controller controllers.Controller) {
	user := router.Group("/api/v1")
	user.GET("/test", controller.TestRoute)
	user.POST("/register", controller.RegisterUser)
	user.POST("/addpost", controller.AddPost)
	user.GET("/getrandompost", controller.GetRandomPost)
	user.GET("/delete", controller.DeleteAccount)
}
