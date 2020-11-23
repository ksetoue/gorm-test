package Models

import (
	"fmt"
	"gorm-test/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllDeals Fetch all user data
func GetAllDeals(deal *[]Deal) (err error) {
	if err = Config.DB.Find(deal).Error; err != nil {
		return err
	}
	return nil
}

//CreateDeal ... Insert New data
func CreateDeal(deal *Deal) (err error) {
	if err = Config.DB.Create(deal).Error; err != nil {
		return err
	}
	return nil
}

//GetDealByID ... Fetch only one Deal by Id
func GetDealByID(deal *Deal, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(deal).Error; err != nil {
		return err
	}
	return nil
}

//UpdateDeal ... Update Deal
func UpdateDeal(deal *Deal, id string) (err error) {
	fmt.Println(deal)
	Config.DB.Save(deal)
	return nil
}

//DeleteDeal ... Delete Deal
func DeleteDeal(deal *Deal, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(deal)
	return nil
}
