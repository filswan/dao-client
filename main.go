package main

import (
	"dao-client/blockchain/browsersync"
	"dao-client/config"
	"dao-client/database"
	"dao-client/logs"
	"dao-client/models"
	"dao-client/scheduler"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	// init database
	db := database.Init()

	initMethod()
	browsersync.Init()

	models.RunAllTheScan()

	scheduler.DAOUnlockPaymentSchedule()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()
}

func initMethod() string {
	config.InitConfig("")
	return ""
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logs.GetLogger().Error(err)
	}
	fmt.Println("name: ", os.Getenv("privateKey"))
}
