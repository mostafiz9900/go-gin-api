package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mostafiz9900/go-gin-api/src/configs"
	"github.com/mostafiz9900/go-gin-api/src/models"

	Entity "github.com/mostafiz9900/go-gin-api/src/models"
)

func GetProducts(c *gin.Context) {
	var product []Entity.ProductModel
	_, err := configs.Dbmap.Select(&product, "select * from product")

	if err == nil {
		c.JSON(200, gin.H{"data": product})
	} else {
		c.JSON(400, gin.H{"error": "productdfdf no found"})
	}
}
func GetMyUser(c *gin.Context) {
	var user []models.UserModel
	_, err := configs.Dbmap.Select(&user, "select * from myuser")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{"error": "users no found"})
	}
}
func GetUsers() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name": "Mostafiz",
			"dept": "Marketin",
		})
	}
}

func UserById() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		myMap := make(map[string]interface {
		})
		myMap["id"] = id
		myMap["name"] = "Mostafizur"
		myMap["age"] = 29

		ctx.JSON(200, gin.H{
			"data": myMap,
		})
	}
}
func PostProduct(c *gin.Context) {
	var product models.ProductModel
	c.Bind(&product)
	log.Println(product)
	if product.Name != "" {

		if insert, _ := configs.Dbmap.Exec(`INSERT INTO product (id,name, price, qty) VALUES (?, ?,?,?)`, product.ID, product.Name, product.Price, product.Qty); insert != nil {
			product_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.ProductModel{
					ID:    product_id,
					Name:  product.Name,
					Price: product.Price,
					Qty:   product.Qty,
				}
				myMap := make(map[string]interface {
				})
				myMap["data"] = content
				c.JSON(201, myMap)

			} else {
				configs.CheckErr(err, "insert failed")
			}
		}
	}
}

func testCode() {
	mapData := map[string]string{}
	fmt.Println(mapData)
}
