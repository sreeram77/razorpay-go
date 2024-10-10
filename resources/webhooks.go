package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Webhook ...
type Webhook struct {
	Request *requests.Request
}

// Create creates a new webhook for the given data.
func (wh *Webhook) Create(ctx context.Context, accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK)
		return wh.Request.Post(ctx, url, data, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.WEBHOOK)
	return wh.Request.Post(ctx, url, data, extraHeaders)
}

// Fetch retrieve and view the details of a webhook.
func (wh *Webhook) Fetch(ctx context.Context, webhookId string, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
	return wh.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Edit updates the details of a webhook.
func (wh *Webhook) Edit(ctx context.Context, webhookId string, accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
		return wh.Request.Patch(ctx, url, data, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.WEBHOOK, url.PathEscape(accountId))
	return wh.Request.Put(ctx, url, data, extraHeaders)
}

// All fetches collection of webhook.
func (wh *Webhook) All(ctx context.Context, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK)
		return wh.Request.Get(ctx, url, queryParams, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.WEBHOOK)
	return wh.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Delete deletes a webhook having the given invoiceID.
func (wh *Webhook) Delete(ctx context.Context, webhookId string, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
	return wh.Request.Delete(ctx, url, queryParams, extraHeaders)
}
