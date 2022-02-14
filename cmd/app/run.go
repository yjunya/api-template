package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"

	userHandler "github.com/yjunya/api-example/domains/user/handler"
	userRepository "github.com/yjunya/api-example/domains/user/repository"
	userUsecase "github.com/yjunya/api-example/domains/user/usecase"
	"github.com/yjunya/api-example/infra"
	"github.com/yjunya/api-example/project/config"
)

func Run() {
	router := gin.New()
	config.InitConfig()

	infra.InitMysql(
		config.Config.App.Env,
		config.Config.Mysql.DBTCPHost,
		config.Config.Mysql.Port,
		config.Config.Mysql.User,
		config.Config.Mysql.Password,
		config.Config.Mysql.Database,
		config.Config.Mysql.InstanceConnectionName,
	)

	userRepo := userRepository.New(infra.DB)
	userUC := userUsecase.New(userRepo)
	userH := userHandler.New(userUC)

	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	if config.Config.App.Env == config.EnvLocal {
		router.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
		boil.DebugMode = false
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users/:id", userH.Get)
	router.POST("/signup", userH.Signup)

	err := router.Run(":" + config.Config.App.Port)
	if err != nil {
		log.Println(err)
	}
}
