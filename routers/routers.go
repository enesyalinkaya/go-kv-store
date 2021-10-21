// Package routers builds the controllers
package routers

import (
	"fmt"
	"net/http"

	"github.com/enesyalinkaya/go-kv-store/controller"
	"github.com/enesyalinkaya/go-kv-store/models"
	"github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"
	"github.com/enesyalinkaya/go-kv-store/services"
)

// It builds the controllers
func BuildController(memoryDBClient *memoryDB.MemoryClient) *http.ServeMux {
	router := http.NewServeMux()

	// added healthcheck handler
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"status\":\"OK\"}")
	})

	// create storeModel instance
	storeModel := models.NewStoreModel(memoryDBClient)

	// create storeService instance with storeModel
	storeService := services.NewStoreService(&services.StoreServiceConfig{
		StoreModel: storeModel,
	})

	controller.NewStoreController(&controller.StoreControllerConfig{
		R:            router,
		StoreService: storeService,
	})

	return router
}
