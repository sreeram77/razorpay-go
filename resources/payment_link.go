package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// PaymentLink ...
type PaymentLink struct {
	Request *requests.Request
}

// All fetches multiple PaymentLink for the given queryParams.
func (pl *PaymentLink) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PaymentLink_URL)
	return pl.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches an PaymentLink having the given paymentLinkID.
func (pl *PaymentLink) Fetch(ctx context.Context, paymentLinkID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PaymentLink_URL, paymentLinkID)
	return pl.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new PaymentLink  for the given data.
func (pl *PaymentLink) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PaymentLink_URL)
	return pl.Request.Post(ctx, url, data, extraHeaders)
}

// Cancel cancels paymentLink
func (pl *PaymentLink) Cancel(ctx context.Context, paymentLinkID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/cancel", constants.VERSION_V1, constants.PaymentLink_URL, paymentLinkID)
	return pl.Request.Post(ctx, url, data, extraHeaders)
}

// NotifyBy Send/re-send notification by given medium
func (pl *PaymentLink) NotifyBy(ctx context.Context, paymentLinkID string, medium string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/notify_by/%s", constants.VERSION_V1, constants.PaymentLink_URL, paymentLinkID, medium)
	return pl.Request.Post(ctx, url, data, extraHeaders)
}

// Update updates the paymentLink for the given data.
func (pl *PaymentLink) Update(ctx context.Context, paymentLinkID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PaymentLink_URL, paymentLinkID)
	return pl.Request.Patch(ctx, url, data, extraHeaders)
}
