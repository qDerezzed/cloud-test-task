package main

import (
	pl "cloud-test-task/internal/playlist"
	"time"
)

type MusicPlayer interface {
	Start() error // запускает работу плеера
	Off()         // выключает работу плеера

	Play() error              // начинает воспроизведение
	Pause() error             // приостанавливает воспроизведение
	AddTrack(*pl.Track) error // добавляет в конец плейлиста песню
	Next()                    // воспроизвести след песню
	Prev()                    // воспроизвести предыдущую песню
}

func main() {
	player := MusicPlayer(pl.NewPlaylist())

	player.AddTrack(pl.NewTrack("abc", 12*time.Second))
	player.AddTrack(pl.NewTrack("qwe", 15*time.Second))
	player.AddTrack(pl.NewTrack("zxc", 4*time.Second))

	player.Start()
	player.Play()
	time.Sleep(5 * time.Second)

	player.Pause()
	time.Sleep(5 * time.Second)

	player.Prev()
	time.Sleep(2 * time.Second)

	player.Next()
	time.Sleep(1 * time.Second)

	player.Prev()
	time.Sleep(1 * time.Second)

	player.Prev()
	time.Sleep(1 * time.Second)

	time.Sleep(15 * time.Second)

	player.Off()
}
