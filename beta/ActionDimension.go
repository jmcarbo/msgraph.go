// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmcarbo/msgraph.go/jsonx"
)

// DimensionValues returns request builder for DimensionValue collection
func (b *DimensionRequestBuilder) DimensionValues() *DimensionDimensionValuesCollectionRequestBuilder {
	bb := &DimensionDimensionValuesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/dimensionValues"
	return bb
}

// DimensionDimensionValuesCollectionRequestBuilder is request builder for DimensionValue collection
type DimensionDimensionValuesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DimensionValue collection
func (b *DimensionDimensionValuesCollectionRequestBuilder) Request() *DimensionDimensionValuesCollectionRequest {
	return &DimensionDimensionValuesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DimensionValue item
func (b *DimensionDimensionValuesCollectionRequestBuilder) ID(id string) *DimensionValueRequestBuilder {
	bb := &DimensionValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DimensionDimensionValuesCollectionRequest is request for DimensionValue collection
type DimensionDimensionValuesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DimensionValue, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DimensionValue
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DimensionValue
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DimensionValue collection, max N pages
func (r *DimensionDimensionValuesCollectionRequest) GetN(ctx context.Context, n int) ([]DimensionValue, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Get(ctx context.Context) ([]DimensionValue, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Add(ctx context.Context, reqObj *DimensionValue) (resObj *DimensionValue, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
