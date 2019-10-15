package router

import (
	"github.com/gin-gonic/gin"
	"github.com/halower/hipay/apis"
)

func paymentInfoRouter(router *gin.Engine) {
	paymentInfoRouter := router.Group("api/pay")
	{
		paymentInfoRouter.GET("list", apis.GetPaysList)
		paymentInfoRouter.POST("pending", apis.PayPending)
		paymentInfoRouter.GET("stream/:id", apis.Sse)
		paymentInfoRouter.GET("status/:id/:trading_status",apis.ConfirmStatus)
	}
}

