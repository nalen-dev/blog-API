package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func AddLikeHandler(c *gin.Context){
	id,ok := c.GetQuery("id")
	id_int,_ := strconv.Atoi(id)

	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "missing id query parameter."})	
		return
	}

	index,blog, err := findBlogById(id_int)
	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "blog not found"})		
		return
	}
	blog.Like += 1;
	db.Blogs[index] = *blog
	c.IndentedJSON(http.StatusOK, blog)
}