// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmcarbo/msgraph.go/jsonx"
)

// PolicySetCollectionGetPolicySetsRequestParameter undocumented
type PolicySetCollectionGetPolicySetsRequestParameter struct {
	// PolicySetIDs undocumented
	PolicySetIDs []string `json:"policySetIds,omitempty"`
}

// PolicySetUpdateRequestParameter undocumented
type PolicySetUpdateRequestParameter struct {
	// AddedPolicySetItems undocumented
	AddedPolicySetItems []PolicySetItem `json:"addedPolicySetItems,omitempty"`
	// UpdatedPolicySetItems undocumented
	UpdatedPolicySetItems []PolicySetItem `json:"updatedPolicySetItems,omitempty"`
	// DeletedPolicySetItems undocumented
	DeletedPolicySetItems []string `json:"deletedPolicySetItems,omitempty"`
	// Assignments undocumented
	Assignments []PolicySetAssignment `json:"assignments,omitempty"`
}

// AppliesTo returns request builder for DirectoryObject collection
func (b *PolicyRequestBuilder) AppliesTo() *PolicyAppliesToCollectionRequestBuilder {
	bb := &PolicyAppliesToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appliesTo"
	return bb
}

// PolicyAppliesToCollectionRequestBuilder is request builder for DirectoryObject collection
type PolicyAppliesToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *PolicyAppliesToCollectionRequestBuilder) Request() *PolicyAppliesToCollectionRequest {
	return &PolicyAppliesToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *PolicyAppliesToCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PolicyAppliesToCollectionRequest is request for DirectoryObject collection
type PolicyAppliesToCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *PolicyAppliesToCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *PolicyAppliesToCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *PolicyAppliesToCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *PolicyAppliesToCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for PolicySetAssignment collection
func (b *PolicySetRequestBuilder) Assignments() *PolicySetAssignmentsCollectionRequestBuilder {
	bb := &PolicySetAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// PolicySetAssignmentsCollectionRequestBuilder is request builder for PolicySetAssignment collection
type PolicySetAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PolicySetAssignment collection
func (b *PolicySetAssignmentsCollectionRequestBuilder) Request() *PolicySetAssignmentsCollectionRequest {
	return &PolicySetAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PolicySetAssignment item
func (b *PolicySetAssignmentsCollectionRequestBuilder) ID(id string) *PolicySetAssignmentRequestBuilder {
	bb := &PolicySetAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PolicySetAssignmentsCollectionRequest is request for PolicySetAssignment collection
type PolicySetAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PolicySetAssignment collection
func (r *PolicySetAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PolicySetAssignment, error) {
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
	var values []PolicySetAssignment
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
			value  []PolicySetAssignment
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

// GetN performs GET request for PolicySetAssignment collection, max N pages
func (r *PolicySetAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]PolicySetAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PolicySetAssignment collection
func (r *PolicySetAssignmentsCollectionRequest) Get(ctx context.Context) ([]PolicySetAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PolicySetAssignment collection
func (r *PolicySetAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *PolicySetAssignment) (resObj *PolicySetAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Items returns request builder for PolicySetItem collection
func (b *PolicySetRequestBuilder) Items() *PolicySetItemsCollectionRequestBuilder {
	bb := &PolicySetItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// PolicySetItemsCollectionRequestBuilder is request builder for PolicySetItem collection
type PolicySetItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PolicySetItem collection
func (b *PolicySetItemsCollectionRequestBuilder) Request() *PolicySetItemsCollectionRequest {
	return &PolicySetItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PolicySetItem item
func (b *PolicySetItemsCollectionRequestBuilder) ID(id string) *PolicySetItemRequestBuilder {
	bb := &PolicySetItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PolicySetItemsCollectionRequest is request for PolicySetItem collection
type PolicySetItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PolicySetItem collection
func (r *PolicySetItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PolicySetItem, error) {
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
	var values []PolicySetItem
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
			value  []PolicySetItem
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

// GetN performs GET request for PolicySetItem collection, max N pages
func (r *PolicySetItemsCollectionRequest) GetN(ctx context.Context, n int) ([]PolicySetItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PolicySetItem collection
func (r *PolicySetItemsCollectionRequest) Get(ctx context.Context) ([]PolicySetItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PolicySetItem collection
func (r *PolicySetItemsCollectionRequest) Add(ctx context.Context, reqObj *PolicySetItem) (resObj *PolicySetItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
