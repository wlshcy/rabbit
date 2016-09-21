package api

import (
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wlshcy/rabbit/src/db"
)

func (r *Router) getAddresses(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		log.Printf("[ERROR] Retrive addresses faile: %s", "uid not in ctxt")
		ctx.JSON(http.StatusUnauthorized, "no uid in ctx")
		return
	}

	addresses, err := r.backend.GetAddresses(uid.(string))
	if err != nil {
		log.Printf("[ERROR] Retrieve addresses failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, addresses)
	return
}

func (r *Router) newAddress(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		log.Printf("[ERROR] Retrive addresses faile: %s", "uid not in ctxt")
		ctx.JSON(http.StatusUnauthorized, "no uid in ctx")
		return
	}

	address := new(db.Address)

	if err := ctx.BindJSON(&address); err != nil {
		log.Printf("[ERROR] Bind json failed: %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	address.Uid = uid.(string)
	address.Default = false

	if err := r.backend.NewAddress(address); err != nil {
		log.Printf("[ERROR] create address failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (r *Router) getAddress(ctx *gin.Context) {
	address, err := r.backend.GetAddress(ctx.Param("id"))
	if err != nil {
		log.Printf("[ERROR] Retrieve address %s failed: %s", ctx.Param("id"), err)
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, address)
	return
}

func (r *Router) deleteAddress(ctx *gin.Context) {
	if err := r.backend.DeleteAddress(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (r *Router) updateAddress(ctx *gin.Context) {
	update := make(map[string]interface{})
	if err := ctx.BindJSON(&update); err != nil {
		log.Printf("[ERROR] Bind json failed: %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("%s", update)
	if err := r.backend.UpdateAddress(ctx.Param("id"), update); err != nil {
		log.Printf("[ERROR] Update address failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}
