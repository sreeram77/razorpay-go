package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestSubscriptionID = "subscription_id"

func TestSubscriptionAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL
	teardown, fixture := utils.StartMockServer(url, "subscription_collection")
	defer teardown()
	body, err := utils.Client.Subscription.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.Fetch(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	params := map[string]interface{}{
		"plan_id":     "plan_00000000000001",
		"total_count": 12,
		"quantity":    1,
		"start_at":    1561852800,
		"expire_by":   1561939199,
	}
	body, err := utils.Client.Subscription.Create(context.Background(), params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionCancel(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/cancel"
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.Cancel(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionPendingUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/retrieve_scheduled_changes"
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.PendingUpdate(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionCancelScheduledChanges(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/cancel_scheduled_changes"
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.CancelScheduledChanges(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionPause(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/pause"
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.Pause(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionResume(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/resume"
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.Resume(context.Background(), TestSubscriptionID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	data := map[string]interface{}{
		"plan_id":            "plan_00000000000002",
		"offer_id":           "offer_JHD834hjbxzhd38d",
		"quantity":           5,
		"remaining_count":    5,
		"start_at":           1496000432,
		"schedule_change_at": "now",
		"customer_notify":    1,
	}
	body, err := utils.Client.Subscription.Update(context.Background(), TestSubscriptionID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSubscriptionDeleteOffer(t *testing.T) {
	TestSubscriptionOfferId := "offer_JGs9gTCOKOUXck"
	url := "/" + constants.VERSION_V1 + constants.SUBSCRIPTION_URL + "/" + TestSubscriptionID + "/" + TestSubscriptionOfferId
	teardown, fixture := utils.StartMockServer(url, "fake_subscription")
	defer teardown()
	body, err := utils.Client.Subscription.DeleteOffer(context.Background(), TestSubscriptionID, TestSubscriptionOfferId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
