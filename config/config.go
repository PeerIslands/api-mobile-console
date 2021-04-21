package config

import "os"

var (
	API_PORT = os.Getenv("PORT")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	DB_URI = os.Getenv("DB_URI")
	PORT =  os.Getenv("PORT")
	GROUPS_PATH = os.Getenv("GROUPS_PATH")
	BASE_PATH = os.Getenv("BASE_URL")
)