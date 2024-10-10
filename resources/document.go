package resources

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sreeram77/razorpay-go/constants"
	"github.com/sreeram77/razorpay-go/requests"
)

// Document ...
type Document struct {
	Request *requests.Request
}

// Create a Document
func (doc *Document) Create(ctx context.Context, params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.DOCUMENT)
	return doc.Request.File(ctx, url, params, extraHeaders)
}

// Fetch document by given documentId.
func (doc *Document) Fetch(ctx context.Context, documentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.DOCUMENT, url.PathEscape(documentId))
	return doc.Request.Get(ctx, url, queryParams, extraHeaders)
}
