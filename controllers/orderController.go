package controllers

import (
	"backend/database"
	"backend/database/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrders(c echo.Context) error {
	var orders []models.Order
	database.GetDatabase().Find(&orders)
	return c.JSON(http.StatusOK, orders)
}

func AddOrder(c echo.Context) error {
	order := new(models.Order)

	requestJsonBody := GetJSONRawBody(c)
	paymentNumber, _ := requestJsonBody["PaymentNumber"]
	userId, _ := requestJsonBody["UserId"]
	stripePaymentId, _ := requestJsonBody["StripePaymentId"]

	if userId != nil {
		order.UserId = int(userId.(float64))
	}
	if paymentNumber != nil && stripePaymentId != nil {
		order.PaymentNumber = int(paymentNumber.(float64))
		order.StripePaymentId = stripePaymentId.(string)
		database.GetDatabase().Create(&order)
	}

	return c.String(http.StatusCreated, "Created")
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		fmt.Println("empty json body")
		return nil
	}

	return jsonBody
}
