package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wlshcy/rabbit/src/db"
)

func (r *Router) getOrders(ctx *gin.Context) {
	return
}

func (r *Router) getOrder(ctx *gin.Context) {
	order, err := r.backend.GetOrder(ctx.Param("id"))
	if err != nil {
		log.Printf("[ERROR] Retrieve order %s failed: %s", ctx.Param("id"), err)
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, order)
	return
}

func (r *Router) newOrder(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		log.Printf("[ERROR] Retrive addresses faile: %s", "uid not in ctxt")
		ctx.JSON(http.StatusUnauthorized, "no uid in ctx")
		return
	}

	order := new(db.Order)
	if err := ctx.BindJSON(&order); err != nil {
		log.Printf("[ERROR] Bind json failed: %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	order.Uid = uid.(string)
	order.Number = "12345678"
	order.Created = time.Now().UTC().Unix()

	if err := r.backend.NewOrder(order); err != nil {
		log.Printf("[ERROR] create order failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}
