package model

import (
	"database/sql"
	"goodvs/server"

	"github.com/jackc/pgtype"
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
	Id            int64          `gorm:"column:id;type:int8;primaryKey;not null;" json:"id"`          // 商品ID, 唯一标识
	Name          string         `gorm:"column:name;type:varchar(100);not null;" json:"name"`         // 商品名称（精简描述）
	Category      string         `gorm:"column:category;type:varchar(100);not null;" json:"category"` // 一级分类
	Type          string         `gorm:"column:type;type:varchar(100);not null;" json:"type"`         // 二级分类
	Img           string         `gorm:"column:img;type:varchar(100);not null;" json:"img"`
	ProductPrices []ProductPrice `gorm:"foreignKey:ProductId;references:Id" json:"product_price"` // 商品价格表
}

// ProductPrice 商品价格表
type ProductPrice struct {
	// 以Product表的Id作为外键
	ProductId int64   `gorm:"column:product_id;type:int8;primaryKey;not null;" json:"product_id"` // 商品ID
	Price     float64 `gorm:"column:price;type:float8;not null;" json:"price"`
	Platform  string  `gorm:"column:platform;type:varchar(100);not null;" json:"platform"` // 商品来源平台
	Url       string  `gorm:"column:url;type:varchar(100);not null;" json:"url"`
	CreatedAt string  `gorm:"column:created_at;type:varchar(100);not null;" json:"created_at"`
}

type Symmetry struct {
	Id     int64  `gorm:"column:id;type:int8;primaryKey;not null;" json:"id"`
	Chirai string `gorm:"column:chirai;type:varchar;not null;" json:"chirai"`
	Pam    bool   `gorm:"column:pam;type:bool;not null;" json:"pam"`
	Name   string `gorm:"column:name;type:varchar(100);not null;" json:"name"`
}

type Material struct {
	Compoundname string         `gorm:"column:compoundname;type:varchar(100);not null;" json:"compoundname"`
	MpID         string         `gorm:"column:mp-ID;type:varchar(32);not null;" json:"mp-ID"`
	Uuid         string         `gorm:"column:uuid;type:varchar(32);primaryKey;not null;" json:"uuid"`
	Band         sql.NullString `gorm:"column:band;type:text;" json:"band"`
	Poscar       sql.NullString `gorm:"column:poscar;type:text;" json:"poscar"`
	Cif          sql.NullString `gorm:"column:cif;type:text;" json:"cif"`
	Nacdos       sql.NullString `gorm:"column:nacdos;type:text;" json:"nacdos"`
	SymmetryId   sql.NullInt64  `gorm:"column:symmetry_id;type:int8;" json:"symmetry_id"`
	Elements     pgtype.JSONB   `gorm:"column:elements;type:jsonb;not null;" json:"elements"`
	BandSoc      sql.NullString `gorm:"column:band_soc;type:text;" json:"band_soc"`
	PoscarSoc    sql.NullString `gorm:"column:poscar_soc;type:text;" json:"poscar_soc"`
	CifSoc       sql.NullString `gorm:"column:cif_soc;type:text;" json:"cif_soc"`
	NacdosSoc    sql.NullString `gorm:"column:nacdos_soc;type:text;" json:"nacdos_soc"`
	Chirai       string         `gorm:"column:chirai;type:varchar;not null;" json:"chirai"`
	Pam          bool           `gorm:"column:pam;type:bool;not null;" json:"pam"`
	SymmetryName string         `gorm:"column:symmetry_name;type:varchar(100);not null;" json:"symmetry_name"`
}

type Datas struct {
	Compoundname string        `gorm:"column:compoundname;type:varchar(100);not null;" json:"compoundname"`
	MpID         string        `gorm:"column:mp-ID;type:varchar(32);not null;" json:"mp-ID"`
	SymmetryId   sql.NullInt64 `gorm:"column:symmetry_id;type:int8;" json:"symmetry_id"`
	Elements     pgtype.JSONB  `gorm:"column:elements;type:jsonb;not null;" json:"elements"`
}

//加三类新的数据cif nac dos ，处理方法和Band Poscar一致
//mp_id band poscar cif nac dos chiral pam （这些是一起上传的，其他的没有）

func (material Material) Unmarshal(data *server.PostUploadJSONBody) (err error) {
	material = Material{
		MpID:   data.ID,
		Pam:    data.Pam,
		Chirai: data.ID,
	}
	material.Band.Scan(data.Data.Band)
	material.BandSoc.Scan(data.SOC)
	material.Cif.Scan(data.Data.Cif)
	material.CifSoc.Scan(data.SOC)
	material.Poscar.Scan(data.Data.Poscar)
	material.PoscarSoc.Scan(data.SOC)
	material.Nacdos.Scan(data.Data.Nacdos)
	material.NacdosSoc.Scan(data.SOC)
	//material.dos.Scan(data.Data.Dos)
	return
}
