package gorm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 模型（用户）
type User struct {
	gorm.Model        // 嵌入 Gorm 基础模型，包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	Username   string `gorm:"size:50;not null;unique"`  // 用户名，唯一且非空
	Email      string `gorm:"size:100;not null;unique"` // 邮箱，唯一且非空
	PostNum    uint
	Posts      []Post // 一对多关系：一个用户可以有多个文章
}

// Post 模型（文章）
type Post struct {
	gorm.Model           // 基础字段
	Title         string `gorm:"size:200;not null"`  // 文章标题
	Content       string `gorm:"type:text;not null"` // 文章内容
	CommentNum    uint
	CommentStatus string
	UserID        uint      `gorm:"not null"` // 外键：关联用户 ID
	User          User      // 关联用户（一对多的反向引用）
	Comments      []Comment // 一对多关系：一篇文章可以有多个评论
}

// Comment 模型（评论）
type Comment struct {
	gorm.Model        // 基础字段
	Content    string `gorm:"size:500;not null"` // 评论内容
	PostID     uint   `gorm:"not null"`          // 外键：关联文章 ID
	Post       Post   // 关联文章（一对多的反向引用）
}

func testA() {
	// 1. 连接 MySQL 数据库
	// 连接字符串格式：用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:xuetao123@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 2. 自动迁移：根据模型创建或更新表结构
	// 会自动处理外键关系（UserID -> User.ID，PostID -> Post.ID）
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("表结构迁移失败: %v", err)
	}

	log.Println("数据库表创建/更新成功")
}
