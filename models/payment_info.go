package models

import (
	"database/sql"
	db "github.com/halower/hipay/databse"
	"log"
	"time"
)

type PaymentInfo struct {
	Id                int64     `json:"id" form:"Id" primaryKey:"true"`
	PayerName         string    `json:"payer_name" form "PayerName" binding:"required"`
	PayerEmail        string    `json:"payer_email" from: "PayerEmail"`
	PayerMobileNumber string    `json:"payer_mobile_number" from: "PayerMobileNumber"`
	Amount            string    `json:"amount" from: "Amount" binding:"required"`
	PayerMessage      string    `json:"payer_message" from: "PayerMessage"`
	PayType           string    `json:"pay_type" from: "PayType binding:"required""`
	PayeeName         string    `json:"payee_name" form:"PayeeName" binding:"required"`
	PayeeEmail        string    `json:"payee_email" from: "PayeeEmail"`
	PayeeMobileNumber string    `json:"payee_mobile_number" from: "PayeeMobileNumber"`
	PayeeQRCodePath   string    `json:"payee_qr_code_path" from: "PayeeQRCodePath" binding:"required"`
	TradingStatus     string    `json:"trading_status" from: "TradingStatus"`
	CreationDate      time.Time `json:"creation_date" from "CreationDate"`
	IdCode            string    `json:"id_code" from "IdCode"`
}

// 查询支付列表
func (p *PaymentInfo) GetPaymentInfoList(page int, pageSize int) (paymentInfoOutputs []PaymentInfoOutputDto, err error) {
	paymentInfoOutputs = make([]PaymentInfoOutputDto, 0)
	offset := pageSize * (page - 1)
	limit := pageSize
	rows, err := db.SqlDB.Query("SELECT Id, PayerName, PayerEmail, PayeeMobileNumber, Amount, PayType, PayeeName,PayeeEmail,PayeeMobileNumber,TradingStatus,CreationDate FROM PaymentInfo LIMIT ?, ?", offset, limit)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		var paymentInfoOutput PaymentInfoOutputDto
		rows.Scan(
			&paymentInfoOutput.Id, &paymentInfoOutput.PayerName, &paymentInfoOutput.PayerEmail, &paymentInfoOutput.PayerMobileNumber,
			&paymentInfoOutput.Amount, &paymentInfoOutput.PayType, &paymentInfoOutput.PayeeName, &paymentInfoOutput.PayeeEmail, &paymentInfoOutput.PayeeMobileNumber,
			&paymentInfoOutput.TradingStatus, &paymentInfoOutput.CreationDate,
		)
		paymentInfoOutputs = append(paymentInfoOutputs, paymentInfoOutput)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *PaymentInfo) FindPaymentInfoById()(payerEmailDto PayerEmailOutputDto, err error) {
	rs:= db.SqlDB.QueryRow("SELECT  Id, PayerName, PayerEmail, PayeeMobileNumber, Amount, PayType,IdCode,PayeeEmail FROM PaymentInfo WHERE id=?", p.Id)
	rs.Scan(
		&payerEmailDto.Id, &payerEmailDto.PayerName,
		&payerEmailDto.PayerEmail, &payerEmailDto.PayerMobileNumber,
		&payerEmailDto.Amount, &payerEmailDto.PayType, &payerEmailDto.IdCode, &payerEmailDto.PayeeEmail,
	)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// 发起支付
func (p *PaymentInfo) Pay() (id int64, err error) {
	rs, err := db.SqlDB.Exec(
		"INSERT INTO HALOWER_EASYPAY.PaymentInfo ("+
			"PayerName, PayerEmail, PayerMobileNumber, Amount, PayerMessage, PayType,"+
			"PayeeName, PayeeEmail, PayeeMobileNumber, PayeeQRCodePath,IdCode)"+
			" VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?,?);",
		p.PayerName, p.PayerEmail, p.PayerMobileNumber, p.Amount, p.PayerMessage, p.PayType,
		p.PayeeName, p.PayeeEmail, p.PayerMobileNumber, p.PayeeQRCodePath, p.IdCode)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

// 判断是否审核通过
func (p *PaymentInfo) HasAudited() (hasAudited bool, err error) {
	var status string
	db.SqlDB.QueryRow("SELECT TradingStatus FROM PaymentInfo WHERE Id=?", p.Id).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("There is not row")
		} else {
			log.Fatalln(err)
		}
		return
	}
	hasAudited = status == "FINISHED"
	return
}

// 更改审核状态
func (p *PaymentInfo) UpdateTradingStatus() (success bool, err error) {
	res, err := db.SqlDB.Exec("UPDATE PaymentInfo SET TradingStatus = ? WHERE Id = ?", p.TradingStatus, p.Id)
	num, err := res.RowsAffected()
	success = num > 0
	return
}
