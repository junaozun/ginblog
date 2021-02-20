package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"角色码"` //角色>=2(1是管理员，大于2以上是用户)
}

// 查询用户是否存在
func CheckUser(userName string) int {
	var data User
	db.Select("id").Where("username = ?", userName).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	// 保存数据库之前需要对密码加密
	//data.Password = ScryptPasswd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func QueryUserLists(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int //页码总数
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	// 不让用户在编辑用户修改密码
	var user User
	// 结构体不能修改0值，这里用map
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	// 软删除
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPasswd(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 22, 66, 222, 111}

	hashPasswd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	realPwd := base64.StdEncoding.EncodeToString(hashPasswd)
	return realPwd
}

// 保存之前调用的钩子函数，用于加密密码
func (u *User) BeforeSave() {
	u.Password = ScryptPasswd(u.Password)
}

// 登录验证
func CheckLogin(userName string, passWord string) int {
	var user User

	db.Where("username = ?", userName).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPasswd(passWord) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	//非管理员不能登录后台
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.SUCCESS
}
