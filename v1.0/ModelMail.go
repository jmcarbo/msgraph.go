// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MailFolder undocumented
type MailFolder struct {
	// Entity is the base model of MailFolder
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// ChildFolderCount undocumented
	ChildFolderCount *int `json:"childFolderCount,omitempty"`
	// UnreadItemCount undocumented
	UnreadItemCount *int `json:"unreadItemCount,omitempty"`
	// TotalItemCount undocumented
	TotalItemCount *int `json:"totalItemCount,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Messages undocumented
	Messages []Message `json:"messages,omitempty"`
	// MessageRules undocumented
	MessageRules []MessageRule `json:"messageRules,omitempty"`
	// ChildFolders undocumented
	ChildFolders []MailFolder `json:"childFolders,omitempty"`
}

// MailSearchFolder undocumented
type MailSearchFolder struct {
	// MailFolder is the base model of MailSearchFolder
	MailFolder
	// IsSupported undocumented
	IsSupported *bool `json:"isSupported,omitempty"`
	// IncludeNestedFolders undocumented
	IncludeNestedFolders *bool `json:"includeNestedFolders,omitempty"`
	// SourceFolderIDs undocumented
	SourceFolderIDs []string `json:"sourceFolderIds,omitempty"`
	// FilterQuery undocumented
	FilterQuery *string `json:"filterQuery,omitempty"`
}

// MailTips undocumented
type MailTips struct {
	// Object is the base model of MailTips
	Object
	// EmailAddress undocumented
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
	// AutomaticReplies undocumented
	AutomaticReplies *AutomaticRepliesMailTips `json:"automaticReplies,omitempty"`
	// MailboxFull undocumented
	MailboxFull *bool `json:"mailboxFull,omitempty"`
	// CustomMailTip undocumented
	CustomMailTip *string `json:"customMailTip,omitempty"`
	// ExternalMemberCount undocumented
	ExternalMemberCount *int `json:"externalMemberCount,omitempty"`
	// TotalMemberCount undocumented
	TotalMemberCount *int `json:"totalMemberCount,omitempty"`
	// DeliveryRestricted undocumented
	DeliveryRestricted *bool `json:"deliveryRestricted,omitempty"`
	// IsModerated undocumented
	IsModerated *bool `json:"isModerated,omitempty"`
	// RecipientScope undocumented
	RecipientScope *RecipientScopeType `json:"recipientScope,omitempty"`
	// RecipientSuggestions undocumented
	RecipientSuggestions []Recipient `json:"recipientSuggestions,omitempty"`
	// MaxMessageSize undocumented
	MaxMessageSize *int `json:"maxMessageSize,omitempty"`
	// Error undocumented
	Error *MailTipsError `json:"error,omitempty"`
}

// MailTipsError undocumented
type MailTipsError struct {
	// Object is the base model of MailTipsError
	Object
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
}
