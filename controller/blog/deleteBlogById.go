package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func DeleteBlogByIdHandler(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))

	for i,data := range db.Blogs{
		if int(data.ID) == id{
			db.Blogs = append(db.Blogs[:i],db.Blogs[i+1:]... )
			c.IndentedJSON(http.StatusOK, gin.H{"message" : "success deleting blog"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "blog not found"})
}