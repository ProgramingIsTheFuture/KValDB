package service

import (
	"KValDB/external/gRPC/proto"
	"KValDB/internal/handlers"
	"KValDB/internal/model"
	"context"
	"fmt"
)

type Service struct {
	kv handlers.KValDB
	proto.UnimplementedKValDBServer
}

func NewService(kv handlers.KValDB) *Service {
	return &Service{kv: kv}
}

func (s *Service) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	kval, err := s.kv.Get(in.Dbname, model.Key(in.Key))
	if err != nil {
		fmt.Println("Err", err)
		return nil, err
	}

	fmt.Printf("Get key: %s\n", kval.Key)
	return &proto.GetResponse{Key: kval.Key, Value: kval.Value}, nil
}

func (s *Service) Set(ctx context.Context, in *proto.SetRequest) (*proto.Empty, error) {
	err := s.kv.Set(in.Dbname, model.Key(in.Key), model.Value(in.Value))
	if err != nil {
		fmt.Println("Err", err)
		return nil, err
	}

	fmt.Printf("Set Key: %s - Value: %s\n", in.Key, in.Value)
	return &proto.Empty{}, nil
}
