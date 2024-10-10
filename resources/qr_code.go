package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// QrCode ...
type QrCode struct {
	Request *requests.Request
}

// Create creates a new qrcode for the given data.
func (q *QrCode) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.QRCODE_URL)
	return q.Request.Post(ctx, url, data, extraHeaders)
}

// FetchPayments fetches the qrcode entity having the given QrCodeID.
func (q *QrCode) FetchPayments(ctx context.Context, QrCodeID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/payments", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches the QrCode entity having the given QrCodeID.
func (q *QrCode) Fetch(ctx context.Context, QrCodeID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Get(ctx, url, queryParams, extraHeaders)
}

// All fetches collection of QrCode for the given queryParams.
func (q *QrCode) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.QRCODE_URL)
	return q.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Close closes a QrCode for the given QrCodeID.
func (q *QrCode) Close(ctx context.Context, QrCodeID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/close", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Post(ctx, url, data, extraHeaders)
}
