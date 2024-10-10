package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPaymentLinkID = "payment_id"

func TestPaymentLinkAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL
	teardown, fixture := utils.StartMockServer(url, "paymentlink_collection")
	defer teardown()
	body, err := utils.Client.PaymentLink.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL + "/" + TestPaymentLinkID
	teardown, fixture := utils.StartMockServer(url, "paymentlink_collection")
	defer teardown()
	body, err := utils.Client.PaymentLink.Fetch(context.Background(), TestPaymentLinkID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL
	teardown, fixture := utils.StartMockServer(url, "fake_paymentlink")
	defer teardown()
	line_item := map[string]interface{}{
		"name":   "name",
		"amount": 1000,
	}
	lineItems := []map[string]interface{}{line_item}
	data := map[string]interface{}{
		"type":        "link",
		"decsription": "test",
		"line_items":  lineItems,
	}
	body, err := utils.Client.PaymentLink.Create(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkCancel(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL + "/" + TestPaymentLinkID + "/cancel"
	teardown, fixture := utils.StartMockServer(url, "fake_paymentlink")
	defer teardown()
	body, err := utils.Client.PaymentLink.Cancel(context.Background(), TestPaymentLinkID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkNotifyBy(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL + "/" + TestPaymentLinkID + "/notify_by/email"
	teardown, fixture := utils.StartMockServer(url, "fake_paymentlink")
	defer teardown()
	body, err := utils.Client.PaymentLink.NotifyBy(context.Background(), TestPaymentLinkID, "email", nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PaymentLink_URL + "/" + TestPaymentLinkID
	teardown, fixture := utils.StartMockServer(url, "fake_paymentlink")
	defer teardown()
	data := map[string]interface{}{
		"reference_id":    "TS35",
		"expire_by":       1653347540,
		"reminder_enable": false,
		"notes": map[string]interface{}{
			"policy_name": "Jeevan Saral",
		},
	}
	body, err := utils.Client.PaymentLink.Update(context.Background(), TestPaymentLinkID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
