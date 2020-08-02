// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsNpvRequestBuilder struct{ BaseRequestBuilder }

// Npv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Npv(reqObj *WorkbookFunctionsNpvRequestParameter) *WorkbookFunctionsNpvRequestBuilder {
	bb := &WorkbookFunctionsNpvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/npv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNpvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNpvRequestBuilder) Request() *WorkbookFunctionsNpvRequest {
	return &WorkbookFunctionsNpvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNpvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
