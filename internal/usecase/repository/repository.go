package repository

import (
	"cloud-test-task/internal/entities"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Track interface {
	AddTrack(track *entities.Track) (int64, error)
	UpdateTrack(name string, newName string, newDuration int64) error
	DeleteTrack(name string) error
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
