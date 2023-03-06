package playlist

import (
	"cloud-test-task/internal/entities"
	"fmt"
	"time"
)

// ожидается запуск этого обработчика в горутине
func (pl *Playlist) playlistWorker() {
	for {
		select {
		case <-pl.offPlaylist:
			pl.stopPlayTrack()
			return
		case <-pl.playChan:
			pl.playHandler()
		case <-pl.pauseChan:
			pl.pauseHandler()
		case <-pl.nextChan:
			pl.nextHandler()
		case <-pl.prevChan:
			pl.prevHandler()
		}
	}
}

func (pl *Playlist) playHandler() {
	pl.wgPlayTrack.Add(1)
	go pl.playTrack()
}

func (pl *Playlist) pauseHandler() {
	pl.stopPlayTrack()
}

func (pl *Playlist) nextHandler() {
	defer pl.wgCommand.Done()
	pl.stopPlayTrack()
	pl.current.Value.(*entities.Track).CurrentTime = 0

	if pl.current.Next() == nil {
		fmt.Println("This is last track in playlist. Playlist starts from beginning")
		pl.current = pl.l.Front()
	} else {
		pl.current = pl.current.Next()
	}

	pl.wgPlayTrack.Add(1)
	go pl.playTrack()
}

func (pl *Playlist) prevHandler() {
	defer pl.wgCommand.Done()
	pl.stopPlayTrack()
	pl.current.Value.(*entities.Track).CurrentTime = 0

	if pl.current.Prev() == nil {
		fmt.Println("This is first track in playlist. Playlist starts from beginning")
		pl.current = pl.l.Front()
	} else {
		pl.current = pl.current.Prev()
	}

	pl.wgPlayTrack.Add(1)
	go pl.playTrack()
}

// завершение обработчика проигрывания трека playTrack
func (pl *Playlist) stopPlayTrack() {
	close(pl.endtPlayTrack)
	pl.wgPlayTrack.Wait()

	pl.endtPlayTrack = make(chan struct{})
}

// выводит в консоль текущее время трека
//
// ожидается запуск этого обработчика в горутине
func (pl *Playlist) playTrack() {
	defer pl.wgPlayTrack.Done()
	pl.current.Value.(*entities.Track).IsPlaying = true
	fmt.Printf("Now playing: %s\n", pl.current.Value.(*entities.Track).Name)
	pl.ticker.Reset(1 * time.Second)
	for {
		select {
		case <-pl.endtPlayTrack:
			pl.current.Value.(*entities.Track).IsPlaying = false
			return
		case <-pl.ticker.C:
			pl.current.Value.(*entities.Track).CurrentTime += 1
			fmt.Printf("Time: %d/%d seconds\n", pl.current.Value.(*entities.Track).CurrentTime, pl.current.Value.(*entities.Track).Duration)

			if pl.current.Value.(*entities.Track).CurrentTime >= pl.current.Value.(*entities.Track).Duration {
				go pl.Next()
			}
		}
	}
}
