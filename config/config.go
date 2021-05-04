package config

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"

	"log"
)

type Environment struct {
	API_PORT      string `env:"PORT"`
	JWT_SECRET_KEY string `env:"JWT_SECRET_KEY"`
	DB_URI  string `env:"DB_URI"`
	PORT string `env:"PORT"`
	GROUPS_PATH     string `env:"GROUPS_PATH"`
	BASE_PATH   string `env:"BASE_URL"`
}

var ENVIRONMENT Environment


func init() {
	err := godotenv.Load("../../.env.example")

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading .env.example file")
	}

	_, err = env.UnmarshalFromEnviron(&ENVIRONMENT)
}