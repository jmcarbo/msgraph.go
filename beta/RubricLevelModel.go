// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RubricLevel undocumented
type RubricLevel struct {
	// LevelID undocumented
	LevelID *string `json:"levelId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
}