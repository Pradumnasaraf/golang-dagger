package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/items", CreateItem)
	r.GET("/items", GetItems)
	r.GET("/items/:id", GetItemByID)
	r.PUT("/items/:id", UpdateItem)
	r.DELETE("/items/:id", DeleteItem)
	return r
}

func TestCreateItem(t *testing.T) {
	r := setupRouter()

	// Reset items slice for a clean test environment
	items = []Item{}

	newItem := Item{ID: "1", Name: "Test Item"}
	jsonValue, _ := json.Marshal(newItem)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdItem Item
	json.Unmarshal(w.Body.Bytes(), &createdItem)
	assert.Equal(t, newItem.ID, createdItem.ID)
	assert.Equal(t, newItem.Name, createdItem.Name)
	assert.Equal(t, 1, len(items))
}

func TestGetItems(t *testing.T) {
	r := setupRouter()

	// Add some items first
	items = []Item{{ID: "1", Name: "Item 1"}, {ID: "2", Name: "Item 2"}}

	req, _ := http.NewRequest("GET", "/items", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedItems []Item
	json.Unmarshal(w.Body.Bytes(), &retrievedItems)
	assert.Equal(t, 2, len(retrievedItems))
	assert.Equal(t, "Item 1", retrievedItems[0].Name)
	assert.Equal(t, "Item 2", retrievedItems[1].Name)
}

func TestGetItemByID(t *testing.T) {
	r := setupRouter()

	// Add an item
	items = []Item{{ID: "1", Name: "Item 1"}}

	// Test existing item
	req, _ := http.NewRequest("GET", "/items/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedItem Item
	json.Unmarshal(w.Body.Bytes(), &retrievedItem)
	assert.Equal(t, "1", retrievedItem.ID)
	assert.Equal(t, "Item 1", retrievedItem.Name)

	// Test non-existing item
	req, _ = http.NewRequest("GET", "/items/2", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateItem(t *testing.T) {
	r := setupRouter()

	// Add an item
	items = []Item{{ID: "1", Name: "Item 1"}}

	// Test updating existing item
	updatedItem := Item{ID: "1", Name: "Updated Item"}
	jsonValue, _ := json.Marshal(updatedItem)

	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Updated Item", items[0].Name)

	// Test updating non-existing item
	updatedItemNonExisting := Item{ID: "2", Name: "Should Not Exist"}
	jsonValueNonExisting, _ := json.Marshal(updatedItemNonExisting)

	req, _ = http.NewRequest("PUT", "/items/2", bytes.NewBuffer(jsonValueNonExisting))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteItem(t *testing.T) {
	r := setupRouter()

	// Add an item
	items = []Item{{ID: "1", Name: "Item 1"}}

	// Test deleting existing item
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 0, len(items))

	// Test deleting non-existing item
	req, _ = http.NewRequest("DELETE", "/items/2", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
