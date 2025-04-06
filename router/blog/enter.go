package blog

import (
	"github.com/gin-gonic/gin"
)

type BlogRouter struct{}

type RouterGroup struct {
	BlogRouter
}

func (s *BlogRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	r := Router.Group("api/v1")

	{
		r.GET("posts", ListPosts)
		r.GET("posts/:slug", GetPostBySlug)
		r.POST("posts", CreatePost)
		r.PUT("posts/:slug", UpdatePost)
		r.DELETE("posts/:slug", DeletePost)
	}

	return r
}
