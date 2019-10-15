package infrastructure

import (
	"crypto/tls"
	"fmt"
	"github.com/Unknwon/goconfig"
	"gopkg.in/gomail.v2"
	"strconv"
)


func SendMail(mailTo []string,subject string, body string) error {
	var smtpInfo = getConfigData("smtp")
	port, _ := strconv.Atoi(smtpInfo["port"])
	m := gomail.NewMessage(gomail.SetCharset("utf-8"))
	m.SetAddressHeader("From",smtpInfo["user"] , "HPay个人收款对账系统")
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpInfo["host"], port, smtpInfo["user"], smtpInfo["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	fmt.Println(err)
	return err
}

func getConfigData(secName string)(map[string] string){
	cfg, err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil{
		panic("错误")
		return nil
	}
	sec, err := cfg.GetSection(secName)
	return sec
}
