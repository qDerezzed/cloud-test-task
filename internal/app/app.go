package app

import (
	"log"
	"net"
	"os"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"cloud-test-task/internal/delivery"
	"cloud-test-task/internal/usecase"
	"cloud-test-task/internal/usecase/playlist"
	"cloud-test-task/internal/usecase/repository"
	"cloud-test-task/pkg/postgres"
)

func Run() {
	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		log.Fatal("$DB_PASSWORD must be set")
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	pg, err := postgres.New(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: databasePassword,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer pg.Close()

	playlistUseCase := usecase.New(
		repository.NewRepository(pg),
		playlist.NewPlaylist(),
	)

	grpcSrv := grpc.NewServer()
	err = delivery.NewPlaylistServerGrpc(grpcSrv, *playlistUseCase)
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	log.Println("starting server at :8080")
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalln("cant listet port", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
