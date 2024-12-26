package dao

import (
	"fmt"
	"goodvs/internal/dao/model"
	"goodvs/server"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSONBQueryExp struct {
	column   string
	keys     []string
	allIn    bool
	oneIn    bool
	notAllIn bool
	notOneIn bool
}

func JSONBQuery(column string) *JSONBQueryExp {
	return &JSONBQueryExp{column: column}
}

func (jsonbQuery *JSONBQueryExp) AllIn(keys []string) *JSONBQueryExp {
	jsonbQuery.keys = keys
	jsonbQuery.allIn = true
	return jsonbQuery
}

func (jsonbQuery *JSONBQueryExp) OneIn(keys []string) *JSONBQueryExp {
	jsonbQuery.keys = keys
	jsonbQuery.oneIn = true
	return jsonbQuery
}

func (jsonbQuery *JSONBQueryExp) NotAllIn(keys []string) *JSONBQueryExp {
	jsonbQuery.keys = keys
	jsonbQuery.notAllIn = true
	return jsonbQuery
}

func (jsonbQuery *JSONBQueryExp) NotOneIn(keys []string) *JSONBQueryExp {
	jsonbQuery.keys = keys
	jsonbQuery.notOneIn = true
	return jsonbQuery
}

func (jsonbQuery *JSONBQueryExp) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		if jsonbQuery.allIn {
			builder.WriteString(jsonbQuery.column + " ?& ARRAY[")
		} else if jsonbQuery.oneIn {
			builder.WriteString(jsonbQuery.column + " ?| ARRAY[")
		} else if jsonbQuery.notAllIn {
			builder.WriteString("NOT " + jsonbQuery.column + " ?& ARRAY[")
		} else if jsonbQuery.notOneIn {
			builder.WriteString("NOT " + jsonbQuery.column + " ?| ARRAY[")
		}
		for index, key := range jsonbQuery.keys {
			if index != 0 {
				builder.WriteString(",")
			}
			builder.AddVar(stmt, key)
		}
		builder.WriteString("]")
	}
}

func (db DBMS) FilterMaterial(params server.Filter, page int, ordered *string, asc bool) (err error, total int, results []model.Material) {
	filter := db.Model(&model.Material{})
	if params.MpID != nil {
		filter = filter.Where(&model.Material{
			MpID: *params.MpID,
		})
	}
	if params.Pam != nil {
		filter = filter.Where("materials.pam = ?", *params.Pam)
		//filter = filter.Joins("JOIN symmetry ON symmetry.id = materials.symmetry_id").Where("symmetry.pam = ?", *params.Pam)
	}
	if params.Type != nil {
		filter = filter.Where("materials.chirai = ?", string(*params.Type))
		//filter = filter.Joins("JOIN symmetry ON symmetry.id = materials.symmetry_id").Where("symmetry.chirai = ?", string((*params.Type).Chiral))
	}
	if params.Symmetry != nil {
		filter = filter.Where("symmetry_id IN ?", *params.Symmetry)
	}
	if params.Elements != nil {
		if !params.Precisely {
			var elementsHas []string
			var elementsFrac []server.Element
			for _, element := range *params.Elements {
				if element.Number == -1 {
					elementsHas = append(elementsHas, element.Name)
				} else {
					elementsFrac = append(elementsFrac, element)
				}
			}
			if len(elementsFrac) == 1 {
				elementsHas = append(elementsHas, elementsFrac[0].Name)
			}
			if len(elementsHas) > 0 {
				filter = filter.Where(JSONBQuery("elements").AllIn(elementsHas))
			}
			if len(elementsFrac) > 1 {
				for index, element := range elementsFrac {
					element1Name := element.Name
					element1Number := element.Number
					element2Name := elementsFrac[(index+1)%len(elementsFrac)].Name
					element2Number := elementsFrac[(index+1)%len(elementsFrac)].Number
					filter = filter.Where("(materials.elements->>?)::int * ? = (materials.elements->>?)::int * ?", element1Name, element2Number, element2Name, element1Number)
				}
			}
		} else {
			for _, element := range *params.Elements {
				filter = filter.Where("elements->>? = ?", element.Name, element.Number)
			}
		}
	}
	if params.ElementsExclude != nil {
		var elements []string
		for _, element := range *params.ElementsExclude {
			elements = append(elements, element.Name)
		}
		filter = filter.Where(JSONBQuery("elements").NotOneIn(elements))
	}
	if ordered != nil {
		if asc {
			filter = filter.Order(*ordered + " ASC")
		} else {
			filter = filter.Order(*ordered + " DESC")
		}
	}
	err = filter.Offset((page - 1) * 20).Limit(20).Find(&results).Error
	total = len(results)
	return err, total, results
}

