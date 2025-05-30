package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{}

// CreateItem creates a new item
func CreateItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// GetItems retrieves all items
func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// GetItemByID retrieves an item by its ID
func GetItemByID(c *gin.Context) {
	id := c.Param("id")
	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

// UpdateItem updates an existing item by its ID
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

// DeleteItem deletes an item by its ID
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func main() {
	r := gin.Default()

	r.POST("/items", CreateItem)
	r.GET("/items", GetItems)
	r.GET("/items/:id", GetItemByID)
	r.PUT("/items/:id", UpdateItem)
	r.DELETE("/items/:id", DeleteItem)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
