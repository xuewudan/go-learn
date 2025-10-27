package main

import (
	"fmt"

	"blog-backend/config"
	"blog-backend/handler"
	"blog-backend/middleware"
	"blog-backend/model"
	"blog-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化日志
	util.InitLogger()

	// 加载配置
	cfg := config.Load()

	// 连接数据库
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		logrus.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化处理器
	authHandler := handler.NewAuthHandler(db)
	postHandler := handler.NewPostHandler(db)
	commentHandler := handler.NewCommentHandler(db)

	// 初始化 Gin 引擎
	r := gin.Default()

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.GET("/posts", postHandler.GetAllPosts)
		public.GET("/posts/:id", postHandler.GetPost)
		public.GET("/posts/:postID/comments", commentHandler.GetPostComments)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		// 文章管理
		protected.POST("/posts", postHandler.CreatePost)
		protected.PUT("/posts/:id", postHandler.UpdatePost)
		protected.DELETE("/posts/:id", postHandler.DeletePost)

		// 评论管理
		protected.POST("/posts/:postID/comments", commentHandler.CreateComment)
	}

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	logrus.Infof("服务器启动在 %s", addr)
	if err := r.Run(addr); err != nil {
		logrus.Fatalf("服务器启动失败: %v", err)
	}
}
