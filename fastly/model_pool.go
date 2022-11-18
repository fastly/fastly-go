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
)

// Pool struct for Pool
type Pool struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TLSCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	// Whether to use TLS.
	UseTLS *int32 `json:"use_tls,omitempty"`
	// Name for the Pool.
	Name *string `json:"name,omitempty"`
	// Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield NullableString `json:"shield,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Maximum number of connections. Optional.
	MaxConnDefault *int32 `json:"max_conn_default,omitempty"`
	// How long to wait for a timeout in milliseconds. Optional.
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`
	// How long to wait for the first byte in milliseconds. Optional.
	FirstByteTimeout *int32 `json:"first_byte_timeout,omitempty"`
	// Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up.
	Quorum *int32 `json:"quorum,omitempty"`
	// List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
	TLSCiphers NullableString `json:"tls_ciphers,omitempty"`
	// SNI hostname. Optional.
	TLSSniHostname NullableString `json:"tls_sni_hostname,omitempty"`
	// Be strict on checking TLS certs. Optional.
	TLSCheckCert NullableInt32 `json:"tls_check_cert,omitempty"`
	// Minimum allowed TLS version on connections to this server. Optional.
	MinTLSVersion NullableInt32 `json:"min_tls_version,omitempty"`
	// Maximum allowed TLS version on connections to this server. Optional.
	MaxTLSVersion NullableInt32 `json:"max_tls_version,omitempty"`
	// Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.
	Healthcheck NullableString `json:"healthcheck,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// What type of load balance group to use.
	Type *string `json:"type,omitempty"`
	// The hostname to [override the Host header](https://docs.fastly.com/en/guides/specifying-an-override-host). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting.
	OverrideHost NullableString `json:"override_host,omitempty"`
	AdditionalProperties map[string]any
}

type _Pool Pool

// NewPool instantiates a new Pool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPool() *Pool {
	this := Pool{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS int32 = 0
	this.UseTLS = &useTLS
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var maxConnDefault int32 = 200
	this.MaxConnDefault = &maxConnDefault
	var quorum int32 = 75
	this.Quorum = &quorum
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// NewPoolWithDefaults instantiates a new Pool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolWithDefaults() *Pool {
	this := Pool{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS int32 = 0
	this.UseTLS = &useTLS
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var maxConnDefault int32 = 200
	this.MaxConnDefault = &maxConnDefault
	var quorum int32 = 75
	this.Quorum = &quorum
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *Pool) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *Pool) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *Pool) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *Pool) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *Pool) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *Pool) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *Pool) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *Pool) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *Pool) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *Pool) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *Pool) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *Pool) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSCertHostname returns the TLSCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSCertHostname() string {
	if o == nil || o.TLSCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCertHostname.Get()
}

// GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSCertHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCertHostname.Get(), o.TLSCertHostname.IsSet()
}

// HasTLSCertHostname returns a boolean if a field has been set.
func (o *Pool) HasTLSCertHostname() bool {
	if o != nil && o.TLSCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSCertHostname gets a reference to the given NullableString and assigns it to the TLSCertHostname field.
func (o *Pool) SetTLSCertHostname(v string) {
	o.TLSCertHostname.Set(&v)
}
// SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil
func (o *Pool) SetTLSCertHostnameNil() {
	o.TLSCertHostname.Set(nil)
}

// UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
func (o *Pool) UnsetTLSCertHostname() {
	o.TLSCertHostname.Unset()
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *Pool) GetUseTLS() int32 {
	if o == nil || o.UseTLS == nil {
		var ret int32
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetUseTLSOk() (*int32, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *Pool) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given int32 and assigns it to the UseTLS field.
func (o *Pool) SetUseTLS(v int32) {
	o.UseTLS = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Pool) SetName(v string) {
	o.Name = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetShieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *Pool) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *Pool) SetShield(v string) {
	o.Shield.Set(&v)
}
// SetShieldNil sets the value for Shield to be an explicit nil
func (o *Pool) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *Pool) UnsetShield() {
	o.Shield.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *Pool) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *Pool) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *Pool) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *Pool) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetMaxConnDefault returns the MaxConnDefault field value if set, zero value otherwise.
func (o *Pool) GetMaxConnDefault() int32 {
	if o == nil || o.MaxConnDefault == nil {
		var ret int32
		return ret
	}
	return *o.MaxConnDefault
}

// GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetMaxConnDefaultOk() (*int32, bool) {
	if o == nil || o.MaxConnDefault == nil {
		return nil, false
	}
	return o.MaxConnDefault, true
}

// HasMaxConnDefault returns a boolean if a field has been set.
func (o *Pool) HasMaxConnDefault() bool {
	if o != nil && o.MaxConnDefault != nil {
		return true
	}

	return false
}

// SetMaxConnDefault gets a reference to the given int32 and assigns it to the MaxConnDefault field.
func (o *Pool) SetMaxConnDefault(v int32) {
	o.MaxConnDefault = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *Pool) GetConnectTimeout() int32 {
	if o == nil || o.ConnectTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *Pool) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *Pool) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetFirstByteTimeout returns the FirstByteTimeout field value if set, zero value otherwise.
