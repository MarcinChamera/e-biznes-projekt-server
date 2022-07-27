package main

import (
	"backend/controllers"
	"backend/database"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env loading error")
	}

	e := echo.New()
	database.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://namelessshop.azurewebsites.net", "http://localhost:3000", "http://localhost:1323"},
	}))

	e.GET("/users", controllers.GetUsers)

	e.DELETE("/users/:id", controllers.DeleteUser)

	e.GET("/products", controllers.GetProducts)

	e.GET("/products/category/:categoryId", controllers.GetProductsByCategory)

	e.GET("/products/:id", controllers.GetProduct)

	e.POST("/products", controllers.AddProduct)

	e.GET("/categories", controllers.GetCategories)

	e.GET("/categories/:id", controllers.GetCategory)

	e.GET("/shopping-carts", controllers.GetCarts)

	e.POST("/shopping-carts", controllers.AddCart)

	e.POST("/shopping-carts/:id", controllers.AddCartProduct)

	e.GET("/payments", controllers.GetPayments)

	e.POST("/payments/stripe/:amountToPay", controllers.AddStripePayment)

	e.GET("/orders", controllers.GetOrders)

	e.POST("/orders", controllers.AddOrder)

	e.GET("/google/callback", controllers.GoogleCallback)

	e.GET("/google/login", controllers.GoogleLogin)

	e.GET("/github/login", controllers.GithubLogin)

	e.GET("/github/callback", controllers.GithubCallback)

	e.GET("/facebook/login", controllers.FacebookLogin)

	e.GET("/facebook/callback", controllers.FacebookCallback)

	e.Logger.Fatal(e.Start(":1323"))
}
