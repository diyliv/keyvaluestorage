package models

type Request struct {
	Key     interface{} `json:"key"`
	Value   interface{} `json:"value"`
	ExpTime int         `json:"exp_time"`
}

type Response struct {
	Key         interface{} `json:"key"`
	Value       interface{} `json:"value"`
	Added       bool        `json:"added_status"`
	AddedTime   string      `json:"added_time"`
	DeletedTime string      `json:"deleted_time"`
	Expiration  Expiration  `json:"expiration_info"`
}

type Storage interface {
	Add(key, data interface{}) (Response, error)
	Get(key interface{}) (*Response, error)
	Delete(key interface{}) (*Response, error)
}
