package http

import (
	"github.com/babon21/redis-impl/internal/app/server/usecase"
	"github.com/babon21/redis-impl/internal/pkg/server/delivery/http/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CacheHandler  represent the httphandler for redis cache
type CacheHandler struct {
	RedisUsecase usecase.RedisUsecase
}

func (h *CacheHandler) GetString(c echo.Context) error {
	key := c.Param("key")

	value, ok, err := h.RedisUsecase.Get(key)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	if !ok {
		return c.JSONPretty(http.StatusNotFound, ResponseError{Message: "key is not found"}, "  ")
	}

	response := api.ValueResponse{Value: value}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (h *CacheHandler) SetString(c echo.Context) error {
	var request api.SetStringRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	h.RedisUsecase.Set(request.Key, request.Value)

	return c.NoContent(http.StatusCreated)
}

func (h *CacheHandler) GetValueByFieldInMap(c echo.Context) error {
	var request api.GetValueByFieldRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	value, ok, err := h.RedisUsecase.HGet(request.Key, request.Field)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	if !ok {
		return c.JSONPretty(http.StatusNotFound, ResponseError{Message: "key or field is not found"}, "  ")
	}

	response := api.ValueResponse{Value: value}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (h *CacheHandler) SetFieldAndValueInMap(c echo.Context) error {
	var request api.SetFieldAndValueRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	count, err := h.RedisUsecase.HSet(request.Key, request.Pairs)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	response := api.SetFieldAndValueResponse{Count: count}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (h *CacheHandler) GetFromList(c echo.Context) error {
	var request api.GetFromListRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	value, err := h.RedisUsecase.LGet(request.Key, request.Index)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	response := api.ValueResponse{Value: value}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (h *CacheHandler) PushToList(c echo.Context) error {
	var request api.PushToListRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	size, err := h.RedisUsecase.LPush(request.Key, request.Values)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	response := api.PushToListResponse{Size: size}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (h *CacheHandler) SetValueInList(c echo.Context) error {
	var request api.SetValueInListRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	err = h.RedisUsecase.LSet(request.Key, request.Index, request.Value)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CacheHandler) Delete(c echo.Context) error {
	key := c.Param("key")
	h.RedisUsecase.Del(key)
	return c.NoContent(http.StatusNoContent)
}

func (h *CacheHandler) ExpireKey(c echo.Context) error {
	var request api.ExpireKeyRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	ok := h.RedisUsecase.Expire(request.Key, request.Ttl)
	if !ok {
		return c.JSONPretty(http.StatusNotFound, ResponseError{Message: "key is not found"}, "  ")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CacheHandler) GetKeys(c echo.Context) error {
	var request api.KeysRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	list, err := h.RedisUsecase.Keys(request.Pattern)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	response := api.KeysResponse{Keys: list}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

// NewCacheHandler will initialize the cache/ resources endpoint
func NewCacheHandler(e *echo.Echo, us usecase.RedisUsecase) {
	handler := &CacheHandler{
		RedisUsecase: us,
	}

	e.GET("/cache/string/:key", handler.GetString)
	e.PUT("/cache/string", handler.SetString)

	e.GET("/cache/map", handler.GetValueByFieldInMap)
	e.PUT("/cache/map", handler.SetFieldAndValueInMap)

	e.GET("/cache/list", handler.GetFromList)
	e.POST("/cache/list", handler.PushToList)
	e.PATCH("/cache/list", handler.SetValueInList)

	e.DELETE("/cache/keys/:key", handler.Delete)
	e.PATCH("/cache/keys/expire", handler.ExpireKey)
	e.GET("/cache/keys", handler.GetKeys)
}
