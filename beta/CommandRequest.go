// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CommandRequestBuilder is request builder for Command
type CommandRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommandRequest
func (b *CommandRequestBuilder) Request() *CommandRequest {
	return &CommandRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommandRequest is request for Command
type CommandRequest struct{ BaseRequest }

// Do performs HTTP request for Command
func (r *CommandRequest) Do(method, path string, reqObj interface{}) (resObj *Command, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Command
func (r *CommandRequest) Get() (*Command, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Command
func (r *CommandRequest) Update(reqObj *Command) (*Command, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Command
func (r *CommandRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Responsepayload is navigation property
func (b *CommandRequestBuilder) Responsepayload() *PayloadResponseRequestBuilder {
	bb := &PayloadResponseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/responsepayload"
	return bb
}