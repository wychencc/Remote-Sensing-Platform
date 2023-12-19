package model

import (
	"github.com/jinzhu/gorm"
	"strings"
	"version1_copy/infer"
	"version1_copy/utils/errmsg"
)

type Inference struct {
	gorm.Model
	RgbImageName string `gorm:"type:varchar(40);not null " json:"rgb_image_name" form:"rgb_image_name"`
	SarImageName string `gorm:"type:varchar(40);not null " json:"sar_image_name" form:"sar_image_name"`
	ResultName   string `gorm:"type:varchar(40);not null " json:"result_name" form:"result_name"`
	ResultPath   string `gorm:"type:varchar(200);not null" json:"result_path" form:"result_path"`
	Uid          int    `gorm:"type:int;not null" json:"uid" form:"uid"`
}

//func GetModelList(typename string) ([]string, int) {
//	var model_name []string
//	err := db.Group("model_name").Select("model_name").Where("type = ?", typename).Find(&ModelZoo{}).Pluck("model_name", &model_name).Error
//	if err != nil {
//		return nil, errmsg.ERROR
//	}
//	return model_name, errmsg.SUCCESS
//}

//func GetDataSetList(type_name, model_name string) ([]string, int) {
//	var data_name []string
//	err := db.Group("dataset_name").Select("dataset_name").Where("type = ? and model_name = ?",
//		type_name, model_name).Find(&ModelZoo{}).Pluck("dataset_name", &data_name).Error
//	if err != nil {
//		return nil, errmsg.ERROR
//	}
//	return data_name, errmsg.SUCCESS
//}

//func GetModelPath(model_name, dataset_name string) (string, int) {
//	var modelzoo ModelZoo
//	err := db.Where("model_name = ? and dataset_name = ?", model_name, dataset_name).First(&modelzoo).Error
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	return modelzoo.ModelPath, errmsg.SUCCESS
//}

// AddInference 添加推理任务
func AddInference(rgb, sar, result, result_path string) int {
	inference := Inference{
		RgbImageName: rgb,
		SarImageName: sar,
		ResultName:   result,
		ResultPath:   result_path,
		Uid:          1,
	}
	err := db.Create(&inference).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func UpdateInference(rgb, sar, result, result_path string) int {
	var infer Inference
	inference := Inference{
		RgbImageName: rgb,
		SarImageName: sar,
		ResultName:   result,
		ResultPath:   result_path,
		Uid:          1,
	}
	err := db.Model(&infer).Where("rgb_image_name = ?", inference.RgbImageName).Updates(&inference).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetResult(rgb_path, sar_path, rgb_name, save_path string) (result_name, result_path string) {
	result_name = "result_" + strings.Split(rgb_name, ".")[0] + ".png"
	result_path = "./" + save_path + result_name
	sar_path = "./" + sar_path
	rgb_path = "./" + rgb_path
	infer.GetInferenceResult(sar_path, rgb_path, result_path)
	return
}

func CheckImageInfer(image_name string, uid int) int {
	var inference Inference
	image_name = "result_" + strings.Split(image_name, ".")[0] + ".png"
	db.Select("id").Where("result_name = ? and uid = ?", image_name, uid).First(&inference)
	if inference.ID > 0 {
		return errmsg.ERROR_IMAGE_EXIT
	}
	return errmsg.SUCCESS
}
