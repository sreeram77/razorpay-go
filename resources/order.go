package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Order ...
type Order struct {
	Request *requests.Request
}

// All fetches multiple orders for the given query params.
func (order *Order) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	return order.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches an order having the given orderID.
func (order *Order) Fetch(ctx context.Context, orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new order for the given data
func (order *Order) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	return order.Request.Post(ctx, url, data, extraHeaders)
}

// Update updates an order having the given orderID.
func (order *Order) Update(ctx context.Context, orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Patch(ctx, url, data, extraHeaders)
}

// Payments fetches the payments for the given orderID.
func (order *Order) Payments(ctx context.Context, orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/payments", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (order *Order) ViewRtoReview(ctx context.Context, orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/rto_review", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Post(ctx, url, data, extraHeaders)
}

func (order *Order) EditFulfillment(ctx context.Context, orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/fulfillment", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Post(ctx, url, data, extraHeaders)
}
