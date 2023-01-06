package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Unbel1evab7e/bank-interface/api"
	"github.com/Unbel1evab7e/bank-interface/api/middleware"
	"github.com/Unbel1evab7e/bank-interface/db/repository"
	_ "github.com/Unbel1evab7e/bank-interface/docs"
	"github.com/Unbel1evab7e/bank-interface/domain"
	"github.com/Unbel1evab7e/bank-interface/domain/properties"
	"github.com/Unbel1evab7e/bank-interface/integration/dadata"
	"github.com/Unbel1evab7e/bank-interface/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
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
	registerDbConnect(container)
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
				environment = domain.EnvironmentStage
			}

			switch keys[1] {
			case domain.EnvironmentStage:
				environment = domain.EnvironmentStage
				break
			case domain.EnvironmentProduction:
				environment = domain.EnvironmentProduction
				break
			default:
				environment = domain.EnvironmentStage
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
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
}
func registerProperties(container *dig.Container) {
	err := container.Provide(func() *properties.DaDataProperties {
		return parseProperties[properties.DaDataProperties]("dadata")
	})

	if err != nil {
		logrus.Fatal("Failed to register dadata properties ", err)
	}

	err = container.Provide(func() *properties.DBProperties {
		return parseProperties[properties.DBProperties]("database")
	})

	if err != nil {
		logrus.Fatal("Failed to register database properties ", err)
	}

}

func registerDbConnect(container *dig.Container) {
	err := container.Invoke(func(dbProperties *properties.DBProperties) {
		connection := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
			dbProperties.Name,
			dbProperties.User,
			dbProperties.Pass,
			dbProperties.Host,
			dbProperties.Port)

		dsn := fmt.Sprintf("%s", connection)
		dbConn, err := sql.Open(`postgres`, dsn)

		if err != nil {
			logrus.Fatal("Failed to connection database ", err)
		}

		err = dbConn.Ping()

		if err != nil {
			logrus.Fatal("Fail to ping database ", err)
		}

		err = container.Provide(func() *sql.DB {
			return dbConn
		})
		if err != nil {
			logrus.Fatal("Fail to provide connection ", err)
		}
	})

	if err != nil {
		logrus.Fatal("Fail to invoke properties ", err)
	}
}
func parseProperties[T interface{}](path string) *T {
	raw, err := json.Marshal(viper.Get(path))

	if err != nil {
		logrus.Fatalf("Cannot get props from config %s with path %s ", err.Error(), path)
	}

	var props T

	err = json.Unmarshal(raw, &props)

	if err != nil {
		logrus.Fatalf("Cannot get props from config %s with path %s ", err.Error(), path)
	}

	return &props
}
func registerClients(container *dig.Container) {
	err := container.Provide(dadata.New)
	if err != nil {
		logrus.Fatal("Failed to register dadataClient ", err)
	}
}
func registerRepositories(container *dig.Container) {
	err := container.Provide(repository.NewPersonRepository)
	if err != nil {
		logrus.Fatal("Fail to register personRepository ", err)
	}
}
func registerServices(container *dig.Container) {
	err := container.Provide(service.NewPersonService)
	if err != nil {
		logrus.Fatal("Fail to register personService ", err)
	}
}

func registerMiddleware(container *dig.Container, engine *gin.Engine) {
	err := container.Provide(middleware.NewLoggerMiddleware)

	if err != nil {
		logrus.Fatal("Failed to provide logging middleware ", err)
	}

	err = container.Invoke(func(loggingMiddleware *middleware.RequestLoggingMiddleware) {
		engine.Use(loggingMiddleware.Logger())
	})

	if err != nil {
		logrus.Fatal("Failed to invoke logging middleware ", err)
	}

	err = container.Provide(middleware.NewSecurityMiddleware)
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a Test.

//	@host		localhost:8080
//	@BasePath	/api/v1

func registerControllers(container *dig.Container, engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := container.Provide(func() *gin.RouterGroup {
		return engine.Group("/api/v1")
	})

	if err != nil {
		logrus.Fatal("Failed to provide engine ", err)
	}

	err = container.Provide(api.NewPersonController)

	if err != nil {
		logrus.Fatal("Failed to provide personController ", err)
	}

	err = container.Invoke(func(personController *api.PersonController) {
		engine.Run(":8080")
	})
}
