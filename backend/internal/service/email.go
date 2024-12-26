package service

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(target string) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "GoodVS <826733088@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{target}
	//设置主题
	e.Subject = "您的关注商品降价提醒"
	//设置文件发送的内容
	e.HTML = []byte(`
    <h2><a href="http://localhost:5173">点击查看</a></h2>    
    `)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "826733088@qq.com", "ofzjlifjolcybcbb", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
