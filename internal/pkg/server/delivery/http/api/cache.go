package api

import (
	"github.com/babon21/redis-impl/internal/app/server/usecase"
)

type SetStringRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ValueResponse struct {
	Value string `json:"value"`
}

type PushToListRequest struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type PushToListResponse struct {
	Size int `json:"size"`
}

type SetFieldAndValueRequest struct {
	Key   string               `json:"key"`
	Pairs []usecase.FieldValue `json:"pairs"`
}

type SetFieldAndValueResponse struct {
	Count int `json:"count"`
}

type GetValueByFieldRequest struct {
	Key   string `json:"key"`
	Field string `json:"field"`
}

type GetFromListRequest struct {
	Key   string `json:"key"`
	Index int    `json:"index"`
}

type SetValueInListRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index int    `json:"index"`
}

type ExpireKeyRequest struct {
	Key string `json:"key"`
	Ttl int    `json:"ttl"`
}

type KeysRequest struct {
	Pattern string `json:"pattern"`
}

type KeysResponse struct {
	Keys []string `json:"keys"`
}
