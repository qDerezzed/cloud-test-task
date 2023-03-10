package main

import (
	"cloud-test-task/internal/entities"
	pl "cloud-test-task/internal/usecase/playlist"
	"time"
)

func main() {
	playlist := pl.NewPlaylist()

	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 12})
	playlist.AddTrack(&entities.Track{Name: "qwe", Duration: 15})
	playlist.AddTrack(&entities.Track{Name: "zxc", Duration: 4})

	playlist.Start()
	playlist.Play()
	time.Sleep(5 * time.Second)

	playlist.Pause()
	time.Sleep(5 * time.Second)

	playlist.Prev()
	time.Sleep(2 * time.Second)

	playlist.Next()
	time.Sleep(1 * time.Second)

	playlist.Prev()
	time.Sleep(1 * time.Second)

	playlist.Prev()
	time.Sleep(1 * time.Second)

	time.Sleep(15 * time.Second)

	playlist.Off()
}
