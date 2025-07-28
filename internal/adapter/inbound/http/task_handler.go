package http


import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"task-api/internal/domain"
	"task-api/internal/port/inbound"
)

type TaskHandler struct {
	service inbound.TaskService
}

func NewTaskHandler(router *gin.Engine, service inbound.TaskService) {
	handler := &TaskHandler{service: service}
	
	router.POST("/tasks", handler.Create)
	router.GET("/tasks", handler.GetAll)
	router.GET("/tasks/:id", handler.GetByID)
	router.PUT("/tasks/:id", handler.Update)
	router.DELETE("/tasks/:id", handler.Delete)
}


func (h *TaskHandler) Create(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}


func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}



func (h *TaskHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	task, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}



func (h *TaskHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = id
	err := h.service.Update(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(http.StatusOK, task)
}


func (h *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.Status(http.StatusNoContent)
}

