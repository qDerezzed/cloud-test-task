package playlist

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"

	"cloud-test-task/internal/entities"
)

type Playlist struct {
	l       *list.List    // список треков
	current *list.Element // текущий трек

	isPlaying bool

	ticker *time.Ticker

	playChan  chan struct{}
	pauseChan chan struct{}
	nextChan  chan struct{}
	prevChan  chan struct{}

	offPlaylist chan struct{}

	wgCommand   *sync.WaitGroup
	wgPlayTrack *sync.WaitGroup

	endtPlayTrack chan struct{}
}

func NewPlaylist() *Playlist {
	return &Playlist{
		l:             list.New(),
		ticker:        time.NewTicker(1 * time.Second),
		playChan:      make(chan struct{}),
		pauseChan:     make(chan struct{}),
		nextChan:      make(chan struct{}),
		prevChan:      make(chan struct{}),
		offPlaylist:   make(chan struct{}),
		wgPlayTrack:   &sync.WaitGroup{},
		wgCommand:     &sync.WaitGroup{},
		endtPlayTrack: make(chan struct{}),
	}
}

func (pl *Playlist) GetTrack() entities.Track {
	return *pl.current.Value.(*entities.Track)
}

func (pl *Playlist) CreatePlaylist(Tracks []*entities.Track) error {
	for _, Track := range Tracks {
		err := pl.AddTrack(entities.NewTrack(Track.Name, Track.Duration))
		if err != nil {
			return err
		}
	}
	return nil
}

func (pl *Playlist) DeleteTrack(name string) error {
	var next *list.Element
	for e := pl.l.Front(); e != nil; e = next {
		if e.Value.(*entities.Track).Name == name {
			if e.Value.(*entities.Track).IsPlaying {
				return errors.New("this Track is playing")
			}
			pl.l.Remove(e)
			return nil
		}
		next = e.Next()
	}
	return errors.New("Track being deleted was not found")
}

func (pl *Playlist) UpdateTrack(name string, track *entities.Track) error {
	var next *list.Element
	for e := pl.l.Front(); e != nil; e = next {
		if e.Value.(*entities.Track).Name == name {
			if e.Value.(*entities.Track).IsPlaying {
				return errors.New("this Track is playing")
			}

			fmt.Println(track.Name)
			e.Value.(*entities.Track).Name = track.Name
			e.Value.(*entities.Track).Duration = track.Duration

			return nil
		}
		next = e.Next()
	}
	return errors.New("Track being updated was not found")
}
