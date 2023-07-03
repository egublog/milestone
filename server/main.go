package main

import (
	"github.com/gin-gonic/gin"
	"github.com/egublog/milestone/server/routes"
	"github.com/egublog/milestone/server/middleware"
)

type Milestone struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	routes.SetupAPIRoutes(r)

	r.Run()
}
