package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fetch_app/constants"
	"github.com/fetch_app/entities/models"
	"github.com/fetch_app/helper"
	"github.com/fetch_app/pkg/logger"
	"github.com/fetch_app/usecases/apicall"

	"github.com/patrickmn/go-cache"
)

type service struct {
	caller        apicall.ResourceCaller
	goCahce       *cache.Cache
	defaultErrMsg string
}

// NewService :nodoc
func NewService() Service {
	return &service{
		caller:        apicall.NewResourceCaller(),
		goCahce:       cache.New(60*time.Minute, 70*time.Minute),
		defaultErrMsg: "Error at Service",
	}
}

func (s *service) Fetch(ctx context.Context) (interface{}, error) {
	var err error
	result := make([]models.Fetch, 0)

	helper.Block{
		Try: func() {
			resources, err := s.caller.CallFetchResource(ctx)

			helper.Throw(err)

			for _, v := range resources {
				newConversion := 0.0
				conversion, found := s.goCahce.Get(constants.ChaceConvUSDKey)

				if found {
					newConversion = conversion.(float64)
				} else {
					response, err := s.caller.CallCurrencyConverter(ctx)
					helper.Throw(err)
					s.goCahce.Set(constants.ChaceConvUSDKey, response[constants.ChaceConvUSDKey].(float64), cache.DefaultExpiration)
					newConversion = response[constants.ChaceConvUSDKey].(float64)
				}

				priceIDR, _ := strconv.ParseFloat(v.PriceIDR, 64)
				priceUSD := newConversion * priceIDR

				v.PriceUSD = fmt.Sprintf("%f", priceUSD)

				result = append(result, v)
			}
		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > Fetch: %v", s.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}
