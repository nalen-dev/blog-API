package blog

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func findBlogById(id int) (int,*db.Blog,error){
	for index,blog := range db.Blogs{
		if id == int(blog.ID){
			return index,&blog, nil
		}
	}
	return 0,nil, errors.New("book not found")
}

func FindBlogByIdHandler(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	_,blog,err := findBlogById(id)
	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "blog not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, blog)
}