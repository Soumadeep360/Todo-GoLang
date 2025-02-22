package routes

import (
	"todo-app/controllers"
	"todo-app/repository"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	repo := &repository.TodoRepository{}
	service := services.NewTodoService(repo)
	controller := controllers.NewTodoController(service)

	todoRoutes := r.Group("/todos")
	{
		todoRoutes.POST("/", controller.CreateTodo)
		todoRoutes.GET("/", controller.GetTodos)
		todoRoutes.GET("/:id", controller.GetTodoByID)
		todoRoutes.PUT("/:id", controller.UpdateTodo)
		todoRoutes.DELETE("/:id", controller.DeleteTodo)
	}

	return r
}
