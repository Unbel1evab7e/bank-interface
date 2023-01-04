package main

import (
	"bank-interface/api"
	"bank-interface/api/middleware"
	_ "bank-interface/docs"
	"bank-interface/domain/properties"
	"bank-interface/integration/dadata"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	"log"
	"os"
	"strings"
)

var environment string
var container *dig.Container

func main() {
	setEnv()
	setConfig()

	container = dig.New()
	engine := gin.New()
	engine.Use(gin.Recovery())

	registerProperties(container)
	registerRepositories(container)
	registerServices(container)
	registerClients(container)
	registerMiddleware(container, engine)
	registerControllers(container, engine)
}

func setEnv() {

	for _, arg := range os.Args {
		if strings.Contains(arg, "environment") {
			keys := strings.Split(arg, ":")

			if len(keys) < 2 {
				environment = EnvironmentStage
			}

			switch keys[1] {
			case EnvironmentStage:
				environment = EnvironmentStage
				break
			case EnvironmentProduction:
				environment = EnvironmentProduction
				break
			default:
				environment = EnvironmentStage
			}
		}
	}
}
func setConfig() {
	viper.SetConfigName(`config`)
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func registerProperties(container *dig.Container) {
	err := container.Provide(func() *properties.DaDataProperties {
		return parseProperties[properties.DaDataProperties]("dadata")
	})

	if err != nil {
		log.Fatal("Failed to register properties")
	}

}
func parseProperties[T interface{}](path string) *T {
	raw, err := json.Marshal(viper.Get(path))

	if err != nil {
		log.Fatalf("Cannot get props from config %s with path %s", err.Error(), path)
	}

	var props T

	err = json.Unmarshal(raw, &props)

	if err != nil {
		log.Fatalf("Cannot get props from config %s with path %s", err.Error(), path)
	}

	return &props
}
func registerClients(container *dig.Container) {
	err := container.Provide(dadata.New)
	if err != nil {
		log.Fatal("Failed to register dadataClient", err)
	}
}
func registerRepositories(container *dig.Container) {

}
func registerServices(container *dig.Container) {

}

func registerMiddleware(container *dig.Container, engine *gin.Engine) {
	err := container.Provide(middleware.NewLoggerMiddleware)

	if err != nil {
		log.Fatal("Failed to provide logging middleware", err)
	}

	err = container.Invoke(func(loggingMiddleware *middleware.RequestLoggingMiddleware) {
		engine.Use(loggingMiddleware.Logger())
	})

	if err != nil {
		log.Fatal("Failed to invoke logging middleware", err)
	}
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a Test.

// @host      localhost:8080
// @BasePath  /api/v1

func registerControllers(container *dig.Container, engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := container.Provide(func() *gin.RouterGroup {
		return engine.Group("/api/v1")
	})

	if err != nil {
		log.Fatal("Failed to provide engine", err)
	}

	err = container.Provide(api.NewPersonController)

	if err != nil {
		log.Fatal("Failed to provide personController", err)
	}

	err = container.Invoke(func(personController *api.PersonController) {
		engine.Run(":8080")
	})
}
