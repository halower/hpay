package models

import "time"

type PaymentInfoInsertDto struct {
	// 支付人姓名
	PayerName string `json:"payer_name"`

	// 支付人邮箱
	PayerEmail string `json:"payer_email"`

	// 支付人电话号码
	PayerMobileNumber string `json:"payer_mobile_number"`

	// 交易金额
	Amount string `json:"amount"`

	// 留言
	PayerMessage string `json:"payer_message"`

	//支付方式
	PayType string `json:"pay_type"`
}

type PaymentInfoOutputDto struct {
	// 流水号
	Id int64 `json:"id"`

	// 支付人姓名
	PayerName string `json:"payer_name"`

	// 支付人邮箱
	PayerEmail string `json:"payer_email"`

	// 支付人电话号码
	PayerMobileNumber string `json:"payer_mobile_number"`

	// 交易金额
	Amount string `json:"amount"`

	//支付方式
	PayType string `json:"pay_type"`

	// 收款人姓名
	PayeeName string `json:"payee_name"`

	// 收款人邮箱
	PayeeEmail string `json:"payee_email"`

	// 收款人电话号码
	PayeeMobileNumber string `json:"payee_mobile_number"`

	// 交易状态
	TradingStatus string `json:"trading_status"`

	// 创建时间
	CreationDate time.Time `json:"creation_date"`
}

type PayerEmailOutputDto struct {
	// 流水号
	Id int64 `json:"id"`

	// 票据号
	IdCode string `json:"id_code"`

	// 支付人姓名
	PayerName string `json:"payer_name"`

	// 支付人邮箱
	PayerEmail string `json:"payer_email"`

	// 支付人电话号码
	PayerMobileNumber string `json:"payer_mobile_number"`

	// 交易金额
	Amount string `json:"amount"`

	//支付方式
	PayType string `json:"pay_type"`

	// 收款人邮箱
	PayeeEmail string `json:"payee_email"`
	
	// mark
	Mark string `json:"mark"`
}

type GetPaysListOutputDto struct {
	// 信息
	Message string `json:"msg" example:"成功"`
	// 业务状态代码
	Code int `json:"code" example:"200"`
	// 数据
	Data []PaymentInfoOutputDto
	// 详细错误
	ErrorMsg string `json:"error_detail" example:"详细错误"`
}

type PayPendingOutputDto struct {
	// 信息
	Message string `json:"msg" example:"成功"`
	// 业务状态代码
	Code int `json:"code" example:"200"`
	// 新增数据的Id
	Data int `json:"data" example: "1"`
	// 详细错误
	ErrorMsg string `json:"error_detail" example:"详细错误"`
}

type PrePaidOutputDto struct {
	// 流水号
	Id int64 `json:"id"`

	// 支付票据号
	IdCode string `json:"id_code"`
}

type TradingStatusInputDto struct {
	Id            int64  `json:"id" binding:"required"`
	TradingStatus string `json:"trading_status"   binding:"required"`
}


