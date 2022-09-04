package grpc

import (
	"context"
	"time"

	"github.com/diyliv/keyvaluestorage/internal/models"
	"github.com/diyliv/keyvaluestorage/pkg/expiration"
	"github.com/diyliv/keyvaluestorage/proto/keyval"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type grpcService struct {
	storage models.Storage
}

func NewGrpcService(storage models.Storage) *grpcService {
	return &grpcService{storage: storage}
}

func (s *grpcService) Add(ctx context.Context, req *keyval.Request) (*keyval.Response, error) {
	key := req.Key
	value := req.Value
	expTime := req.ExpTime

	var expResult models.Expiration

	if expTime != 0 {
		expResult = expiration.ExpTime(int(expTime), key, s.storage)
		expResult.Value = value
	}

	resp, err := s.storage.Add(key, value)
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong ):")
	}

	resp.Expiration = expResult

	return &keyval.Response{
		Key:       &anypb.Any{Value: key.Value},
		Value:     &anypb.Any{Value: value.Value},
		Added:     true,
		AddedTime: time.Now().Local().String(),
		Expiration: &keyval.Expiration{
			Key:     &anypb.Any{Value: key.Value},
			Value:   &anypb.Any{Value: value.Value},
			ExpTime: expTime,
		},
	}, nil
}
