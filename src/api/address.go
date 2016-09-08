package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wlshcy/rabbit/src/db"
)

func (r *Router) getAddresses(ctx *gin.Context) {
	addresses, err := r.backend.GetAddresses()
	if err != nil {
		log.Printf("[ERROR] Retrieve addresses failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, addresses)
	return
}

func (r *Router) newAddress(ctx *gin.Context) {
	var address db.Address

	if err := ctx.BindJSON(&address); err != nil {
		log.Printf("[ERROR] Bind json failed: %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := r.backend.NewAddress(&address); err != nil {
		log.Printf("[ERROR] create address failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (r *Router) deleteAddress(ctx *gin.Context) {
	return
}

func (r *Router) updateAddress(ctx *gin.Context) {
	return
}
