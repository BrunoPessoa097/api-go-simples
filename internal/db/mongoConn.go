package db

import (
	"context"
	"fmt"
	"log"

	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func MongoDB() *mongo.Database {
	uri, err := utils.Dotenv("MONGO_DB")
	if err != nil {
		fmt.Println(err)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("MongoDB conectado com sucesso")

	return client.Database("golang")
}
