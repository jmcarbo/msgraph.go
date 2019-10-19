// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidWorkProfilePkcsCertificateProfile Android Work Profile PKCS certificate profile
type AndroidWorkProfilePkcsCertificateProfile struct {
	AndroidWorkProfileCertificateProfileBase
	// CertificationAuthority PKCS Certification Authority
	CertificationAuthority *string `json:"certificationAuthority,omitempty"`
	// CertificationAuthorityName PKCS Certification Authority Name
	CertificationAuthorityName *string `json:"certificationAuthorityName,omitempty"`
	// CertificateTemplateName PKCS Certificate Template Name
	CertificateTemplateName *string `json:"certificateTemplateName,omitempty"`
	// SubjectAlternativeNameFormatString Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString *string `json:"subjectAlternativeNameFormatString,omitempty"`
}