package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestRefundID = "fake_refund_id"

func TestRefundAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.REFUND_URL
	teardown, fixture := utils.StartMockServer(url, "refund_collection")
	defer teardown()
	body, err := utils.Client.Refund.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRefundFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.REFUND_URL + "/" + TestRefundID
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	body, err := utils.Client.Refund.Fetch(context.Background(), TestRefundID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRefundCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.REFUND_URL
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	params := map[string]interface{}{
		"payment_id": TestPaymentID,
	}
	body, err := utils.Client.Refund.Create(context.Background(), params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRefundUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.REFUND_URL + "/" + TestRefundID
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	params := map[string]interface{}{
		"notes": map[string]interface{}{
			"notes_key_1": "Beam me up Scotty.",
			"notes_key_2": "Engage",
		},
	}
	body, err := utils.Client.Refund.Update(context.Background(), TestRefundID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
