// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AppleDeviceFeaturesConfigurationBase Apple device features configuration profile.
type AppleDeviceFeaturesConfigurationBase struct {
	// DeviceConfiguration is the base model of AppleDeviceFeaturesConfigurationBase
	DeviceConfiguration
	// AirPrintDestinations An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
	AirPrintDestinations []AirPrintDestination `json:"airPrintDestinations,omitempty"`
}

// AppleEnrollmentProfileAssignment An assignment of an Apple profile.
type AppleEnrollmentProfileAssignment struct {
	// Entity is the base model of AppleEnrollmentProfileAssignment
	Entity
	// Target The assignment target for the Apple user initiated deployment profile.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// AppleOwnerTypeEnrollmentType undocumented
type AppleOwnerTypeEnrollmentType struct {
	// Object is the base model of AppleOwnerTypeEnrollmentType
	Object
	// OwnerType The owner type.
	OwnerType *ManagedDeviceOwnerType `json:"ownerType,omitempty"`
	// EnrollmentType The enrollment type.
	EnrollmentType *AppleUserInitiatedEnrollmentType `json:"enrollmentType,omitempty"`
}

// ApplePushNotificationCertificate Apple push notification certificate.
type ApplePushNotificationCertificate struct {
	// Entity is the base model of ApplePushNotificationCertificate
	Entity
	// AppleIdentifier Apple Id of the account used to create the MDM push certificate.
	AppleIdentifier *string `json:"appleIdentifier,omitempty"`
	// TopicIdentifier Topic Id.
	TopicIdentifier *string `json:"topicIdentifier,omitempty"`
	// LastModifiedDateTime Last modified date and time for Apple push notification certificate.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ExpirationDateTime The expiration date and time for Apple push notification certificate.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// CertificateUploadStatus The certificate upload status.
	CertificateUploadStatus *string `json:"certificateUploadStatus,omitempty"`
	// CertificateUploadFailureReason The reason the certificate upload failed.
	CertificateUploadFailureReason *string `json:"certificateUploadFailureReason,omitempty"`
	// Certificate undocumented
	Certificate *string `json:"certificate,omitempty"`
}

// AppleUserInitiatedEnrollmentProfile The enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
	// Entity is the base model of AppleUserInitiatedEnrollmentProfile
	Entity
	// DefaultEnrollmentType The default profile enrollment type.
	DefaultEnrollmentType *AppleUserInitiatedEnrollmentType `json:"defaultEnrollmentType,omitempty"`
	// AvailableEnrollmentTypeOptions List of available enrollment type options
	AvailableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentType `json:"availableEnrollmentTypeOptions,omitempty"`
	// DisplayName Name of the profile
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the profile
	Description *string `json:"description,omitempty"`
	// Priority Priority, 0 is highest
	Priority *int `json:"priority,omitempty"`
	// Platform The platform of the Device.
	Platform *DevicePlatformType `json:"platform,omitempty"`
	// CreatedDateTime Profile creation time
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime Profile last modified time
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Assignments undocumented
	Assignments []AppleEnrollmentProfileAssignment `json:"assignments,omitempty"`
}

// AppleVPNConfiguration Apple VPN configuration profile.
type AppleVPNConfiguration struct {
	// DeviceConfiguration is the base model of AppleVPNConfiguration
	DeviceConfiguration
	// ConnectionName Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`
	// ConnectionType Connection type.
	ConnectionType *AppleVPNConnectionType `json:"connectionType,omitempty"`
	// LoginGroupOrDomain Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
	LoginGroupOrDomain *string `json:"loginGroupOrDomain,omitempty"`
	// Role Role when connection type is set to Pulse Secure.
	Role *string `json:"role,omitempty"`
	// Realm Realm when connection type is set to Pulse Secure.
	Realm *string `json:"realm,omitempty"`
	// Server VPN Server on the network. Make sure end users can access this network location.
	Server *VPNServer `json:"server,omitempty"`
	// Identifier Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
	Identifier *string `json:"identifier,omitempty"`
	// CustomData Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
	CustomData []KeyValue `json:"customData,omitempty"`
	// CustomKeyValueData Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
	CustomKeyValueData []KeyValuePair `json:"customKeyValueData,omitempty"`
	// EnableSplitTunneling Send all network traffic through VPN.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`
	// AuthenticationMethod Authentication method for this VPN connection.
	AuthenticationMethod *VPNAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// EnablePerApp Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
	EnablePerApp *bool `json:"enablePerApp,omitempty"`
	// SafariDomains Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
	SafariDomains []string `json:"safariDomains,omitempty"`
	// OnDemandRules On-Demand Rules. This collection can contain a maximum of 500 elements.
	OnDemandRules []VPNOnDemandRule `json:"onDemandRules,omitempty"`
	// ProxyServer Proxy Server.
	ProxyServer *VPNProxyServer `json:"proxyServer,omitempty"`
	// OptInToDeviceIDSharing Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
	OptInToDeviceIDSharing *bool `json:"optInToDeviceIdSharing,omitempty"`
}

// AppleVPPTokenTroubleshootingEvent Event representing an Apple Vpp Token Troubleshooting Event.
type AppleVPPTokenTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of AppleVPPTokenTroubleshootingEvent
	DeviceManagementTroubleshootingEvent
	// TokenID Apple Volume Purchase Program Token Identifier.
	TokenID *string `json:"tokenId,omitempty"`
}
