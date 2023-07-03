package main

import (
	"github.com/gin-gonic/gin"
)

type Milestone struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	r := gin.Default()

	// CORS設定
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Reactアプリのオリジンを許可
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")                  // 許可するHTTPメソッドを設定
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // 許可するヘッダーを設定
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")                      // プリフライトリクエストのキャッシュ時間を設定
		c.Next()
	})

	r.GET("/api/milestones", func(c *gin.Context) {
		milestone := []Milestone{
			{ID: 1, Title: "Milestone 1"},
			{ID: 2, Title: "Milestone 2"},
			{ID: 3, Title: "Milestone 3"},
		}

		c.JSON(200, milestone)
	})

	r.Run()
}
