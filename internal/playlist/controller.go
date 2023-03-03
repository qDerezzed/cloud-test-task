package playlist

import (
	"fmt"
	"log"
)

// запускает работу плейлиста
func (pl *Playlist) Start() error {
	if pl.l == nil || pl.playChan == nil || pl.pauseChan == nil || pl.nextChan == nil || pl.prevChan == nil {
		log.Println(ErrorEmptyPlaylistStruct)
		return ErrorEmptyPlaylistStruct
	}

	if pl.l.Front() == nil {
		log.Println(ErrorEmptyPlaylist)
		return ErrorEmptyPlaylist
	}

	if pl.current == nil {
		pl.current = pl.l.Front()
	}

	go pl.playlistWorker()

	return nil
}

// добавляет трек в конец плейлиста
func (pl *Playlist) AddTrack(track *Track) error {
	if int(track.duration.Seconds()) < 1 {
		return ErrorNotValidTrackDuration
	}

	pl.l.PushBack(track)

	return nil
}

// начинает воспроизведение трека
func (pl *Playlist) Play() error {
	if pl.isPlaying {
		log.Println(ErrorAlreadyPlay)
		return ErrorAlreadyPlay
	}

	pl.playChan <- struct{}{}

	pl.isPlaying = true

	return nil
}

// ставит трек на паузу
func (pl *Playlist) Pause() error {
	if !pl.isPlaying {
		log.Println(ErrorIsNotPlaying)
		return ErrorIsNotPlaying
	}

	pl.pauseChan <- struct{}{}

	pl.isPlaying = false
	fmt.Println("Pause")

	return nil
}

// переключает на следующий трек
func (pl *Playlist) Next() {
	pl.wgCommand.Add(1)
	pl.nextChan <- struct{}{}
	pl.wgCommand.Wait()

	fmt.Println("Next")
}

// переключает на предыдущий трек
func (pl *Playlist) Prev() {
	pl.wgCommand.Add(1)
	pl.prevChan <- struct{}{}
	pl.wgCommand.Wait()

	fmt.Println("Prev")
}

// полностью останавливает работу плейлиста
func (pl *Playlist) Off() {
	pl.offPlaylist <- struct{}{}

	fmt.Println("Off")
}
