package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestProductId = "acc_prd_M4y7gSYczESUDF"
const TestProductAccountId = "acc_M28vQMUgbIBo1N"

func TestProductFetch(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestProductAccountId + constants.PRODUCT_URL + "/" + TestProductId
	teardown, fixture := utils.StartMockServer(url, "fake_product")
	defer teardown()
	body, err := utils.Client.Product.Fetch(context.Background(), TestProductAccountId, TestProductId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestProductCreate(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestProductAccountId + constants.PRODUCT_URL

	teardown, fixture := utils.StartMockServer(url, "fake_product")
	defer teardown()

	param := map[string]interface{}{
		"product_name": "payment_gateway",
		"tnc_accepted": true,
		"ip":           "233.233.233.234",
	}

	body, err := utils.Client.Product.RequestProductConfiguration(context.Background(), "acc_M28vQMUgbIBo1N", param, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestProductEdit(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestProductAccountId + constants.PRODUCT_URL + "/" + TestProductId
	teardown, fixture := utils.StartMockServer(url, "fake_product")
	defer teardown()
	params := map[string]interface{}{
		"phone": map[string]interface{}{
			"primary":   "9898989898",
			"secondary": "9898989898",
		},
	}
	body, err := utils.Client.Product.Edit(context.Background(), TestProductAccountId, TestProductId, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
