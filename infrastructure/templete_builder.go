package infrastructure

import (
	"bytes"
	"github.com/halower/hipay/models"
	"html/template"
	"log"
)

func GetAuditMailBody(input models.PaymentInfo) string {
	var buf bytes.Buffer
	tpl, _ := template.ParseFiles("templates/pending.html")
	if err := tpl.Execute(&buf, input); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func GetFeedBackMailBody(input models.PayerEmailOutputDto) string {
	var buf bytes.Buffer
	tpl, _ := template.ParseFiles("templates/feedback.html")
	if err := tpl.Execute(&buf, input); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
