package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Refund ...
type Refund struct {
	Request *requests.Request
}

// All fetches colelction of Refund for the given queryParams
func (refund *Refund) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.REFUND_URL)
	return refund.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches the Refund having the given refundID.
func (refund *Refund) Fetch(ctx context.Context, refundID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.REFUND_URL, url.PathEscape(refundID))
	return refund.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new Refund for the given data.
func (refund *Refund) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.REFUND_URL)
	return refund.Request.Post(ctx, url, data, extraHeaders)
}

// Update updates the Refund having the given refundID.
func (refund *Refund) Update(ctx context.Context, refundID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.REFUND_URL, url.PathEscape(refundID))
	return refund.Request.Patch(ctx, url, queryParams, extraHeaders)
}
