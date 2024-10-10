package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Settlement ...
type Settlement struct {
	Request *requests.Request
}

// All fetches collection of settlements for the given queryParams.
func (settlement *Settlement) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.SETTLEMENT_URL)
	return settlement.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Fetch fetches a settlement having the given settlementID.
func (settlement *Settlement) Fetch(ctx context.Context, settlementID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.SETTLEMENT_URL, url.PathEscape(settlementID))
	return settlement.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Reports to get combined report of settlements.
func (settlement *Settlement) Reports(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/recon/combined", constants.VERSION_V1, constants.SETTLEMENT_URL)
	return settlement.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchAllOnDemandSettlement fetches all On-demand Settlements.
func (settlement *Settlement) FetchAllOnDemandSettlement(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/ondemand", constants.VERSION_V1, constants.SETTLEMENT_URL)
	return settlement.Request.Get(ctx, url, queryParams, extraHeaders)
}

// FetchOnDemandSettlementById fetches On-demand Settlement by settlementID.
func (settlement *Settlement) FetchOnDemandSettlementById(ctx context.Context, settlementID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/ondemand/%s", constants.VERSION_V1, constants.SETTLEMENT_URL, url.PathEscape(settlementID))
	return settlement.Request.Get(ctx, url, queryParams, extraHeaders)
}

// CreateOnDemandSettlement creates a On-demand Settlement for the given data.
func (settlement *Settlement) CreateOnDemandSettlement(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/ondemand", constants.VERSION_V1, constants.SETTLEMENT_URL)
	return settlement.Request.Post(ctx, url, data, extraHeaders)
}
