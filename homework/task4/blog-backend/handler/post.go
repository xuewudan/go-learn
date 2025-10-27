package handler

import (
	"net/http"

	"blog-backend/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// PostHandler 文章处理器
type PostHandler struct {
	db *gorm.DB
}

// NewPostHandler 创建文章处理器
func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{db: db}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// 绑定并验证输入
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("创建文章参数错误: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建文章
	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID.(uint),
	}
	if err := h.db.Create(&post).Error; err != nil {
		logrus.Errorf("创建文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	logrus.Infof("用户 %d 创建文章成功: %d", userID, post.ID)
	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// GetAllPosts 获取所有文章
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	var posts []model.Post
	if err := h.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username") // 只加载用户必要字段
	}).Find(&posts).Error; err != nil {
		logrus.Errorf("获取文章列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// GetPost 获取单篇文章
func (h *PostHandler) GetPost(c *gin.Context) {
	postID := c.Param("id")
	var post model.Post

	// 查询文章并预加载用户和评论
	if err := h.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		logrus.Errorf("获取文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("id")
	var post model.Post

	// 检查文章是否存在
	if err := h.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		logrus.Errorf("查询文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	// 检查权限（只有作者可以更新）
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	// 绑定更新内容
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("更新文章参数错误: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新文章
	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Content != "" {
		post.Content = input.Content
	}
	if err := h.db.Save(&post).Error; err != nil {
		logrus.Errorf("更新文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	logrus.Infof("用户 %d 更新文章成功: %s", userID, postID)
	c.JSON(http.StatusOK, gin.H{"post": post})
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("id")
	var post model.Post

	// 检查文章是否存在
	if err := h.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		logrus.Errorf("查询文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	// 检查权限（只有作者可以删除）
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	// 删除文章（级联删除评论，需在数据库设置外键级联）
	if err := h.db.Delete(&post).Error; err != nil {
		logrus.Errorf("删除文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	logrus.Infof("用户 %d 删除文章成功: %s", userID, postID)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
