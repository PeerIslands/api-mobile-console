package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	API_PORT          string `env:"PORT"`
	JWT_SECRET_KEY    string `env:"JWT_SECRET_KEY"`
	DB_URI            string `env:"DB_URI"`
	PORT              string `env:"PORT"`
	GROUPS_PATH       string `env:"GROUPS_PATH"`
	PROCESS_PATH      string `env:"PROCESS_PATH"`
	MEASUREMENTS_PATH string `env:"MEASUREMENTS_PATH"`
	BASE_PATH         string `env:"BASE_URL"`
}

const (
	PATH_PROCESS     = "processes"
	PATH_GROUP       = "groups"
	PATH_MEASUREMENT = "measurements"

	PARAM_GROUP       = "group_id"
	PARAM_PROCESS     = "process_id"
	PARAM_GRANULARITY = "granularity"
	PARAM_PERIOD      = "period"
	PARAM_ST_DATE     = "start"
	PARAM_END_DATE    = "end"
	PARAM_MEASUREMENT = "m"

	STR_REQ_STATUS_OPEN         = "OPEN"
	STR_REQ_STATUS_CLOSED       = "CLOSED"
	STR_REQ_STATUS_REJECTED     = "DELETED"
	STR_REQ_TYPE_NETWORK_ACCESS = "NETWORK-ACCESS"
	STR_REQ_TYPE_DB_ACCESS      = "USER-ACCESS"
)

var ENVIRONMENT Environment

func Start() {
	err := godotenv.Load("../../.env")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error loading .env.example file")
	}
	_, err = env.UnmarshalFromEnviron(&ENVIRONMENT)
}
