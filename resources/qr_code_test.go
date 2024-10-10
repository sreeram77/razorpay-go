package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestQrCodeID = "qr_code_id"

func TestQrCodeAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.QRCODE_URL
	teardown, fixture := utils.StartMockServer(url, "qr_code_collection")
	defer teardown()
	body, err := utils.Client.QrCode.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestQrCodeFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.QRCODE_URL + "/" + TestQrCodeID
	teardown, fixture := utils.StartMockServer(url, "fake_qr_code")
	defer teardown()
	body, err := utils.Client.QrCode.Fetch(context.Background(), TestQrCodeID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestQrCodeFetchPayments(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.QRCODE_URL + "/" + TestQrCodeID + "/payments"
	teardown, fixture := utils.StartMockServer(url, "fake_qr_code_payments")
	defer teardown()
	body, err := utils.Client.QrCode.FetchPayments(context.Background(), TestQrCodeID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestQrCodeClose(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.QRCODE_URL + "/" + TestQrCodeID + "/close"
	teardown, fixture := utils.StartMockServer(url, "fake_qr_code_payments")
	defer teardown()
	body, err := utils.Client.QrCode.Close(context.Background(), TestQrCodeID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestQrCodeCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.QRCODE_URL
	teardown, fixture := utils.StartMockServer(url, "fake_qr_code")
	defer teardown()
	data := map[string]interface{}{
		"type":           "upi_qr",
		"name":           "Store_1",
		"usage":          "single_use",
		"fixed_amount":   true,
		"payment_amount": 300,
		"description":    "For Store 1",
		"customer_id":    "cust_JFz35u2L3c6KJl",
	}
	body, err := utils.Client.QrCode.Create(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
