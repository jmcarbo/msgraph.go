// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TelecomExpenseManagementPartnerRequestBuilder is request builder for TelecomExpenseManagementPartner
type TelecomExpenseManagementPartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns TelecomExpenseManagementPartnerRequest
func (b *TelecomExpenseManagementPartnerRequestBuilder) Request() *TelecomExpenseManagementPartnerRequest {
	return &TelecomExpenseManagementPartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TelecomExpenseManagementPartnerRequest is request for TelecomExpenseManagementPartner
type TelecomExpenseManagementPartnerRequest struct{ BaseRequest }

// Get performs GET request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Get(ctx context.Context) (resObj *TelecomExpenseManagementPartner, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Update(ctx context.Context, reqObj *TelecomExpenseManagementPartner) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
