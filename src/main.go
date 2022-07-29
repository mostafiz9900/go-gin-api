package main

import (
	"fmt"

	"github.com/mostafiz9900/go-gin-api/src/routes"

	"github.com/gin-gonic/gin"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mostafiz9900/go-gin-api/src/configs"
)

func endpointHandler(c *gin.Context) {
	c.String(200, "%s %s", c.Request.Method, c.Request.URL.Path)
}
func main() {
	fmt.Printf("Alhamdullilah")
	db, err := sql.Open("mysql", "user:password@/gin_db")
if err != nil {
	panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
	initRouter()
	configs.Cors()
}
func initRouter() {
	
	router := gin.Default()
	v1 := router.Group("/v1")
	
	// product := v1.Group("product")
	// sales := v1.Group("sales")
	// product.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello Golang",
	// 	})
	// })
	// sales.GET("/", endpointHandler)
	// sales.GET("/:id", endpointHandler)
	routes.UserRouter(v1.Group("user"))
	routes.ProductRouter(router)
	routes.SalesRouter(router)  

	router.Run(":8080")
}

type User struct{
	Name string
	Age int

}