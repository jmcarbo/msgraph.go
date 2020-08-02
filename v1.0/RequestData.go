// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DataPolicyOperationRequestBuilder is request builder for DataPolicyOperation
type DataPolicyOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataPolicyOperationRequest
func (b *DataPolicyOperationRequestBuilder) Request() *DataPolicyOperationRequest {
	return &DataPolicyOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataPolicyOperationRequest is request for DataPolicyOperation
type DataPolicyOperationRequest struct{ BaseRequest }

// Get performs GET request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Get(ctx context.Context) (resObj *DataPolicyOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Update(ctx context.Context, reqObj *DataPolicyOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
