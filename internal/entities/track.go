package entities

type Track struct {
	ID       int    `db:"track_id"`
	Name     string `db:"track_name"`
	Duration int    `db:"track_duration"`

	IsPlaying bool

	CurrentTime int
}

func NewTrack(ID int, name string, duration int) *Track {
	return &Track{
		ID:       ID,
		Name:     name,
		Duration: duration,
	}
}
