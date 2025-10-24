package gorm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 扩展 User 模型，增加文章数量统计字段
// Post 模型的创建前钩子（BeforeCreate）
// 在文章创建时自动更新用户的文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 1. 先查询用户当前的文章数
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return fmt.Errorf("查询用户失败: %w", err)
	}

	// 2. 更新用户的文章数（+1）
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_num", user.PostNum+1).Error; err != nil {
		return fmt.Errorf("更新用户文章数失败: %w", err)
	}
	return nil
}

// Comment 模型的删除后钩子（AfterDelete）
// 在评论删除后检查文章评论数，若为0则更新文章状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 查询当前文章的剩余评论数
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return fmt.Errorf("查询评论数失败: %w", err)
	}

	// 2. 若评论数为0，更新文章的评论状态
	if commentCount == 0 {
		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error; err != nil {
			return fmt.Errorf("更新文章评论状态失败: %w", err)
		}
	}
	return nil
}

func testC() {
	// 连接数据库
	dsn := "root:xuetao123@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	//// 自动迁移表结构（确保新增字段生效）
	//err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	//if err != nil {
	//	log.Fatalf("表结构迁移失败: %v", err)
	//}

	// 测试 Post 钩子：创建一篇文章，检查用户文章数是否自动增加
	//testCreatePost(db)

	// 测试 Comment 钩子：删除最后一条评论，检查文章状态是否更新
	testDeleteComment(db)
}

// 测试文章创建钩子
func testCreatePost(db *gorm.DB) {
	// 假设存在 ID=1 的用户
	userID := uint(1)

	// 创建一篇新文章
	newPost := Post{
		Title:   "测试钩子函数的文章",
		Content: "这是一篇用于测试 BeforeCreate 钩子的文章",
		UserID:  userID,
	}

	if err := db.Create(&newPost).Error; err != nil {
		log.Printf("创建文章失败: %v", err)
		return
	}

	// 验证用户文章数是否增加
	var user User
	db.First(&user, userID)
	fmt.Printf("\n创建文章后，用户 %d 的文章数: %d\n", userID, user.PostNum)
}

// 测试评论删除钩子
func testDeleteComment(db *gorm.DB) {
	// 假设存在 ID=1 的文章，且该文章有评论
	postID := uint(2)

	// 查询该文章的一条评论（取第一条）
	var comment Comment
	if err := db.Where("post_id = ?", postID).First(&comment).Error; err != nil {
		log.Printf("查询评论失败: %v", err)
		return
	}

	// 删除该评论
	if err := db.Delete(&comment).Error; err != nil {
		log.Printf("删除评论失败: %v", err)
		return
	}

	// 验证文章评论状态
	var post Post
	db.First(&post, postID)
	fmt.Printf("删除评论后，文章 %d 的评论状态: %s\n", postID, post.CommentStatus)
}
