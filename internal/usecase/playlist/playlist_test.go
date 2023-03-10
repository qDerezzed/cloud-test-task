package playlist

import (
	"cloud-test-task/internal/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartPlaylist(t *testing.T) {
	pl := Playlist{}
	assert.Equal(t, entities.ErrorEmptyPlaylistStruct, pl.Start(), "empty playlist struct")

	playlist := NewPlaylist()
	assert.Equal(t, entities.ErrorEmptyPlaylist, playlist.Start(), "empty playlist")

	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 12})
	err := playlist.Start()
	playlist.Off()
	assert.Nil(t, err, "good off")
}

func TestPlay(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 2})
	playlist.Start()

	playlist.Play()
	assert.True(t, playlist.isPlaying, "track must playing")
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")

	time.Sleep(2 * time.Second)

	assert.True(t, playlist.GetTrack().CurrentTime <= playlist.GetTrack().Duration, "track must playing again")

	playlist.Off()
}

func TestPause(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 12})
	playlist.Start()

	playlist.Pause()
	assert.False(t, playlist.isPlaying, "track must not playing")
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")
	expectedCurrentTime := playlist.GetTrack().CurrentTime

	playlist.Play()
	assert.True(t, playlist.isPlaying, "track must playing")
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")
	assert.Equal(t, expectedCurrentTime, playlist.GetTrack().CurrentTime, "track played from the same moment")

	playlist.Off()
}

func TestNext(t *testing.T) {
	playlist := NewPlaylist()

	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 3})
	playlist.AddTrack(&entities.Track{Name: "qwe", Duration: 4})
	playlist.AddTrack(&entities.Track{Name: "zxc", Duration: 1})

	playlist.Start()

	playlist.Play()
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "qwe", playlist.GetTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "zxc", playlist.GetTrack().Name, "check current track name")

	time.Sleep(2 * time.Second)

	assert.Equal(t, "abc", playlist.GetTrack().Name, "play first track")
}

func TestPrev(t *testing.T) {
	playlist := NewPlaylist()

	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 3})
	playlist.AddTrack(&entities.Track{Name: "qwe", Duration: 4})
	playlist.AddTrack(&entities.Track{Name: "zxc", Duration: 3})

	playlist.Start()

	playlist.Play()
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "qwe", playlist.GetTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "zxc", playlist.GetTrack().Name, "check current track name")

	playlist.Prev()
	assert.Equal(t, "qwe", playlist.GetTrack().Name, "check current track name")
	playlist.Prev()
	assert.Equal(t, "abc", playlist.GetTrack().Name, "check current track name")
	playlist.Prev()
	assert.Equal(t, "abc", playlist.GetTrack().Name, "play again frist track")

	playlist.Off()

}

func TestErrorAdd(t *testing.T) {
	playlist := NewPlaylist()
	err := playlist.AddTrack(&entities.Track{Name: "abc", Duration: 0})
	assert.Equal(t, entities.ErrorNotValidTrackDuration, err, "track duration is not valid")
}

func TestErrorPlay(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 12})
	playlist.Start()
	err := playlist.Play()
	assert.Nil(t, err, "track is normal playing")
	err = playlist.Play()
	assert.Equal(t, entities.ErrorAlreadyPlay, err, "track is already playing")

	playlist.Off()
}

func TestErrorPause(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(&entities.Track{Name: "abc", Duration: 12})
	playlist.Start()
	playlist.Play()

	err := playlist.Pause()
	assert.Nil(t, err, "normal pause")

	err = playlist.Pause()
	assert.Equal(t, entities.ErrorIsNotPlaying, err, "track is already pause")

	playlist.Off()
}
