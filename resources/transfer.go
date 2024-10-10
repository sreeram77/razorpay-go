package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Transfer ...
type Transfer struct {
	Request *requests.Request
}

// All fetches collection of transfer for the given queryParams.
func (t *Transfer) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.TRANSFER_URL)
	return t.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches a transfer having the given transferID.
func (t *Transfer) Fetch(ctx context.Context, transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new transfer for the given data.
func (t *Transfer) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.TRANSFER_URL)
	return t.Request.Post(ctx, url, data, extraHeaders)
}

// Edit edits the transfer having the given transferID.
func (t *Transfer) Edit(ctx context.Context, transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Patch(ctx, url, data, extraHeaders)
}

// Reverse reverses the transfer having the given transferID.
func (t *Transfer) Reverse(ctx context.Context, transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/reversals", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Post(ctx, url, data, extraHeaders)
}

// Reversals fetches a collection of transfer associated with the given transferID.
func (t *Transfer) Reversals(ctx context.Context, transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/reversals", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Get(ctx, url, queryParams, extraHeaders)
}
