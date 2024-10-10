package resources_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestAccountID = "acc_M28vQMUgbIBo1N"

func TestFetchAccount(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestAccountID
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	body, err := utils.Client.Account.Fetch(context.Background(), TestAccountID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAccountCreate(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	b, err := ioutil.ReadFile("../testdata/fake_account_request.json")
	if err != nil {
		panic(err)
	}
	fixture = strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)

	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(fixture), &mapData); err != nil {
		fmt.Println(err)
	}

	body, err := utils.Client.Account.Create(context.Background(), mapData, nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, body["email"], "gauriagain.kumar@example.org")
}

func TestAccountEdit(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestAccountID
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	params := map[string]interface{}{
		"name":  "test",
		"email": "test@test.com",
	}
	body, err := utils.Client.Account.Edit(context.Background(), TestAccountID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAccountDelete(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestAccountID
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	body, err := utils.Client.Account.Delete(context.Background(), TestAccountID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAccountFetchDoc(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestAccountID + "/documents"
	teardown, fixture := utils.StartMockServer(url, "fake_doc_data")
	defer teardown()
	body, err := utils.Client.Account.FetchAccountDoc(context.Background(), TestAccountID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
