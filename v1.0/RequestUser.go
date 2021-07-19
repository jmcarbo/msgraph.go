// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmcarbo/msgraph.go/jsonx"
)

// UserRequestBuilder is request builder for User
type UserRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserRequest
func (b *UserRequestBuilder) Request() *UserRequest {
	return &UserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserRequest is request for User
type UserRequest struct{ BaseRequest }

// Get performs GET request for User
func (r *UserRequest) Get(ctx context.Context) (resObj *User, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for User
func (r *UserRequest) Update(ctx context.Context, reqObj *User) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for User
func (r *UserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserActivityRequestBuilder is request builder for UserActivity
type UserActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserActivityRequest
func (b *UserActivityRequestBuilder) Request() *UserActivityRequest {
	return &UserActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserActivityRequest is request for UserActivity
type UserActivityRequest struct{ BaseRequest }

// Get performs GET request for UserActivity
func (r *UserActivityRequest) Get(ctx context.Context) (resObj *UserActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserActivity
func (r *UserActivityRequest) Update(ctx context.Context, reqObj *UserActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserActivity
func (r *UserActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserInstallStateSummaryRequestBuilder is request builder for UserInstallStateSummary
type UserInstallStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserInstallStateSummaryRequest
func (b *UserInstallStateSummaryRequestBuilder) Request() *UserInstallStateSummaryRequest {
	return &UserInstallStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserInstallStateSummaryRequest is request for UserInstallStateSummary
type UserInstallStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Get(ctx context.Context) (resObj *UserInstallStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Update(ctx context.Context, reqObj *UserInstallStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSettingsRequestBuilder is request builder for UserSettings
type UserSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSettingsRequest
func (b *UserSettingsRequestBuilder) Request() *UserSettingsRequest {
	return &UserSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSettingsRequest is request for UserSettings
type UserSettingsRequest struct{ BaseRequest }

// Get performs GET request for UserSettings
func (r *UserSettingsRequest) Get(ctx context.Context) (resObj *UserSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSettings
func (r *UserSettingsRequest) Update(ctx context.Context, reqObj *UserSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSettings
func (r *UserSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type UserAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumented
func (b *UserRequestBuilder) AssignLicense(reqObj *UserAssignLicenseRequestParameter) *UserAssignLicenseRequestBuilder {
	bb := &UserAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserAssignLicenseRequest struct{ BaseRequest }

//
func (b *UserAssignLicenseRequestBuilder) Request() *UserAssignLicenseRequest {
	return &UserAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserAssignLicenseRequest) Post(ctx context.Context) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserChangePasswordRequestBuilder struct{ BaseRequestBuilder }

// ChangePassword action undocumented
func (b *UserRequestBuilder) ChangePassword(reqObj *UserChangePasswordRequestParameter) *UserChangePasswordRequestBuilder {
	bb := &UserChangePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/changePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserChangePasswordRequest struct{ BaseRequest }

//
func (b *UserChangePasswordRequestBuilder) Request() *UserChangePasswordRequest {
	return &UserChangePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserChangePasswordRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserRevokeSignInSessionsRequestBuilder struct{ BaseRequestBuilder }

// RevokeSignInSessions action undocumented
func (b *UserRequestBuilder) RevokeSignInSessions(reqObj *UserRevokeSignInSessionsRequestParameter) *UserRevokeSignInSessionsRequestBuilder {
	bb := &UserRevokeSignInSessionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeSignInSessions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRevokeSignInSessionsRequest struct{ BaseRequest }

//
func (b *UserRevokeSignInSessionsRequestBuilder) Request() *UserRevokeSignInSessionsRequest {
	return &UserRevokeSignInSessionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRevokeSignInSessionsRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserFindMeetingTimesRequestBuilder struct{ BaseRequestBuilder }

// FindMeetingTimes action undocumented
func (b *UserRequestBuilder) FindMeetingTimes(reqObj *UserFindMeetingTimesRequestParameter) *UserFindMeetingTimesRequestBuilder {
	bb := &UserFindMeetingTimesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/findMeetingTimes"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserFindMeetingTimesRequest struct{ BaseRequest }

//
func (b *UserFindMeetingTimesRequestBuilder) Request() *UserFindMeetingTimesRequest {
	return &UserFindMeetingTimesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserFindMeetingTimesRequest) Post(ctx context.Context) (resObj *MeetingTimeSuggestionsResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserSendMailRequestBuilder struct{ BaseRequestBuilder }

// SendMail action undocumented
func (b *UserRequestBuilder) SendMail(reqObj *UserSendMailRequestParameter) *UserSendMailRequestBuilder {
	bb := &UserSendMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserSendMailRequest struct{ BaseRequest }

//
func (b *UserSendMailRequestBuilder) Request() *UserSendMailRequest {
	return &UserSendMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserSendMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserGetMailTipsRequestBuilder struct{ BaseRequestBuilder }

// GetMailTips action undocumented
func (b *UserRequestBuilder) GetMailTips(reqObj *UserGetMailTipsRequestParameter) *UserGetMailTipsRequestBuilder {
	bb := &UserGetMailTipsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getMailTips"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserGetMailTipsRequest struct{ BaseRequest }

//
func (b *UserGetMailTipsRequestBuilder) Request() *UserGetMailTipsRequest {
	return &UserGetMailTipsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserGetMailTipsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MailTips, error) {
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
	var values []MailTips
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
			value  []MailTips
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

//
func (r *UserGetMailTipsRequest) PostN(ctx context.Context, n int) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *UserGetMailTipsRequest) Post(ctx context.Context) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type UserTranslateExchangeIDsRequestBuilder struct{ BaseRequestBuilder }

// TranslateExchangeIDs action undocumented
func (b *UserRequestBuilder) TranslateExchangeIDs(reqObj *UserTranslateExchangeIDsRequestParameter) *UserTranslateExchangeIDsRequestBuilder {
	bb := &UserTranslateExchangeIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/translateExchangeIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserTranslateExchangeIDsRequest struct{ BaseRequest }

//
func (b *UserTranslateExchangeIDsRequestBuilder) Request() *UserTranslateExchangeIDsRequest {
	return &UserTranslateExchangeIDsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserTranslateExchangeIDsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConvertIDResult, error) {
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
	var values []ConvertIDResult
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
			value  []ConvertIDResult
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

//
func (r *UserTranslateExchangeIDsRequest) PostN(ctx context.Context, n int) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *UserTranslateExchangeIDsRequest) Post(ctx context.Context) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type UserRemoveAllDevicesFromManagementRequestBuilder struct{ BaseRequestBuilder }

// RemoveAllDevicesFromManagement action undocumented
func (b *UserRequestBuilder) RemoveAllDevicesFromManagement(reqObj *UserRemoveAllDevicesFromManagementRequestParameter) *UserRemoveAllDevicesFromManagementRequestBuilder {
	bb := &UserRemoveAllDevicesFromManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeAllDevicesFromManagement"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRemoveAllDevicesFromManagementRequest struct{ BaseRequest }

//
func (b *UserRemoveAllDevicesFromManagementRequestBuilder) Request() *UserRemoveAllDevicesFromManagementRequest {
	return &UserRemoveAllDevicesFromManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRemoveAllDevicesFromManagementRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationsByDeviceTag action undocumented
func (b *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag(reqObj *UserWipeManagedAppRegistrationsByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipeManagedAppRegistrationsByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequest struct{ BaseRequest }

//
func (b *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationsByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationsByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserWipeManagedAppRegistrationsByDeviceTagRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserExportPersonalDataRequestBuilder struct{ BaseRequestBuilder }

// ExportPersonalData action undocumented
func (b *UserRequestBuilder) ExportPersonalData(reqObj *UserExportPersonalDataRequestParameter) *UserExportPersonalDataRequestBuilder {
	bb := &UserExportPersonalDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/exportPersonalData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserExportPersonalDataRequest struct{ BaseRequest }

//
func (b *UserExportPersonalDataRequestBuilder) Request() *UserExportPersonalDataRequest {
	return &UserExportPersonalDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserExportPersonalDataRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
