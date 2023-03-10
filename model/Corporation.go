package model

import (
	"fmt"
	"userservice/database"
	"userservice/schema"
)

type Corporation schema.Corporation
type UpdateCorporation schema.UpdateCorporation

func (emp *Corporation) Save() (*Corporation, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindCorporateAll() ([]Corporation, error) {
	var emp []Corporation
	err := database.Database.Find(&emp).Error
	fmt.Println(emp)
	return emp, err
}

func FindCorporateById(id string) (Corporation, error) {
	var emp Corporation
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}

func FindCorporateByName(id string) (Corporation, error) {
	var emp Corporation
	err := database.Database.Where("name=?", id).First(&emp).Error
	return emp, err
}

func FindCorporateByCode(id string) (Corporation, error) {
	var emp Corporation
	err := database.Database.Where("code=?", id).First(&emp).Error
	return emp, err
}

func FindCorporationMapById(params []string) ([]Corporation, error) {
	var corp []Corporation
	err := database.Database.Where("id IN ?", params).Find(&corp).Error
	fmt.Println(corp)
	return corp, err
}
func FindCorporationMapByName(params []string) ([]Corporation, error) {
	var corp []Corporation
	err := database.Database.Where("name IN ?", params).Find(&corp).Error
	fmt.Println(corp)
	return corp, err
}

func (update_data *Corporation) ChangeData(id string, ua UpdateCorporation) (Corporation, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindCorporateById(id)
	return res, nil
}
