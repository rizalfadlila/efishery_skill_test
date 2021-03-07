package service

import "context"

// Service :nodoc:
type Service interface {
	Fetch(ctx context.Context) (interface{}, error)
	Aggregate(ctx context.Context) (interface{}, error)
}
