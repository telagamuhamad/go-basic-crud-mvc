package routes

import (
	"testing-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("views/*")

	// Routes untuk User
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/new", controllers.ShowUserForm)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.ShowUser)
	r.GET("/users/edit/:id", controllers.EditUser)
	r.POST("/users/update/:id", controllers.UpdateUser)
	r.GET("/users/delete/:id", controllers.DeleteUser)
}
