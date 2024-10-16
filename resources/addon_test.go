package resources_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestAddonID = "ao_8sg8LU73Y3ieav"

func TestFetchAddon(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ADDON_URL, TestAddonID)
	teardown, fixture := utils.StartMockServer(url, "fake_addon")
	defer teardown()
	body, err := utils.Client.Addon.Fetch(context.Background(), TestAddonID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAddonDelete(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ADDON_URL, TestAddonID)
	teardown, fixture := utils.StartMockServer(url, "fake_addon")
	defer teardown()
	body, err := utils.Client.Addon.Delete(context.Background(), TestAddonID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAddonAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ADDON_URL)
	teardown, fixture := utils.StartMockServer(url, "addon_collection")
	defer teardown()
	body, err := utils.Client.Addon.All(context.Background(), nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
