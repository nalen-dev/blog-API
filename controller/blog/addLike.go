package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func AddLikeHandler(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	index,blog, err := findBlogById(id)
	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "blog not found"})		
		return
	}
	blog.Like += 1;
	db.Blogs[index] = *blog
	c.IndentedJSON(http.StatusOK, blog)
}