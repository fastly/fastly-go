// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"time"
)

// BackendResponse struct for BackendResponse
type BackendResponse struct {
	// A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend.
	Address *string `json:"address,omitempty"`
	// Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don't have a `request_condition` will be selected based on their `weight`.
	AutoLoadbalance *bool `json:"auto_loadbalance,omitempty"`
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *int32 `json:"between_bytes_timeout,omitempty"`
	// Unused.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`.
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`
	// Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`.
	FirstByteTimeout *int32 `json:"first_byte_timeout,omitempty"`
	// The name of the healthcheck to use with this backend.
	Healthcheck NullableString `json:"healthcheck,omitempty"`
	// The hostname of the backend. May be used as an alternative to `address` to set the backend location.
	Hostname NullableString `json:"hostname,omitempty"`
	// IPv4 address of the backend. May be used as an alternative to `address` to set the backend location.
	Ipv4 NullableString `json:"ipv4,omitempty"`
	// IPv6 address of the backend. May be used as an alternative to `address` to set the backend location.
	Ipv6 NullableString `json:"ipv6,omitempty"`
	// How long in seconds to keep a persistent connection to the backend between requests.
	KeepaliveTime NullableInt32 `json:"keepalive_time,omitempty"`
	// Maximum number of concurrent connections this backend will accept.
	MaxConn *int32 `json:"max_conn,omitempty"`
	// Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	MaxTLSVersion NullableString `json:"max_tls_version,omitempty"`
	// Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	MinTLSVersion NullableString `json:"min_tls_version,omitempty"`
	// The name of the backend.
	Name *string `json:"name,omitempty"`
	// If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL.
	OverrideHost NullableString `json:"override_host,omitempty"`
	// Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request.
	Port *int32 `json:"port,omitempty"`
	// Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests.
	RequestCondition *string `json:"request_condition,omitempty"`
	// Value that when shared across backends will enable those backends to share the same health check.
	ShareKey NullableString `json:"share_key,omitempty"`
	// Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding).
	Shield NullableString `json:"shield,omitempty"`
	// CA certificate attached to origin.
	SslCaCert NullableString `json:"ssl_ca_cert,omitempty"`
	// Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all.
	SslCertHostname NullableString `json:"ssl_cert_hostname,omitempty"`
	// Be strict on checking SSL certs.
	SslCheckCert NullableBool `json:"ssl_check_cert,omitempty"`
	// List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	SslCiphers NullableString `json:"ssl_ciphers,omitempty"`
	// Client certificate attached to origin.
	SslClientCert NullableString `json:"ssl_client_cert,omitempty"`
	// Client key attached to origin.
	SslClientKey NullableString `json:"ssl_client_key,omitempty"`
	// Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation.
	// Deprecated
	SslHostname NullableString `json:"ssl_hostname,omitempty"`
	// Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all.
	SslSniHostname NullableString `json:"ssl_sni_hostname,omitempty"`
	// Whether or not to require TLS for connections to this backend.
	UseSsl *bool `json:"use_ssl,omitempty"`
	// Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true.
	Weight *int32 `json:"weight,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version *int32 `json:"version,omitempty"`
	// Indicates whether the version of the service this backend is attached to accepts edits.
	Locked *bool `json:"locked,omitempty"`
	AdditionalProperties map[string]any
}

type _BackendResponse BackendResponse

// NewBackendResponse instantiates a new BackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackendResponse() *BackendResponse {
	this := BackendResponse{}
	var sslCheckCert bool = true
	this.SslCheckCert = *NewNullableBool(&sslCheckCert)
	return &this
}

// NewBackendResponseWithDefaults instantiates a new BackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackendResponseWithDefaults() *BackendResponse {
	this := BackendResponse{}
	var sslCheckCert bool = true
	this.SslCheckCert = *NewNullableBool(&sslCheckCert)
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *BackendResponse) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *BackendResponse) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *BackendResponse) SetAddress(v string) {
	o.Address = &v
}

// GetAutoLoadbalance returns the AutoLoadbalance field value if set, zero value otherwise.
func (o *BackendResponse) GetAutoLoadbalance() bool {
	if o == nil || o.AutoLoadbalance == nil {
		var ret bool
		return ret
	}
	return *o.AutoLoadbalance
}

// GetAutoLoadbalanceOk returns a tuple with the AutoLoadbalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetAutoLoadbalanceOk() (*bool, bool) {
	if o == nil || o.AutoLoadbalance == nil {
		return nil, false
	}
	return o.AutoLoadbalance, true
}

// HasAutoLoadbalance returns a boolean if a field has been set.
func (o *BackendResponse) HasAutoLoadbalance() bool {
	if o != nil && o.AutoLoadbalance != nil {
		return true
	}

	return false
}

// SetAutoLoadbalance gets a reference to the given bool and assigns it to the AutoLoadbalance field.
func (o *BackendResponse) SetAutoLoadbalance(v bool) {
	o.AutoLoadbalance = &v
}

// GetBetweenBytesTimeout returns the BetweenBytesTimeout field value if set, zero value otherwise.
func (o *BackendResponse) GetBetweenBytesTimeout() int32 {
	if o == nil || o.BetweenBytesTimeout == nil {
		var ret int32
		return ret
	}
	return *o.BetweenBytesTimeout
}

// GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetBetweenBytesTimeoutOk() (*int32, bool) {
	if o == nil || o.BetweenBytesTimeout == nil {
		return nil, false
	}
	return o.BetweenBytesTimeout, true
}

// HasBetweenBytesTimeout returns a boolean if a field has been set.
func (o *BackendResponse) HasBetweenBytesTimeout() bool {
	if o != nil && o.BetweenBytesTimeout != nil {
		return true
	}

	return false
}

// SetBetweenBytesTimeout gets a reference to the given int32 and assigns it to the BetweenBytesTimeout field.
func (o *BackendResponse) SetBetweenBytesTimeout(v int32) {
	o.BetweenBytesTimeout = &v
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetClientCert() string {
	if o == nil || o.ClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *BackendResponse) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *BackendResponse) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *BackendResponse) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *BackendResponse) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *BackendResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *BackendResponse) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *BackendResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *BackendResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *BackendResponse) GetConnectTimeout() int32 {
	if o == nil || o.ConnectTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *BackendResponse) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *BackendResponse) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetFirstByteTimeout returns the FirstByteTimeout field value if set, zero value otherwise.
func (o *BackendResponse) GetFirstByteTimeout() int32 {
	if o == nil || o.FirstByteTimeout == nil {
		var ret int32
		return ret
	}
	return *o.FirstByteTimeout
}

// GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetFirstByteTimeoutOk() (*int32, bool) {
	if o == nil || o.FirstByteTimeout == nil {
		return nil, false
	}
	return o.FirstByteTimeout, true
}

// HasFirstByteTimeout returns a boolean if a field has been set.
func (o *BackendResponse) HasFirstByteTimeout() bool {
	if o != nil && o.FirstByteTimeout != nil {
		return true
	}

	return false
}

// SetFirstByteTimeout gets a reference to the given int32 and assigns it to the FirstByteTimeout field.
func (o *BackendResponse) SetFirstByteTimeout(v int32) {
	o.FirstByteTimeout = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetHealthcheck() string {
	if o == nil || o.Healthcheck.Get() == nil {
		var ret string
		return ret
	}
	return *o.Healthcheck.Get()
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetHealthcheckOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Healthcheck.Get(), o.Healthcheck.IsSet()
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *BackendResponse) HasHealthcheck() bool {
	if o != nil && o.Healthcheck.IsSet() {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given NullableString and assigns it to the Healthcheck field.
func (o *BackendResponse) SetHealthcheck(v string) {
	o.Healthcheck.Set(&v)
}
// SetHealthcheckNil sets the value for Healthcheck to be an explicit nil
func (o *BackendResponse) SetHealthcheckNil() {
	o.Healthcheck.Set(nil)
}

// UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
func (o *BackendResponse) UnsetHealthcheck() {
	o.Healthcheck.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *BackendResponse) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *BackendResponse) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *BackendResponse) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *BackendResponse) UnsetHostname() {
	o.Hostname.Unset()
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetIpv4() string {
	if o == nil || o.Ipv4.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ipv4.Get()
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetIpv4Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ipv4.Get(), o.Ipv4.IsSet()
}

// HasIpv4 returns a boolean if a field has been set.
func (o *BackendResponse) HasIpv4() bool {
	if o != nil && o.Ipv4.IsSet() {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NullableString and assigns it to the Ipv4 field.
func (o *BackendResponse) SetIpv4(v string) {
	o.Ipv4.Set(&v)
}
// SetIpv4Nil sets the value for Ipv4 to be an explicit nil
func (o *BackendResponse) SetIpv4Nil() {
	o.Ipv4.Set(nil)
}

// UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
func (o *BackendResponse) UnsetIpv4() {
	o.Ipv4.Unset()
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetIpv6() string {
	if o == nil || o.Ipv6.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ipv6.Get()
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetIpv6Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ipv6.Get(), o.Ipv6.IsSet()
}

// HasIpv6 returns a boolean if a field has been set.
func (o *BackendResponse) HasIpv6() bool {
	if o != nil && o.Ipv6.IsSet() {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NullableString and assigns it to the Ipv6 field.
func (o *BackendResponse) SetIpv6(v string) {
	o.Ipv6.Set(&v)
}
// SetIpv6Nil sets the value for Ipv6 to be an explicit nil
func (o *BackendResponse) SetIpv6Nil() {
	o.Ipv6.Set(nil)
}

// UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
func (o *BackendResponse) UnsetIpv6() {
	o.Ipv6.Unset()
}

// GetKeepaliveTime returns the KeepaliveTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetKeepaliveTime() int32 {
	if o == nil || o.KeepaliveTime.Get() == nil {
		var ret int32
		return ret
	}
	return *o.KeepaliveTime.Get()
}

// GetKeepaliveTimeOk returns a tuple with the KeepaliveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetKeepaliveTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KeepaliveTime.Get(), o.KeepaliveTime.IsSet()
}

// HasKeepaliveTime returns a boolean if a field has been set.
func (o *BackendResponse) HasKeepaliveTime() bool {
	if o != nil && o.KeepaliveTime.IsSet() {
		return true
	}

	return false
}

// SetKeepaliveTime gets a reference to the given NullableInt32 and assigns it to the KeepaliveTime field.
func (o *BackendResponse) SetKeepaliveTime(v int32) {
	o.KeepaliveTime.Set(&v)
}
// SetKeepaliveTimeNil sets the value for KeepaliveTime to be an explicit nil
func (o *BackendResponse) SetKeepaliveTimeNil() {
	o.KeepaliveTime.Set(nil)
}

// UnsetKeepaliveTime ensures that no value is present for KeepaliveTime, not even an explicit nil
func (o *BackendResponse) UnsetKeepaliveTime() {
	o.KeepaliveTime.Unset()
}

// GetMaxConn returns the MaxConn field value if set, zero value otherwise.
func (o *BackendResponse) GetMaxConn() int32 {
	if o == nil || o.MaxConn == nil {
		var ret int32
		return ret
	}
	return *o.MaxConn
}

// GetMaxConnOk returns a tuple with the MaxConn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetMaxConnOk() (*int32, bool) {
	if o == nil || o.MaxConn == nil {
		return nil, false
	}
	return o.MaxConn, true
}

// HasMaxConn returns a boolean if a field has been set.
func (o *BackendResponse) HasMaxConn() bool {
	if o != nil && o.MaxConn != nil {
		return true
	}

	return false
}

// SetMaxConn gets a reference to the given int32 and assigns it to the MaxConn field.
func (o *BackendResponse) SetMaxConn(v int32) {
	o.MaxConn = &v
}

// GetMaxTLSVersion returns the MaxTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetMaxTLSVersion() string {
	if o == nil || o.MaxTLSVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxTLSVersion.Get()
}

// GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetMaxTLSVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxTLSVersion.Get(), o.MaxTLSVersion.IsSet()
}

// HasMaxTLSVersion returns a boolean if a field has been set.
func (o *BackendResponse) HasMaxTLSVersion() bool {
	if o != nil && o.MaxTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTLSVersion gets a reference to the given NullableString and assigns it to the MaxTLSVersion field.
func (o *BackendResponse) SetMaxTLSVersion(v string) {
	o.MaxTLSVersion.Set(&v)
}
// SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil
func (o *BackendResponse) SetMaxTLSVersionNil() {
	o.MaxTLSVersion.Set(nil)
}

// UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
func (o *BackendResponse) UnsetMaxTLSVersion() {
	o.MaxTLSVersion.Unset()
}

// GetMinTLSVersion returns the MinTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetMinTLSVersion() string {
	if o == nil || o.MinTLSVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.MinTLSVersion.Get()
}

// GetMinTLSVersionOk returns a tuple with the MinTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetMinTLSVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinTLSVersion.Get(), o.MinTLSVersion.IsSet()
}

// HasMinTLSVersion returns a boolean if a field has been set.
func (o *BackendResponse) HasMinTLSVersion() bool {
	if o != nil && o.MinTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTLSVersion gets a reference to the given NullableString and assigns it to the MinTLSVersion field.
func (o *BackendResponse) SetMinTLSVersion(v string) {
	o.MinTLSVersion.Set(&v)
}
// SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil
func (o *BackendResponse) SetMinTLSVersionNil() {
	o.MinTLSVersion.Set(nil)
}

// UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
func (o *BackendResponse) UnsetMinTLSVersion() {
	o.MinTLSVersion.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackendResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackendResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackendResponse) SetName(v string) {
	o.Name = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetOverrideHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *BackendResponse) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *BackendResponse) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}
// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *BackendResponse) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *BackendResponse) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *BackendResponse) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *BackendResponse) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *BackendResponse) SetPort(v int32) {
	o.Port = &v
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise.
func (o *BackendResponse) GetRequestCondition() string {
	if o == nil || o.RequestCondition == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetRequestConditionOk() (*string, bool) {
	if o == nil || o.RequestCondition == nil {
		return nil, false
	}
	return o.RequestCondition, true
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *BackendResponse) HasRequestCondition() bool {
	if o != nil && o.RequestCondition != nil {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given string and assigns it to the RequestCondition field.
func (o *BackendResponse) SetRequestCondition(v string) {
	o.RequestCondition = &v
}

// GetShareKey returns the ShareKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetShareKey() string {
	if o == nil || o.ShareKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShareKey.Get()
}

// GetShareKeyOk returns a tuple with the ShareKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetShareKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShareKey.Get(), o.ShareKey.IsSet()
}

// HasShareKey returns a boolean if a field has been set.
func (o *BackendResponse) HasShareKey() bool {
	if o != nil && o.ShareKey.IsSet() {
		return true
	}

	return false
}

// SetShareKey gets a reference to the given NullableString and assigns it to the ShareKey field.
func (o *BackendResponse) SetShareKey(v string) {
	o.ShareKey.Set(&v)
}
// SetShareKeyNil sets the value for ShareKey to be an explicit nil
func (o *BackendResponse) SetShareKeyNil() {
	o.ShareKey.Set(nil)
}

// UnsetShareKey ensures that no value is present for ShareKey, not even an explicit nil
func (o *BackendResponse) UnsetShareKey() {
	o.ShareKey.Unset()
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetShieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *BackendResponse) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *BackendResponse) SetShield(v string) {
	o.Shield.Set(&v)
}
// SetShieldNil sets the value for Shield to be an explicit nil
func (o *BackendResponse) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *BackendResponse) UnsetShield() {
	o.Shield.Unset()
}

// GetSslCaCert returns the SslCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslCaCert() string {
	if o == nil || o.SslCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslCaCert.Get()
}

// GetSslCaCertOk returns a tuple with the SslCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslCaCert.Get(), o.SslCaCert.IsSet()
}

// HasSslCaCert returns a boolean if a field has been set.
func (o *BackendResponse) HasSslCaCert() bool {
	if o != nil && o.SslCaCert.IsSet() {
		return true
	}

	return false
}

// SetSslCaCert gets a reference to the given NullableString and assigns it to the SslCaCert field.
func (o *BackendResponse) SetSslCaCert(v string) {
	o.SslCaCert.Set(&v)
}
// SetSslCaCertNil sets the value for SslCaCert to be an explicit nil
func (o *BackendResponse) SetSslCaCertNil() {
	o.SslCaCert.Set(nil)
}

// UnsetSslCaCert ensures that no value is present for SslCaCert, not even an explicit nil
func (o *BackendResponse) UnsetSslCaCert() {
	o.SslCaCert.Unset()
}

// GetSslCertHostname returns the SslCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslCertHostname() string {
	if o == nil || o.SslCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslCertHostname.Get()
}

// GetSslCertHostnameOk returns a tuple with the SslCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslCertHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslCertHostname.Get(), o.SslCertHostname.IsSet()
}

// HasSslCertHostname returns a boolean if a field has been set.
func (o *BackendResponse) HasSslCertHostname() bool {
	if o != nil && o.SslCertHostname.IsSet() {
		return true
	}

	return false
}

// SetSslCertHostname gets a reference to the given NullableString and assigns it to the SslCertHostname field.
func (o *BackendResponse) SetSslCertHostname(v string) {
	o.SslCertHostname.Set(&v)
}
// SetSslCertHostnameNil sets the value for SslCertHostname to be an explicit nil
func (o *BackendResponse) SetSslCertHostnameNil() {
	o.SslCertHostname.Set(nil)
}

// UnsetSslCertHostname ensures that no value is present for SslCertHostname, not even an explicit nil
func (o *BackendResponse) UnsetSslCertHostname() {
	o.SslCertHostname.Unset()
}

// GetSslCheckCert returns the SslCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslCheckCert() bool {
	if o == nil || o.SslCheckCert.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SslCheckCert.Get()
}

// GetSslCheckCertOk returns a tuple with the SslCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslCheckCertOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslCheckCert.Get(), o.SslCheckCert.IsSet()
}

// HasSslCheckCert returns a boolean if a field has been set.
func (o *BackendResponse) HasSslCheckCert() bool {
	if o != nil && o.SslCheckCert.IsSet() {
		return true
	}

	return false
}

// SetSslCheckCert gets a reference to the given NullableBool and assigns it to the SslCheckCert field.
func (o *BackendResponse) SetSslCheckCert(v bool) {
	o.SslCheckCert.Set(&v)
}
// SetSslCheckCertNil sets the value for SslCheckCert to be an explicit nil
func (o *BackendResponse) SetSslCheckCertNil() {
	o.SslCheckCert.Set(nil)
}

// UnsetSslCheckCert ensures that no value is present for SslCheckCert, not even an explicit nil
func (o *BackendResponse) UnsetSslCheckCert() {
	o.SslCheckCert.Unset()
}

// GetSslCiphers returns the SslCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslCiphers() string {
	if o == nil || o.SslCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslCiphers.Get()
}

// GetSslCiphersOk returns a tuple with the SslCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslCiphersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslCiphers.Get(), o.SslCiphers.IsSet()
}

// HasSslCiphers returns a boolean if a field has been set.
func (o *BackendResponse) HasSslCiphers() bool {
	if o != nil && o.SslCiphers.IsSet() {
		return true
	}

	return false
}

// SetSslCiphers gets a reference to the given NullableString and assigns it to the SslCiphers field.
func (o *BackendResponse) SetSslCiphers(v string) {
	o.SslCiphers.Set(&v)
}
// SetSslCiphersNil sets the value for SslCiphers to be an explicit nil
func (o *BackendResponse) SetSslCiphersNil() {
	o.SslCiphers.Set(nil)
}

// UnsetSslCiphers ensures that no value is present for SslCiphers, not even an explicit nil
func (o *BackendResponse) UnsetSslCiphers() {
	o.SslCiphers.Unset()
}

// GetSslClientCert returns the SslClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslClientCert() string {
	if o == nil || o.SslClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslClientCert.Get()
}

// GetSslClientCertOk returns a tuple with the SslClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslClientCert.Get(), o.SslClientCert.IsSet()
}

// HasSslClientCert returns a boolean if a field has been set.
func (o *BackendResponse) HasSslClientCert() bool {
	if o != nil && o.SslClientCert.IsSet() {
		return true
	}

	return false
}

// SetSslClientCert gets a reference to the given NullableString and assigns it to the SslClientCert field.
func (o *BackendResponse) SetSslClientCert(v string) {
	o.SslClientCert.Set(&v)
}
// SetSslClientCertNil sets the value for SslClientCert to be an explicit nil
func (o *BackendResponse) SetSslClientCertNil() {
	o.SslClientCert.Set(nil)
}

// UnsetSslClientCert ensures that no value is present for SslClientCert, not even an explicit nil
func (o *BackendResponse) UnsetSslClientCert() {
	o.SslClientCert.Unset()
}

// GetSslClientKey returns the SslClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslClientKey() string {
	if o == nil || o.SslClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslClientKey.Get()
}

// GetSslClientKeyOk returns a tuple with the SslClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslClientKey.Get(), o.SslClientKey.IsSet()
}

// HasSslClientKey returns a boolean if a field has been set.
func (o *BackendResponse) HasSslClientKey() bool {
	if o != nil && o.SslClientKey.IsSet() {
		return true
	}

	return false
}

// SetSslClientKey gets a reference to the given NullableString and assigns it to the SslClientKey field.
func (o *BackendResponse) SetSslClientKey(v string) {
	o.SslClientKey.Set(&v)
}
// SetSslClientKeyNil sets the value for SslClientKey to be an explicit nil
func (o *BackendResponse) SetSslClientKeyNil() {
	o.SslClientKey.Set(nil)
}

// UnsetSslClientKey ensures that no value is present for SslClientKey, not even an explicit nil
func (o *BackendResponse) UnsetSslClientKey() {
	o.SslClientKey.Unset()
}

// GetSslHostname returns the SslHostname field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *BackendResponse) GetSslHostname() string {
	if o == nil || o.SslHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslHostname.Get()
}

// GetSslHostnameOk returns a tuple with the SslHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *BackendResponse) GetSslHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslHostname.Get(), o.SslHostname.IsSet()
}

// HasSslHostname returns a boolean if a field has been set.
func (o *BackendResponse) HasSslHostname() bool {
	if o != nil && o.SslHostname.IsSet() {
		return true
	}

	return false
}

// SetSslHostname gets a reference to the given NullableString and assigns it to the SslHostname field.
// Deprecated
func (o *BackendResponse) SetSslHostname(v string) {
	o.SslHostname.Set(&v)
}
// SetSslHostnameNil sets the value for SslHostname to be an explicit nil
func (o *BackendResponse) SetSslHostnameNil() {
	o.SslHostname.Set(nil)
}

// UnsetSslHostname ensures that no value is present for SslHostname, not even an explicit nil
func (o *BackendResponse) UnsetSslHostname() {
	o.SslHostname.Unset()
}

// GetSslSniHostname returns the SslSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetSslSniHostname() string {
	if o == nil || o.SslSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslSniHostname.Get()
}

// GetSslSniHostnameOk returns a tuple with the SslSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetSslSniHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslSniHostname.Get(), o.SslSniHostname.IsSet()
}

// HasSslSniHostname returns a boolean if a field has been set.
func (o *BackendResponse) HasSslSniHostname() bool {
	if o != nil && o.SslSniHostname.IsSet() {
		return true
	}

	return false
}

// SetSslSniHostname gets a reference to the given NullableString and assigns it to the SslSniHostname field.
func (o *BackendResponse) SetSslSniHostname(v string) {
	o.SslSniHostname.Set(&v)
}
// SetSslSniHostnameNil sets the value for SslSniHostname to be an explicit nil
func (o *BackendResponse) SetSslSniHostnameNil() {
	o.SslSniHostname.Set(nil)
}

// UnsetSslSniHostname ensures that no value is present for SslSniHostname, not even an explicit nil
func (o *BackendResponse) UnsetSslSniHostname() {
	o.SslSniHostname.Unset()
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *BackendResponse) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *BackendResponse) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *BackendResponse) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *BackendResponse) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *BackendResponse) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *BackendResponse) SetWeight(v int32) {
	o.Weight = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BackendResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *BackendResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *BackendResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *BackendResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *BackendResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *BackendResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *BackendResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *BackendResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackendResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackendResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BackendResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *BackendResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *BackendResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *BackendResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *BackendResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *BackendResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *BackendResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BackendResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BackendResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BackendResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BackendResponse) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponse) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BackendResponse) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BackendResponse) SetLocked(v bool) {
	o.Locked = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.AutoLoadbalance != nil {
		toSerialize["auto_loadbalance"] = o.AutoLoadbalance
	}
	if o.BetweenBytesTimeout != nil {
		toSerialize["between_bytes_timeout"] = o.BetweenBytesTimeout
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.ConnectTimeout != nil {
		toSerialize["connect_timeout"] = o.ConnectTimeout
	}
	if o.FirstByteTimeout != nil {
		toSerialize["first_byte_timeout"] = o.FirstByteTimeout
	}
	if o.Healthcheck.IsSet() {
		toSerialize["healthcheck"] = o.Healthcheck.Get()
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Ipv4.IsSet() {
		toSerialize["ipv4"] = o.Ipv4.Get()
	}
	if o.Ipv6.IsSet() {
		toSerialize["ipv6"] = o.Ipv6.Get()
	}
	if o.KeepaliveTime.IsSet() {
		toSerialize["keepalive_time"] = o.KeepaliveTime.Get()
	}
	if o.MaxConn != nil {
		toSerialize["max_conn"] = o.MaxConn
	}
	if o.MaxTLSVersion.IsSet() {
		toSerialize["max_tls_version"] = o.MaxTLSVersion.Get()
	}
	if o.MinTLSVersion.IsSet() {
		toSerialize["min_tls_version"] = o.MinTLSVersion.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverrideHost.IsSet() {
		toSerialize["override_host"] = o.OverrideHost.Get()
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.RequestCondition != nil {
		toSerialize["request_condition"] = o.RequestCondition
	}
	if o.ShareKey.IsSet() {
		toSerialize["share_key"] = o.ShareKey.Get()
	}
	if o.Shield.IsSet() {
		toSerialize["shield"] = o.Shield.Get()
	}
	if o.SslCaCert.IsSet() {
		toSerialize["ssl_ca_cert"] = o.SslCaCert.Get()
	}
	if o.SslCertHostname.IsSet() {
		toSerialize["ssl_cert_hostname"] = o.SslCertHostname.Get()
	}
	if o.SslCheckCert.IsSet() {
		toSerialize["ssl_check_cert"] = o.SslCheckCert.Get()
	}
	if o.SslCiphers.IsSet() {
		toSerialize["ssl_ciphers"] = o.SslCiphers.Get()
	}
	if o.SslClientCert.IsSet() {
		toSerialize["ssl_client_cert"] = o.SslClientCert.Get()
	}
	if o.SslClientKey.IsSet() {
		toSerialize["ssl_client_key"] = o.SslClientKey.Get()
	}
	if o.SslHostname.IsSet() {
		toSerialize["ssl_hostname"] = o.SslHostname.Get()
	}
	if o.SslSniHostname.IsSet() {
		toSerialize["ssl_sni_hostname"] = o.SslSniHostname.Get()
	}
	if o.UseSsl != nil {
		toSerialize["use_ssl"] = o.UseSsl
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BackendResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBackendResponse := _BackendResponse{}

	if err = json.Unmarshal(bytes, &varBackendResponse); err == nil {
		*o = BackendResponse(varBackendResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "auto_loadbalance")
		delete(additionalProperties, "between_bytes_timeout")
		delete(additionalProperties, "client_cert")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "connect_timeout")
		delete(additionalProperties, "first_byte_timeout")
		delete(additionalProperties, "healthcheck")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ipv6")
		delete(additionalProperties, "keepalive_time")
		delete(additionalProperties, "max_conn")
		delete(additionalProperties, "max_tls_version")
		delete(additionalProperties, "min_tls_version")
		delete(additionalProperties, "name")
		delete(additionalProperties, "override_host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "share_key")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "ssl_ca_cert")
		delete(additionalProperties, "ssl_cert_hostname")
		delete(additionalProperties, "ssl_check_cert")
		delete(additionalProperties, "ssl_ciphers")
		delete(additionalProperties, "ssl_client_cert")
		delete(additionalProperties, "ssl_client_key")
		delete(additionalProperties, "ssl_hostname")
		delete(additionalProperties, "ssl_sni_hostname")
		delete(additionalProperties, "use_ssl")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "locked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBackendResponse is a helper abstraction for handling nullable backendresponse types. 
type NullableBackendResponse struct {
	value *BackendResponse
	isSet bool
}

// Get returns the value.
func (v NullableBackendResponse) Get() *BackendResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBackendResponse) Set(val *BackendResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBackendResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBackendResponse returns a pointer to a new instance of NullableBackendResponse.
func NewNullableBackendResponse(val *BackendResponse) *NullableBackendResponse {
	return &NullableBackendResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
