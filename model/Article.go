package model

import (
	"go_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid;reference:id"`
	Title    string   `gorm:"type:varchar(100); not null;" json:"title"`
	Cid      int      `gorm:"type:int; not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200);" json:"desc"`
	Content  string   `gorm:"type:longtext;" json:"content"`
	Img      string   `gorm:"type:varchar(100);" json:"img"`
}

func (art *Article) TableName() string {
	return "articles"
}

// CreateArt
// 新增文章
func CreateArt(data *Article) int {
	//data.PassWord = ScryptPw(data.PassWord)
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArtInfo
// 查询单个文章信息
func GetArtInfo(id int) (*Article, int) {
	var art *Article
	err := Db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return nil, errmsg.ERROR_ART_NOT_EXITS
	}
	return art, errmsg.SUCCESS

}

// GetCateArt
// 查询分类下的文章
func GetCateArt(cid, pageSize, pageNum int) ([]Article, int) {
	var arts []Article
	err := Db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Cid = ?", cid).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXITS
	}
	return arts, errmsg.SUCCESS
}

// GetArts
// 查询文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int) {
	var arts []Article
	err := Db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return arts, errmsg.SUCCESS
}

// UpdateArt
// 编辑文章
func UpdateArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := Db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// DeleteArt
// 删除分类
func DeleteArt(id int) int {
	var art Article
	err := Db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
