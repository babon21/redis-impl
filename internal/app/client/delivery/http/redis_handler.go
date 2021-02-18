package http

import (
	"github.com/babon21/redis-impl/internal/app/client/usecase"
	"github.com/labstack/echo"
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

	response, err := h.RedisUsecase.Get(key)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) SetString(c echo.Context) error {
	response, err := h.RedisUsecase.Set(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) GetValueByFieldInMap(c echo.Context) error {
	response, err := h.RedisUsecase.HGet(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) SetFieldAndValueInMap(c echo.Context) error {
	response, err := h.RedisUsecase.HSet(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) GetFromList(c echo.Context) error {
	response, err := h.RedisUsecase.LGet(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) PushToList(c echo.Context) error {
	response, err := h.RedisUsecase.LPush(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) SetValueInList(c echo.Context) error {
	response, err := h.RedisUsecase.LSet(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) Delete(c echo.Context) error {
	key := c.Param("key")
	response, err := h.RedisUsecase.Del(key)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) ExpireKey(c echo.Context) error {
	response, err := h.RedisUsecase.Expire(c.Request().Body)
	return returnServerResponse(c, response, err)
}

func (h *CacheHandler) GetKeys(c echo.Context) error {
	response, err := h.RedisUsecase.Keys(c.Request().Body)
	return returnServerResponse(c, response, err)
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

func returnServerResponse(c echo.Context, response *http.Response, err error) error {
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, ResponseError{Message: err.Error()}, "  ")
	}
	defer response.Body.Close()

	return c.Stream(response.StatusCode, "application/json", response.Body)
}
