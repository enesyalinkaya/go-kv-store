package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/enesyalinkaya/go-kv-store/services"
)

type StoreController struct {
	StoreService services.StoreService
}

type StoreControllerConfig struct {
	R            *http.ServeMux
	StoreService services.StoreService
}

func NewStoreController(c *StoreControllerConfig) {
	StoreController := &StoreController{
		StoreService: c.StoreService,
	}
	c.R.HandleFunc("/kv/", StoreController.GetSetHandler)
	c.R.HandleFunc("/flush", StoreController.FlushHandler)
}

type ResponseBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetRequestBody struct {
	Value string `json:"value"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetSet handler, it handles GET requests to get value and PUT requests to set value
func (c *StoreController) GetSetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Getting Key variable from path
	key := strings.TrimPrefix(r.URL.Path, "/kv/")
	ctx := r.Context()

	switch r.Method {

	// HTTP GET method provides getting data with key
	case "GET":
		value := c.StoreService.Get(ctx, key)
		json.NewEncoder(w).Encode(&ResponseBody{
			Key:   key,
			Value: value,
		})
		return

	// HTTP PUT method provides setting data with key
	case "PUT":
		var body SetRequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		value := c.StoreService.Set(ctx, key, body.Value)
		json.NewEncoder(w).Encode(&ResponseBody{
			Key:   key,
			Value: value,
		})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Error: "Sorry, only GET and PUT methods are supported.",
		})
		return
	}
}

// Flush Handler, it flushes in-memory store
func (c *StoreController) FlushHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	switch r.Method {
	case "POST":
		c.StoreService.Flush(ctx)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Error: "Sorry, only POST method is supported.",
		})
		return
	}
}
