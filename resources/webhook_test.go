package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestWebhookID = "fake_webhook_id"

func TestWebhookAll(t *testing.T) {
	url := "/" + constants.VERSION_V2 + constants.ACCOUNT_URL + "/" + TestAccountID + constants.WEBHOOK
	teardown, fixture := utils.StartMockServer(url, "webhook_collection")
	defer teardown()
	body, err := utils.Client.Webhook.All(context.Background(), TestAccountID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestWebhookFetch(t *testing.T) {
	url := "/" + constants.VERSION_V2 + constants.ACCOUNT_URL + "/" + TestAccountID + constants.WEBHOOK + "/" + TestWebhookID
	teardown, fixture := utils.StartMockServer(url, "fake_webhook")
	defer teardown()
	body, err := utils.Client.Webhook.Fetch(context.Background(), TestWebhookID, TestAccountID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestWebhookCreate(t *testing.T) {
	url := "/" + constants.VERSION_V2 + constants.ACCOUNT_URL + "/" + TestAccountID + constants.WEBHOOK
	teardown, fixture := utils.StartMockServer(url, "fake_webhook")
	defer teardown()

	events := make(map[string]string)
	events["0"] = "payment.authorized"
	events["1"] = "payment.failed"
	events["2"] = "payment.captured"
	events["3"] = "payment.dispute.created"
	events["4"] = "refund.failed"
	events["5"] = "refund.created"

	data := map[string]interface{}{
		"url":         "https://google.com",
		"alert_email": "gaurav.kumar@example.com",
		"secret":      "12345",
		"events":      events,
	}
	body, err := utils.Client.Webhook.Create(context.Background(), TestAccountID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestWebhookEdit(t *testing.T) {
	url := "/" + constants.VERSION_V2 + constants.ACCOUNT_URL + "/" + TestAccountID + constants.WEBHOOK + "/" + TestWebhookID
	teardown, fixture := utils.StartMockServer(url, "fake_webhook")
	defer teardown()

	events := make(map[string]string)
	events["0"] = "payment.authorized"
	events["1"] = "payment.failed"

	data := map[string]interface{}{
		"url":    "https://www.linkedin.com",
		"events": events,
	}
	body, err := utils.Client.Webhook.Edit(context.Background(), TestWebhookID, TestAccountID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
