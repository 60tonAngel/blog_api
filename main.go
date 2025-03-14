package main

import (
	"blog_api/config"
	"blog_api/database"
	"blog_api/model"
	"blog_api/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	// 设置 gin 模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 初始化数据库
	if err := database.Connect(); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移
	if err := database.DB.AutoMigrate(&model.TypechoContent{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 创建 Gin 实例
	r := gin.Default()

	// 配置路由
	routes.SetupRouter(r)

	// 启动服务
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
