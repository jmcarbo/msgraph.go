// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ClassificationInnerError undocumented
type ClassificationInnerError struct {
	// ErrorDateTime undocumented
	ErrorDateTime *time.Time `json:"errorDateTime,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// ClientRequestID undocumented
	ClientRequestID *string `json:"clientRequestId,omitempty"`
	// ActivityID undocumented
	ActivityID *string `json:"activityId,omitempty"`
}