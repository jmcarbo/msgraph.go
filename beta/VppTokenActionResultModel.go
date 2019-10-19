// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// VppTokenActionResult undocumented
type VppTokenActionResult struct {
	// ActionName Action name
	ActionName *string `json:"actionName,omitempty"`
	// ActionState State of the action
	ActionState *ActionState `json:"actionState,omitempty"`
	// StartDateTime Time the action was initiated
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// LastUpdatedDateTime Time the action state was last updated
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
}