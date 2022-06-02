package register

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type Db_Config struct {
	Url      string
	Username string
	Pwd      string
	Table    string
}

type DB_Controller struct {
	Db        *gorm.DB
	Encryptor *Encryptor
}

func (register *DB_Controller) CheckUser(loguser *LogUser) int32 {
	var user User
	result := register.Db.Where("user_name = ? AND user_pwd = ?", loguser.UserName, loguser.UserPwd).First(&user)
	if result.RowsAffected == 0 {
		return -1
	} else {
		return user.UserId
	}
}

func (register *DB_Controller) RegisterUser(user *User) int {
	if register.CheckUserName(user.UserName) {
		fmt.Println("Duplicated username --->", user.UserName)
		return -1
	}
	user.UserPwd = register.Encryptor.Encrypt(user.UserPwd)
	register.Db.Create(user)
	return 1
}

func (register *DB_Controller) CheckUserName(s string) bool {
	var users []User
	result := register.Db.Where("user_name = ? ", s).Find(&users)
	return result.RowsAffected > 0
}

func (register *DB_Controller) Initialize(config *Db_Config) (bool, error) {
	fmt.Println("database connecting to table ", config.Table)
	var build strings.Builder
	build.WriteString(config.Username)
	build.WriteString(":")
	build.WriteString(config.Pwd)
	build.WriteString("@tcp(localhost:3306)/")
	build.WriteString(config.Table)
	build.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := build.String()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errors.As(err, "err")
		return false, err
	}
	fmt.Println("database connected!")
	register.Db = db
	return true, nil
}

func (register *DB_Controller) Table2User() {
	register.Db.AutoMigrate(&User{})
}
