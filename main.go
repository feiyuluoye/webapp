package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"tarot-app/config"
	"tarot-app/database"
	"tarot-app/routes"
)

func main() {

	// 设置Gin的模式为开发者模式
	gin.SetMode(gin.DebugMode)
	// 初始化配置
	config.LoadConfig()

	// 初始化数据库
	db := database.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB from gorm.DB")
	}
	defer sqlDB.Close() // 关闭数据库连接池

	// 初始化Gin引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)
	// 注册路由
	//
	// 获取环境变量 PORT，如果没有定义则使用默认端口 :8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("Environment variable PORT is undefined. Using port :8080 by default")
	} else {
		fmt.Printf("Using port :%s from environment variable PORT\n", port)
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
		fmt.Println("Environment variable PORT is undefined. Using port :8080 by default")
	} else {
		fmt.Printf("Using port :%s from environment variable PORT\n", host)
	}
	r.Run(host + ":" + port)
}
