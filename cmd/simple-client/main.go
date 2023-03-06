package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"cloud-test-task/internal/delivery/pb"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	playlistManager := pb.NewPlaylistClient(grcpConn)
	_, err = playlistManager.Start(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	_, err = playlistManager.Play(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(5 * time.Second)
	fmt.Println(playlistManager.GetTrack(context.Background(), &pb.Nothing{Dummy: true}))

	_, err = playlistManager.Pause(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(5 * time.Second)

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(1 * time.Second)

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(1 * time.Second)

	time.Sleep(20 * time.Second)
	_, err = playlistManager.UpdateTrack(context.Background(), &pb.UpdateTrackRequest{Name: "Queen - «Bohemian Rhapsody»", Track: &pb.Track{Name: "qwe", Duration: 5}})
	if err != nil {
		log.Println(err.Error())
	}

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}

	//time.Sleep(30 * time.Second)
	//playlistManager.Off(context.Background(), &pb.Nothing{Dummy: true})
}
