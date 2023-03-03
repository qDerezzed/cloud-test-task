package entities

import (
	"time"
)

type Track struct {
	Name     string
	Duration time.Duration
}

func NewTrack(name string, duration time.Duration) *Track {
	return &Track{
		Name:     name,
		Duration: duration,
	}
}

// type Playlist struct {
// 	l       *list.List    // список треков
// 	current *list.Element // текущий трек
// }

// func NewPlaylist() *Playlist {
// 	return &Playlist{
// 		l: list.New(),
// 	}
// }
