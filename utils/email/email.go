package email

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)
import "gopkg.in/gomail.v2"

func SendMail(mailTo, subject, body string) error {
	defer func() {
		if p := recover(); p != nil {
			logs.Error("send email recover error:", p)
		}
	}()
	host := beego.AppConfig.String("email_host")
	portStr := beego.AppConfig.String("email_port")
	senderName := beego.AppConfig.String("email_senderName")
	userName := beego.AppConfig.String("email_user")
	password := beego.AppConfig.String("email_password")
	port, _ := strconv.Atoi(portStr)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, senderName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, senderName, password)
	err := d.DialAndSend(m)
	if err != nil {
		logs.Error("send email error:" + err.Error())
	}
	return err
}