func (o *Pool) GetFirstByteTimeout() int32 {
	if o == nil || o.FirstByteTimeout == nil {
		var ret int32
		return ret
	}
	return *o.FirstByteTimeout
}

// GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetFirstByteTimeoutOk() (*int32, bool) {
	if o == nil || o.FirstByteTimeout == nil {
		return nil, false
	}
	return o.FirstByteTimeout, true
}

// HasFirstByteTimeout returns a boolean if a field has been set.
func (o *Pool) HasFirstByteTimeout() bool {
	if o != nil && o.FirstByteTimeout != nil {
		return true
	}

	return false
}

// SetFirstByteTimeout gets a reference to the given int32 and assigns it to the FirstByteTimeout field.
func (o *Pool) SetFirstByteTimeout(v int32) {
	o.FirstByteTimeout = &v
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *Pool) GetQuorum() int32 {
	if o == nil || o.Quorum == nil {
		var ret int32
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetQuorumOk() (*int32, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *Pool) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given int32 and assigns it to the Quorum field.
func (o *Pool) SetQuorum(v int32) {
	o.Quorum = &v
}

// GetTLSCiphers returns the TLSCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSCiphers() string {
	if o == nil || o.TLSCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCiphers.Get()
}

// GetTLSCiphersOk returns a tuple with the TLSCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSCiphersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCiphers.Get(), o.TLSCiphers.IsSet()
}

// HasTLSCiphers returns a boolean if a field has been set.
func (o *Pool) HasTLSCiphers() bool {
	if o != nil && o.TLSCiphers.IsSet() {
		return true
	}

	return false
}

// SetTLSCiphers gets a reference to the given NullableString and assigns it to the TLSCiphers field.
func (o *Pool) SetTLSCiphers(v string) {
	o.TLSCiphers.Set(&v)
}
// SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil
func (o *Pool) SetTLSCiphersNil() {
	o.TLSCiphers.Set(nil)
}

// UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
func (o *Pool) UnsetTLSCiphers() {
	o.TLSCiphers.Unset()
}

// GetTLSSniHostname returns the TLSSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSSniHostname() string {
	if o == nil || o.TLSSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSSniHostname.Get()
}

// GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSSniHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSSniHostname.Get(), o.TLSSniHostname.IsSet()
}

// HasTLSSniHostname returns a boolean if a field has been set.
func (o *Pool) HasTLSSniHostname() bool {
	if o != nil && o.TLSSniHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSSniHostname gets a reference to the given NullableString and assigns it to the TLSSniHostname field.
func (o *Pool) SetTLSSniHostname(v string) {
	o.TLSSniHostname.Set(&v)
}
// SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil
func (o *Pool) SetTLSSniHostnameNil() {
	o.TLSSniHostname.Set(nil)
}

// UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
func (o *Pool) UnsetTLSSniHostname() {
	o.TLSSniHostname.Unset()
}

// GetTLSCheckCert returns the TLSCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetTLSCheckCert() int32 {
	if o == nil || o.TLSCheckCert.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TLSCheckCert.Get()
}

// GetTLSCheckCertOk returns a tuple with the TLSCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetTLSCheckCertOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCheckCert.Get(), o.TLSCheckCert.IsSet()
}

// HasTLSCheckCert returns a boolean if a field has been set.
func (o *Pool) HasTLSCheckCert() bool {
	if o != nil && o.TLSCheckCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCheckCert gets a reference to the given NullableInt32 and assigns it to the TLSCheckCert field.
func (o *Pool) SetTLSCheckCert(v int32) {
	o.TLSCheckCert.Set(&v)
}
// SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil
func (o *Pool) SetTLSCheckCertNil() {
	o.TLSCheckCert.Set(nil)
}

// UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
func (o *Pool) UnsetTLSCheckCert() {
	o.TLSCheckCert.Unset()
}

// GetMinTLSVersion returns the MinTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetMinTLSVersion() int32 {
	if o == nil || o.MinTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinTLSVersion.Get()
}

// GetMinTLSVersionOk returns a tuple with the MinTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetMinTLSVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinTLSVersion.Get(), o.MinTLSVersion.IsSet()
}

// HasMinTLSVersion returns a boolean if a field has been set.
func (o *Pool) HasMinTLSVersion() bool {
	if o != nil && o.MinTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTLSVersion gets a reference to the given NullableInt32 and assigns it to the MinTLSVersion field.
func (o *Pool) SetMinTLSVersion(v int32) {
	o.MinTLSVersion.Set(&v)
}
// SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil
func (o *Pool) SetMinTLSVersionNil() {
	o.MinTLSVersion.Set(nil)
}

// UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
func (o *Pool) UnsetMinTLSVersion() {
	o.MinTLSVersion.Unset()
}

// GetMaxTLSVersion returns the MaxTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetMaxTLSVersion() int32 {
	if o == nil || o.MaxTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxTLSVersion.Get()
}

// GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetMaxTLSVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxTLSVersion.Get(), o.MaxTLSVersion.IsSet()
}

// HasMaxTLSVersion returns a boolean if a field has been set.
func (o *Pool) HasMaxTLSVersion() bool {
	if o != nil && o.MaxTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTLSVersion gets a reference to the given NullableInt32 and assigns it to the MaxTLSVersion field.
func (o *Pool) SetMaxTLSVersion(v int32) {
	o.MaxTLSVersion.Set(&v)
}
// SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil
func (o *Pool) SetMaxTLSVersionNil() {
	o.MaxTLSVersion.Set(nil)
}

// UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
func (o *Pool) UnsetMaxTLSVersion() {
	o.MaxTLSVersion.Unset()
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetHealthcheck() string {
	if o == nil || o.Healthcheck.Get() == nil {
		var ret string
		return ret
	}
	return *o.Healthcheck.Get()
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetHealthcheckOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Healthcheck.Get(), o.Healthcheck.IsSet()
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *Pool) HasHealthcheck() bool {
	if o != nil && o.Healthcheck.IsSet() {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given NullableString and assigns it to the Healthcheck field.
func (o *Pool) SetHealthcheck(v string) {
	o.Healthcheck.Set(&v)
}
// SetHealthcheckNil sets the value for Healthcheck to be an explicit nil
func (o *Pool) SetHealthcheckNil() {
	o.Healthcheck.Set(nil)
}

// UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
func (o *Pool) UnsetHealthcheck() {
	o.Healthcheck.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Pool) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Pool) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Pool) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Pool) UnsetComment() {
	o.Comment.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Pool) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Pool) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Pool) SetType(v string) {
	o.Type = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pool) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pool) GetOverrideHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *Pool) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *Pool) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}
// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *Pool) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *Pool) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Pool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TLSCaCert.Get()
	}
	if o.TLSClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TLSClientCert.Get()
	}
	if o.TLSClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TLSClientKey.Get()
	}
	if o.TLSCertHostname.IsSet() {
		toSerialize["tls_cert_hostname"] = o.TLSCertHostname.Get()
	}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Shield.IsSet() {
		toSerialize["shield"] = o.Shield.Get()
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.MaxConnDefault != nil {
		toSerialize["max_conn_default"] = o.MaxConnDefault
	}
	if o.ConnectTimeout != nil {
		toSerialize["connect_timeout"] = o.ConnectTimeout
	}
	if o.FirstByteTimeout != nil {
		toSerialize["first_byte_timeout"] = o.FirstByteTimeout
	}
	if o.Quorum != nil {
		toSerialize["quorum"] = o.Quorum
	}
	if o.TLSCiphers.IsSet() {
		toSerialize["tls_ciphers"] = o.TLSCiphers.Get()
	}
	if o.TLSSniHostname.IsSet() {
		toSerialize["tls_sni_hostname"] = o.TLSSniHostname.Get()
	}
	if o.TLSCheckCert.IsSet() {
		toSerialize["tls_check_cert"] = o.TLSCheckCert.Get()
	}
	if o.MinTLSVersion.IsSet() {
		toSerialize["min_tls_version"] = o.MinTLSVersion.Get()
	}
	if o.MaxTLSVersion.IsSet() {
		toSerialize["max_tls_version"] = o.MaxTLSVersion.Get()
	}
	if o.Healthcheck.IsSet() {
		toSerialize["healthcheck"] = o.Healthcheck.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.OverrideHost.IsSet() {
		toSerialize["override_host"] = o.OverrideHost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Pool) UnmarshalJSON(bytes []byte) (err error) {
	varPool := _Pool{}

	if err = json.Unmarshal(bytes, &varPool); err == nil {
		*o = Pool(varPool)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_cert_hostname")
		delete(additionalProperties, "use_tls")
		delete(additionalProperties, "name")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "max_conn_default")
		delete(additionalProperties, "connect_timeout")
		delete(additionalProperties, "first_byte_timeout")
		delete(additionalProperties, "quorum")
		delete(additionalProperties, "tls_ciphers")
		delete(additionalProperties, "tls_sni_hostname")
		delete(additionalProperties, "tls_check_cert")
		delete(additionalProperties, "min_tls_version")
		delete(additionalProperties, "max_tls_version")
		delete(additionalProperties, "healthcheck")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "type")
		delete(additionalProperties, "override_host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePool is a helper abstraction for handling nullable pool types. 
type NullablePool struct {
	value *Pool
	isSet bool
}

// Get returns the value.
func (v NullablePool) Get() *Pool {
	return v.value
}

// Set modifies the value.
func (v *NullablePool) Set(val *Pool) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePool) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePool) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePool returns a pointer to a new instance of NullablePool.
func NewNullablePool(val *Pool) *NullablePool {
	return &NullablePool{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
