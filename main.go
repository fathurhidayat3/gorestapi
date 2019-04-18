package main

import (
	"github.com/fathurhidayat3/myrestapi/config"
	"github.com/fathurhidayat3/myrestapi/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/buku/:id", inDB.GetBuku)
	router.GET("/bukus", inDB.GetBukus)
	router.POST("/buku", inDB.CreateBuku)
	router.PUT("/buku/:id", inDB.UpdateBuku)
	router.DELETE("/buku/:id", inDB.DeleteBuku)
	router.Run(":3000")
}
