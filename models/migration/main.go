package main

import (
	"github.com/provider-go/district/models"
	"github.com/provider-go/pkg/go-logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:13306)/pyrethrum?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = DB.AutoMigrate(models.District{})
	if err != nil {
		logger.Error(err)
	}

	// 读取SQL文件
	content, err := os.ReadFile("./models/migration/district.sql")
	if err != nil {
		logger.Fatal("failed to read SQL file")
	}
	// 将SQL文件内容分割成多个语句
	statements := strings.Split(string(content), ";\n")

	for _, statement := range statements {
		logger.Warn(statement)
		DB.Exec(statement)
	}

	logger.Info("SQL file imported successfully")
}
