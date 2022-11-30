package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/responses"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type Payments struct {
	CartItemService *models.CartItemService
	UserService     *models.UserService
}

func (p Payments) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.CategoryResponse[string]{
		Message: "",
	}

	token, err := readCookie(r, COOKIE_TOKEN)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	userId, _, err := ParseJWT(token)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	products, err := p.CartItemService.GetAll(userId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams
	for i := 0; i < len(products); i++ {
		var imageUrls []*string
		priceData := &stripe.CheckoutSessionLineItemPriceDataParams{
			UnitAmountDecimal: &products[i].Price,
			Currency:          stripe.String("usd"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name:   stripe.String(products[i].ProductName),
				Images: append(imageUrls, stripe.String(products[i].ImageUrl)),
			},
		}

		lineItemParam := stripe.CheckoutSessionLineItemParams{
			PriceData: priceData,
			Quantity:  stripe.Int64(products[i].Quantity),
		}

		lineItems = append(lineItems, &lineItemParam)

	}

	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	userEmail, err := p.UserService.GetUserEmail(userId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(userEmail),
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: []*string{stripe.String("US")},
		},
		PaymentMethodTypes: []*string{stripe.String("card")},
		SuccessURL:         stripe.String("https://localhost:3000/checkout/order-success"),
		CancelURL:          stripe.String("https://localhost:3000/cart"),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	s, _ := session.New(params)

	response.Data = s.URL
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
