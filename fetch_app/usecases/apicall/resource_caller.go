package apicall

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fetch_app/constants"
	"github.com/fetch_app/entities/models"
	"github.com/fetch_app/helper"
	"github.com/fetch_app/pkg/logger"
	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
)

type resourceCaller struct {
	baseCaller
	httpClient    heimdall.Client
	defaultErrMsg string
}

// NewResourceCaller :nodoc:
func NewResourceCaller() ResourceCaller {
	return &resourceCaller{
		httpClient:    httpclient.NewClient(),
		defaultErrMsg: "Error at ResourceCaller",
	}
}

func (c *resourceCaller) CallFetchResource(ctx context.Context) ([]models.Fetch, error) {
	var err error
	result := make([]models.Fetch, 0)

	helper.Block{
		Try: func() {
			apiURL := "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"

			response, errRes := c.get(ctx, apiURL, c.httpClient)

			helper.Throw(errRes)
			helper.Throw(json.Unmarshal([]byte(response), &result))
		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > CallFetchResource: %v", c.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}

func (c *resourceCaller) CallCurrencyConverter(ctx context.Context) (interface{}, error) {
	var err error
	var result interface{}

	helper.Block{
		Try: func() {
			apiURL := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=%s", constants.ConverterAPIKey)

			response, errRes := c.get(ctx, apiURL, c.httpClient)

			helper.Throw(errRes)
			helper.Throw(json.Unmarshal([]byte(response), &result))
		},
		Catch: func(e helper.Exception) {
			logger.Error(fmt.Sprintf("%v > CallCurrencyConverter: %v", c.defaultErrMsg, e))
			err = e
		},
	}.Do()

	return result, err
}
