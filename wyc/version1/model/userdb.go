package model

import (
	"github.com/jinzhu/gorm"
	"version1_copy/utils/errmsg"
)

type UserDb struct {
	gorm.Model
	Image_Name string `gorm:"type:varchar(40);not null" json:"image_name" form:"image_name"`
	Save_Path  string `gorm:"type:varchar(100);not null" json:"save_path" form:"save_path"`
	Uid        int    `gorm:"type:int;not null" json:"uid" form:"uid"`
}

func GetGeeList() ([]string, []string, int) {
	var gee_names []string
	var gee_paths []string
	err1 := db.Group("image_name").Select("image_name").Find(&Gee{}).Pluck("image_name", &gee_names).Error
	err2 := db.Group("save_path").Select("save_path").Find(&Gee{}).Pluck("save_path", &gee_paths).Error
	if err1 != nil || err2 != nil {
		return nil, nil, errmsg.ERROR
	}
	return gee_names, gee_paths, errmsg.SUCCESS
}
func GetImageList() ([]string, []string, int) {
	var image_names []string
	var image_paths []string
	err1 := db.Group("image_name").Select("image_name").Where("uid = ?", 1).Find(&UserDb{}).Pluck("image_name", &image_names).Error
	err2 := db.Group("save_path").Select("save_path").Where("uid = ?", 1).Find(&UserDb{}).Pluck("save_path", &image_paths).Error
	if err1 != nil || err2 != nil {
		return nil, nil, errmsg.ERROR
	}
	return image_names, image_paths, errmsg.SUCCESS
}

func GetResultList() ([]string, []string, int) {
	var result_names []string
	var result_paths []string
	err1 := db.Group("result_name").Select("result_name").Find(&Inference{}).Pluck("result_name", &result_names).Error
	err2 := db.Group("result_path").Select("result_path").Find(&Inference{}).Pluck("result_path", &result_paths).Error
	if err1 != nil || err2 != nil {
		return nil, nil, errmsg.ERROR
	}
	return result_names, result_paths, errmsg.SUCCESS
}

func AddImage(image_name, save_path string) int {
	userDb := UserDb{
		Image_Name: image_name,
		Save_Path:  save_path,
		Uid:        1,
	}
	err2 := db.Create(&userDb).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateImageUserdb(image_name, save_path string) int {
	var userdb UserDb
	ud := UserDb{
		Image_Name: image_name,
		Save_Path:  save_path,
		Uid:        1,
	}
	err := db.Model(&userdb).Where("image_name = ?", ud.Image_Name).Updates(&ud).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckImageUserdb(image_name string, uid int) int {
	var userdb UserDb
	db.Select("id").Where("image_name = ? and uid = ?", image_name, uid).First(&userdb)
	if userdb.ID > 0 {
		return errmsg.ERROR_IMAGE_EXIT
	}
	return errmsg.SUCCESS
}
