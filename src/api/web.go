package api

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) adminHandler(ctx *gin.Context) {
	ctx.File("./src/web/index.html")
}
