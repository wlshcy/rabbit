package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wlshcy/rabbit/src/auth"
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

	authorized := router.Group("/v1")
	authorized.Use(auth.JWTAuthMiddleware())
	{
		// orders
		authorized.GET("/orders", r.getOrders)
		authorized.POST("/orders", r.newOrder)

		// address
		authorized.GET("/addresses", r.getAddresses)
		authorized.GET("/addresses/:id", r.getAddress)
		authorized.POST("/addresses", r.newAddress)
		authorized.DELETE("/addresses/:id", r.deleteAddress)
		authorized.POST("/addresses/:id", r.updateAddress)
	}

	unauthorized := router.Group("/v1")
	{
		// login
		unauthorized.POST("/login", r.login)

		// SMS
		unauthorized.POST("/sms/:phone", r.sendSMS)

		// items
		unauthorized.GET("/items/:id", r.getItem)
		unauthorized.GET("/items", r.getItems)
		unauthorized.GET("/onsales/:id", r.getOnSale)
		unauthorized.GET("/onsales", r.getOnSales)
	}

	web := router.Group("/v1/admin")
	{
		web.GET("/", r.adminHandler)
		web.GET("/summaries", r.getSummaries)
	}

	router.Static("/static", "./src/web/static")

	r.router = router
}
