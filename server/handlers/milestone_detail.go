package handlers

import (
	"github.com/egublog/milestone/server/db"
	"github.com/egublog/milestone/server/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetMilestoneDetail(c *gin.Context) {
	milestoneIDStr := c.Param("id")

	milestoneID, err := strconv.Atoi(milestoneIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid milestone id"})
	}

	milestones := db.GetDummyMilestones()

	// マイルストーンIDに該当するデータを取得
	var milestone models.Milestone
	for _, m := range milestones {
		if m.ID == milestoneID {
			milestone = m
			break
		}
	}

	// マイルストーンIDに該当するデータがない場合は404を返す
	if milestone.ID == 0 {
		c.JSON(404, gin.H{"error": "Milestone not found"})
		return
	}

	c.JSON(200, milestone)
}
