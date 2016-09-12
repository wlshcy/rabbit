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
		// login
		v1.POST("/login", r.login)

		// items
		v1.GET("/items/:id", r.getItem)
		v1.GET("/items", r.getItems)
		v1.GET("/onsales/:id", r.getOnSale)
		v1.GET("/onsales", r.getOnSales)

		// orders
		v1.GET("/orders", r.getOrders)
		v1.POST("/orders", r.newOrder)

		// address
		v1.GET("/addresses", r.getAddresses)
		v1.GET("/addresses/:id", r.getAddress)
		v1.POST("/addresses", r.newAddress)
		v1.DELETE("/addresses/:id", r.deleteAddress)
		v1.POST("/addresses/:id", r.updateAddress)
	}

	r.router = router
}
