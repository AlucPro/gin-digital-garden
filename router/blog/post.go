package blog

import (
	"gin-digital-garden/global"
	"gin-digital-garden/model/blog"
	"gin-digital-garden/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ListPosts(ctx *gin.Context) {
	postModel := blog.NewPostModel(global.GLOBAL_DB)
	posts, err := postModel.GetAll()
	if err != nil {
		global.GLOBAL_LOG.Error("获取文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func GetPostBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	postModel := blog.NewPostModel(global.GLOBAL_DB)
	post, err := postModel.GetBySlug(slug)
	if err != nil {
		global.GLOBAL_LOG.Error("获取文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func CreatePost(ctx *gin.Context) {
	var post blog.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		global.GLOBAL_LOG.Error("创建文章失败", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 生成 slug
	post.Slug = utils.GenerateSlug(post.Title)
	// 检查 slug 是否已存在
	postModel := blog.NewPostModel(global.GLOBAL_DB)
	existingPost, err := postModel.GetBySlug(post.Slug)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.GLOBAL_LOG.Error("检查 slug 失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "检查 slug 失败",
		})
		return
	}
	if existingPost != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "slug 已存在",
		})
		return
	}
	// 创建文章
	err = postModel.Create(&post)
	if err != nil {
		global.GLOBAL_LOG.Error("创建文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, post)
}

func UpdatePost(ctx *gin.Context) {
	// 修改为使用slug更新
	slug := ctx.Param("slug")
	postModel := blog.NewPostModel(global.GLOBAL_DB)
	post, err := postModel.GetBySlug(slug)
	if err != nil {
		global.GLOBAL_LOG.Error("获取文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "文章不存在",
		})
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		global.GLOBAL_LOG.Error("更新文章失败", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = postModel.UpdateBySlug(slug, updateData)
	if err != nil {
		global.GLOBAL_LOG.Error("更新文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	updatedPost, err := postModel.GetBySlug(slug)
	if err != nil {
		global.GLOBAL_LOG.Error("获取更新后的文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, updatedPost)
}

func DeletePost(ctx *gin.Context) {
	// 修改为使用slug删除
	slug := ctx.Param("slug")
	postModel := blog.NewPostModel(global.GLOBAL_DB)
	post, err := postModel.GetBySlug(slug)
	if err != nil {
		global.GLOBAL_LOG.Error("获取文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "文章不存在",
		})
		return
	}
	err = postModel.DeleteBySlug(slug)
	if err != nil {
		global.GLOBAL_LOG.Error("删除文章失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
	})
}
