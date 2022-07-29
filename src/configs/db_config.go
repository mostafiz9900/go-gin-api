package configs

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:9900@tcp(localhost:3306)/gin_db")
	CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")
	return dbmap

}
func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}
