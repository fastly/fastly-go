// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"encoding/json"
	"time"
)

// PoolResponsePost struct for PoolResponsePost
type PoolResponsePost struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TLSCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	// Whether to use TLS.
	UseTLS *string `json:"use_tls,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string      `json:"service_id,omitempty"`
	Version   *string      `json:"version,omitempty"`
	// Name for the Pool.
	Name *string `json:"name,omitempty"`
	// Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield NullableString `json:"shield,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
	TLSCiphers NullableString `json:"tls_ciphers,omitempty"`
	// SNI hostname. Optional.
	TLSSniHostname NullableString `json:"tls_sni_hostname,omitempty"`
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
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *string `json:"between_bytes_timeout,omitempty"`
	// How long to wait for a timeout in milliseconds.
	ConnectTimeout *string `json:"connect_timeout,omitempty"`
	// How long to wait for the first byte in milliseconds.
	FirstByteTimeout *string `json:"first_byte_timeout,omitempty"`
	// Maximum number of connections.
	MaxConnDefault *string `json:"max_conn_default,omitempty"`
	// Be strict on checking TLS certs.
	TLSCheckCert NullableString `json:"tls_check_cert,omitempty"`
	ID           *string        `json:"id,omitempty"`
	// Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up.
	Quorum               *int32 `json:"quorum,omitempty"`
	AdditionalProperties map[string]any
}

type _PoolResponsePost PoolResponsePost

// NewPoolResponsePost instantiates a new PoolResponsePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolResponsePost() *PoolResponsePost {
	this := PoolResponsePost{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS string = "0"
	this.UseTLS = &useTLS
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	var quorum int32 = 75
	this.Quorum = &quorum
	return &this
}

// NewPoolResponsePostWithDefaults instantiates a new PoolResponsePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolResponsePostWithDefaults() *PoolResponsePost {
	this := PoolResponsePost{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS string = "0"
	this.UseTLS = &useTLS
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	var quorum int32 = 75
	this.Quorum = &quorum
	return &this
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *PoolResponsePost) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}

// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *PoolResponsePost) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *PoolResponsePost) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}

// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *PoolResponsePost) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *PoolResponsePost) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}

// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *PoolResponsePost) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSCertHostname returns the TLSCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSCertHostname() string {
	if o == nil || o.TLSCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCertHostname.Get()
}

// GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSCertHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCertHostname.Get(), o.TLSCertHostname.IsSet()
}

// HasTLSCertHostname returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSCertHostname() bool {
	if o != nil && o.TLSCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSCertHostname gets a reference to the given NullableString and assigns it to the TLSCertHostname field.
func (o *PoolResponsePost) SetTLSCertHostname(v string) {
	o.TLSCertHostname.Set(&v)
}

// SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil
func (o *PoolResponsePost) SetTLSCertHostnameNil() {
	o.TLSCertHostname.Set(nil)
}

// UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSCertHostname() {
	o.TLSCertHostname.Unset()
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *PoolResponsePost) GetUseTLS() string {
	if o == nil || o.UseTLS == nil {
		var ret string
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetUseTLSOk() (*string, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *PoolResponsePost) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given string and assigns it to the UseTLS field.
func (o *PoolResponsePost) SetUseTLS(v string) {
	o.UseTLS = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PoolResponsePost) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *PoolResponsePost) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PoolResponsePost) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PoolResponsePost) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PoolResponsePost) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *PoolResponsePost) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *PoolResponsePost) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *PoolResponsePost) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PoolResponsePost) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *PoolResponsePost) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *PoolResponsePost) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *PoolResponsePost) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *PoolResponsePost) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *PoolResponsePost) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *PoolResponsePost) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PoolResponsePost) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PoolResponsePost) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PoolResponsePost) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolResponsePost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolResponsePost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolResponsePost) SetName(v string) {
	o.Name = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetShieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *PoolResponsePost) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *PoolResponsePost) SetShield(v string) {
	o.Shield.Set(&v)
}

