// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RejectReason undocumented
type RejectReason int

const (
	// RejectReasonVNone undocumented
	RejectReasonVNone RejectReason = 0
	// RejectReasonVBusy undocumented
	RejectReasonVBusy RejectReason = 1
	// RejectReasonVForbidden undocumented
	RejectReasonVForbidden RejectReason = 2
)

// RejectReasonPNone returns a pointer to RejectReasonVNone
func RejectReasonPNone() *RejectReason {
	v := RejectReasonVNone
	return &v
}

// RejectReasonPBusy returns a pointer to RejectReasonVBusy
func RejectReasonPBusy() *RejectReason {
	v := RejectReasonVBusy
	return &v
}

// RejectReasonPForbidden returns a pointer to RejectReasonVForbidden
func RejectReasonPForbidden() *RejectReason {
	v := RejectReasonVForbidden
	return &v
}