package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter()  *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/pay")
	})
	paymentInfoRouter(router)
	return router
}




