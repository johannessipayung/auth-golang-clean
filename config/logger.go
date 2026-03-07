package config

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {

	Logger = log.New(
		os.Stdout,
		"[AUTH SERVICE] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
