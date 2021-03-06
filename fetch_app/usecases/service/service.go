package service

import (
	"context"
	"fmt"

	"github.com/fetch_app/entities/models"
	"github.com/fetch_app/helper"
	"github.com/fetch_app/pkg/logger"
	"github.com/fetch_app/usecases/apicall"
)

type service struct {
	caller        apicall.ResourceCaller
	defaultErrMsg string
}

// NewService :nodoc
func NewService() Service {
	return &service{
		caller:        apicall.NewResourceCaller(),
		defaultErrMsg: "Error at Service",
	}
}

func (s *service) Fetch(ctx context.Context) (interface{}, error) {
	var err error
	result := make([]models.Fetch, 0)

	helper.Block{
		Try: func() {

		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > Fetch: %v", s.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}
