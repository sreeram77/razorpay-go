package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Stakeholder ...
type Stakeholder struct {
	Request *requests.Request
}

// Create creates a new stakeholder for the given data.
func (stake *Stakeholder) Create(ctx context.Context, accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL)
	return stake.Request.Post(ctx, url, data, extraHeaders)
}

// Fetch fetches stakeholder having the given accountId & stakeholderId.
func (stake *Stakeholder) Fetch(ctx context.Context, accountId string, stakeholderId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL, url.PathEscape(stakeholderId))
	return stake.Request.Get(ctx, url, queryParams, extraHeaders)
}

// All fetches collection of stakeholder for the given accountId.
func (stake *Stakeholder) All(ctx context.Context, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL)
	return stake.Request.Get(ctx, url, queryParams, extraHeaders)
}

// Edit updates the stakeholder having the given accountId & stakeholderId.
func (stake *Stakeholder) Edit(ctx context.Context, accountId string, stakeholderId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL, url.PathEscape(stakeholderId))
	return stake.Request.Patch(ctx, url, data, extraHeaders)
}

// UploadStakeholderDoc upload stakeholder documents.
func (stake *Stakeholder) UploadStakeholderDoc(ctx context.Context, accountId string, stakeholderId string, params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/stakeholders/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), url.PathEscape(stakeholderId))
	return stake.Request.File(ctx, url, params, extraHeaders)
}

// FetchStakeholderDoc fetches the stakeholder documents.
func (stake *Stakeholder) FetchStakeholderDoc(ctx context.Context, accountId string, stakeholderId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/stakeholders/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), url.PathEscape(stakeholderId))
	return stake.Request.Get(ctx, url, queryParams, extraHeaders)
}
