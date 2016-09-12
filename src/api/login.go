package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wlshcy/rabbit/src/db"
)

func (r *Router) login(ctx *gin.Context) {
	credential := db.Credential{}
	if err := ctx.BindJSON(&credential); err != nil {
		log.Printf("[ERROR] Bind json failed: %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := r.backend.Login(&credential)
	if err != nil {
		log.Printf("[ERROR] login failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
	return
}
