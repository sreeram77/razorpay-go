package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPaymentID = "fake_payment_id"
const TestCaptureAmount = 500
const TestRefundAmount = 2000
const DowntimeId = "down_F7LroRQAAFuswd"
const RefundId = "fake_refund_id"

func TestPaymentAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL
	teardown, fixture := utils.StartMockServer(url, "payment_collection")
	defer teardown()
	body, err := utils.Client.Payment.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentAllWithOptions(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL
	teardown, fixture := utils.StartMockServer(url, "payment_collection_with_one_payment")
	defer teardown()
	queryParams := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Payment.All(context.Background(), queryParams, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID
	teardown, fixture := utils.StartMockServer(url, "payment_collection_with_one_payment")
	defer teardown()
	body, err := utils.Client.Payment.Fetch(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentCapture(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/capture"
	teardown, fixture := utils.StartMockServer(url, "fake_captured_payment")
	defer teardown()
	body, err := utils.Client.Payment.Capture(context.Background(), TestPaymentID, TestCaptureAmount, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentRefundCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/refund"
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	body, err := utils.Client.Payment.Refund(context.Background(), TestPaymentID, TestRefundAmount, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
func TestPaymentTransfer(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	teardown, fixture := utils.StartMockServer(url, "transfers_collection_with_payment_id")
	defer teardown()
	params := map[string]interface{}{
		"transfers": map[string]interface{}{
			"currency": map[string]interface{}{
				"amount":   100,
				"currency": "INR",
				"account":  "dummy_acc",
			},
		},
	}
	body, err := utils.Client.Payment.Transfer(context.Background(), TestPaymentID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentTransferFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	teardown, fixture := utils.StartMockServer(url, "transfers_collection_with_payment_id")
	defer teardown()
	body, err := utils.Client.Payment.Transfers(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentBankTransferFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/bank_transfer"
	teardown, fixture := utils.StartMockServer(url, "fake_bank_transfer")
	defer teardown()
	body, err := utils.Client.Payment.BankTransfer(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentCreateJsonPayment(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/create/json"
	teardown, fixture := utils.StartMockServer(url, "fake_payment_json")
	defer teardown()
	data := map[string]interface{}{
		"amount":   100,
		"currency": "INR",
		"order_id": "order_EAkbvXiCJlwhHR",
		"email":    "gaurav.kumar@example.com",
		"contact":  9090909090,
		"method":   "upi",
	}
	body, err := utils.Client.Payment.CreatePaymentJson(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentCreateRecurringPayment(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/create/recurring"
	teardown, fixture := utils.StartMockServer(url, "fake_recurring_payment")
	defer teardown()

	data := map[string]interface{}{
		"email":       "gaurav.kumar@example.com",
		"contact":     "9876543567",
		"amount":      10000,
		"currency":    "INR",
		"order_id":    "order_JIPONOfVuBVB4B",
		"customer_id": "cust_DzYEzfJLV03rkp",
		"token":       "token_JIOj6ss6Ug43Gg",
		"recurring":   "1",
		"notes": map[string]interface{}{
			"note_key_1": "Tea. Earl grey. Hot.",
			"note_key_2": "Tea. Earl grey. Decaf.",
		},
		"description": "Creating recurring payment for Gaurav Kumar",
	}

	body, err := utils.Client.Payment.CreateRecurringPayment(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentEdit(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID
	teardown, fixture := utils.StartMockServer(url, "fake_payment_json")
	defer teardown()

	data := map[string]interface{}{
		"notes": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	body, err := utils.Client.Payment.Edit(context.Background(), TestPaymentID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetchCardDetails(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/card"
	teardown, fixture := utils.StartMockServer(url, "fetch_card_details")
	defer teardown()
	body, err := utils.Client.Payment.FetchCardDetails(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetchPaymentDowntime(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/downtimes"
	teardown, fixture := utils.StartMockServer(url, "downtimes_collection")
	defer teardown()
	body, err := utils.Client.Payment.FetchPaymentDowntime(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetchPaymentDowntimeById(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/downtimes" + "/" + DowntimeId
	teardown, fixture := utils.StartMockServer(url, "fake_downtime")
	defer teardown()
	body, err := utils.Client.Payment.FetchPaymentDowntimeById(context.Background(), DowntimeId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetchMultipleRefund(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/refunds"
	teardown, fixture := utils.StartMockServer(url, "refund_collection")
	defer teardown()
	data := map[string]interface{}{
		"count": 2,
	}
	body, err := utils.Client.Payment.FetchMultipleRefund(context.Background(), TestPaymentID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetchRefund(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/refunds" + "/" + RefundId
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	body, err := utils.Client.Payment.FetchRefund(context.Background(), TestPaymentID, RefundId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOtpGenerate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/otp_generate"
	teardown, fixture := utils.StartMockServer(url, "fake_otp_generate")
	defer teardown()
	body, err := utils.Client.Payment.OtpGenerate(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOtpSubmit(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/otp/submit"
	teardown, fixture := utils.StartMockServer(url, "fake_otp_submit")
	defer teardown()
	data := map[string]interface{}{
		"otp": "123456",
	}
	body, err := utils.Client.Payment.OtpSubmit(context.Background(), TestPaymentID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOtpResend(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYMENT_URL + "/" + TestPaymentID + "/otp/resend"
	teardown, fixture := utils.StartMockServer(url, "fake_otp_resend")
	defer teardown()
	body, err := utils.Client.Payment.OtpResend(context.Background(), TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
