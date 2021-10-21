package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enesyalinkaya/go-kv-store/models"
	"github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"
	"github.com/enesyalinkaya/go-kv-store/services"
)

var memoryDBClient = memoryDB.NewMemoryClient("tmp", "test.txt")
var storeModel = models.NewStoreModel(memoryDBClient)
var storeService = services.NewStoreService(&services.StoreServiceConfig{
	StoreModel: storeModel,
})

func TestStoreController_Flush(t *testing.T) {
	storyController := StoreController{
		StoreService: storeService,
	}
	req, err := http.NewRequest("POST", "/flush", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(storyController.FlushHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestStoreController_GetSetHandler(t *testing.T) {
	testKey := "testKey"
	testValue := "testValue"

	storyController := StoreController{
		StoreService: storeService,
	}

	requestBody := &SetRequestBody{
		Value: testValue,
	}

	val, _ := json.Marshal(requestBody)
	path := fmt.Sprintf("/kv/%s", testKey)
	req, err := http.NewRequest("PUT", path, bytes.NewBufferString(string(val)))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(storyController.GetSetHandler)
	handler.ServeHTTP(rr, req)
	var responseBody ResponseBody
	_ = json.NewDecoder(rr.Body).Decode(&responseBody)

	// Check the response is what we expect.
	if responseBody.Value != requestBody.Value || testKey != responseBody.Key {
		t.Errorf("PUT method = %v, want %v", requestBody, responseBody)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	path = fmt.Sprintf("/kv/%s", testKey)
	req, err = http.NewRequest("GET", path, bytes.NewBufferString(string(val)))
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	_ = json.NewDecoder(rr.Body).Decode(&responseBody)

	// Check the response is what we expect.
	if testValue != responseBody.Value || testKey != responseBody.Key {
		t.Errorf("GET method = %v, want value: %v key: %v", responseBody, testValue, testKey)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
