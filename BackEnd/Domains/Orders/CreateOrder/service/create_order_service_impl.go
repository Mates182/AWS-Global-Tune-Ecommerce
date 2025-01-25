package service

import (
	"context"
	"create-order-service/data/request"
	"create-order-service/data/response"
	"create-order-service/model"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/google/uuid"
)

type CreateOrderServiceImpl struct {
	DBClient *redis.Client
}

func NewCreateOrderServiceImpl(dbClient *redis.Client) CreateOrderService {
	return &CreateOrderServiceImpl{
		DBClient: dbClient,
	}
}

func (service *CreateOrderServiceImpl) CreateOrder(serviceRequest request.Request) (int, response.Response) {
	orderID := uuid.New().String()

	order := model.Order{
		ID: orderID,
		Items: model.OrderItems{
			Number:            0, // TODO: generate order number
			Date:              serviceRequest.Data.Date,
			BillingDetailsID:  serviceRequest.Data.BillingDetailsID,
			ShippingDetailsID: serviceRequest.Data.ShippingDetailsID,
			InvoiceID:         serviceRequest.Data.InvoiceID,
			PaymentID:         serviceRequest.Data.PaymentID,
		},
	}
	client := service.DBClient
	orderKey := "order:" + order.ID

	err := client.HSet(context.Background(), orderKey, map[string]interface{}{
		"number":              order.Items.Number,
		"date":                order.Items.Date,
		"billing_details_id":  order.Items.BillingDetailsID,
		"shipping_details_id": order.Items.ShippingDetailsID,
		"invoice_id":          order.Items.InvoiceID,
		"payment_id":          order.Items.PaymentID,
	}).Err()

	if err != nil {
		return http.StatusInternalServerError, response.Response{Message: "Failed to create order"}
	}

	return http.StatusOK, response.Response{
		Message: "Order Created Successfully",
		Data: response.ResponseData{
			Order: order,
		},
	}
}
