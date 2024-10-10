package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Payment ...
type Payment struct {
	Request *requests.Request
}

// All fetches multiple payment entities for the given queryParams.
func (p *Payment) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches the payment entity for the given paymentID.
func (p *Payment) Fetch(ctx context.Context, paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))

	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Capture captures the payment having the given paymentID.
func (p *Payment) Capture(ctx context.Context, paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s/capture", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(ctx, url, data, extraHeaders)
}

// Refund initiates a refund for the given paymentID.
func (p *Payment) Refund(ctx context.Context, paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s/refund", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(ctx, url, data, extraHeaders)
}

// Transfer creates a transfer of the payment having the given paymentID.
func (p *Payment) Transfer(ctx context.Context, paymentID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/transfers", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Post(ctx, url, data, extraHeaders)
}

// Transfers fetches collection of all transfers associated with the given paymentID.
func (p *Payment) Transfers(ctx context.Context, paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/transfers", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// BankTransfer fetches BankTransfer associated with the given paymentID.
func (p *Payment) BankTransfer(ctx context.Context, paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/bank_transfer", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// CreatePaymentJson creates a json payment for the given data.
func (p *Payment) CreatePaymentJson(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/create/json", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Post(ctx, url, data, extraHeaders)
}

// CreateRecurringPayment creates a recurring payment for the given data.
func (p *Payment) CreateRecurringPayment(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/create/recurring", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Post(ctx, url, data, extraHeaders)
}

// Edit updates the payment having the given paymentID.
func (p *Payment) Edit(ctx context.Context, paymentID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Patch(ctx, url, data, extraHeaders)
}

// FetchCardDetails fetches card details with the given paymentID.
func (p *Payment) FetchCardDetails(ctx context.Context, paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/card", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchPaymentDowntime fetches downtime details.
func (p *Payment) FetchPaymentDowntime(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/downtimes", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchPaymentDowntimeById fetches downtime details with the given downtimeID.
func (p *Payment) FetchPaymentDowntimeById(ctx context.Context, downtimeId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/downtimes/%s", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(downtimeId))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchMultipleRefund fetches multiple refunds details with the given paymentID.
func (p *Payment) FetchMultipleRefund(ctx context.Context, paymentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/refunds", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentId))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchRefund fetches refund detail with the given paymentId and refundId
func (p *Payment) FetchRefund(ctx context.Context, paymentId string, refundId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/refunds/%s", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentId), url.PathEscape(refundId))
	return p.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (p *Payment) CreateUpi(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/create/upi", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Post(ctx, url, data, extraHeaders)
}

func (p *Payment) ValidateVpa(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/validate/vpa", constants.VERSION_V1, constants.PAYMENT_URL)
	return p.Request.Post(ctx, url, data, extraHeaders)
}

func (p *Payment) FetchMethods(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.METHODS_URL)
	return p.Request.Get(ctx, url, data, extraHeaders)
}

func (p *Payment) OtpGenerate(ctx context.Context, paymentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/otp_generate", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentId))
	return p.Request.Post(ctx, url, queryParams, extraHeaders)
}

func (p *Payment) OtpSubmit(ctx context.Context, paymentId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/otp/submit", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentId))
	return p.Request.Post(ctx, url, data, extraHeaders)
}

func (p *Payment) OtpResend(ctx context.Context, paymentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/otp/resend", constants.VERSION_V1, constants.PAYMENT_URL, url.PathEscape(paymentId))
	return p.Request.Post(ctx, url, queryParams, extraHeaders)
}
