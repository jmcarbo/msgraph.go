// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsErf_PreciseRequestBuilder struct{ BaseRequestBuilder }

// Erf_Precise action undocumented
func (b *WorkbookFunctionsRequestBuilder) Erf_Precise(reqObj *WorkbookFunctionsErf_PreciseRequestParameter) *WorkbookFunctionsErf_PreciseRequestBuilder {
	bb := &WorkbookFunctionsErf_PreciseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/erf_Precise"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsErf_PreciseRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsErf_PreciseRequestBuilder) Request() *WorkbookFunctionsErf_PreciseRequest {
	return &WorkbookFunctionsErf_PreciseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsErf_PreciseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
