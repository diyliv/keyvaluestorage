package realize

import (
	"sync"
	"time"

	"github.com/diyliv/keyvaluestorage/internal/models"
)

type storage struct {
	mu    sync.Mutex
	cloud map[interface{}]models.Response
}

func NewStorage(cloud map[interface{}]models.Response) *storage {
	return &storage{cloud: cloud}
}

func (s *storage) Add(key, val interface{}) (models.Response, error) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.cloud[key] = models.Response{
		Key:       key,
		Value:     val,
		Added:     true,
		AddedTime: time.Now().Local().String()}

	return models.Response{
		Key:       key,
		Value:     val,
		Added:     true,
		AddedTime: time.Now().Local().String(),
	}, nil
}

func (s *storage) Get(key interface{}) (*models.Response, error) {
	defer s.mu.Unlock()
	s.mu.Lock()
	val, ok := s.cloud[key]
	if !ok {
		return nil, ErrNotFound
	}
	return &val, nil
}

func (s *storage) Delete(key interface{}) (*models.Response, error) {
	defer s.mu.Unlock()
	s.mu.Lock()

	if val, ok := s.cloud[key]; ok {
		delete(s.cloud, key)
		val.Added = false
		return &val, nil
	} else {
		return nil, ErrNotFound
	}
}
