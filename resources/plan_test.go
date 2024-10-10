package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPlanID = "fake_plan_id"

func TestPlanAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PLAN_URL
	teardown, fixture := utils.StartMockServer(url, "plan_collection")
	defer teardown()
	body, err := utils.Client.Plan.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPlanFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PLAN_URL + "/" + TestPlanID
	teardown, fixture := utils.StartMockServer(url, "fake_plan")
	defer teardown()
	body, err := utils.Client.Plan.Fetch(context.Background(), TestPlanID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPlanCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PLAN_URL
	teardown, fixture := utils.StartMockServer(url, "fake_plan")
	defer teardown()
	data := map[string]interface{}{
		"period":   "weekly",
		"interval": 1,
		"item": map[string]interface{}{
			"name":        "Test plan - Weekly",
			"amount":      69900,
			"currency":    "INR",
			"description": "Description for the test plan",
		},
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}
	body, err := utils.Client.Plan.Create(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
