package httpy

import (
	"log"
	"net/http"
	"task-api/internal/domain"
	"task-api/internal/port/inbound"

	"github.com/gin-gonic/gin"
)

// Struct for connect to service in application
type httpHandler struct {
	service inbound.Connect
}

// Handler : For starting a server
func Handler(service inbound.Connect) {
	handler := &httpHandler{service: service}
	router := gin.Default()

	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.GetAll)
	router.GET("/tasks/:id", handler.GetByID)
	router.PUT("/tasks", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.Delete)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Server can't start: %v", err)
	}

}

func (h *httpHandler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create_task task"})
		return
	}
	task.ID = id
	c.IndentedJSON(http.StatusCreated, task)
}

func (h *httpHandler) GetAll(c *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (h *httpHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (h *httpHandler) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update the requested task"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (h *httpHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the requested task"})
		return
	}
	c.Status(http.StatusOK)
}