// SetShieldNil sets the value for Shield to be an explicit nil
func (o *PoolResponsePost) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *PoolResponsePost) UnsetShield() {
	o.Shield.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetRequestConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *PoolResponsePost) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *PoolResponsePost) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}

// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *PoolResponsePost) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *PoolResponsePost) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetTLSCiphers returns the TLSCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSCiphers() string {
	if o == nil || o.TLSCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCiphers.Get()
}

// GetTLSCiphersOk returns a tuple with the TLSCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSCiphersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCiphers.Get(), o.TLSCiphers.IsSet()
}

// HasTLSCiphers returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSCiphers() bool {
	if o != nil && o.TLSCiphers.IsSet() {
		return true
	}

	return false
}

// SetTLSCiphers gets a reference to the given NullableString and assigns it to the TLSCiphers field.
func (o *PoolResponsePost) SetTLSCiphers(v string) {
	o.TLSCiphers.Set(&v)
}

// SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil
func (o *PoolResponsePost) SetTLSCiphersNil() {
	o.TLSCiphers.Set(nil)
}

// UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSCiphers() {
	o.TLSCiphers.Unset()
}

// GetTLSSniHostname returns the TLSSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSSniHostname() string {
	if o == nil || o.TLSSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSSniHostname.Get()
}

// GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSSniHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSSniHostname.Get(), o.TLSSniHostname.IsSet()
}

// HasTLSSniHostname returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSSniHostname() bool {
	if o != nil && o.TLSSniHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSSniHostname gets a reference to the given NullableString and assigns it to the TLSSniHostname field.
func (o *PoolResponsePost) SetTLSSniHostname(v string) {
	o.TLSSniHostname.Set(&v)
}

// SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil
func (o *PoolResponsePost) SetTLSSniHostnameNil() {
	o.TLSSniHostname.Set(nil)
}

// UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSSniHostname() {
	o.TLSSniHostname.Unset()
}

// GetMinTLSVersion returns the MinTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetMinTLSVersion() int32 {
	if o == nil || o.MinTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinTLSVersion.Get()
}

// GetMinTLSVersionOk returns a tuple with the MinTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetMinTLSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTLSVersion.Get(), o.MinTLSVersion.IsSet()
}

// HasMinTLSVersion returns a boolean if a field has been set.
func (o *PoolResponsePost) HasMinTLSVersion() bool {
	if o != nil && o.MinTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTLSVersion gets a reference to the given NullableInt32 and assigns it to the MinTLSVersion field.
func (o *PoolResponsePost) SetMinTLSVersion(v int32) {
	o.MinTLSVersion.Set(&v)
}

// SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil
func (o *PoolResponsePost) SetMinTLSVersionNil() {
	o.MinTLSVersion.Set(nil)
}

// UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
func (o *PoolResponsePost) UnsetMinTLSVersion() {
	o.MinTLSVersion.Unset()
}

// GetMaxTLSVersion returns the MaxTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetMaxTLSVersion() int32 {
	if o == nil || o.MaxTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxTLSVersion.Get()
}

// GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetMaxTLSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTLSVersion.Get(), o.MaxTLSVersion.IsSet()
}

// HasMaxTLSVersion returns a boolean if a field has been set.
func (o *PoolResponsePost) HasMaxTLSVersion() bool {
	if o != nil && o.MaxTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTLSVersion gets a reference to the given NullableInt32 and assigns it to the MaxTLSVersion field.
func (o *PoolResponsePost) SetMaxTLSVersion(v int32) {
	o.MaxTLSVersion.Set(&v)
}

// SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil
func (o *PoolResponsePost) SetMaxTLSVersionNil() {
	o.MaxTLSVersion.Set(nil)
}

// UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
func (o *PoolResponsePost) UnsetMaxTLSVersion() {
	o.MaxTLSVersion.Unset()
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetHealthcheck() string {
	if o == nil || o.Healthcheck.Get() == nil {
		var ret string
		return ret
	}
	return *o.Healthcheck.Get()
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetHealthcheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Healthcheck.Get(), o.Healthcheck.IsSet()
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *PoolResponsePost) HasHealthcheck() bool {
	if o != nil && o.Healthcheck.IsSet() {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given NullableString and assigns it to the Healthcheck field.
func (o *PoolResponsePost) SetHealthcheck(v string) {
	o.Healthcheck.Set(&v)
}

// SetHealthcheckNil sets the value for Healthcheck to be an explicit nil
func (o *PoolResponsePost) SetHealthcheckNil() {
	o.Healthcheck.Set(nil)
}

// UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
func (o *PoolResponsePost) UnsetHealthcheck() {
	o.Healthcheck.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *PoolResponsePost) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *PoolResponsePost) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *PoolResponsePost) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *PoolResponsePost) UnsetComment() {
	o.Comment.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PoolResponsePost) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PoolResponsePost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PoolResponsePost) SetType(v string) {
	o.Type = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetOverrideHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *PoolResponsePost) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *PoolResponsePost) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}

// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *PoolResponsePost) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *PoolResponsePost) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// GetBetweenBytesTimeout returns the BetweenBytesTimeout field value if set, zero value otherwise.
func (o *PoolResponsePost) GetBetweenBytesTimeout() string {
	if o == nil || o.BetweenBytesTimeout == nil {
		var ret string
		return ret
	}
	return *o.BetweenBytesTimeout
}

// GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetBetweenBytesTimeoutOk() (*string, bool) {
	if o == nil || o.BetweenBytesTimeout == nil {
		return nil, false
	}
	return o.BetweenBytesTimeout, true
}

// HasBetweenBytesTimeout returns a boolean if a field has been set.
func (o *PoolResponsePost) HasBetweenBytesTimeout() bool {
	if o != nil && o.BetweenBytesTimeout != nil {
		return true
	}

	return false
}

// SetBetweenBytesTimeout gets a reference to the given string and assigns it to the BetweenBytesTimeout field.
func (o *PoolResponsePost) SetBetweenBytesTimeout(v string) {
	o.BetweenBytesTimeout = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *PoolResponsePost) GetConnectTimeout() string {
	if o == nil || o.ConnectTimeout == nil {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PoolResponsePost) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *PoolResponsePost) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetFirstByteTimeout returns the FirstByteTimeout field value if set, zero value otherwise.
func (o *PoolResponsePost) GetFirstByteTimeout() string {
	if o == nil || o.FirstByteTimeout == nil {
		var ret string
		return ret
	}
	return *o.FirstByteTimeout
}

// GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetFirstByteTimeoutOk() (*string, bool) {
	if o == nil || o.FirstByteTimeout == nil {
		return nil, false
	}
	return o.FirstByteTimeout, true
}

// HasFirstByteTimeout returns a boolean if a field has been set.
func (o *PoolResponsePost) HasFirstByteTimeout() bool {
	if o != nil && o.FirstByteTimeout != nil {
		return true
	}

	return false
}

// SetFirstByteTimeout gets a reference to the given string and assigns it to the FirstByteTimeout field.
func (o *PoolResponsePost) SetFirstByteTimeout(v string) {
	o.FirstByteTimeout = &v
}

// GetMaxConnDefault returns the MaxConnDefault field value if set, zero value otherwise.
func (o *PoolResponsePost) GetMaxConnDefault() string {
	if o == nil || o.MaxConnDefault == nil {
		var ret string
		return ret
	}
	return *o.MaxConnDefault
}

// GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetMaxConnDefaultOk() (*string, bool) {
	if o == nil || o.MaxConnDefault == nil {
		return nil, false
	}
	return o.MaxConnDefault, true
}

// HasMaxConnDefault returns a boolean if a field has been set.
func (o *PoolResponsePost) HasMaxConnDefault() bool {
	if o != nil && o.MaxConnDefault != nil {
		return true
	}

	return false
}

// SetMaxConnDefault gets a reference to the given string and assigns it to the MaxConnDefault field.
func (o *PoolResponsePost) SetMaxConnDefault(v string) {
	o.MaxConnDefault = &v
}

// GetTLSCheckCert returns the TLSCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponsePost) GetTLSCheckCert() string {
	if o == nil || o.TLSCheckCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCheckCert.Get()
}

