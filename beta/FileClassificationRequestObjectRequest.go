// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FileClassificationRequestObjectRequestBuilder is request builder for FileClassificationRequestObject
type FileClassificationRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns FileClassificationRequestObjectRequest
func (b *FileClassificationRequestObjectRequestBuilder) Request() *FileClassificationRequestObjectRequest {
	return &FileClassificationRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FileClassificationRequestObjectRequest is request for FileClassificationRequestObject
type FileClassificationRequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *FileClassificationRequestObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Get() (*FileClassificationRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Update(reqObj *FileClassificationRequestObject) (*FileClassificationRequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}