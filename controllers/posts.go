package controllers

import (
	"blog_api/database"
	"blog_api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePost 创建新文章
func CreatePost(c *gin.Context) {
	var post model.TypechoContent

	// 绑定请求体到结构体
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 设置创建时间和修改时间
	currentTime := time.Now().Unix()
	post.Created = currentTime
	post.Modified = currentTime

	// 创建文章
	result := database.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetAllPosts 获取所有文章
func GetAllPosts(c *gin.Context) {
	var posts []model.TypechoContent

	// 可以添加分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 查询文章列表
	result := database.DB.Where("type = ?", "post").
		Order("created DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch posts",
		})
		return
	}

	// 获取总数
	var total int64
	database.DB.Model(&model.TypechoContent{}).Where("type = ?", "post").Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"data":  posts,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetPost 获取单个文章
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post model.TypechoContent

	result := database.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post model.TypechoContent

	// 检查文章是否存在
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// 绑定更新数据
	var updateData model.TypechoContent
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 更新修改时间
	updateData.Modified = time.Now().Unix()

	// 更新文章
	result := database.DB.Model(&post).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update post",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&model.TypechoContent{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete post",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
