package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestCardID = "fake_card_id"

func TestCardFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CARD_URL + "/" + TestCardID
	teardown, fixture := utils.StartMockServer(url, "fake_card")
	defer teardown()
	body, err := utils.Client.Card.Fetch(context.Background(), TestCardID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCardRequestCardReference(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CARD_URL + "/fingerprints"
	teardown, fixture := utils.StartMockServer(url, "fake_card_reference")
	defer teardown()
	data := map[string]interface{}{
		"number": "4854980604708430",
	}
	body, err := utils.Client.Card.RequestCardReference(context.Background(), data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
