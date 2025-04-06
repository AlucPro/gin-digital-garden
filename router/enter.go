package router

import (
	"gin-digital-garden/router/blog"
	"gin-digital-garden/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Blog   blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
