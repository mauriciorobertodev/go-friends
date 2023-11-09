package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	ConnectionString = ""
	Port             = 0
)

func Load() {
	var err error

	if err = gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DBPort, err := strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 5000
	}

	Port = DBPort

	ConnectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
