package playlist

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartPlaylist(t *testing.T) {
	pl := Playlist{}
	assert.Equal(t, ErrorEmptyPlaylistStruct, pl.Start(), "empty playlist struct")

	playlist := NewPlaylist()
	assert.Equal(t, ErrorEmptyPlaylist, playlist.Start(), "empty playlist")

	playlist.AddTrack(NewTrack("abc", 12*time.Second))
	err := playlist.Start()
	playlist.Off()
	assert.Nil(t, err, "good off")
}

func TestPlay(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(NewTrack("abc", 2*time.Second))
	playlist.Start()

	playlist.Play()
	assert.True(t, playlist.IsPlaying(), "track must playing")
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")

	time.Sleep(2 * time.Second)

	assert.True(t, playlist.GetCurrentTrack().GetCurrentTime() <= playlist.GetCurrentTrack().duration, "track must playing again")

	playlist.Off()
}

func TestPause(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(NewTrack("abc", 12*time.Second))
	playlist.Start()

	playlist.Pause()
	assert.False(t, playlist.IsPlaying(), "track must not playing")
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")
	expectedCurrentTime := playlist.GetCurrentTrack().currentTime

	playlist.Play()
	assert.True(t, playlist.IsPlaying(), "track must playing")
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")
	assert.Equal(t, expectedCurrentTime, playlist.GetCurrentTrack().GetCurrentTime(), "track played from the same moment")

	playlist.Off()
}

func TestNext(t *testing.T) {
	playlist := NewPlaylist()

	playlist.AddTrack(NewTrack("abc", 3*time.Second))
	playlist.AddTrack(NewTrack("qwe", 4*time.Second))
	playlist.AddTrack(NewTrack("zxc", 1*time.Second))

	playlist.Start()

	playlist.Play()
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "qwe", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "zxc", playlist.GetCurrentTrack().Name, "check current track name")

	time.Sleep(2 * time.Second)

	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "play first track")
}

func TestPrev(t *testing.T) {
	playlist := NewPlaylist()

	playlist.AddTrack(NewTrack("abc", 3*time.Second))
	playlist.AddTrack(NewTrack("qwe", 4*time.Second))
	playlist.AddTrack(NewTrack("zxc", 3*time.Second))

	playlist.Start()

	playlist.Play()
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "qwe", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Next()
	assert.Equal(t, "zxc", playlist.GetCurrentTrack().Name, "check current track name")

	playlist.Prev()
	assert.Equal(t, "qwe", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Prev()
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "check current track name")
	playlist.Prev()
	assert.Equal(t, "abc", playlist.GetCurrentTrack().Name, "play again frist track")

	playlist.Off()

}

func TestErrorAdd(t *testing.T) {
	playlist := NewPlaylist()
	err := playlist.AddTrack(NewTrack("abc", 0*time.Second))
	assert.Equal(t, ErrorNotValidTrackDuration, err, "track duration is not valid")
}

func TestErrorPlay(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(NewTrack("abc", 12*time.Second))
	playlist.Start()
	err := playlist.Play()
	assert.Nil(t, err, "track is normal playing")
	err = playlist.Play()
	assert.Equal(t, ErrorAlreadyPlay, err, "track is already playing")

	playlist.Off()
}

func TestErrorPause(t *testing.T) {
	playlist := NewPlaylist()
	playlist.AddTrack(NewTrack("abc", 12*time.Second))
	playlist.Start()
	playlist.Play()

	err := playlist.Pause()
	assert.Nil(t, err, "normal pause")

	err = playlist.Pause()
	assert.Equal(t, ErrorIsNotPlaying, err, "track is already pause")

	playlist.Off()
}
