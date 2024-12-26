package model

import (
	"database/sql"
	"goodvs/server"

	"github.com/jackc/pgtype"
)

type User struct {
	Id       int64  `gorm:"column:id;type:int8;primaryKey;not null;" json:"id"`
	Name     string `gorm:"column:username;type:varchar(100);not null;" json:"name"`
	Password string `gorm:"column:password;type:varchar(100);not null;" json:"password"`
	Email    string `gorm:"column:email;type:varchar(100);not null;" json:"email"`
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
