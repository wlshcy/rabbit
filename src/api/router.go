package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	backend Backend
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func NewRouter(b Backend) *Router {
	router := &Router{
		backend: b,
	}

	router.initRoutes()

	return router
}

func (r *Router) initRoutes() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// items
		v1.GET("/items", r.getItems)
		v1.GET("/onsale", r.getOnSale)

		// orders
		v1.GET("/orders", r.getOrders)
		v1.POST("/orders", r.newOrder)

		// address
		v1.GET("/address", r.getAddress)
		v1.POST("/address", r.newAddress)
		v1.DELETE("/address/:id", r.deleteAddress)
		v1.PATCH("/address/:id", r.updateAddress)
	}

	r.router = router
}
