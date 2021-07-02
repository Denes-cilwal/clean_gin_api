package infrastructure

import (
	"clean_gin_api/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

//NewRouter : all the routes are defined here
func NewRouter(env lib.Env) Router {

	httpRouter := gin.Default()
	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "clean_gin_api 📺 API Up and Running"})
	})

	return Router{
		httpRouter,
	}
}
