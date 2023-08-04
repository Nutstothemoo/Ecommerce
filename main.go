package main

import (
	"ginapp/middleware"
	"ginapp/routes"
	"os"
	"ginapp/database"
	"ginapp/controllers"
	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "Products")

func main() {
	port:= os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app:= controllers.NewApplication(database.ProductData(database.Client, "Products",database.UserData(database.Client, "Users")))
	
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// r.Use(cors.Default())
	routes.UserRoutes(r)
	r.Use(middleware.Authentification())

	r.GET("/products", app.AddtoCart())
	r.GET("/removeitem", app.RemoveItem())
	r.GET("/cartcheckout", app.BuyFromCart())
	r.GET("/instantbuy", app.InstantBuy())
	

	r.Run(":"+ port ) // listen and serve on 0.0.0.0:8080
}