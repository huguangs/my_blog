package model

import (
	"encoding/base64"
	"go_blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"user_name"`
	PassWord string `gorm:"type:varchar(20);not null;" json:"pass_word"`
	Role     int    `gorm:"type:int;" json:"role"`
}

func (table *User) TableName() string {
	return "users"
}

// CheckUser
// 检测用户是否存在
func CheckUser(name string) int {
	users := &User{}
	Db.Select("id").Where("user_name = ? ", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser
// 新增用户
func CreateUser(data *User) int {
	data.PassWord = ScryptPw(data.PassWord)
	err := Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers
// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// UpdateUser
// 编辑用户
func UpdateUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["user_name"] = data.UserName
	maps["role"] = data.Role
	err := Db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// SelectUser
// 查询用户
func SelectUser(id int) (*User, error) {
	user := &User{}
	err := Db.Omit("pass_word").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

// DeleteUser
// 删除用户
func DeleteUser(id int) int {
	var user User
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

//func (u *User) BeforeSave() {
//	u.PassWord = ScryptPw(u.PassWord)
//}

// ScryptPw
// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// CheckLogin
// 登录验证
func CheckLogin(username, password string) int {
	var user User
	Db.Where("user_name = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.PassWord {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
