package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Category struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(categoryName string) int {
	var data Category
	db.Select("id").Where("name = ?", categoryName).First(&data)
	if data.Id > 0 {
		return errmsg.ERROR_CATENAME_EXIST
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) int {
	// 保存数据库之前需要对密码加密
	//data.Password = ScryptPasswd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func QueryCategoryLists(pageSize int, pageNum int) ([]Category, int) {
	var cates []Category
	var total int //分类总数
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// 编辑分类信息
func EditCategory(id int, data *Category) int {
	// 不让用户在编辑用户修改密码
	var cate Category
	// 结构体不能修改0值，这里用map
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	// 软删除
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下所有文章
