// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PhysicalAddress undocumented
type PhysicalAddress struct {
	// Type undocumented
	Type *PhysicalAddressType `json:"type,omitempty"`
	// PostOfficeBox undocumented
	PostOfficeBox *string `json:"postOfficeBox,omitempty"`
	// Street undocumented
	Street *string `json:"street,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
}