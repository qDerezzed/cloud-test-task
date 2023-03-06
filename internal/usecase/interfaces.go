package usecase

import (
	"cloud-test-task/internal/entities"
)

type (
	Playlist interface {
		Init() error // добавляет в плейлист треки из БД

		Start() error // запускает работу плеера
		Off()         // выключает работу плеера

		Play() error  // начинает воспроизведение
		Pause() error // приостанавливает воспроизведение
		Next()        // воспроизвести след песню
		Prev()        // воспроизвести предыдущую песню

		AddTrack(name string, duration int64) (int64, error)
		DeleteTrack(name string) error
		UpdateTrack(name string, newName string, newDuration int64) error
	}

	PlaylistRepo interface {
		AddTrack(*entities.Track) (int64, error)
		UpdateTrack(name string, newName string, newDuration int64) error
		DeleteTrack(name string) error
		GetAllTracks() ([]*entities.Track, error)
	}

	PlaylistLibrary interface {
		Start() error // запускает работу плеера
		Off()         // выключает работу плеера

		Play() error                    // начинает воспроизведение
		Pause() error                   // приостанавливает воспроизведение
		AddTrack(*entities.Track) error // добавляет в конец плейлиста песню
		Next()                          // воспроизвести след песню
		Prev()                          // воспроизвести предыдущую песню

		CreatePlaylist(tracks []*entities.Track) error
		DeleteTrack(name string) error
		UpdateTrack(name string, track *entities.Track) error
		GetTrack() entities.Track
	}
)