func (db DBMS) AddMaterial(material model.Material) (err error) {
	err = db.Create(&material).Error
	return err
}

func (db DBMS) GetMaterialByUUID(uuid string) (err error, material model.Material) {
	err = db.Where(&model.Material{
		Uuid: uuid,
	}).First(&material).Error
	return err, material
}

func (db DBMS) GetMaterialByMpID(mpID string) (err error, material model.Material) {
	err = db.Where(&model.Material{
		MpID: mpID,
	}).First(&material).Error
	return err, material
}

func (db DBMS) GetSymmetryByID(id int64) (err error, symmetry model.Symmetry) {
	err = db.Where(&model.Symmetry{
		Id: symmetry.Id,
	}).First(&symmetry).Error
	return err, symmetry
}

func (db DBMS) PostUpload(uploads []server.PostUploadJSONBody) (err error) {
	for _, result := range uploads {
		newMaterial := model.Material{}
		newMaterial.Unmarshal(&result)
		newData := model.Datas{}
		err = db.Model(&model.Datas{}).Where(&model.Datas{
			MpID: newMaterial.MpID,
		}).First(&newData).Error
		if err != nil {
			logrus.Error("fail to find the compound in table Datas")
			return
		}
		newMaterial.Compoundname = newData.Compoundname
		newMaterial.SymmetryId = newData.SymmetryId
		newMaterial.Elements = newData.Elements
		err = db.AddMaterial(newMaterial)
		if err != nil {
			logrus.Error("fail to create in table Material")
			return
		}
	}
	return
}

// ---------------

// AddUser create a new user
func (db DBMS) AddUser(user server.UserRegisterReq) (token string, err error) {
	var tmp model.User
	// check if user already exist
	err = db.Where(&model.User{
		Email: user.Email,
	}).First(&tmp).Error
	if tmp != (model.User{}) || err == nil {
		return "", fmt.Errorf("user already exist")
	}
	result := model.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	// create new user
	err = db.Create(&result).Error
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(result.Id, 10), nil
}

// ValidateUser validate user and return token
func (db DBMS) ValidateUser(user server.UserLoginReq) (token string, err error) {
	var result model.User
	err = db.Where(&model.User{
		Email:    user.Email,
		Password: user.Password,
	}).First(&result).Error
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(result.Id, 10), nil
}

// AddProductItem create a new product
func (db DBMS) AddProductItem(product model.Product) (err error) {
	// check if product already exist
	var tmp model.Product
	err = db.Where(&model.Product{
		Name: product.Name,
	}).First(&tmp).Error
	if err == nil {
		return fmt.Errorf("product already exist")
	}
	err = db.Create(&product).Error
	return err
}

// GetProductItem get product by id
func (db DBMS) GetProductItem(id int64) (err error, product model.Product) {
	err = db.Where(&model.Product{
		Id: id,
	}).First(&product).Error
	return err, product
}

// PutProductPriceList get product list
func (db DBMS) PutProductPriceList(productPriceList []model.ProductPrice) (err error) {
	for _, productPrice := range productPriceList {
		err = db.Create(&productPrice).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPriceListByProductID get price list by product id
func (db DBMS) GetPriceListByProductID(productID int64) (err error, priceList []model.ProductPrice) {
	err = db.Where(&model.Product{
		Id: productID,
	}).First(&model.Product{}).Error
	if err != nil {
		return err, nil
	}
	err = db.Where(&model.ProductPrice{
		ProductId: productID,
	}).Find(&priceList).Error
	return err, priceList
}
