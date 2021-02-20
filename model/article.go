package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:" varchar(100);not null json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type: varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	// 保存数据库之前需要对密码加密
	//data.Password = ScryptPasswd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func QueryCategoryAllArt(pageSize int, pageNum int, id int) ([]Article, int, int) {
	var cateArtLists []Article
	var total int //页码总数
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtLists).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtLists, errmsg.SUCCESS, total
}

// 查询单个文章
func QueryArtInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// 查询文章列表
func QueryArtLists(pageSize int, pageNum int) ([]Article, int, int) {
	var artList []Article
	var total int //页码总数
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total

}

// 编辑文章
func EditArticle(id int, data *Article) int {
	// 不让用户在编辑用户修改密码
	var art Article
	// 结构体不能修改0值，这里用map
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var cate Article
	// 软删除
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
