		package routes

	import (
		controllers "ginapp/api/controllers"
		"ginapp/api/middleware"
		"github.com/gin-gonic/gin"
	)
	
	func EventRoutes (incomingRoutes * gin.Engine){
	
		incomingRoutes.GET("/events", controllers.GetEvents())
		incomingRoutes.GET("/events/:id", controllers.GetEventByID())
		incomingRoutes.GET("/events/date/:fromDate/:toDate", controllers.GetEventFromDateToDate())		
		dealerRoutes := incomingRoutes.Group("/dealer")
    dealerRoutes.Use(middleware.DealerAuthentication())
    dealerRoutes.POST("/events", controllers.AddEvent())
    dealerRoutes.PUT("/events/:id", controllers.UpdateEvent())
    dealerRoutes.DELETE("/events/:id", controllers.DeleteEvent())
	}
	
	func AdminEventRoutes (incomingRoutes * gin.Engine){
	
		incomingRoutes.GET("/events", controllers.GetEvents())
		incomingRoutes.GET("/events/:id", controllers.GetEventByID())
		incomingRoutes.GET("/events/date/:fromDate/:toDate", controllers.GetEventFromDateToDate())		
		incomingRoutes.POST("/events", controllers.AddEvent())
		incomingRoutes.PUT("/events/:id", controllers.UpdateEvent())
		incomingRoutes.DELETE("/events/:id", controllers.DeleteEvent())
	}


	