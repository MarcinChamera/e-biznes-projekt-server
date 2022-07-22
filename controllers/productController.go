package controllers

import (
	"backend/database"
	"backend/database/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	database.GetDatabase().Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProductsByCategory(c echo.Context) error {
	categoryId := c.Param("categoryId")
	var products []models.Product
	database.GetDatabase().Where("category_id = ?", categoryId).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product

	if result := database.GetDatabase().First(&product, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, product)
}

func AddProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	database.GetDatabase().Create(&product)
	return c.String(http.StatusCreated, "Created")
}
