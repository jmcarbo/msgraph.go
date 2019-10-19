// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidWorkProfileCertificateProfileBaseRequestBuilder is request builder for AndroidWorkProfileCertificateProfileBase
type AndroidWorkProfileCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidWorkProfileCertificateProfileBaseRequest
func (b *AndroidWorkProfileCertificateProfileBaseRequestBuilder) Request() *AndroidWorkProfileCertificateProfileBaseRequest {
	return &AndroidWorkProfileCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidWorkProfileCertificateProfileBaseRequest is request for AndroidWorkProfileCertificateProfileBase
type AndroidWorkProfileCertificateProfileBaseRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidWorkProfileCertificateProfileBase, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Get() (*AndroidWorkProfileCertificateProfileBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Update(reqObj *AndroidWorkProfileCertificateProfileBase) (*AndroidWorkProfileCertificateProfileBase, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidWorkProfileCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidWorkProfileTrustedRootCertificateRequestBuilder {
	bb := &AndroidWorkProfileTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}