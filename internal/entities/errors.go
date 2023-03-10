package entities

import "errors"

var (
	ErrorEmptyPlaylistStruct = errors.New("Playlist structure is not initialized")
	ErrorEmptyPlaylist       = errors.New("Playlist is empty")

	ErrorAlreadyPlay  = errors.New("Track is already playing")
	ErrorIsNotPlaying = errors.New("Track is not playing")

	ErrorNotValidTrackDuration = errors.New("Track duration is not valid")

	ErrorDelTrackNotFound = errors.New("track being deleted was not found")
	ErrorUpdTrackNotFound = errors.New("track being updated was not found")
)
