package main

import (
	"github.com/egublog/milestone/server/middleware"
	"github.com/egublog/milestone/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	routes.SetupAPIRoutes(r)

	r.Run()
}
