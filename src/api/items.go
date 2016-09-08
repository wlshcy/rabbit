package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) getItems(ctx *gin.Context) {
	length, _ := strconv.Atoi(ctx.DefaultQuery("length", "10"))
	lastid := ctx.DefaultQuery("lastid", "0")

	items, _ := r.backend.GetItems(length, lastid)

	ctx.JSON(http.StatusOK, items)
	return
}

func (r *Router) getItem(ctx *gin.Context) {
	// get item from mongodb
	item, err := r.backend.GetItem(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, item)
	return
}

func (r *Router) getOnSale(ctx *gin.Context) {
	return
}

func (r *Router) getOnSales(ctx *gin.Context) {
	return
}
