package playlist

import (
	"container/list"
	"errors"
	"sync"
	"time"

	"cloud-test-task/internal/entities"
)

type Track struct {
	Name     string
	duration time.Duration

	isPlaying bool

	currentTime time.Duration
	quit        chan struct{}
}

func NewTrack(name string, duration time.Duration) *Track {
	return &Track{
		Name:     name,
		duration: duration,
		quit:     make(chan struct{}),
	}
}

func (track Track) GetCurrentTime() time.Duration {
	return track.currentTime
}

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
}

func NewPlaylist() *Playlist {
	return &Playlist{
		l:           list.New(),
		ticker:      time.NewTicker(1 * time.Second),
		playChan:    make(chan struct{}),
		pauseChan:   make(chan struct{}),
		nextChan:    make(chan struct{}),
		prevChan:    make(chan struct{}),
		offPlaylist: make(chan struct{}),
		wgPlayTrack: &sync.WaitGroup{},
		wgCommand:   &sync.WaitGroup{},
	}
}

func (pl *Playlist) IsPlaying() bool {
	return pl.isPlaying
}

func (pl *Playlist) GetCurrentTrack() Track {
	return *pl.current.Value.(*Track)
}

func (pl *Playlist) GetPlaylistFromRep(tracks []*entities.Track) {
	for _, track := range tracks {
		pl.AddTrack(NewTrack(track.Name, track.Duration))
	}
}

func (pl *Playlist) DeleteTrack(name string) error {
	var next *list.Element
	for e := pl.l.Front(); e != nil; e = next {
		if e.Value.(*Track).Name == name {
			if e.Value.(*Track).isPlaying {
				return errors.New("this track is playing")
			}
			pl.l.Remove(e)
			return nil
		}
		next = e.Next()
	}
	return errors.New("track being deleted was not found")
}

func (pl *Playlist) UpdateTrack(name string, track *Track) error {
	var next *list.Element
	for e := pl.l.Front(); e != nil; e = next {
		if e.Value.(*Track).Name == name {
			if e.Value.(*Track).isPlaying {
				return errors.New("this track is playing")
			}

			e.Value.(*Track).Name = track.Name
			e.Value.(*Track).duration = track.duration

			return nil
		}
		next = e.Next()
	}
	return errors.New("track being updated was not found")
}
