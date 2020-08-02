// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TiIndicator undocumented
type TiIndicator struct {
	// Entity is the base model of TiIndicator
	Entity
	// Action undocumented
	Action *TiAction `json:"action,omitempty"`
	// ActivityGroupNames undocumented
	ActivityGroupNames []string `json:"activityGroupNames,omitempty"`
	// AdditionalInformation undocumented
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DiamondModel undocumented
	DiamondModel *DiamondModel `json:"diamondModel,omitempty"`
	// DomainName undocumented
	DomainName *string `json:"domainName,omitempty"`
	// EmailEncoding undocumented
	EmailEncoding *string `json:"emailEncoding,omitempty"`
	// EmailLanguage undocumented
	EmailLanguage *string `json:"emailLanguage,omitempty"`
	// EmailRecipient undocumented
	EmailRecipient *string `json:"emailRecipient,omitempty"`
	// EmailSenderAddress undocumented
	EmailSenderAddress *string `json:"emailSenderAddress,omitempty"`
	// EmailSenderName undocumented
	EmailSenderName *string `json:"emailSenderName,omitempty"`
	// EmailSourceDomain undocumented
	EmailSourceDomain *string `json:"emailSourceDomain,omitempty"`
	// EmailSourceIPAddress undocumented
	EmailSourceIPAddress *string `json:"emailSourceIpAddress,omitempty"`
	// EmailSubject undocumented
	EmailSubject *string `json:"emailSubject,omitempty"`
	// EmailXMailer undocumented
	EmailXMailer *string `json:"emailXMailer,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// FileCompileDateTime undocumented
	FileCompileDateTime *time.Time `json:"fileCompileDateTime,omitempty"`
	// FileCreatedDateTime undocumented
	FileCreatedDateTime *time.Time `json:"fileCreatedDateTime,omitempty"`
	// FileHashType undocumented
	FileHashType *FileHashType `json:"fileHashType,omitempty"`
	// FileHashValue undocumented
	FileHashValue *string `json:"fileHashValue,omitempty"`
	// FileMutexName undocumented
	FileMutexName *string `json:"fileMutexName,omitempty"`
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// FilePacker undocumented
	FilePacker *string `json:"filePacker,omitempty"`
	// FilePath undocumented
	FilePath *string `json:"filePath,omitempty"`
	// FileSize undocumented
	FileSize *int `json:"fileSize,omitempty"`
	// FileType undocumented
	FileType *string `json:"fileType,omitempty"`
	// IngestedDateTime undocumented
	IngestedDateTime *time.Time `json:"ingestedDateTime,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// KillChain undocumented
	KillChain []string `json:"killChain,omitempty"`
	// KnownFalsePositives undocumented
	KnownFalsePositives *string `json:"knownFalsePositives,omitempty"`
	// LastReportedDateTime undocumented
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// MalwareFamilyNames undocumented
	MalwareFamilyNames []string `json:"malwareFamilyNames,omitempty"`
	// NetworkCIDRBlock undocumented
	NetworkCIDRBlock *string `json:"networkCidrBlock,omitempty"`
	// NetworkDestinationAsn undocumented
	NetworkDestinationAsn *int `json:"networkDestinationAsn,omitempty"`
	// NetworkDestinationCIDRBlock undocumented
	NetworkDestinationCIDRBlock *string `json:"networkDestinationCidrBlock,omitempty"`
	// NetworkDestinationIPv4 undocumented
	NetworkDestinationIPv4 *string `json:"networkDestinationIPv4,omitempty"`
	// NetworkDestinationIPv6 undocumented
	NetworkDestinationIPv6 *string `json:"networkDestinationIPv6,omitempty"`
	// NetworkDestinationPort undocumented
	NetworkDestinationPort *int `json:"networkDestinationPort,omitempty"`
	// NetworkIPv4 undocumented
	NetworkIPv4 *string `json:"networkIPv4,omitempty"`
	// NetworkIPv6 undocumented
	NetworkIPv6 *string `json:"networkIPv6,omitempty"`
	// NetworkPort undocumented
	NetworkPort *int `json:"networkPort,omitempty"`
	// NetworkProtocol undocumented
	NetworkProtocol *int `json:"networkProtocol,omitempty"`
	// NetworkSourceAsn undocumented
	NetworkSourceAsn *int `json:"networkSourceAsn,omitempty"`
	// NetworkSourceCIDRBlock undocumented
	NetworkSourceCIDRBlock *string `json:"networkSourceCidrBlock,omitempty"`
	// NetworkSourceIPv4 undocumented
	NetworkSourceIPv4 *string `json:"networkSourceIPv4,omitempty"`
	// NetworkSourceIPv6 undocumented
	NetworkSourceIPv6 *string `json:"networkSourceIPv6,omitempty"`
	// NetworkSourcePort undocumented
	NetworkSourcePort *int `json:"networkSourcePort,omitempty"`
	// PassiveOnly undocumented
	PassiveOnly *bool `json:"passiveOnly,omitempty"`
	// Severity undocumented
	Severity *int `json:"severity,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// TargetProduct undocumented
	TargetProduct *string `json:"targetProduct,omitempty"`
	// ThreatType undocumented
	ThreatType *string `json:"threatType,omitempty"`
	// TlpLevel undocumented
	TlpLevel *TlpLevel `json:"tlpLevel,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// UserAgent undocumented
	UserAgent *string `json:"userAgent,omitempty"`
}
