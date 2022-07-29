package routes

import (
	"github.com/mostafiz9900/go-gin-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/products", controllers.GetProducts)
	router.GET("/my-users", controllers.GetMyUser)
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:id", controllers.UserById())
	router.POST("/products", controllers.PostProduct)
}
func ProductRouter(router *gin.Engine) {
	p :=
		router.Group("v1")
	p.GET("product", controllers.GetProducts)
	p.POST("/insert", controllers.PostProduct)
}

func SalesRouter(router *gin.Engine) {
	p :=
		router.Group("v1")
	p.GET("sales", controllers.GetUsers())
}
