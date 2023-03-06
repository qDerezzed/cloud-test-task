package delivery

import (
	"context"

	"cloud-test-task/internal/delivery/pb"
	"cloud-test-task/internal/usecase"

	"google.golang.org/grpc"
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

func (s *server) AddTrack(ctx context.Context, req *pb.Track) (*pb.TrackID, error) {
	id, err := s.usecase.AddTrack(req.Name, req.Duration)
	if err != nil {
		return nil, err
	}

	return &pb.TrackID{ID: id}, nil
}

func (s *server) DeleteTrack(ctx context.Context, req *pb.Track) (*pb.Nothing, error) {
	err := s.usecase.DeleteTrack(req.Name)
	return &pb.Nothing{Dummy: true}, err
}

func (s *server) UpdateTrack(ctx context.Context, req *pb.UpdateTrackRequest) (*pb.Nothing, error) {
	err := s.usecase.UpdateTrack(req.Name, req.Track.Name, req.Track.Duration)
	if err != nil {
		return nil, err
	}

	return &pb.Nothing{Dummy: true}, nil
}

func (s *server) GetTrack(ctx context.Context, n *pb.Nothing) (*pb.Track, error) {
	track := s.usecase.GetTrack()
	return &pb.Track{Name: track.Name, Duration: int64(track.CurrentTime)}, nil
}
