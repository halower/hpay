package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/halower/hipay/infrastructure"
	"github.com/halower/hipay/models"
	"io"
	"net/http"
	"strconv"
	"time"
)

// @Summary 获取所有支付信息
// @version 1.0
// @Tags  支付接口
// @Param page query int true  "页码"
// @Param page_size  query int true "页容"
// @Success 200 {object} models.GetPaysListOutputDto
// @Router /api/pay/list [get]
func GetPaysList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	paymentInfo := models.PaymentInfo{}
	paymentInfoOutputs, err := paymentInfo.GetPaymentInfoList(pageInt, pageSizeInt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":         1000,
			"msg":          infrastructure.FindAllFail,
			"error_detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  infrastructure.FindAllSuccess,
		"data": paymentInfoOutputs,
	})
}

// @Summary 发起支付
// @version 1.0
// @Tags  支付接口
// @Accept  json
// @Produce  json
// @Param payment_info body models.PaymentInfoInsertDto true "交易信息"
// @Success 200 {object} models.PrePaidOutputDto
// @Router /api/pay/pending  [post]
func PayPending(ctx *gin.Context) {
	var p models.PaymentInfo
	err := ctx.ShouldBindJSON(&p)

	p.PayeeEmail = "121625933@qq.com"
	p.PayeeMobileNumber = "123456789"
	p.PayeeName = "OldXie"
	p.PayeeQRCodePath = "/"
	p.TradingStatus= "PENDING"
	p.IdCode = infrastructure.RandomString(6)

	id, err := p.Pay()
    p.Id = id
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":         1000,
			"msg":          infrastructure.PayPendingFail,
			"error_detail": err.Error(),
		})
		return
	}


	body := infrastructure.GetAuditMailBody(p)
	_ = infrastructure.SendMail([]string{p.PayeeEmail}, "【HPay个人收款对账系统】--收款通知", body)
	fmt.Println("发送付款邮件给:", p.PayeeEmail, "发送成功")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  infrastructure.PayPendingSuccess,
		"data": models.PrePaidOutputDto{Id: id, IdCode: p.IdCode},
	})
}

// @Summary 服务推送
// @version 1.0
// @Tags  支付接口
// @Router /api/pay/stream  [get]
func Sse(ctx  *gin.Context) {
	var p models.PaymentInfo
	id :=ctx.Param("id")
	p.Id, _ = strconv.ParseInt(id, 10, 64)
	chanStream := make(chan string, 1)
	go func() {
		defer close(chanStream)
		var hasAudited = false
		for {
			if hasAudited {
				chanStream <- "FINISHED"
				return
			} else {
				hasAudited , _ = p.HasAudited()
				chanStream <- "PROCESSING"
			}
			time.Sleep(time.Second * 5)
		}
	}()
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}


// @Summary 确认状态
// @version 1.0
// @Tags  支付接口
// @Accept  json
// @Produce  json
// @Param payment_info body models.TradingStatusInputDto true "状态"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/pay/status/:id/:trading_status [get]
func ConfirmStatus (ctx * gin.Context) {
	var p models.PaymentInfo
	p.Id, _ =strconv.ParseInt(ctx.Param("id"), 10, 64)
	p.TradingStatus =ctx.Param("trading_status")
	success, _:=p.UpdateTradingStatus()
	if success {
		payerEmailDto, _ :=p.FindPaymentInfoById()

		switch {
		case payerEmailDto.Amount == "10" :
			payerEmailDto.Mark = "程序源码及部署下载地址:链接: https://pan.baidu.com/s/1Oh14hZsFm7vd7JNiSuY0sQ 提取码: 2g58 "
		case payerEmailDto.Amount == "20" :
			payerEmailDto.Mark = "你将获得程序的持续升级服务和解答,程序源码及部署下载地址:链接: https://pan.baidu.com/s/1Oh14hZsFm7vd7JNiSuY0sQ 提取码: 2g58 "
		default:
			payerEmailDto.Mark  = "感谢使用"
		}

		body := infrastructure.GetFeedBackMailBody(payerEmailDto)
		_ = infrastructure.SendMail([]string{payerEmailDto.PayerEmail}, "【HPay个人收款对账系统】--到账通知", body)
		fmt.Println("发送到账邮件给:", payerEmailDto.PayeeEmail, "发送成功")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  infrastructure.ChangeTradingStatusSuccess,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  infrastructure.ChangeTradingStatusFail,
	})
}