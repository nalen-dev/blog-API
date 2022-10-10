package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func CreateBlogHandler(c *gin.Context){
	var newBlog db.Blog
	if err := c.BindJSON(&newBlog);err != nil{
		return
	}
	newBlog.ID = db.Blogs[len(db.Blogs)-1].ID + 1
	db.Blogs = append(db.Blogs, newBlog)
	c.IndentedJSON(http.StatusOK,newBlog)
}