package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) adminHandler(ctx *gin.Context) {
	//ctx.File("./src/web/index.html")
	ctx.Redirect(http.StatusMovedPermanently, "/v1/admin/summaries")
}

func (r *Router) getSummaries(ctx *gin.Context) {
	ctx.File("./src/web/templates/summaries/summaries.html")
}
