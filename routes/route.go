package routes

import (
	"github.com/gin-gonic/gin"
	"library_management/controllers"
	"library_management/global"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	bookController := controllers.NewBookController(global.PDB)

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", bookController.CreateBook)
		bookRoutes.GET("/", bookController.GetBooks)
		bookRoutes.GET("/:id", bookController.GetBook)
		bookRoutes.PUT("/:id", bookController.UpdateBook)
		bookRoutes.DELETE("/:id", bookController.DeleteBook)
	}

	return r
}
