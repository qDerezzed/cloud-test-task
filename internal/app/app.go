package app

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"cloud-test-task/config"
	"cloud-test-task/internal/delivery"
	"cloud-test-task/internal/usecase"
	"cloud-test-task/internal/usecase/playlist"
	"cloud-test-task/internal/usecase/repository"
	"cloud-test-task/pkg/postgres"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(*cfg)
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
