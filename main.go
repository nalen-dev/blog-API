package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nlndrpraja/blog-API/controller/blog"
)

func main() {
	r := gin.Default()
	r.GET("/blog", blog.FindAllBlogsHandler)
	r.GET("/blog/:id", blog.FindBlogByIdHandler)
	r.POST("/blog", blog.CreateBlogHandler)
	r.PUT("/blog/like/:id", blog.AddLikeHandler)
	r.DELETE("/blog/:id", blog.DeleteBlogByIdHandler)
	r.Run() 
}
