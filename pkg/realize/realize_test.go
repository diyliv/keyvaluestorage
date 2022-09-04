package realize

import (
	"testing"

	"github.com/diyliv/keyvaluestorage/internal/models"
)

func TestAdd(t *testing.T) {
	storage := NewStorage(make(map[interface{}]models.Response))

	test := []struct {
		key   interface{}
		value interface{}
	}{
		{
			key:   "hello",
			value: "world",
		},
		{
			key:   "secret-data",
			value: "some-secret-value",
		}, {key: "idk", value: "idk-data"},
	}

	for _, value := range test {
		resp, err := storage.Add(value.key, value.value)
		if err != nil {
			t.Error(err)
		}

		if resp.Added != true {
			t.Errorf("Something went wrong while adding values.")
		}
	}
}

func TestGet(t *testing.T) {
	storage := NewStorage(make(map[interface{}]models.Response))

	test := []struct {
		key   interface{}
		value interface{}
	}{{
		key:   "hello",
		value: "world",
	}, {
		key:   "world",
		value: "hello",
	}}

	for _, value := range test {
		_, err := storage.Add(value.key, value.value)
		if err != nil {
			t.Error(err)
		}

		resp, err := storage.Get(value.key)
		if err != nil {
			t.Error(err)
		}
		if resp.Value != value.value {
			t.Errorf("Unexpected values. Got: %v want %v\n", resp.Value, value.value)
		}
	}
}

func TestDelete(t *testing.T) {
	storage := NewStorage(make(map[interface{}]models.Response))

	test := []struct {
		key   interface{}
		value interface{}
	}{{
		key:   "hello",
		value: "world",
	}, {
		key:   "delete-me",
		value: "deleted-value",
	}}

	for _, value := range test {
		_, err := storage.Add(value.key, value.value)
		if err != nil {
			t.Error(err)
		}

		delete, err := storage.Delete(value.key)
		if err != nil {
			t.Error(err)
		}

		if delete.Added != false {
			t.Errorf("Error while deleting.")
		}
	}
}
