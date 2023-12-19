package model

import (
	"github.com/jinzhu/gorm"
	"version1_copy/utils/errmsg"
)

type ModelZoo struct {
	gorm.Model
	ModelName   string `gorm:"type:varchar(40);not null " json:"model_name" form:"model_name"`
	DatasetName string `gorm:"type:varchar(40);not null " json:"dataset_name" form:"dataset_name"`
	ModelPath   string `gorm:"type:varchar(100);not null" json:"model_path" form:"model_path"`
	Type        string `gorm:"type:varchar(20);not null" json:"type" form:"type"`
}

// CheckModel 检查模型是否存在
func CheckModel(modelname, datasetname string) int {
	var model ModelZoo
	db.Select("id").Where("model_name = ? and dataset_name = ?", modelname, datasetname).First(&model)
	if model.ID > 0 {
		return errmsg.ERROR_MODEL_EXISTED
	}
	return errmsg.SUCCESS
}

func AddModel(model *ModelZoo) int {
	err := db.Create(&model).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ModelZooList 获取模型列表
func ModelZooList() ([]ModelZoo, int) {
	var modelzoolst []ModelZoo
	err := db.Find(&modelzoolst).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return modelzoolst, errmsg.SUCCESS
}
