package repository

import (
	"cloud-test-task/internal/entities"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Track interface {
	AddTrack(track *entities.Track) (int64, error)
	GetTrack(trackID int64) (string, int64, error)
	UpdateTrack(trackID int64, name string, duration int64) error
	DeleteTrack(trackID int) error
	GetAllTracks() ([]*entities.Track, error)
}

type Repository struct {
	Track
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Track: NewATrackPostgres(dbPool),
	}
}
