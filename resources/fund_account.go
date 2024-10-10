package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// FundAccount ...
type FundAccount struct {
	Request *requests.Request
}

// All fetches collection of fund accounts for the given queryParams.
func (fa *FundAccount) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.FUND_ACCOUNT_URL)
	return fa.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Create creates a fund account for the given data.
func (fa *FundAccount) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.FUND_ACCOUNT_URL)
	return fa.Request.Post(ctx, url, data, extraHeaders)
}
