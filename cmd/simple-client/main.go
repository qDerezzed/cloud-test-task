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

	pl, err := playlistManager.GetPlaylist(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Playlist: ")
	for _, v := range pl.Playlist {
		fmt.Printf("track: %v\n", v)
	}

	_, err = playlistManager.Start(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}

	_, err = playlistManager.Play(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	track, err := playlistManager.GetTrack(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Now playing: ", track)
	time.Sleep(5 * time.Second)

	_, err = playlistManager.Pause(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(2 * time.Second)

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	track, err = playlistManager.GetTrack(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Now playing: ", track)
	time.Sleep(1 * time.Second)

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	track, err = playlistManager.GetTrack(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Now playing: ", track)
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
	_, err = playlistManager.UpdateTrack(context.Background(), &pb.UpdateTrackRequest{ID: 3, NewName: "qwe", NewDuration: 5})
	if err != nil {
		log.Println(err.Error())
	}
	pl, err = playlistManager.GetPlaylist(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Playlist: ")
	for _, v := range pl.Playlist {
		fmt.Printf("track: %v\n", v)
	}

	id, err := playlistManager.AddTrack(context.Background(), &pb.AddTrackRequest{Name: "abc", Duration: 10})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("add track with id = %d\n", id.ID)
	pl, err = playlistManager.GetPlaylist(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Playlist: ")
	for _, v := range pl.Playlist {
		fmt.Printf("track: %v\n", v)
	}

	_, err = playlistManager.Next(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}

	_, err = playlistManager.DeleteTrack(context.Background(), &pb.TrackID{ID: id.ID})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Track with id = %d is deleted\n", id.ID)

	pl, err = playlistManager.GetPlaylist(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Playlist: ")
	for _, v := range pl.Playlist {
		fmt.Printf("track: %v\n", v)
	}

	track, err = playlistManager.GetTrack(context.Background(), &pb.Nothing{Dummy: true})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Now playing: ", track)
	fmt.Println("Try delete this track")
	_, err = playlistManager.DeleteTrack(context.Background(), &pb.TrackID{ID: track.ID})
	if err != nil {
		log.Println(err.Error())
	}

	time.Sleep(30 * time.Second)
	playlistManager.Off(context.Background(), &pb.Nothing{Dummy: true})
}
