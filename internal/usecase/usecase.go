package usecase

import (
	"fmt"

	"cloud-test-task/internal/entities"
)

type PlaylistUseCase struct {
	repo     PlaylistRepo
	playlist PlaylistLibrary
}

func New(r PlaylistRepo, l PlaylistLibrary) *PlaylistUseCase {
	return &PlaylistUseCase{
		repo:     r,
		playlist: l,
	}
}

func (uc *PlaylistUseCase) Init() error {
	tracks, err := uc.repo.GetAllTracks()
	if err != nil {
		fmt.Println("Need add tracks to playlist")
		return err
	}
	err = uc.playlist.CreatePlaylist(tracks)
	if err != nil {
		return err
	}
	fmt.Println("Playlist created")
	return nil
}

func (uc *PlaylistUseCase) Start() error {
	err := uc.playlist.Start()
	return err
}

func (uc *PlaylistUseCase) Off() {
	uc.playlist.Off()
}

func (uc *PlaylistUseCase) Play() error {
	err := uc.playlist.Play()
	return err
}

func (uc *PlaylistUseCase) Pause() error {
	err := uc.playlist.Pause()
	return err
}

func (uc *PlaylistUseCase) Next() {
	uc.playlist.Next()
}

func (uc *PlaylistUseCase) Prev() {
	uc.playlist.Prev()
}

func (uc *PlaylistUseCase) AddTrack(name string, duration int64) (int64, error) {
	id, err := uc.repo.AddTrack(entities.NewTrack(name, int(duration)))
	if err != nil {
		return id, err
	}

	err = uc.playlist.AddTrack(entities.NewTrack(name, int(duration)))
	if err != nil {
		return id, err
	}
	return id, nil
}

func (uc *PlaylistUseCase) DeleteTrack(name string) error {
	err := uc.playlist.DeleteTrack(name)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteTrack(name)
	if err != nil {
		return err
	}

	return err
}

func (uc *PlaylistUseCase) UpdateTrack(name string, newName string, newDuration int64) error {
	track := entities.NewTrack(newName, int(newDuration))
	err := uc.playlist.UpdateTrack(name, track)
	if err != nil {
		return err
	}

	err = uc.repo.UpdateTrack(name, newName, newDuration)
	if err != nil {
		return err
	}

	return err
}

func (uc *PlaylistUseCase) GetTrack() entities.Track {
	return uc.playlist.GetTrack()
}
