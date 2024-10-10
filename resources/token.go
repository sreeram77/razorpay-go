package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Token ...
type Token struct {
	Request *requests.Request
}

// Create creates a new token for the given data.
func (t *Token) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens", constants.VERSION_V1)
	return t.Request.Post(ctx, url, data, extraHeaders)
}

// Fetch fetches a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Fetch(ctx context.Context, customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Get(ctx, url, queryParams, extraHeaders)
}

// All fetches collection of tokens for associated with a customer having the given customerID.
func (t *Token) All(ctx context.Context, customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens", constants.VERSION_V1, constants.CUSTOMER_URL, customerID)
	return t.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Delete deletes a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Delete(ctx context.Context, customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Delete(ctx, url, queryParams, extraHeaders)
}

// FetchCardPropertiesByToken fetches card properties of an existing token.
func (t *Token) FetchCardPropertiesByToken(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/fetch", constants.VERSION_V1)
	return t.Request.Post(ctx, url, data, extraHeaders)
}

// DeleteToken deletes a token.
func (t *Token) DeleteToken(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/delete", constants.VERSION_V1)
	return t.Request.Post(ctx, url, data, extraHeaders)
}

// ProcessPaymentOnAlternatePAorPG process a payment on the tokenised card on another PA/PG with token
// created on razorpay
func (t *Token) ProcessPaymentOnAlternatePAorPG(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/service_provider_tokens/token_transactional_data", constants.VERSION_V1)
	return t.Request.Post(ctx, url, data, extraHeaders)
}
