package server

import (
	"github.com/gin-gonic/gin"
	"../controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	u := r.Group("/users")
	p := r.Group("/posts")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}
	{
		ctrl := post.Controller{}
		p.GET("", ctrl.Index)
		p.GET("/:id", ctrl.Show)
		p.POST("", ctrl.Create)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}

	return r
}