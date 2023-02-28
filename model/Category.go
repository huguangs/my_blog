package model

import (
	"go_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment;" json:"id"`
	Name string `gorm:"type:varchar(20); not null;" json:"name"`
}

func (cate *Category) TableName() string {
	return "category"
}

// CheckCategory
// 检测分类是否存在
func CheckCategory(name string) int {
	var cate Category
	Db.Select("id").Where("name = ? ", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED // 1009
	}
	return errmsg.SUCCESS
}

// CreateCategory
// 新增分类
func CreateCategory(data *Category) int {
	//data.PassWord = ScryptPw(data.PassWord)
	err := Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategories
// 查询分类列表
func GetCategories(pageSize int, pageNum int) []Category {
	var cates []Category
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// UpdateCategory
// 编辑分类
func UpdateCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := Db.Model(&Category{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// DeleteCategory
// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := Db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
