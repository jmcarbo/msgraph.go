// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LookupColumn undocumented
type LookupColumn struct {
	// Object is the base model of LookupColumn
	Object
	// AllowMultipleValues undocumented
	AllowMultipleValues *bool `json:"allowMultipleValues,omitempty"`
	// AllowUnlimitedLength undocumented
	AllowUnlimitedLength *bool `json:"allowUnlimitedLength,omitempty"`
	// ColumnName undocumented
	ColumnName *string `json:"columnName,omitempty"`
	// ListID undocumented
	ListID *string `json:"listId,omitempty"`
	// PrimaryLookupColumnID undocumented
	PrimaryLookupColumnID *string `json:"primaryLookupColumnId,omitempty"`
}
