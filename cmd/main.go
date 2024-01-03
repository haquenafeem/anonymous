package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/haquenafeem/anonymous/api/v1"
	"github.com/haquenafeem/anonymous/db"
	"github.com/haquenafeem/anonymous/internal"
	"github.com/haquenafeem/anonymous/repository"
	"github.com/haquenafeem/anonymous/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	panicOnErr(err)

	err = v1.New(
		getInjectionParamsMust(),
	).Run(getServerPortMust())

	panicOnErr(err)
}

func getInjectionParamsMust() (*gin.Engine, *service.Service) {
	db := db.DBMust(os.Getenv(internal.DB_PATH))
	repo := repository.NewMust(db)
	svc := service.New(repo)

	return gin.New(), svc
}

func getServerPortMust() int {
	portStr := os.Getenv(internal.SERVER_PORT)
	if portStr == "" {
		return 3002
	}

	port, err := strconv.Atoi(portStr)
	panicOnErr(err)

	return port
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
