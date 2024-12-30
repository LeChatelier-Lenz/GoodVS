package service

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"goodvs/server"
	"log"
	"net/smtp"
	"strconv"
)

type EmailService struct {
	email *email.Email
	cfg   EmailCfg
}

type EmailCfg struct {
	From string `mapstructure:"From"`
	Pass string `mapstructure:"Pass"`
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

var (
	ES EmailService
)

func InitES() {
	var cfg EmailCfg
	err := viper.Sub("EmailService").UnmarshalExact(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling email config")
		fmt.Println(err)
	}
	fmt.Println("cfg: ", cfg)
	ES = *NewEmailService(cfg)
}

func NewEmailService(C EmailCfg) *EmailService {
	e := email.NewEmail()
	e.From = "GoodVS <" + C.From + ">"
	e.Subject = "您的关注商品降价提醒"
	return &EmailService{
		email: e,
		cfg:   C,
	}
}

// SendEmail send email to target
func (es *EmailService) SendEmail(req server.EmailReq) (err error) {
	// 设置接收方的邮箱
	es.email.To = []string{req.Target}
	fmt.Println("from: ", es.email.From)
	fmt.Println("to: ", es.email.To)

	//设置主题
	//设置文件发送的内容
	oldPriceStr := strconv.FormatFloat(req.OldPrice, 'f', 2, 64)
	newPriceStr := strconv.FormatFloat(req.NewPrice, 'f', 2, 64)
	es.email.HTML = []byte(`
		<!DOCTYPE html>
		<html lang="zh">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>降价提醒</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f7f7f7;
					margin: 0;
					padding: 0;
				}
				.email-container {
					width: 100%;
					max-width: 600px;
					margin: 20px auto;
					background-color: #ffffff;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					padding: 20px;
				}
				.email-header {
					text-align: center;
					padding-bottom: 20px;
					border-bottom: 2px solid #f0f0f0;
				}
				.email-header h1 {
					color: #333;
					font-size: 24px;
					margin: 0;
				}
				.product-info {
					padding: 20px 0;
					text-align: center;
				}
				.product-info img {
					max-width: 200px;
					margin-bottom: 10px;
				}
				.product-info h2 {
					font-size: 20px;
					margin-bottom: 10px;
					color: #333;
				}
				.product-info p {
					font-size: 16px;
					color: #777;
					margin-bottom: 5px;
				}
				.price {
					font-size: 18px;
					color: #e74c3c;
					font-weight: bold;
				}
				.old-price {
					text-decoration: line-through;
					color: #aaa;
					margin-left: 10px;
				}
				.discount {
					color: #2ecc71;
					font-weight: bold;
				}
				.cta-button {
					display: block;
					width: 100%;
					padding: 15px;
					background-color: #3498db;
					color: #ffffff;
					text-align: center;
					text-decoration: none;
					border-radius: 5px;
					font-size: 18px;
					margin-top: 20px;
				}
				.cta-button:hover {
					background-color: #2980b9;
				}
				.footer {
					text-align: center;
					font-size: 12px;
					color: #888;
					margin-top: 30px;
				}
			</style>
		</head>
		<body>
			<div class="email-container">
				<div class="email-header">
					<h1>您的降价提醒</h1>
				</div>
		
				<div class="product-info">
					<img src="` + req.ImageUrl + `" alt="产品图片">
					<h2>` + req.ProductName + `</h2>
					<p>原价: <span class="old-price">` + oldPriceStr + `</span></p>
					<p>现价: <span class="price">` + newPriceStr + `</span></p>
				</div>
				<a href="` + req.Url + `" class="cta-button">立即查看商品</a>
				<div class="footer">
					<p>如果您想到我们的网页了解更多信息，<a href="https://localhost:5173" style="color: #3498db;">请点击</a></p>
				</div>
			</div>
		</body>
		</html>
    `)
	//设置服务器相关的配置
	err = es.email.Send(es.cfg.Host+":"+es.cfg.Port, smtp.PlainAuth("", es.cfg.From, es.cfg.Pass, es.cfg.Host))
	if err != nil {
		log.Println("send email failed ", err)
		return err
	}
	return nil
}
