package model
import "github.com/jinzhu/gorm"

type Zone struct {
	gorm.Model
	Name string
	Info string
	Prefectures []Prefecture 
	Url string
	Path string
	Sha string
}
type Prefecture struct {
	gorm.Model
	ZoneID uint
	Name string
	Info string
	Url string
	Path string
	Sha string
}

 
