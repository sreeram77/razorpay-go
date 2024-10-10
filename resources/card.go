package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Card ...
type Card struct {
	Request *requests.Request
}

// Fetch fetches card having the given cardID.
func (card *Card) Fetch(ctx context.Context, cardID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.CARD_URL, cardID)
	return card.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (card *Card) RequestCardReference(ctx context.Context, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/fingerprints", constants.VERSION_V1, constants.CARD_URL)
	return card.Request.Post(ctx, url, data, extraHeaders)
}
