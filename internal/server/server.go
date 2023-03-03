package server

import (
	"context"
	"fmt"
	"time"

	"cloud-test-task/internal/entities"
	pl "cloud-test-task/internal/playlist"
	"cloud-test-task/internal/proto"
	"cloud-test-task/internal/repository"
)

type Server struct {
	proto.PlaylistServer

	pl *pl.Playlist
	db *repository.PostgresDB
}

func NewServer(db *repository.PostgresDB) *Server {
	return &Server{pl: pl.NewPlaylist(), db: db}
}

func (s *Server) Init() {
	tracks, err := s.db.GetAllTracks()
	if err != nil {
		s.pl.GetPlaylistFromRep(tracks)
		fmt.Println("Playlist created")
	} else {
		fmt.Println("Need add tracks to playlist")
	}
}

func (s *Server) Start(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	s.pl.Start()
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Off(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	s.pl.Off()
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Play(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	err := s.pl.Play()
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Pause(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	err := s.pl.Pause()
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Next(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	s.pl.Next()
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Prev(ctx context.Context, n *proto.Nothing) (*proto.Nothing, error) {
	s.pl.Prev()
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) AddTrack(ctx context.Context, req *proto.Track) (*proto.TrackID, error) {
	id, err := s.db.AddTrack(entities.NewTrack(req.Name, time.Duration(req.Duration)*time.Second))
	trackID := &proto.TrackID{
		ID: id,
	}

	if err != nil {
		return nil, err
	}

	err = s.pl.AddTrack(pl.NewTrack(req.Name, time.Duration(req.Duration)*time.Second))
	if err != nil {
		return nil, err
	}
	return trackID, nil
}

func (s *Server) Delete(ctx context.Context, req *proto.Track) (*proto.Nothing, error) {
	err := s.db.DeleteTrack(req.Name)
	if err != nil {
		return nil, err
	}

	err = s.pl.DeleteTrack(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, nil
}

func (s *Server) Update(ctx context.Context, req *proto.UpdateTrackRequest) (*proto.Nothing, error) {
	err := s.db.UpdateTrack(req.Track.Name, req.Track.Duration, req.Name)
	if err != nil {
		return nil, err
	}

	track := pl.NewTrack(req.Track.Name, time.Duration(req.Track.Duration)*time.Second)
	err = s.pl.UpdateTrack(req.Name, track)
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, nil
}
