// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DateTimeColumn undocumented
type DateTimeColumn struct {
	// Object is the base model of DateTimeColumn
	Object
	// DisplayAs undocumented
	DisplayAs *string `json:"displayAs,omitempty"`
	// Format undocumented
	Format *string `json:"format,omitempty"`
}

// DateTimeTimeZone undocumented
type DateTimeTimeZone struct {
	// Object is the base model of DateTimeTimeZone
	Object
	// DateTime undocumented
	DateTime *string `json:"dateTime,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}
