// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmcarbo/msgraph.go/jsonx"
)

// Tabs returns request builder for TeamsTab collection
func (b *ChannelRequestBuilder) Tabs() *ChannelTabsCollectionRequestBuilder {
	bb := &ChannelTabsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tabs"
	return bb
}

// ChannelTabsCollectionRequestBuilder is request builder for TeamsTab collection
type ChannelTabsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsTab collection
func (b *ChannelTabsCollectionRequestBuilder) Request() *ChannelTabsCollectionRequest {
	return &ChannelTabsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsTab item
func (b *ChannelTabsCollectionRequestBuilder) ID(id string) *TeamsTabRequestBuilder {
	bb := &TeamsTabRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelTabsCollectionRequest is request for TeamsTab collection
type ChannelTabsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsTab, error) {
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
	var values []TeamsTab
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
			value  []TeamsTab
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

// GetN performs GET request for TeamsTab collection, max N pages
func (r *ChannelTabsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsTab, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Get(ctx context.Context) ([]TeamsTab, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Add(ctx context.Context, reqObj *TeamsTab) (resObj *TeamsTab, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
