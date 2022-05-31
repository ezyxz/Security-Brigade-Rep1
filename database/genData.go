package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Url struct {
	ID  int32  `gorm:"primaryKey;autoIncrement:true"`
	Url string `gorm:"not null""`
	Dec string `gorm:"not null""`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/searchdb2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errors.As(err, "err")
	}
	db.AutoMigrate(&Url{})
	filepath := "src/wukong50k_release.csv"

	opencast, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer opencast.Close()

	//创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)

	//获取一行内容，一般为第一行内容
	read, _ := ReadCsv.Read() //返回切片类型：[chen  hai wei]
	log.Println(read)

	//读取所有内容
	ReadAll, err := ReadCsv.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	n := len(ReadAll)
	for i := 0; i < n; i++ {
		db.Create(&Url{Url: ReadAll[i][0], Dec: ReadAll[i][1]})
		fmt.Println(i)
	}

}
