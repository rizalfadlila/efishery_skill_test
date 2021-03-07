package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fetch_app/constants"
	"github.com/fetch_app/entities/models"
	"github.com/fetch_app/helper"
	"github.com/fetch_app/pkg/logger"
	"github.com/fetch_app/pkg/util"
	"github.com/fetch_app/usecases/apicall"
	"github.com/patrickmn/go-cache"

	linq "github.com/ahmetb/go-linq/v3"
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
	result := make([]models.Resource, 0)

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

func (s *service) Aggregate(ctx context.Context) (interface{}, error) {
	var err error = nil
	result := make([]interface{}, 0)

	helper.Block{
		Try: func() {
			resources, err := s.groupingByWeekly(ctx)

			helper.Throw(err)

			result = resources
		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > Aggregate: %v", s.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}

func (s *service) groupingByWeekly(ctx context.Context) ([]interface{}, error) {
	var err error = nil
	result := make([]interface{}, 0)

	helper.Block{
		Try: func() {
			resources, err := s.caller.CallFetchResource(ctx)

			helper.Throw(err)

			linq.From(resources).
				WhereT(func(r models.Resource) bool {
					return r.Province != "" && r.Date != ""
				}).
				GroupByT(
					func(r models.Resource) string {
						return util.GetWeekByDateString(r.Date)
					},
					func(r models.Resource) string {
						return fmt.Sprintf("%v_%v", r.Province, r.Date)
					}).
				SelectIndexedT(func(index int, resourceGroup linq.Group) map[string]interface{} {
					var result map[string]interface{}
					if resourceGroup.Key.(string) != "" {

						max := linq.From(resourceGroup.Group).Max()
						min := linq.From(resourceGroup.Group).Min()

						aggregate := map[string]interface{}{
							"min": setAggregateResult(min),
							"max": setAggregateResult(max),
						}

						keys := strings.Split(resourceGroup.Key.(string), "_")

						result = map[string]interface{}{
							"week":      keys[0],
							"year":      keys[1],
							"aggregate": aggregate,
						}
					}
					return result
				}).
				ToSlice(&result)
		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > groupByWeekly: %v", s.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}

func extractValue(value interface{}) []string {
	values := strings.Split(value.(string), "_")
	return values
}

func setAggregateResult(values interface{}) map[string]interface{} {
	return map[string]interface{}{
		"province": extractValue(values)[0],
		"date":     extractValue(values)[1],
	}
}
