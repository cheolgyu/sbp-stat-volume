package main

import (
	"fmt"
	"os"

	"log"

	"github.com/cheolgyu/stock-write-project-trading-volume/utils/local_log"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		panic("디비 유알엘 없다.")
	}
}

func main() {
	logPath := "logs/api/development.log"
	local_log.OpenLogFile(logPath)
	defer local_log.ElapsedTime("걸린시간", "start")()

	fmt.Println("hello world ")
	project_run()
}

func project_run() {

}
