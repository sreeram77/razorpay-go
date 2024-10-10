package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Customer ...
type Customer struct {
	Request *requests.Request
}

// Fetch fetches customer having the given cutomerID.
func (cust *Customer) Fetch(ctx context.Context, customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))
	return cust.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a new customer for the given data.
func (cust *Customer) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Post(ctx, url, data, extraHeaders)
}

// Edit updates the customer having the given customerID.
func (cust *Customer) Edit(ctx context.Context, customerID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))

	return cust.Request.Put(ctx, url, data, extraHeaders)
}

// All fetches collection of customer for the given queryParams.
func (cust *Customer) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (cust *Customer) AddBankAccount(ctx context.Context, customerID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/bank_account", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))
	return cust.Request.Post(ctx, url, data, extraHeaders)
}

func (cust *Customer) DeleteBankAccount(ctx context.Context, customerID string, bankAccountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/bank_account/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID), url.PathEscape(bankAccountId))
	return cust.Request.Delete(ctx, url, queryParams, extraHeaders)
}

func (cust *Customer) RequestEligibilityCheck(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/eligibility", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Post(ctx, url, data, extraHeaders)
}

func (cust *Customer) FetchEligibility(ctx context.Context, eligibilityId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/eligibility/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(eligibilityId))
	return cust.Request.Get(ctx, url, queryParams, extraHeaders)
}
