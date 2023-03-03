package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"cloud-test-task/internal/proto"
	"cloud-test-task/internal/repository"
	srv "cloud-test-task/internal/server"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	storage := repository.New(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: databasePassword,
	})

	err = storage.Open()
	if err != nil {
		log.Panic(err)
	}
	defer storage.Close()

	server := srv.NewServer(storage)
	server.Init()

	grpcSrv := grpc.NewServer()
	proto.RegisterPlaylistServer(grpcSrv, server)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	fmt.Println("starting server at :8080")
	grpcSrv.Serve(lis)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
