package kvstoreservice

import (
	"context"
)

func (s *kvStoreService) Set(ctx context.Context, sr *SetRequest) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, err := s.storage.Set(sr.Key, sr.Value)
		if err != nil {
			return nil, err
		}
		return &ItemResponse{
			Key:   sr.Key,
			Value: value,
		}, nil
	}
}
