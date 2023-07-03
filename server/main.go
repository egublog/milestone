package main

import (
	"github.com/gin-gonic/gin"
)

type Milestone struct {
	ID    int 	 `json:"id"`
	Title string `json:"title"`
}

func main() {
	r := gin.Default()

	r.GET("/milestones", func(c *gin.Context) {
		milestone := []Milestone{
			{ID: 1, Title: "Milestone 1"},
			{ID: 2, Title: "Milestone 2"},
			{ID: 3, Title: "Milestone 3"},
		}

		c.JSON(200, milestone)
	})

	r.Run()
}