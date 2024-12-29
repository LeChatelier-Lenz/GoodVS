// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

//import "goodvs/internal/dao/model"

// Defines values for BriefSearchResType.
const (
	BriefSearchResTypeAchiral BriefSearchResType = "achiral"
	BriefSearchResTypeChiral  BriefSearchResType = "chiral"
)

// Defines values for FilterType.
const (
	FilterTypeAchiral FilterType = "achiral"
	FilterTypeChiral  FilterType = "chiral"
)

// Defines values for TypeChiral.
const (
	Achiral TypeChiral = "achiral"
	Chiral  TypeChiral = "chiral"
)

// BriefSearchReq defines model for BriefSearchReq.
type BriefSearchReq = Filter

// BriefSearchRes defines model for BriefSearchRes.
type BriefSearchRes struct {
	Symmetry *[]int `json:"symmetry"`
	Total    int    `json:"total"`

	// Type 绑定sg
	Type *BriefSearchResType `json:"type"`
}

// BriefSearchResType 绑定sg
type BriefSearchResType string

// DetailsRes defines model for DetailsRes.
type DetailsRes = Material

// Element defines model for Element.
type Element struct {
	Name string `json:"name"`

	// Number 搜索时为比例系数，-1表示无限制
	Number int `json:"number"`
}

// Filter defines model for Filter.
type Filter struct {
	Elements        *[]Element `json:"elements"`
	ElementsExclude *[]Element `json:"elements exclude"`

	// MpID 最多9位数字，无前导0
	MpID *string `json:"mp-ID"`
	Pam  *bool   `json:"pam"`

	// Precisely 精确搜索
	Precisely bool   `json:"precisely"`
	Symmetry  *[]int `json:"symmetry"`

	// Type 绑定sg
	Type *FilterType `json:"type"`
}

// FilterType 绑定sg
type FilterType string

// Material defines model for Material.
type Material struct {
	Compound []Element `json:"compound"`

	// CompoundName compound 的string版
	CompoundName string `json:"compound_name"`

	// MpID mp-开头的字符串，最多9位，储存时不存mp-
	MpID string `json:"mp-ID"`

	// Pam 有无PAM定义
	Pam      bool `json:"pam"`
	Symmetry int  `json:"symmetry"`
	Type     Type `json:"type"`

	// Uuid 仅在我们的数据库中使用
	Uuid string `json:"uuid"`
}

//// SearchReq defines model for SearchReq.
//type SearchReq struct {
//	// Asc 升序 or 降序
//	Asc     bool    `json:"asc"`
//	Filter  Filter  `json:"filter"`
//	Ordered *string `json:"ordered,omitempty"`
//
//	// Page 默认为1,每一页20个结果
//	Page *int `json:"page,omitempty"`
//}
//
//// SearchRes defines model for SearchRes.
//type SearchRes struct {
//	Results []Material `json:"results"`
//	Total   int        `json:"total"`
//}

// Type defines model for Type.
type Type struct {
	// Chiral 绑定sg
	Chiral TypeChiral `json:"chiral"`
}

// TypeChiral 绑定sg
type TypeChiral string

// GetMaterialBandParams defines parameters for GetMaterialBand.
type GetMaterialBandParams struct {
	Id string `form:"id" json:"id"`

	// SOC 是否考虑修正
	SOC int `form:"SOC" json:"SOC"`
}

// GetMaterialCifParams defines parameters for GetMaterialCif.
type GetMaterialCifParams struct {
	Id string `form:"id" json:"id"`

	// SOC 是否考虑修正
	SOC int `form:"SOC" json:"SOC"`
}

// GetMaterialDetailParams defines parameters for GetMaterialDetail.
type GetMaterialDetailParams struct {
	Id string `form:"id" json:"id"`
}

// GetMaterialNacdosParams defines parameters for GetMaterialNacdos.
type GetMaterialNacdosParams struct {
	Id string `form:"id" json:"id"`

	// SOC 是否考虑修正
	SOC int `form:"SOC" json:"SOC"`
}

// GetMaterialPoscarParams defines parameters for GetMaterialPoscar.
type GetMaterialPoscarParams struct {
	Id string `form:"id" json:"id"`

	// SOC 是否考虑修正
	SOC int `form:"SOC" json:"SOC"`
}

// PostUploadJSONBody defines parameters for PostUpload.
type PostUploadJSONBody struct {
	ID     string `json:"ID"`
	SOC    string `json:"SOC"`
	Chirai string `json:"chirai"`
	Data   struct {
		Band   string `json:"band"`
		Cif    string `json:"cif"`
		Dos    string `json:"dos"`
		Nacdos string `json:"nacdos"`
		Poscar string `json:"poscar"`
	} `json:"data"`
	Pam bool `json:"pam"`
}

// PostSearchBriefJSONRequestBody defines body for PostSearchBrief for application/json ContentType.
type PostSearchBriefJSONRequestBody = BriefSearchReq

// PostSearchResultJSONRequestBody defines body for PostSearchResult for application/json ContentType.
//type PostSearchResultJSONRequestBody = SearchReq

// PostUploadJSONRequestBody defines body for PostUpload for application/json ContentType.
type PostUploadJSONRequestBody PostUploadJSONBody


type UserRegisterReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UserLoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SearchReq struct {
	ProductStr string `json:"product"`
}

type ProductByCraw struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Url 	 string  `json:"url"`
	ImgUrl   string  `json:"img_url"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Platform string  `json:"platform"`

}


//type ProductByCraw struct {
//	Id string `json:"id"`
//	Name string `json:"name"`
//	ImgUrl string `json:"img_url"`
//	Price float64 `json:"price"`
//	Category string `json:"category"`
//	Platform string `json:"platform"`
//}


type SearchRes struct {
	Results []ProductByCraw `json:"results"`
	Msg string `json:"msg"`
}