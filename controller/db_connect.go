package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 连接数据库，并创建各种需要的表
// 目前已完成用户信息表

var DB *gorm.DB

func ConnectDB() {
	// 连接数据库
	dsn := "root:root@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败")
	}
	//创建Video表
	err = DB.AutoMigrate(&Video{})
	//=左边的变量是已经声明过的变量，:=是声明并赋值
	if err != nil {
		log.Fatalln("Video表创建失败")
	}
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalln("User表创建失败")
	}
	if err := DB.AutoMigrate(&Favorite{}); err != nil {
		log.Fatalln("Favorite表创建失败")
	}
	if err := DB.AutoMigrate(&Comment{}); err != nil {
		log.Fatalln("Comment表创建失败")
	}
	if err := DB.AutoMigrate(&Relation{}); err != nil {
		log.Fatalln("Relation表创建失败")
	}
}
