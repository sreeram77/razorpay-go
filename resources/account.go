package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Account ...
type Account struct {
	Request *requests.Request
}

// Create creates a new account for the given data.
func (acc *Account) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL)
	return acc.Request.Post(ctx, url, data, extraHeaders)
}

// Fetch fetches account having the given accountId.
func (acc *Account) Fetch(ctx context.Context, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Edit updates the account having the given accountId.
func (acc *Account) Edit(ctx context.Context, accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Patch(ctx, url, data, extraHeaders)
}

// Delete deletes the account having the given accountId.
func (acc *Account) Delete(ctx context.Context, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Delete(ctx, url, queryParams, extraHeaders)
}

// UploadAccountDoc upload the account documents.
func (acc *Account) UploadAccountDoc(ctx context.Context, accountId string, params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.File(ctx, url, params, extraHeaders)
}

// FetchAccountDoc fetches the account documents.
func (acc *Account) FetchAccountDoc(ctx context.Context, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(ctx, url, queryParams, extraHeaders)
}
