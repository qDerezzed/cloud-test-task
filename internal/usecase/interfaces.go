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
		DeleteTrack(ID int64) error
		UpdateTrack(ID int64, newName string, newDuration int64) error
		GetTrack() entities.Track

		GetPlaylist() ([]*entities.Track, error)
	}

	PlaylistRepo interface {
		AddTrack(track *entities.Track) (int64, error)
		GetTrack(trackID int64) (string, int64, error)
		UpdateTrack(trackID int64, name string, duration int64) error
		DeleteTrack(trackID int) error
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
		DeleteTrack(ID int) error
		UpdateTrack(ID int, track *entities.Track) error
		GetTrack() entities.Track // получитьтрек, который сейчас играет
	}
)
