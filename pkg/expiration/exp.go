package expiration

import (
	"fmt"
	"time"

	"github.com/diyliv/keyvaluestorage/internal/models"
)

func ExpTime(expTime int, key interface{}, storage models.Storage) models.Expiration {
	time.AfterFunc(time.Duration(expTime)*time.Second, func() {
		resp, err := storage.Delete(key)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	})
	return models.Expiration{
		Key:            key,
		ExpirationTime: expTime,
		Deleted:        time.Now().Add(time.Duration(expTime) * time.Second).String()}
}
