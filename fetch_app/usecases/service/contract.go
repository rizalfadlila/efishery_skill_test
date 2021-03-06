package service

import "context"

// Service :nodoc:
type Service interface {
	Fetch(ctx context.Context) (interface{}, error)
}
