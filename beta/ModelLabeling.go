// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LabelingOptions undocumented
type LabelingOptions struct {
	// Object is the base model of LabelingOptions
	Object
	// LabelID undocumented
	LabelID *UUID `json:"labelId,omitempty"`
	// AssignmentMethod undocumented
	AssignmentMethod *AssignmentMethod `json:"assignmentMethod,omitempty"`
	// DowngradeJustification undocumented
	DowngradeJustification *DowngradeJustification `json:"downgradeJustification,omitempty"`
	// ExtendedProperties undocumented
	ExtendedProperties []KeyValuePair `json:"extendedProperties,omitempty"`
}
