package main

import (
	"Lab1/intermal/app/config"
	"Lab1/intermal/app/dsn"
	"Lab1/intermal/app/handler"
	"Lab1/intermal/app/repository"
	"Lab1/intermal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error loading config: %v", err)
	}

	postgresString := dsn.FromEnv()
	logrus.Info("Connecting to database...")

	rep, errRep := repository.New(
		postgresString,
		conf.MinIO.Endpoint,
		conf.MinIO.AccessKeyID,
		conf.MinIO.SecretAccessKey,
		conf.MinIO.BucketName,
		conf.MinIO.UseSSL,
	)
	if errRep != nil {
		logrus.Fatalf("error initializing repository: %v", errRep)
	}

	if err := rep.ResetSequences(); err != nil {
		logrus.Warnf("failed to reset sequences: %v", err)
	}

	hand := handler.NewHandler(rep)

	application := pkg.NewApp(conf, router, hand)
	application.RunApp()
}