// GetTLSCheckCertOk returns a tuple with the TLSCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponsePost) GetTLSCheckCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCheckCert.Get(), o.TLSCheckCert.IsSet()
}

// HasTLSCheckCert returns a boolean if a field has been set.
func (o *PoolResponsePost) HasTLSCheckCert() bool {
	if o != nil && o.TLSCheckCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCheckCert gets a reference to the given NullableString and assigns it to the TLSCheckCert field.
func (o *PoolResponsePost) SetTLSCheckCert(v string) {
	o.TLSCheckCert.Set(&v)
}

// SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil
func (o *PoolResponsePost) SetTLSCheckCertNil() {
	o.TLSCheckCert.Set(nil)
}

// UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
func (o *PoolResponsePost) UnsetTLSCheckCert() {
	o.TLSCheckCert.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *PoolResponsePost) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *PoolResponsePost) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *PoolResponsePost) SetID(v string) {
	o.ID = &v
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *PoolResponsePost) GetQuorum() int32 {
	if o == nil || o.Quorum == nil {
		var ret int32
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponsePost) GetQuorumOk() (*int32, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *PoolResponsePost) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given int32 and assigns it to the Quorum field.
func (o *PoolResponsePost) SetQuorum(v int32) {
	o.Quorum = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PoolResponsePost) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Shield.IsSet() {
		toSerialize["shield"] = o.Shield.Get()
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.TLSCiphers.IsSet() {
		toSerialize["tls_ciphers"] = o.TLSCiphers.Get()
	}
	if o.TLSSniHostname.IsSet() {
		toSerialize["tls_sni_hostname"] = o.TLSSniHostname.Get()
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
	if o.BetweenBytesTimeout != nil {
		toSerialize["between_bytes_timeout"] = o.BetweenBytesTimeout
	}
	if o.ConnectTimeout != nil {
		toSerialize["connect_timeout"] = o.ConnectTimeout
	}
	if o.FirstByteTimeout != nil {
		toSerialize["first_byte_timeout"] = o.FirstByteTimeout
	}
	if o.MaxConnDefault != nil {
		toSerialize["max_conn_default"] = o.MaxConnDefault
	}
	if o.TLSCheckCert.IsSet() {
		toSerialize["tls_check_cert"] = o.TLSCheckCert.Get()
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Quorum != nil {
		toSerialize["quorum"] = o.Quorum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PoolResponsePost) UnmarshalJSON(bytes []byte) (err error) {
	varPoolResponsePost := _PoolResponsePost{}

	if err = json.Unmarshal(bytes, &varPoolResponsePost); err == nil {
		*o = PoolResponsePost(varPoolResponsePost)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_cert_hostname")
		delete(additionalProperties, "use_tls")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "name")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "tls_ciphers")
		delete(additionalProperties, "tls_sni_hostname")
		delete(additionalProperties, "min_tls_version")
		delete(additionalProperties, "max_tls_version")
		delete(additionalProperties, "healthcheck")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "type")
		delete(additionalProperties, "override_host")
		delete(additionalProperties, "between_bytes_timeout")
		delete(additionalProperties, "connect_timeout")
		delete(additionalProperties, "first_byte_timeout")
		delete(additionalProperties, "max_conn_default")
		delete(additionalProperties, "tls_check_cert")
		delete(additionalProperties, "id")
		delete(additionalProperties, "quorum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePoolResponsePost is a helper abstraction for handling nullable poolresponsepost types.
type NullablePoolResponsePost struct {
	value *PoolResponsePost
	isSet bool
}

// Get returns the value.
func (v NullablePoolResponsePost) Get() *PoolResponsePost {
	return v.value
}

// Set modifies the value.
func (v *NullablePoolResponsePost) Set(val *PoolResponsePost) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePoolResponsePost) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePoolResponsePost) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePoolResponsePost returns a pointer to a new instance of NullablePoolResponsePost.
func NewNullablePoolResponsePost(val *PoolResponsePost) *NullablePoolResponsePost {
	return &NullablePoolResponsePost{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePoolResponsePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePoolResponsePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
