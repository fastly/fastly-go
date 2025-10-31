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
)

// PoolAdditional struct for PoolAdditional
type PoolAdditional struct {
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
	OverrideHost         NullableString `json:"override_host,omitempty"`
	AdditionalProperties map[string]any
}

type _PoolAdditional PoolAdditional

// NewPoolAdditional instantiates a new PoolAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAdditional() *PoolAdditional {
	this := PoolAdditional{}
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// NewPoolAdditionalWithDefaults instantiates a new PoolAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAdditionalWithDefaults() *PoolAdditional {
	this := PoolAdditional{}
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolAdditional) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAdditional) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolAdditional) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolAdditional) SetName(v string) {
	o.Name = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetShieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *PoolAdditional) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *PoolAdditional) SetShield(v string) {
	o.Shield.Set(&v)
}

// SetShieldNil sets the value for Shield to be an explicit nil
func (o *PoolAdditional) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *PoolAdditional) UnsetShield() {
	o.Shield.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetRequestConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *PoolAdditional) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *PoolAdditional) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}

// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *PoolAdditional) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *PoolAdditional) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetTlsCiphers returns the TlsCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetTlsCiphers() string {
	if o == nil || o.TlsCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCiphers.Get()
}

// GetTlsCiphersOk returns a tuple with the TlsCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetTlsCiphersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCiphers.Get(), o.TlsCiphers.IsSet()
}

// HasTlsCiphers returns a boolean if a field has been set.
func (o *PoolAdditional) HasTlsCiphers() bool {
	if o != nil && o.TlsCiphers.IsSet() {
		return true
	}

	return false
}

// SetTlsCiphers gets a reference to the given NullableString and assigns it to the TlsCiphers field.
func (o *PoolAdditional) SetTlsCiphers(v string) {
	o.TlsCiphers.Set(&v)
}

// SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil
func (o *PoolAdditional) SetTlsCiphersNil() {
	o.TlsCiphers.Set(nil)
}

// UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
func (o *PoolAdditional) UnsetTlsCiphers() {
	o.TlsCiphers.Unset()
}

// GetTlsSniHostname returns the TlsSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetTlsSniHostname() string {
	if o == nil || o.TlsSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsSniHostname.Get()
}

// GetTlsSniHostnameOk returns a tuple with the TlsSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetTlsSniHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsSniHostname.Get(), o.TlsSniHostname.IsSet()
}

// HasTlsSniHostname returns a boolean if a field has been set.
func (o *PoolAdditional) HasTlsSniHostname() bool {
	if o != nil && o.TlsSniHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsSniHostname gets a reference to the given NullableString and assigns it to the TlsSniHostname field.
func (o *PoolAdditional) SetTlsSniHostname(v string) {
	o.TlsSniHostname.Set(&v)
}

// SetTlsSniHostnameNil sets the value for TlsSniHostname to be an explicit nil
func (o *PoolAdditional) SetTlsSniHostnameNil() {
	o.TlsSniHostname.Set(nil)
}

// UnsetTlsSniHostname ensures that no value is present for TlsSniHostname, not even an explicit nil
func (o *PoolAdditional) UnsetTlsSniHostname() {
	o.TlsSniHostname.Unset()
}

// GetMinTlsVersion returns the MinTlsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetMinTlsVersion() int32 {
	if o == nil || o.MinTlsVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinTlsVersion.Get()
}

// GetMinTlsVersionOk returns a tuple with the MinTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetMinTlsVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTlsVersion.Get(), o.MinTlsVersion.IsSet()
}

// HasMinTlsVersion returns a boolean if a field has been set.
func (o *PoolAdditional) HasMinTlsVersion() bool {
	if o != nil && o.MinTlsVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTlsVersion gets a reference to the given NullableInt32 and assigns it to the MinTlsVersion field.
func (o *PoolAdditional) SetMinTlsVersion(v int32) {
	o.MinTlsVersion.Set(&v)
}

// SetMinTlsVersionNil sets the value for MinTlsVersion to be an explicit nil
func (o *PoolAdditional) SetMinTlsVersionNil() {
	o.MinTlsVersion.Set(nil)
}

// UnsetMinTlsVersion ensures that no value is present for MinTlsVersion, not even an explicit nil
func (o *PoolAdditional) UnsetMinTlsVersion() {
	o.MinTlsVersion.Unset()
}

// GetMaxTlsVersion returns the MaxTlsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetMaxTlsVersion() int32 {
	if o == nil || o.MaxTlsVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxTlsVersion.Get()
}

// GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetMaxTlsVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTlsVersion.Get(), o.MaxTlsVersion.IsSet()
}

// HasMaxTlsVersion returns a boolean if a field has been set.
func (o *PoolAdditional) HasMaxTlsVersion() bool {
	if o != nil && o.MaxTlsVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTlsVersion gets a reference to the given NullableInt32 and assigns it to the MaxTlsVersion field.
func (o *PoolAdditional) SetMaxTlsVersion(v int32) {
	o.MaxTlsVersion.Set(&v)
}

// SetMaxTlsVersionNil sets the value for MaxTlsVersion to be an explicit nil
func (o *PoolAdditional) SetMaxTlsVersionNil() {
	o.MaxTlsVersion.Set(nil)
}

// UnsetMaxTlsVersion ensures that no value is present for MaxTlsVersion, not even an explicit nil
func (o *PoolAdditional) UnsetMaxTlsVersion() {
	o.MaxTlsVersion.Unset()
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetHealthcheck() string {
	if o == nil || o.Healthcheck.Get() == nil {
		var ret string
		return ret
	}
	return *o.Healthcheck.Get()
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetHealthcheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Healthcheck.Get(), o.Healthcheck.IsSet()
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *PoolAdditional) HasHealthcheck() bool {
	if o != nil && o.Healthcheck.IsSet() {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given NullableString and assigns it to the Healthcheck field.
func (o *PoolAdditional) SetHealthcheck(v string) {
	o.Healthcheck.Set(&v)
}

// SetHealthcheckNil sets the value for Healthcheck to be an explicit nil
func (o *PoolAdditional) SetHealthcheckNil() {
	o.Healthcheck.Set(nil)
}

// UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
func (o *PoolAdditional) UnsetHealthcheck() {
	o.Healthcheck.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *PoolAdditional) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *PoolAdditional) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *PoolAdditional) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *PoolAdditional) UnsetComment() {
	o.Comment.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PoolAdditional) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAdditional) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PoolAdditional) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PoolAdditional) SetType(v string) {
	o.Type = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetOverrideHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *PoolAdditional) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *PoolAdditional) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}

// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *PoolAdditional) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *PoolAdditional) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PoolAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PoolAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAdditional := _PoolAdditional{}

	if err = json.Unmarshal(bytes, &varPoolAdditional); err == nil {
		*o = PoolAdditional(varPoolAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePoolAdditional is a helper abstraction for handling nullable pooladditional types.
type NullablePoolAdditional struct {
	value *PoolAdditional
	isSet bool
}

// Get returns the value.
func (v NullablePoolAdditional) Get() *PoolAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullablePoolAdditional) Set(val *PoolAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePoolAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePoolAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePoolAdditional returns a pointer to a new instance of NullablePoolAdditional.
func NewNullablePoolAdditional(val *PoolAdditional) *NullablePoolAdditional {
	return &NullablePoolAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePoolAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePoolAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
