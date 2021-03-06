package apicall

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gojektech/heimdall"
)

type baseCaller struct{}

func (b *baseCaller) get(ctx context.Context, url string, c heimdall.Client) (string, error) {
	var result string

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	res, err := c.Get(url, headers)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	result = string(body)

	if res.StatusCode != 200 {
		return result, errors.New(result)
	}

	return result, nil
}
