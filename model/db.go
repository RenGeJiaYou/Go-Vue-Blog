package model

import (
	"fmt"
	"go-vue-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

// 配置数据库，以便被其它 model/ 文件调用
var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true, //采用逻辑外键,而不是物理外键,可以优化性能
	})
	if err != nil {
		fmt.Println("MySQL 连接出错；", err)
		os.Exit(1) //code !=0 表示出错
	}

	//配置连接池属性
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(10)                  //设置连接池最大空闲连接数->并发时可以同时获取的链接
	mysqlDB.SetMaxOpenConns(100)                 //设置连接池最大连接数
	mysqlDB.SetConnMaxLifetime(10 * time.Second) //设置连接的最长时间,到时则断开

	//自动迁移
	err := db.AutoMigrate(&User{}, &Category{}, &Article{}, &Profile{})
	if err != nil {
		fmt.Println(err)
	}
}
