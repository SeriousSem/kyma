package model

// ServiceDefinition is an internal representation of a service.
type ServiceDefinition struct {
	// ID of service
	ID string
	// Name of a service
	Name string
	// Provider of a service
	Provider string
	// Description of a service
	Description string
	// Api of a service
	Api *API
	// Events of a service
	Events *Events
	// Documentation of service
	Documentation []byte
}

// API is an internal representation of a service's API.
type API struct {
	// TargetUrl points to API.
	TargetUrl string
	// Credentials is a credentials of API.
	Credentials *Credentials
	// Spec contains specification of an API.
	Spec []byte
}

// Credentials contains OAuth or BasicAuth configuration.
type Credentials struct {
	// OAuth is OAuth configuration.
	OAuth *OAuth
	// BasicAuth is BasicAuth configuration.
	BasicAuth *BasicAuth
	// CertificateGen is CertificateGen configuration.
	CertificateGen *CertificateGen
	// CSRFTokenEndpointURL (optional) to fetch CSRF token
	CSRFTokenEndpointURL string
}

// BasicAuth contains details of BasicAuth Auth configuration
type BasicAuth struct {
	// Username to use for authentication
	Username string
	// Password to use for authentication
	Password string
}

// OAuth contains details of OAuth configuration
type OAuth struct {
	// URL to OAuth token provider.
	URL string
	// ClientID to use for authorization.
	ClientID string
	// ClientSecret to use for authorization.
	ClientSecret string
}

// CertificateGen details of CertificateGen configuration
type CertificateGen struct {
	// CommonName of the certificate
	CommonName string
	// Certificate generated by Application Registry
	Certificate []byte
	// PrivateKey generated by Application Registry
	PrivateKey []byte
}

// Events contains specification for events.
type Events struct {
	// Spec contains data of events specification.
	Spec []byte
}
