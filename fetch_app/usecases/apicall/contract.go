package apicall

import (
	"context"

	"github.com/fetch_app/entities/models"
)

// ResourceCaller :nodoc:
type ResourceCaller interface {
	CallFetchResource(ctx context.Context) ([]models.Fetch, error)
	CallCurrencyConverter(ctx context.Context) (map[string]interface{}, error)
}
