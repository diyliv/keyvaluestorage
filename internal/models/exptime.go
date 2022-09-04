package models

type Expiration struct {
	Key            interface{} `json:"exp_key"`
	Value          interface{} `json:"key_value"`
	ExpirationTime int         `json:"exp_time"`
	Deleted        string      `json:"deleted_time"`
}
