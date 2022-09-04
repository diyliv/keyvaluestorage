package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diyliv/keyvaluestorage/internal/models"
)

type httphandler struct {
	storage models.Storage
}

func NewHttpHandler(storage models.Storage) *httphandler {
	return &httphandler{storage: storage}
}

func (h *httphandler) Add(w http.ResponseWriter, r *http.Request) {
	var req models.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	var expResult models.Expiration

	if req.ExpTime != 0 {
		expResult = h.expTime(req.ExpTime, req.Key)
		expResult.Value = req.Value
	}

	resp, err := h.storage.Add(req.Key, req.Value)
	if err != nil {
		h.writeResponse(http.StatusBadRequest, w, "something went wrong :(")
	}

	resp.Expiration = expResult

	h.writeResponse(http.StatusOK, w, resp)
}

func (h *httphandler) Get(w http.ResponseWriter, r *http.Request) {
	var req models.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	resp, err := h.storage.Get(req.Key)
	if err != nil {
		h.writeResponse(http.StatusBadRequest, w, err.Error())
		return
	}

	h.writeResponse(http.StatusOK, w, resp)
}

func (h *httphandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req models.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}

	resp, err := h.storage.Delete(req.Key)
	if err != nil {
		h.writeResponse(http.StatusBadRequest, w, err.Error())
		return
	}

	resp.DeletedTime = time.Now().Local().String()

	h.writeResponse(http.StatusOK, w, resp)

}

func (h *httphandler) writeResponse(code int, w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		panic(err)
	}
}

func (h *httphandler) expTime(expTime int, key interface{}) models.Expiration {
	time.AfterFunc(time.Duration(expTime)*time.Second, func() {
		resp, err := h.storage.Delete(key)
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
