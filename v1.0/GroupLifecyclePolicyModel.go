// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupLifecyclePolicy undocumented
type GroupLifecyclePolicy struct {
	Entity
	// GroupLifetimeInDays undocumented
	GroupLifetimeInDays *int `json:"groupLifetimeInDays,omitempty"`
	// ManagedGroupTypes undocumented
	ManagedGroupTypes *string `json:"managedGroupTypes,omitempty"`
	// AlternateNotificationEmails undocumented
	AlternateNotificationEmails *string `json:"alternateNotificationEmails,omitempty"`
}