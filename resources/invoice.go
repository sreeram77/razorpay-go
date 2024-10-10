package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Invoice ...
type Invoice struct {
	Request *requests.Request
}

// All fetches multiple invoices for the given queryParams.
func (inv *Invoice) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.INVOICE_URL)
	return inv.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches an invoice having the given invoiceID.
func (inv *Invoice) Fetch(ctx context.Context, invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return inv.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new invoice for the given data.
func (inv *Invoice) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.INVOICE_URL)
	return inv.Request.Post(ctx, url, data, extraHeaders)
}

// Cancel cancels a invoice having the given invoiceID.
func (inv *Invoice) Cancel(ctx context.Context, invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/cancel", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return inv.Request.Post(ctx, url, queryParams, extraHeaders)
}

// Delete delete a draft invoice having the given invoiceID.
func (inv *Invoice) Delete(ctx context.Context, invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return inv.Request.Delete(ctx, url, queryParams, extraHeaders)
}

// Update updates the invoice having the given invoiceID.
func (inv *Invoice) Update(ctx context.Context, invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return inv.Request.Patch(ctx, url, data, extraHeaders)
}

// Issue issues an invoice having the given invoiceID.
func (inv *Invoice) Issue(ctx context.Context, invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/issue", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return inv.Request.Post(ctx, url, data, extraHeaders)
}

// Notify notify to customer via email or sms by the given invoiceID.
func (inv *Invoice) Notify(ctx context.Context, invoiceID string, medium string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/notify_by/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID, medium)
	return inv.Request.Post(ctx, url, data, extraHeaders)
}

// CreateRegistrationLink creates a registration link
func (inv *Invoice) CreateRegistrationLink(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/subscription_registration/auth_links", constants.VERSION_V1)
	return inv.Request.Post(ctx, url, data, extraHeaders)
}
