package playlist

import "errors"

var (
	ErrorEmptyPlaylistStruct = errors.New("Playlist structure is not initialized")
	ErrorEmptyPlaylist       = errors.New("Playlist is empty")

	ErrorAlreadyPlay  = errors.New("Track is already playing")
	ErrorIsNotPlaying = errors.New("Track is not playing")

	ErrorNotValidTrackDuration = errors.New("Track duration is not valid")
)
