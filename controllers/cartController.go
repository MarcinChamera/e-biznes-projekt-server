package controllers

import (
	"backend/database"
	"backend/database/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return err
	}
	database.GetDatabase().Create(&cart)
	return c.String(http.StatusCreated, strconv.Itoa(int(cart.ID)))
}

func AddCartProduct(c echo.Context) error {
	cartId := c.Param("id")
	var cart models.Cart
	result := database.GetDatabase().Find(&cart, cartId)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Cart doesn't exist")
	}

	cartProduct := new(models.Cart)
	if err := c.Bind(cartProduct); err != nil {
		return err
	}
	cartProduct.CartNumber, _ = strconv.Atoi(cartId)
	database.GetDatabase().Create(&cartProduct)
	return c.String(http.StatusCreated, strconv.Itoa(int(cartProduct.ID)))
}

func GetCarts(c echo.Context) error {
	var carts []models.Cart
	database.GetDatabase().Find(&carts)
	return c.JSON(http.StatusOK, carts)
}
