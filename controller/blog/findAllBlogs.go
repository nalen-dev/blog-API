package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nlndrpraja/blog-API/resource/datablog"
)

func FindAllBlogsHandler(c *gin.Context){
	c.IndentedJSON(http.StatusOK,db.Blogs)
}
