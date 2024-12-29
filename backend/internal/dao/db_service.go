package dao

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goodvs/internal/dao/model"
	"goodvs/server"
	"strconv"
	"time"
)

// AddUser create a new user
func (db DBMS) AddUser(user server.UserRegisterReq) (token string, err error) {
	var tmp []model.User
	// check if user already exist
	users := db.Where(&model.User{
		Name:  user.Name,
		Email: user.Email,
	}).Find(&tmp)
	if users.RowsAffected != 0 || users.Error == nil {
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
		Email: user.Email,
		//Password: user.Password,
	}).First(&result).Error
	if err != nil {
		thisErr := fmt.Errorf("user not exist")
		return "", thisErr
	}
	if result.Password != user.Password {
		thisErr := fmt.Errorf("password not correct")
		return "", thisErr
	}
	return strconv.FormatInt(result.Id, 10), nil
}

// AddProductItem purely add a new product item into database
func (db DBMS) AddProductItem(products server.ProductByCraw) (productId string, err error) {
	var product model.Product
	// check if product already exist
	result := db.Where(&model.Product{
		Name: products.Name,
		Id:   products.Id,
	}).Find(&product)
	if result.Error == nil && result.RowsAffected != 0 {
		// product already exist
		logrus.Info("product already exist")
		return "###", fmt.Errorf("product already exist")
	}
	newProduct := model.Product{
		Id:       products.Id,
		Name:     products.Name,
		Url:      products.Url,
		ImgUrl:   products.ImgUrl,
		Detail:   products.Title,
		Category: products.Category,
		Platform: products.Platform,
	}
	err = db.Create(&newProduct).Error
	if err != nil {
		/// logrus.Error("fail to create in table Product")
		return productId, err
	}
	return newProduct.Id, nil
}

// AddProductPrice add a new price into database
func (db DBMS) AddProductPrice(productID string, price float64) (err error) {
	newProductPrice := model.ProductPrice{
		ProductId: productID,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err = db.Create(&newProductPrice).Error
	if err != nil {
		//logrus.Error("fail to create in table ProductPrice")
		return err
	}
	return nil
}

// GetProductItem get product item by name and id
func (db DBMS) GetProductItem(name string, id string) (product model.Product, err error) {
	err = db.Where(&model.Product{
		Name: name,
		Id:   id,
	}).First(&product).Error
	return product, err
}

// GetProductItemByID get product
func (db DBMS) GetProductItemByID(productID string) (product model.Product, err error) {
	err = db.Where(&model.Product{
		Id: productID,
	}).First(&product).Error
	return product, err
}

// GetProductPriceList get price list by product id
func (db DBMS) GetProductPriceList(productID string) (err error, priceList []model.ProductPrice) {
	err = db.Where(&model.Product{
		Id: productID,
	}).First(&model.Product{}).Error
	if err != nil {
		logrus.Error("fail to find product")
		return err, nil
	}
	// get price list，根据商品ID获取价格列表，按照时间排序
	err = db.Order("created_at").Where(&model.ProductPrice{
		ProductId: productID,
	}).Find(&priceList, "created_at").Error
	if len(priceList) == 0 {
		return fmt.Errorf("no price list"), nil
	}
	return err, priceList
}

func (db DBMS) GetLatestProductPrice(productID string) (price float64, err error) {
	var productPrice model.ProductPrice
	err = db.Where(&model.ProductPrice{
		ProductId: productID,
	}).Order("created_at desc").First(&productPrice).Error
	if err != nil {
		return -1, err
	}
	return productPrice.Price, nil
}

// AddFollow add a new follow
func (db DBMS) AddFollow(productID string, userID int64) (err error) {
	// check if already follow
	var follow model.Follow
	result := db.Where(&model.Follow{
		ProductId: productID,
		UserId:    userID,
	}).Find(&follow)
	if result.Error == nil && result.RowsAffected != 0 {
		// already follow
		return fmt.Errorf("already follow")
	}
	newFollow := model.Follow{
		ProductId: productID,
		UserId:    userID,
	}
	err = db.Create(&newFollow).Error
	if err != nil {
		logrus.Error("fail to create in table Follow")
		return err
	}
	return nil
}

// GetUserFollowList get user follow list
func (db DBMS) GetUserFollowList(userID int64) (productList []server.ProductByCraw, err error) {
	var followList []model.Follow
	userSearch := db.Where(&model.User{
		Id: userID,
	}).Find(&model.User{})
	if userSearch.RowsAffected == 0 {
		return nil, fmt.Errorf("user not exist")
	}
	err = db.Where(&model.Follow{
		UserId: userID,
	}).Find(&followList).Error
	if len(followList) == 0 {
		return nil, fmt.Errorf("you have not follow any product")
	}
	for _, follow := range followList {
		product, err := db.GetProductItemByID(follow.ProductId)
		if err != nil {
			logrus.Error("fail to get product item")
			return nil, err
		}
		resp := product.Marshal()
		resp.Price, err = db.GetLatestProductPrice(follow.ProductId)
		if err != nil || resp.Price == -1 {
			logrus.Error("fail to get latest price")
			return nil, err
		}
		productList = append(productList, resp)
	}
	return productList, err
}

// RemoveFollow remove a follow
func (db DBMS) RemoveFollow(productID string, userID int64) (err error) {
	err = db.Where(&model.Follow{
		ProductId: productID,
		UserId:    userID,
	}).Delete(&model.Follow{}).Error
	if err != nil {
		logrus.Error("fail to delete in table Follow")
		return err
	}
	return err
}

//// PutProductPriceList get product list
//func (db DBMS) PutProductPriceList(productPriceList []server.ProductByCraw) (err error) {
//	// 先添加商品项条目
//	productID, err := db.AddProductItem(productPriceList)
//	if err != nil {
//		logrus.Error("fail to add product item")
//		return err
//	}
//	if productID == "" {
//		return fmt.Errorf("
//	}
//	// 无论是否已经存在，都会获得商品ID，接下来添加价格列表
//	for _, product := range productPriceList {
//		newProductPrice := model.ProductPrice{
//			ProductId: productID,
//			Price:     product.Price,
//			//Url:       product.ImgUrl,
//			CreatedAt: time.Now(),
//		}
//		err = db.Create(&newProductPrice).Error
//		if err != nil {
//			logrus.Error("fail to create in table ProductPrice")
//			return err
//		}
//	}
//	return nil
//}

//// GetPriceListByProductID get price list by product id
//func (db DBMS) GetPriceListByProductID(productID int64) (err error, priceList []model.ProductPrice) {
//	err = db.Where(&model.Product{
//		Id: productID,
//	}).First(&model.Product{}).Error
//	if err != nil {
//		return err, nil
//	}
//	// get price list，根据商品ID获取价格列表，按照时间排序
//	err = db.Where(&model.ProductPrice{
//		ProductId: productID,
//	}).Find(&priceList).Error
//	return err, priceList
//}
