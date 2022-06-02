package register

//登录验证时的User
type LogUser struct {
	UserName string
	UserPwd  string
}
type User struct {
	//注册时为空
	UserId   int32  `gorm:"primaryKey;autoIncrement:true"`
	UserName string `gorm:"not null""`
	UserPwd  string `gorm:"not null""`
}
