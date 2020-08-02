// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// PasswordCredential undocumented
type PasswordCredential struct {
	// Object is the base model of PasswordCredential
	Object
	// CustomKeyIdentifier undocumented
	CustomKeyIdentifier *Binary `json:"customKeyIdentifier,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// SecretText undocumented
	SecretText *string `json:"secretText,omitempty"`
	// Hint undocumented
	Hint *string `json:"hint,omitempty"`
}

// PasswordProfile undocumented
type PasswordProfile struct {
	// Object is the base model of PasswordProfile
	Object
	// Password undocumented
	Password *string `json:"password,omitempty"`
	// ForceChangePasswordNextSignIn undocumented
	ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
	// ForceChangePasswordNextSignInWithMFA undocumented
	ForceChangePasswordNextSignInWithMFA *bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
}
