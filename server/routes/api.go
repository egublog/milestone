package routes

import (
	"github.com/egublog/milestone/server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(r *gin.Engine) {
	api := r.Group("/api/milestones")

	api.GET("", handlers.GetMilestones)
	api.GET(":id", handlers.GetMilestoneDetail)
}
