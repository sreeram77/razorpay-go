package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Item ...
type Item struct {
	Request *requests.Request
}

// All fetches collection of items for the given queryParams.
func (item *Item) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ITEM_URL)
	return item.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches an item having the given itemID.
func (item *Item) Fetch(ctx context.Context, itemID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new item for the given data.
func (item *Item) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ITEM_URL)
	return item.Request.Post(ctx, url, data, extraHeaders)
}

// Update updates an item for the given data.
func (item *Item) Update(ctx context.Context, itemID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Patch(ctx, url, data, extraHeaders)
}

// Delete deletes an item having the given itemID.
func (item *Item) Delete(ctx context.Context, itemID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Delete(ctx, url, queryParams, extraHeaders)
}
