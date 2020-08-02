// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AuditActivityInitiator undocumented
type AuditActivityInitiator struct {
	// Object is the base model of AuditActivityInitiator
	Object
	// User undocumented
	User *UserIdentity `json:"user,omitempty"`
	// App undocumented
	App *AppIdentity `json:"app,omitempty"`
}

// AuditLogRoot undocumented
type AuditLogRoot struct {
	// Entity is the base model of AuditLogRoot
	Entity
	// SignIns undocumented
	SignIns []SignIn `json:"signIns,omitempty"`
	// DirectoryAudits undocumented
	DirectoryAudits []DirectoryAudit `json:"directoryAudits,omitempty"`
	// RestrictedSignIns undocumented
	RestrictedSignIns []RestrictedSignIn `json:"restrictedSignIns,omitempty"`
}
