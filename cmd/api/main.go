package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/NhanNT-VNG/go-auth/internal/config"
	"github.com/NhanNT-VNG/go-auth/internal/helpers"
	"github.com/go-playground/validator"
)

const portNumber = ":3000"

func main() {
	fmt.Println("Hello world")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("DB connected")

	validate := validator.New()
	validate.RegisterValidation("strong-password", helpers.StrongPassword)

	config := config.NewConfigurations(validate)
	srv := &http.Server{
		Handler:      routes(config),
		Addr:         portNumber,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("App listen on port", portNumber)

	log.Fatal(srv.ListenAndServe())

}
