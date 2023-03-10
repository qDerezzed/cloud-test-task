package delivery

import (
	"context"
	"errors"

	"cloud-test-task/internal/delivery/pb"
	"cloud-test-task/internal/entities"
	"cloud-test-task/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedPlaylistServer

	usecase usecase.PlaylistUseCase
}

func NewPlaylistServerGrpc(gserver *grpc.Server, playlistUcase usecase.PlaylistUseCase) error {

	playlistServer := &server{
		usecase: playlistUcase,
	}

	pb.RegisterPlaylistServer(gserver, playlistServer)

	err := playlistUcase.Init()

	return err
}

func (s *server) Start(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	err := s.usecase.Start()
	return &pb.Nothing{Dummy: true}, err
}

func (s *server) Off(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	s.usecase.Off()
	return &pb.Nothing{Dummy: true}, nil
}

func (s *server) Play(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	err := s.usecase.Play()
	return &pb.Nothing{Dummy: true}, err
}

func (s *server) Pause(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	err := s.usecase.Pause()
	return &pb.Nothing{Dummy: true}, err
}

func (s *server) Next(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	s.usecase.Next()
	return &pb.Nothing{Dummy: true}, nil
}

func (s *server) Prev(ctx context.Context, n *pb.Nothing) (*pb.Nothing, error) {
	s.usecase.Prev()
	return &pb.Nothing{Dummy: true}, nil
}

func (s *server) AddTrack(ctx context.Context, req *pb.AddTrackRequest) (*pb.TrackID, error) {
	id, err := s.usecase.AddTrack(req.Name, req.Duration)
	if err != nil {
		if errors.Is(err, entities.ErrorNotValidTrackDuration) {
			err = status.Error(codes.InvalidArgument, err.Error())
		} else {
			err = status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, err
	}

	return &pb.TrackID{ID: id}, nil
}

func (s *server) DeleteTrack(ctx context.Context, req *pb.TrackID) (*pb.Nothing, error) {
	err := s.usecase.DeleteTrack(req.ID)
	if err != nil {
		if errors.Is(err, entities.ErrorDelTrackNotFound) {
			err = status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, entities.ErrorAlreadyPlay) {
			err = status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &pb.Nothing{Dummy: true}, err
}

func (s *server) UpdateTrack(ctx context.Context, req *pb.UpdateTrackRequest) (*pb.Nothing, error) {
	err := s.usecase.UpdateTrack(req.ID, req.NewName, req.NewDuration)
	if err != nil {
		if errors.Is(err, entities.ErrorUpdTrackNotFound) {
			err = status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, entities.ErrorAlreadyPlay) {
			err = status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}

	return &pb.Nothing{Dummy: true}, nil
}

func (s *server) GetTrack(ctx context.Context, n *pb.Nothing) (*pb.Track, error) {
	track := s.usecase.GetTrack()
	return TrackToProto(track), nil
}

func (s *server) GetPlaylist(ctx context.Context, n *pb.Nothing) (*pb.TrackList, error) {
	playlist, err := s.usecase.GetPlaylist()
	if err != nil {
		return nil, err
	}

	result := &pb.TrackList{}
	for _, v := range playlist {
		result.Playlist = append(result.Playlist, TrackToProto(*v))
	}

	return result, nil
}

func TrackToProto(track entities.Track) *pb.Track {
	return &pb.Track{ID: int64(track.ID), Name: track.Name, Duration: int64(track.Duration)}
}
