package controllers

import (
	"backend/database"
	"backend/database/models"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func GetPayments(c echo.Context) error {
	var payments []models.Payment
	database.GetDatabase().Find(&payments)
	return c.JSON(http.StatusOK, payments)
}

func AddStripePayment(c echo.Context) error {
	stripeKey, success := os.LookupEnv("STRIPE_SECRET_KEY")
	if success == false {
		return c.JSON(http.StatusInternalServerError, "Stripe key error")
	}
	stripe.Key = stripeKey

	amount, err := strconv.Atoi(c.Param("amountToPay"))
	if err != nil {
		fmt.Println(err)
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Payment intent initialization error")
	}

	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return err
	}
	database.GetDatabase().Create(&payment)
	return c.JSON(http.StatusCreated, map[string]string{"paymentNumber": strconv.Itoa(int(payment.ID)), "clientSecret": pi.ClientSecret})
}
