package repository

import (
	"context"

	"cloud-test-task/internal/entities"

	"github.com/georgysavva/scany/pgxscan"
)

func (db *PostgresDB) AddTrack(track *entities.Track) (int64, error) {
	var trackID int64
	err := db.dbPool.QueryRow(context.Background(),
		`INSERT INTO playlist (track_name, track_duration)
		VALUES 
		($1, $2)
		RETURNING track_id;`,
		track.Name, track.Duration).Scan(&trackID)
	return trackID, err
}

func (db *PostgresDB) GetTrack(trackID int64) (string, int64, error) {
	var name string
	var duration int64
	err := db.dbPool.QueryRow(
		context.Background(),
		"SELECT track_name, track_duration FROM playlist WHERE track_id = $1;",
		trackID).Scan(&name)

	return name, duration, err
}

// func (db *PostgresDB) UpdateTrack(name string, duration int64, trackID int64) error {
// 	_, err := db.dbPool.Exec(context.Background(),
// 		`UPDATE playlist SET track_name = $1 AND track_duration = $2 WHERE track_id = $3;`,
// 		name, duration, trackID)
// 	return err
// }

// func (db *PostgresDB) DeleteTrack(trackID int) error {
// 	_, err := db.dbPool.Exec(context.Background(),
// 		`DELETE FROM playlist WHERE track_id = $1;`, trackID)
// 	return err
// }

func (db *PostgresDB) UpdateTrack(newName string, duration int64, name string) error {
	_, err := db.dbPool.Exec(context.Background(),
		`UPDATE playlist SET track_name = $1 AND track_duration = $2 WHERE track_name = $3;`,
		newName, duration, name)
	return err
}

func (db *PostgresDB) DeleteTrack(name string) error {
	_, err := db.dbPool.Exec(context.Background(),
		`DELETE FROM playlist WHERE track_name = $1;`, name)
	return err
}

func (db *PostgresDB) GetAllTracks() ([]*entities.Track, error) {
	var tracks []*entities.Track
	err := pgxscan.Select(context.Background(), db.dbPool, &tracks, "SELECT * FROM playlist")

	return tracks, err
}
