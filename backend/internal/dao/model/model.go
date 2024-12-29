package model

import (
	"goodvs/server"
	"time"
)

// User 用户表
type User struct {
	Id       int64  `gorm:"column:id;type:int8;primaryKey;not null;" json:"id"`
	Name     string `gorm:"column:name;type:varchar(100);not null;" json:"name"`
	Password string `gorm:"column:password;type:varchar(100);not null;" json:"password"`
	Email    string `gorm:"column:email;type:varchar(100);not null;" json:"email"`
}

// Product 商品表
type Product struct {
	Id            string         `gorm:"column:id;type:varchar(100);primaryKey;not null;" json:"id"`  // 商品ID
	Name          string         `gorm:"column:name;type:varchar(100);not null;" json:"name"`         // 商品名称（精简描述）
	Detail        string         `gorm:"column:detail;type:varchar(100);not null;" json:"detail"`     // 商品详情
	Category      string         `gorm:"column:category;type:varchar(100);not null;" json:"category"` // 一级分类
	Type          string         `gorm:"column:type;type:varchar(100);" json:"type"`                  // 二级分类
	ImgUrl        string         `gorm:"column:img_url;type:varchar(400);not null;" json:"img_url"`   // 商品图片链接
	Url           string         `gorm:"column:url;type:varchar(400);not null;" json:"url"`           // 商品链接
	Platform      string         `gorm:"column:platform;type:varchar(100);not null;" json:"platform"` // 商品来源平台
	ProductPrices []ProductPrice `gorm:"foreignKey:ProductId;references:Id" json:"product_price"`     // 商品价格表
}

// ProductPrice 商品价格表
type ProductPrice struct {
	// 以Product表的Id作为外键
	ProductId string    `gorm:"column:product_id;type:varchar(100);not null;" json:"product_id"` // 商品ID
	Price     float64   `gorm:"column:price;type:float8;not null;" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
}

// Follow 关注表
type Follow struct {
	Id        int64  `gorm:"column:id;type:int8;primaryKey;not null;" json:"id"`
	ProductId string `gorm:"column:product_id;type:varchar(100);not null;" json:"product_id"`
	UserId    int64  `gorm:"column:user_id;type:int8;not null;" json:"user_id"`
}

// Unmarshal 将ProductByCraw转换为Product
func (product Product) Unmarshal(data *server.ProductByCraw) {
	product = Product{
		Id:       data.Id,
		Name:     data.Name,
		Category: data.Category,
		Detail:   data.Title,
		ImgUrl:   data.ImgUrl,
		Url:      data.Url,
		Platform: data.Platform,
	}
}

// Marshal 将Product转换为ProductByCraw
func (product Product) Marshal() (data server.ProductByCraw) {
	data = server.ProductByCraw{
		Id:       product.Id,
		Name:     product.Name,
		Url:      product.Url,
		ImgUrl:   product.ImgUrl,
		Title:    product.Detail,
		Category: product.Category,
		Platform: product.Platform,
	}
	return data
}
