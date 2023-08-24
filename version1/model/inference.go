package model

import (
	"github.com/jinzhu/gorm"
	"version1/utils/errmsg"
)

type Inference struct {
	gorm.Model
	ImageName  string `gorm:"type:varchar(40);not null " json:"image_name" form:"image_name"`
	ResultName string `gorm:"type:varchar(40);not null " json:"result_name" form:"result_name"`
	ImagePath  string `gorm:"type:varchar(100);not null" json:"image_path" form:"image_path"`
	ResultPath string `gorm:"type:varchar(100);not null" json:"result_path" form:"result_path"`
	Uid        int    `gorm:"type:int;not null" json:"uid" form:"uid"`
}

func GetModelList(typename string) ([]string, int) {
	var model_name []string
	err := db.Group("model_name").Select("model_name").Where("type = ?", typename).Find(&ModelZoo{}).Pluck("model_name", &model_name).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return model_name, errmsg.SUCCESS
}

func GetDataSetList(type_name, model_name string) ([]string, int) {
	var data_name []string
	err := db.Group("dataset_name").Select("dataset_name").Where("type = ? and model_name = ?",
		type_name, model_name).Find(&ModelZoo{}).Pluck("dataset_name", &data_name).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return data_name, errmsg.SUCCESS
}

func GetModelPath(model_name, dataset_name string) (string, int) {
	var modelzoo ModelZoo
	err := db.Where("model_name = ? and dataset_name = ?", model_name, dataset_name).First(&modelzoo).Error
	if err != nil {
		return "", errmsg.ERROR
	}
	return modelzoo.ModelPath, errmsg.SUCCESS
}

// GetResult 执行推理任务接口
func GetResult(model_path, image_name, img_save_path, result_save_path string) (string, string) {

	return result_save_path + "result_" + image_name, "result_" + image_name
}

// AddInference 添加推理任务
func AddInference(image_name, result_name, image_path, result_path string) int {
	inference := &Inference{
		ImageName:  image_name,
		ResultName: result_name,
		ImagePath:  image_path,
		ResultPath: result_path,
		Uid:        1,
	}
	err := db.Create(&inference).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
