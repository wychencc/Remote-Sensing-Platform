package model

import "time"

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

// GetImage Geez采集图像接口
func GetImage(geeparam *GeeParam) {

}
