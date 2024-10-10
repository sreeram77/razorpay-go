package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Dispute ...
type Dispute struct {
	Request *requests.Request
}

func (acc *Dispute) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.DISPUTE)
	return acc.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (acc *Dispute) Fetch(ctx context.Context, disputeId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (acc *Dispute) Accept(ctx context.Context, disputeId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/accept", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Post(ctx, url, data, extraHeaders)
}

func (acc *Dispute) Contest(ctx context.Context, disputeId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/contest", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Patch(ctx, url, data, extraHeaders)
}
