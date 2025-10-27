package handler

import (
	"net/http"

	"blog-backend/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CommentHandler 评论处理器
type CommentHandler struct {
	db *gorm.DB
}

// NewCommentHandler 创建评论处理器
func NewCommentHandler(db *gorm.DB) *CommentHandler {
	return &CommentHandler{db: db}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("postID")
	var input struct {
		Content string `json:"content" binding:"required"`
	}

	// 绑定并验证输入
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("创建评论参数错误: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查文章是否存在
	var post model.Post
	if err := h.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		logrus.Errorf("查询文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	// 创建评论
	comment := model.Comment{
		Content: input.Content,
		UserID:  userID.(uint),
		PostID:  post.ID,
	}
	if err := h.db.Create(&comment).Error; err != nil {
		logrus.Errorf("创建评论失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// 关联用户信息返回
	h.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).First(&comment, comment.ID)

	logrus.Infof("用户 %d 为文章 %s 创建评论成功", userID, postID)
	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

// GetPostComments 获取文章的所有评论
func (h *CommentHandler) GetPostComments(c *gin.Context) {
	postID := c.Param("postID")

	// 检查文章是否存在
	var post model.Post
	if err := h.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		logrus.Errorf("查询文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	// 获取评论列表
	var comments []model.Comment
	if err := h.db.Where("post_id = ?", postID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Order("created_at DESC").Find(&comments).Error; err != nil {
		logrus.Errorf("获取评论列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
