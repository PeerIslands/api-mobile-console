package database

import (
	"context"
	"fmt"
	"log"
	"mongo-admin-backend/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is exported Mongo Database client
var Client *mongo.Client

type DBCrendentials struct {
	dbUserName string
	dbPassword string
	authDB     string
}

var DBCredentialsVal *DBCrendentials

func (ret *DBCrendentials) GetDBCredentials() (string, string, string) {
	return ret.authDB, ret.dbUserName, ret.dbPassword
}

// ConnectDatabase is used to connect the MongoDB database
func ConnectDatabase() {
	log.Println("Database connecting...")
	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn := options.Client().ApplyURI(config.ENVIRONMENT.DB_URI)
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	DBCredentialsVal = &DBCrendentials{
		dbUserName: conn.Auth.Username,
		dbPassword: conn.Auth.Password,
		authDB:     conn.Auth.AuthSource,
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database Connected.")
}
