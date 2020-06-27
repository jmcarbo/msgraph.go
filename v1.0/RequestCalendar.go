// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CalendarRequestBuilder is request builder for Calendar
type CalendarRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarRequest
func (b *CalendarRequestBuilder) Request() *CalendarRequest {
	return &CalendarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarRequest is request for Calendar
type CalendarRequest struct{ BaseRequest }

// Get performs GET request for Calendar
func (r *CalendarRequest) Get(ctx context.Context) (resObj *Calendar, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Calendar
func (r *CalendarRequest) Update(ctx context.Context, reqObj *Calendar) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Calendar
func (r *CalendarRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarGroupRequestBuilder is request builder for CalendarGroup
type CalendarGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarGroupRequest
func (b *CalendarGroupRequestBuilder) Request() *CalendarGroupRequest {
	return &CalendarGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarGroupRequest is request for CalendarGroup
type CalendarGroupRequest struct{ BaseRequest }

// Get performs GET request for CalendarGroup
func (r *CalendarGroupRequest) Get(ctx context.Context) (resObj *CalendarGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarGroup
func (r *CalendarGroupRequest) Update(ctx context.Context, reqObj *CalendarGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarGroup
func (r *CalendarGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type CalendarGetScheduleRequestBuilder struct{ BaseRequestBuilder }

// GetSchedule action undocumented
func (b *CalendarRequestBuilder) GetSchedule(reqObj *CalendarGetScheduleRequestParameter) *CalendarGetScheduleRequestBuilder {
	bb := &CalendarGetScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getSchedule"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CalendarGetScheduleRequest struct{ BaseRequest }

//
func (b *CalendarGetScheduleRequestBuilder) Request() *CalendarGetScheduleRequest {
	return &CalendarGetScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CalendarGetScheduleRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ScheduleInformation, error) {
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
	var values []ScheduleInformation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ScheduleInformation
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
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

//
func (r *CalendarGetScheduleRequest) PostN(ctx context.Context, n int) ([]ScheduleInformation, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *CalendarGetScheduleRequest) Post(ctx context.Context) ([]ScheduleInformation, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}