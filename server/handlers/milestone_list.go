package handlers

import	(
	"github.com/gin-gonic/gin"
	"github.com/egublog/milestone/server/models"
)

func GetMilestones(c *gin.Context) {
	milestones := []models.Milestone{
		{ID: 1, Title: "Milestone 1"},
		{ID: 2, Title: "Milestone 2"},
		{ID: 3, Title: "Milestone 3"},
	}

	c.JSON(200, milestones)
}