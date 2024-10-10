package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Plan ...
type Plan struct {
	Request *requests.Request
}

// Create creates a new plan for the given data.
func (plan *Plan) Create(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PLAN_URL)
	return plan.Request.Post(ctx, url, data, extraHeaders)
}

// Fetch fetches the plan entity having the given planID.
func (plan *Plan) Fetch(ctx context.Context, planID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PLAN_URL, planID)
	return plan.Request.Get(ctx, url, queryParams, extraHeaders)
}

// All fetches collection of plans for the given queryParams.
func (plan *Plan) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PLAN_URL)
	return plan.Request.Get(ctx, url, queryParams, extraHeaders)
}
