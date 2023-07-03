package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Reactアプリのオリジンを許可
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")                  // 許可するHTTPメソッドを設定
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // 許可するヘッダーを設定
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")                      // プリフライトリクエストのキャッシュ時間を設定
		c.Next()
	}
}