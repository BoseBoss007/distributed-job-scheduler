package main

import (
	"encoding/json"
	"net/http"

	"scheduler/shared"

	"github.com/gin-gonic/gin"
)

func main() {
	shared.InitRedis()

	r := gin.Default()

	r.POST("/job", submitJobHandler)

	r.Run(":8080")
}

// 🔥 Handler stays inside API (NOT shared)
func submitJobHandler(c *gin.Context) {
	var job shared.Job

	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	jobData, _ := json.Marshal(job)

	var queue string

	switch job.Priority {
	case "HIGH":
		queue = "high_priority_queue"
	case "MEDIUM":
		queue = "medium_priority_queue"
	default:
		queue = "low_priority_queue"
	}

	err := shared.Rdb.LPush(shared.Ctx, queue, jobData).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job added"})
}