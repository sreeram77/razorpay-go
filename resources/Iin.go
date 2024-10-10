package resources

import (
	"context"
	"fmt"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Iin ...
type Iin struct {
	Request *requests.Request
}

func (i *Iin) Fetch(ctx context.Context, tokenIin string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.IIN, tokenIin)
	return i.Request.Get(ctx, url, queryParams, extraHeaders)
}

func (i *Iin) All(ctx context.Context, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/list", constants.VERSION_V1, constants.IIN)
	return i.Request.Get(ctx, url, queryParams, extraHeaders)
}
