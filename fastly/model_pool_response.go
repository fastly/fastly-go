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

// PoolResponse struct for PoolResponse
type PoolResponse struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TlsCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TlsClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TlsClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TlsCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	// Whether to use TLS.
	UseTls *string `json:"use_tls,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceId *string      `json:"service_id,omitempty"`
	Version   *string      `json:"version,omitempty"`
	// Name for the Pool.
	Name *string `json:"name,omitempty"`
	// Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield NullableString `json:"shield,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
	TlsCiphers NullableString `json:"tls_ciphers,omitempty"`
	// SNI hostname. Optional.
	TlsSniHostname NullableString `json:"tls_sni_hostname,omitempty"`
	// Minimum allowed TLS version on connections to this server. Optional.
	MinTlsVersion NullableInt32 `json:"min_tls_version,omitempty"`
	// Maximum allowed TLS version on connections to this server. Optional.
	MaxTlsVersion NullableInt32 `json:"max_tls_version,omitempty"`
	// Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.
	Healthcheck NullableString `json:"healthcheck,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// What type of load balance group to use.
	Type *string `json:"type,omitempty"`
	// The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting.
	OverrideHost NullableString `json:"override_host,omitempty"`
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *string `json:"between_bytes_timeout,omitempty"`
	// How long to wait for a timeout in milliseconds.
	ConnectTimeout *string `json:"connect_timeout,omitempty"`
	// How long to wait for the first byte in milliseconds.
	FirstByteTimeout *string `json:"first_byte_timeout,omitempty"`
	// Maximum number of connections.
	MaxConnDefault *string `json:"max_conn_default,omitempty"`
	// Be strict on checking TLS certs.
	TlsCheckCert NullableString `json:"tls_check_cert,omitempty"`
	Id           *string        `json:"id,omitempty"`
	// Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up.
	Quorum               *string `json:"quorum,omitempty"`
	AdditionalProperties map[string]any
}

type _PoolResponse PoolResponse

// NewPoolResponse instantiates a new PoolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolResponse() *PoolResponse {
	this := PoolResponse{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TlsCertHostname = *NewNullableString(&tlsCertHostname)
	var useTls string = "0"
	this.UseTls = &useTls
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	var quorum string = "75"
	this.Quorum = &quorum
	return &this
}

// NewPoolResponseWithDefaults instantiates a new PoolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolResponseWithDefaults() *PoolResponse {
	this := PoolResponse{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TlsCertHostname = *NewNullableString(&tlsCertHostname)
	var useTls string = "0"
	this.UseTls = &useTls
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	var quorum string = "75"
	this.Quorum = &quorum
	return &this
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *PoolResponse) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *PoolResponse) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *PoolResponse) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *PoolResponse) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *PoolResponse) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *PoolResponse) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *PoolResponse) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *PoolResponse) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *PoolResponse) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsCertHostname returns the TlsCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsCertHostname() string {
	if o == nil || o.TlsCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCertHostname.Get()
}

// GetTlsCertHostnameOk returns a tuple with the TlsCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsCertHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCertHostname.Get(), o.TlsCertHostname.IsSet()
}

// HasTlsCertHostname returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsCertHostname() bool {
	if o != nil && o.TlsCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsCertHostname gets a reference to the given NullableString and assigns it to the TlsCertHostname field.
func (o *PoolResponse) SetTlsCertHostname(v string) {
	o.TlsCertHostname.Set(&v)
}

// SetTlsCertHostnameNil sets the value for TlsCertHostname to be an explicit nil
func (o *PoolResponse) SetTlsCertHostnameNil() {
	o.TlsCertHostname.Set(nil)
}

// UnsetTlsCertHostname ensures that no value is present for TlsCertHostname, not even an explicit nil
func (o *PoolResponse) UnsetTlsCertHostname() {
	o.TlsCertHostname.Unset()
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *PoolResponse) GetUseTls() string {
	if o == nil || o.UseTls == nil {
		var ret string
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetUseTlsOk() (*string, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *PoolResponse) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given string and assigns it to the UseTls field.
func (o *PoolResponse) SetUseTls(v string) {
	o.UseTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PoolResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *PoolResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PoolResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PoolResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PoolResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *PoolResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *PoolResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *PoolResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PoolResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *PoolResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *PoolResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *PoolResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PoolResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PoolResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *PoolResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PoolResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PoolResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PoolResponse) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolResponse) SetName(v string) {
	o.Name = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetShieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *PoolResponse) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *PoolResponse) SetShield(v string) {
	o.Shield.Set(&v)
}

// SetShieldNil sets the value for Shield to be an explicit nil
func (o *PoolResponse) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *PoolResponse) UnsetShield() {
	o.Shield.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetRequestConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *PoolResponse) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *PoolResponse) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}

// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *PoolResponse) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *PoolResponse) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetTlsCiphers returns the TlsCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsCiphers() string {
	if o == nil || o.TlsCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCiphers.Get()
}

// GetTlsCiphersOk returns a tuple with the TlsCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsCiphersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCiphers.Get(), o.TlsCiphers.IsSet()
}

// HasTlsCiphers returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsCiphers() bool {
	if o != nil && o.TlsCiphers.IsSet() {
		return true
	}

	return false
}

// SetTlsCiphers gets a reference to the given NullableString and assigns it to the TlsCiphers field.
func (o *PoolResponse) SetTlsCiphers(v string) {
	o.TlsCiphers.Set(&v)
}

// SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil
func (o *PoolResponse) SetTlsCiphersNil() {
	o.TlsCiphers.Set(nil)
}

// UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
func (o *PoolResponse) UnsetTlsCiphers() {
	o.TlsCiphers.Unset()
}

// GetTlsSniHostname returns the TlsSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsSniHostname() string {
	if o == nil || o.TlsSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsSniHostname.Get()
}

// GetTlsSniHostnameOk returns a tuple with the TlsSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsSniHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsSniHostname.Get(), o.TlsSniHostname.IsSet()
}

// HasTlsSniHostname returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsSniHostname() bool {
	if o != nil && o.TlsSniHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsSniHostname gets a reference to the given NullableString and assigns it to the TlsSniHostname field.
func (o *PoolResponse) SetTlsSniHostname(v string) {
	o.TlsSniHostname.Set(&v)
}

// SetTlsSniHostnameNil sets the value for TlsSniHostname to be an explicit nil
func (o *PoolResponse) SetTlsSniHostnameNil() {
	o.TlsSniHostname.Set(nil)
}

// UnsetTlsSniHostname ensures that no value is present for TlsSniHostname, not even an explicit nil
func (o *PoolResponse) UnsetTlsSniHostname() {
	o.TlsSniHostname.Unset()
}

// GetMinTlsVersion returns the MinTlsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetMinTlsVersion() int32 {
	if o == nil || o.MinTlsVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinTlsVersion.Get()
}

// GetMinTlsVersionOk returns a tuple with the MinTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetMinTlsVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTlsVersion.Get(), o.MinTlsVersion.IsSet()
}

// HasMinTlsVersion returns a boolean if a field has been set.
func (o *PoolResponse) HasMinTlsVersion() bool {
	if o != nil && o.MinTlsVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTlsVersion gets a reference to the given NullableInt32 and assigns it to the MinTlsVersion field.
func (o *PoolResponse) SetMinTlsVersion(v int32) {
	o.MinTlsVersion.Set(&v)
}

// SetMinTlsVersionNil sets the value for MinTlsVersion to be an explicit nil
func (o *PoolResponse) SetMinTlsVersionNil() {
	o.MinTlsVersion.Set(nil)
}

// UnsetMinTlsVersion ensures that no value is present for MinTlsVersion, not even an explicit nil
func (o *PoolResponse) UnsetMinTlsVersion() {
	o.MinTlsVersion.Unset()
}

// GetMaxTlsVersion returns the MaxTlsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetMaxTlsVersion() int32 {
	if o == nil || o.MaxTlsVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxTlsVersion.Get()
}

// GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetMaxTlsVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTlsVersion.Get(), o.MaxTlsVersion.IsSet()
}

// HasMaxTlsVersion returns a boolean if a field has been set.
func (o *PoolResponse) HasMaxTlsVersion() bool {
	if o != nil && o.MaxTlsVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTlsVersion gets a reference to the given NullableInt32 and assigns it to the MaxTlsVersion field.
func (o *PoolResponse) SetMaxTlsVersion(v int32) {
	o.MaxTlsVersion.Set(&v)
}

// SetMaxTlsVersionNil sets the value for MaxTlsVersion to be an explicit nil
func (o *PoolResponse) SetMaxTlsVersionNil() {
	o.MaxTlsVersion.Set(nil)
}

// UnsetMaxTlsVersion ensures that no value is present for MaxTlsVersion, not even an explicit nil
func (o *PoolResponse) UnsetMaxTlsVersion() {
	o.MaxTlsVersion.Unset()
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetHealthcheck() string {
	if o == nil || o.Healthcheck.Get() == nil {
		var ret string
		return ret
	}
	return *o.Healthcheck.Get()
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetHealthcheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Healthcheck.Get(), o.Healthcheck.IsSet()
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *PoolResponse) HasHealthcheck() bool {
	if o != nil && o.Healthcheck.IsSet() {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given NullableString and assigns it to the Healthcheck field.
func (o *PoolResponse) SetHealthcheck(v string) {
	o.Healthcheck.Set(&v)
}

// SetHealthcheckNil sets the value for Healthcheck to be an explicit nil
func (o *PoolResponse) SetHealthcheckNil() {
	o.Healthcheck.Set(nil)
}

// UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
func (o *PoolResponse) UnsetHealthcheck() {
	o.Healthcheck.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *PoolResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *PoolResponse) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *PoolResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *PoolResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PoolResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PoolResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PoolResponse) SetType(v string) {
	o.Type = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetOverrideHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *PoolResponse) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *PoolResponse) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}

// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *PoolResponse) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *PoolResponse) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// GetBetweenBytesTimeout returns the BetweenBytesTimeout field value if set, zero value otherwise.
func (o *PoolResponse) GetBetweenBytesTimeout() string {
	if o == nil || o.BetweenBytesTimeout == nil {
		var ret string
		return ret
	}
	return *o.BetweenBytesTimeout
}

// GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetBetweenBytesTimeoutOk() (*string, bool) {
	if o == nil || o.BetweenBytesTimeout == nil {
		return nil, false
	}
	return o.BetweenBytesTimeout, true
}

// HasBetweenBytesTimeout returns a boolean if a field has been set.
func (o *PoolResponse) HasBetweenBytesTimeout() bool {
	if o != nil && o.BetweenBytesTimeout != nil {
		return true
	}

	return false
}

// SetBetweenBytesTimeout gets a reference to the given string and assigns it to the BetweenBytesTimeout field.
func (o *PoolResponse) SetBetweenBytesTimeout(v string) {
	o.BetweenBytesTimeout = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *PoolResponse) GetConnectTimeout() string {
	if o == nil || o.ConnectTimeout == nil {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PoolResponse) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *PoolResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetFirstByteTimeout returns the FirstByteTimeout field value if set, zero value otherwise.
func (o *PoolResponse) GetFirstByteTimeout() string {
	if o == nil || o.FirstByteTimeout == nil {
		var ret string
		return ret
	}
	return *o.FirstByteTimeout
}

// GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetFirstByteTimeoutOk() (*string, bool) {
	if o == nil || o.FirstByteTimeout == nil {
		return nil, false
	}
	return o.FirstByteTimeout, true
}

// HasFirstByteTimeout returns a boolean if a field has been set.
func (o *PoolResponse) HasFirstByteTimeout() bool {
	if o != nil && o.FirstByteTimeout != nil {
		return true
	}

	return false
}

// SetFirstByteTimeout gets a reference to the given string and assigns it to the FirstByteTimeout field.
func (o *PoolResponse) SetFirstByteTimeout(v string) {
	o.FirstByteTimeout = &v
}

// GetMaxConnDefault returns the MaxConnDefault field value if set, zero value otherwise.
func (o *PoolResponse) GetMaxConnDefault() string {
	if o == nil || o.MaxConnDefault == nil {
		var ret string
		return ret
	}
	return *o.MaxConnDefault
}

// GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetMaxConnDefaultOk() (*string, bool) {
	if o == nil || o.MaxConnDefault == nil {
		return nil, false
	}
	return o.MaxConnDefault, true
}

// HasMaxConnDefault returns a boolean if a field has been set.
func (o *PoolResponse) HasMaxConnDefault() bool {
	if o != nil && o.MaxConnDefault != nil {
		return true
	}

	return false
}

// SetMaxConnDefault gets a reference to the given string and assigns it to the MaxConnDefault field.
func (o *PoolResponse) SetMaxConnDefault(v string) {
	o.MaxConnDefault = &v
}

// GetTlsCheckCert returns the TlsCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponse) GetTlsCheckCert() string {
	if o == nil || o.TlsCheckCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCheckCert.Get()
}

// GetTlsCheckCertOk returns a tuple with the TlsCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponse) GetTlsCheckCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCheckCert.Get(), o.TlsCheckCert.IsSet()
}

// HasTlsCheckCert returns a boolean if a field has been set.
func (o *PoolResponse) HasTlsCheckCert() bool {
	if o != nil && o.TlsCheckCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCheckCert gets a reference to the given NullableString and assigns it to the TlsCheckCert field.
func (o *PoolResponse) SetTlsCheckCert(v string) {
	o.TlsCheckCert.Set(&v)
}

// SetTlsCheckCertNil sets the value for TlsCheckCert to be an explicit nil
func (o *PoolResponse) SetTlsCheckCertNil() {
	o.TlsCheckCert.Set(nil)
}

// UnsetTlsCheckCert ensures that no value is present for TlsCheckCert, not even an explicit nil
func (o *PoolResponse) UnsetTlsCheckCert() {
	o.TlsCheckCert.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolResponse) SetId(v string) {
	o.Id = &v
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *PoolResponse) GetQuorum() string {
	if o == nil || o.Quorum == nil {
		var ret string
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponse) GetQuorumOk() (*string, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *PoolResponse) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given string and assigns it to the Quorum field.
func (o *PoolResponse) SetQuorum(v string) {
	o.Quorum = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PoolResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TlsCaCert.Get()
	}
	if o.TlsClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TlsClientCert.Get()
	}
	if o.TlsClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TlsClientKey.Get()
	}
	if o.TlsCertHostname.IsSet() {
		toSerialize["tls_cert_hostname"] = o.TlsCertHostname.Get()
	}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
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
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
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
	if o.TlsCiphers.IsSet() {
		toSerialize["tls_ciphers"] = o.TlsCiphers.Get()
	}
	if o.TlsSniHostname.IsSet() {
		toSerialize["tls_sni_hostname"] = o.TlsSniHostname.Get()
	}
	if o.MinTlsVersion.IsSet() {
		toSerialize["min_tls_version"] = o.MinTlsVersion.Get()
	}
	if o.MaxTlsVersion.IsSet() {
		toSerialize["max_tls_version"] = o.MaxTlsVersion.Get()
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
	if o.TlsCheckCert.IsSet() {
		toSerialize["tls_check_cert"] = o.TlsCheckCert.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *PoolResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPoolResponse := _PoolResponse{}

	if err = json.Unmarshal(bytes, &varPoolResponse); err == nil {
		*o = PoolResponse(varPoolResponse)
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

// NullablePoolResponse is a helper abstraction for handling nullable poolresponse types.
type NullablePoolResponse struct {
	value *PoolResponse
	isSet bool
}

// Get returns the value.
func (v NullablePoolResponse) Get() *PoolResponse {
	return v.value
}

// Set modifies the value.
func (v *NullablePoolResponse) Set(val *PoolResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePoolResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePoolResponse returns a pointer to a new instance of NullablePoolResponse.
func NewNullablePoolResponse(val *PoolResponse) *NullablePoolResponse {
	return &NullablePoolResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
