package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Model 定义
type Model struct {
	ID          uint      `gorm:"primary_key; <-:create; autoIncrement"`
	CreatedTime time.Time `json:"created_time; autoCreateTime"`
	UpdatedTime time.Time `json:"updated_time; autoUpdateTime"`
	DeletedTime time.Time `json:"deleted_time"`
}

// 初始化
func init() {
	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "47.108.140.163", 8166, "gin_test")
	//constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "lee")
	db, err = gorm.Open(mysql.Open(constr), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connect mysql server error !  \n err: %s", err))
	}
	err = db.AutoMigrate(&User{})
	isExist := db.Migrator().HasTable(&User{})
	fmt.Println("是否存在数据库： ", isExist)
	if isExist { //判断表是否存在
		db.AutoMigrate(&User{}) //存在就自动适配表，也就说原先没字段的就增加字段
	} else {
		db.Migrator().CreateTable(&User{}) //不存在就创建新表
	}
	if err != nil {
		panic(fmt.Sprintf("DB auto migrate err: %s", err))
	}
	sqlDB, err := db.DB()

	if err != nil {
		panic(fmt.Sprintf("DB auto migrate err: %s", err))
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
