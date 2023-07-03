package handlers

import (
	"github.com/egublog/milestone/server/db"
	"github.com/gin-gonic/gin"
)

func GetMilestones(c *gin.Context) {
	milestones := db.GetDummyMilestones()

	c.JSON(200, milestones)
}
