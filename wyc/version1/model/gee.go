package model

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"version1_copy/utils/errmsg"
)

const (
	base_path       = "http://localhost:6060/"
	image_save_path = "static/trans/"
	save_path       = "./static/inference/"
	trans_path      = "./static/trans/"
)

type Gee struct {
	gorm.Model
	Image_Name string `gorm:"type:varchar(40);not null" json:"image_name" form:"image_name"`
	Save_Path  string `gorm:"type:varchar(100);not null" json:"save_path" form:"save_path"`
	Uid        int    `gorm:"type:int;not null" json:"uid" form:"uid"`
}

type GeeParam struct {
	Start_date time.Time `json:"start_date" form:"start_date" time_format:"2006-01-02"`
	End_date   time.Time `json:"end_date" form:"end_date" time_format:"2006-01-02"`
	Max_pixel  float32   `json:"max_pixel" form:"max_pixel"`
	Min_pixel  float32   `json:"min_pixel" form:"min_pixel"`
	Point_lon1 float32   `json:"lon1" form:"lon1"`
	Point_lat1 float32   `json:"lat1" form:"lat1"`
	Point_lon2 float32   `json:"lon2" form:"lon2"`
	Point_lat2 float32   `json:"lat2" form:"lat2"`
	Dataset    string    `json:"dataset" form:"dataset"`
	Image_name string    `json:"image_name" form:"image_name"`
}

// GetImage Gee采集图像接口
func GetImage(geeparam *GeeParam) (int, string) {

	// 先做trans
	file_name := strings.Split(geeparam.Image_name, ".")[0]
	if geeparam.Dataset == "sar" {
		geeparam.Dataset = "1"
	} else if geeparam.Dataset == "rgb" {
		geeparam.Dataset = "3"
	}
	GetTrans(save_path+geeparam.Image_name, trans_path+file_name, geeparam.Dataset)
	return errmsg.SUCCESS, base_path + image_save_path + strings.Split(geeparam.Image_name, ".")[0] + ".png"
}
