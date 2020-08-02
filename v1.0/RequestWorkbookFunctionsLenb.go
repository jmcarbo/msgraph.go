// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLenbRequestBuilder struct{ BaseRequestBuilder }

// Lenb action undocumented
func (b *WorkbookFunctionsRequestBuilder) Lenb(reqObj *WorkbookFunctionsLenbRequestParameter) *WorkbookFunctionsLenbRequestBuilder {
	bb := &WorkbookFunctionsLenbRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/lenb"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLenbRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLenbRequestBuilder) Request() *WorkbookFunctionsLenbRequest {
	return &WorkbookFunctionsLenbRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLenbRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
