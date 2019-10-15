package middleware

import "github.com/gin-gonic/gin"
import  "github.com/unrolled/secure"
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			return
		}
		c.Next()
	}
}
