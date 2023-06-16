package handler

import (
	"crm_admin/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	InitRoutes(router *gin.Engine)
}

type handler struct {
	services *services.Services
}

func (h *handler) InitRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/get-all", h.getUsers)
			users.GET("/get/:id", h.getUser)
			users.POST("/add", h.addUser)
			users.PUT("/edit", h.editUser)
		}

		groups := api.Group("/groups")
		{
			groups.GET("/get-all", h.getGroups)
			groups.PUT("/edit", h.editGroup)
		}

		plots := api.Group("/plots")
		{
			plots.GET("/get-all", h.getPlots)
			plots.POST("/add", h.addPlot)
			plots.PUT("/edit", h.editPlot)
		}

		filters := api.Group("/filters")
		{
			filters.GET("/get-all", h.getFilters)
			filters.POST("/add", h.addFilter)
			filters.PUT("/edit", h.editFilter)
		}

		clients := api.Group("/clients")
		{
			clients.GET("/get-all", h.getClients)
			clients.POST("/add", h.addClient)
			clients.PUT("/edit", h.editClient)
		}
	}
}

func NewHandler(services *services.Services) Handler {
	return &handler{
		services: services,
	}
}
