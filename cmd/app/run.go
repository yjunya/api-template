package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/yjunya/api-example/project/config"
)

func Run() {
	config.InitConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}
