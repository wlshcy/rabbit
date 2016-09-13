package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) sendSMS(ctx *gin.Context) {
	if err := r.backend.SendSMS(ctx.Param("phone")); err != nil {
		log.Printf("[ERROR] Retrieve addresses failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return
}
