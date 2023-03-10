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
	id, err := uc.repo.AddTrack(&entities.Track{Name: name, Duration: int(duration)})
	if err != nil {
		return id, err
	}

	err = uc.playlist.AddTrack(entities.NewTrack(int(id), name, int(duration)))
	if err != nil {
		return id, err
	}
	return id, nil
}

func (uc *PlaylistUseCase) DeleteTrack(ID int64) error {
	err := uc.playlist.DeleteTrack(int(ID))
	if err != nil {
		return err
	}

	err = uc.repo.DeleteTrack(int(ID))
	if err != nil {
		return err
	}

	return err
}

func (uc *PlaylistUseCase) UpdateTrack(ID int64, newName string, newDuration int64) error {
	track := entities.NewTrack(int(ID), newName, int(newDuration))
	err := uc.playlist.UpdateTrack(int(ID), track)
	if err != nil {
		return err
	}

	err = uc.repo.UpdateTrack(ID, newName, newDuration)
	if err != nil {
		return err
	}

	return err
}

// получить трек, который сейчас играет
func (uc *PlaylistUseCase) GetTrack() entities.Track {
	return uc.playlist.GetTrack()
}

func (uc *PlaylistUseCase) GetPlaylist() ([]*entities.Track, error) {
	tracks, err := uc.repo.GetAllTracks()
	if err != nil {
		return nil, err
	}

	return tracks, err
}
