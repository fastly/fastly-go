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
	if o == nil  {
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
	if o == nil  {
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

// GetTLSCiphers returns the TLSCiphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetTLSCiphers() string {
	if o == nil || o.TLSCiphers.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCiphers.Get()
}

// GetTLSCiphersOk returns a tuple with the TLSCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetTLSCiphersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCiphers.Get(), o.TLSCiphers.IsSet()
}

// HasTLSCiphers returns a boolean if a field has been set.
func (o *PoolAdditional) HasTLSCiphers() bool {
	if o != nil && o.TLSCiphers.IsSet() {
		return true
	}

	return false
}

// SetTLSCiphers gets a reference to the given NullableString and assigns it to the TLSCiphers field.
func (o *PoolAdditional) SetTLSCiphers(v string) {
	o.TLSCiphers.Set(&v)
}
// SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil
func (o *PoolAdditional) SetTLSCiphersNil() {
	o.TLSCiphers.Set(nil)
}

// UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
func (o *PoolAdditional) UnsetTLSCiphers() {
	o.TLSCiphers.Unset()
}

// GetTLSSniHostname returns the TLSSniHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetTLSSniHostname() string {
	if o == nil || o.TLSSniHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSSniHostname.Get()
}

// GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetTLSSniHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSSniHostname.Get(), o.TLSSniHostname.IsSet()
}

// HasTLSSniHostname returns a boolean if a field has been set.
func (o *PoolAdditional) HasTLSSniHostname() bool {
	if o != nil && o.TLSSniHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSSniHostname gets a reference to the given NullableString and assigns it to the TLSSniHostname field.
func (o *PoolAdditional) SetTLSSniHostname(v string) {
	o.TLSSniHostname.Set(&v)
}
// SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil
func (o *PoolAdditional) SetTLSSniHostnameNil() {
	o.TLSSniHostname.Set(nil)
}

// UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
func (o *PoolAdditional) UnsetTLSSniHostname() {
	o.TLSSniHostname.Unset()
}

// GetMinTLSVersion returns the MinTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetMinTLSVersion() int32 {
	if o == nil || o.MinTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinTLSVersion.Get()
}

// GetMinTLSVersionOk returns a tuple with the MinTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetMinTLSVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinTLSVersion.Get(), o.MinTLSVersion.IsSet()
}

// HasMinTLSVersion returns a boolean if a field has been set.
func (o *PoolAdditional) HasMinTLSVersion() bool {
	if o != nil && o.MinTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMinTLSVersion gets a reference to the given NullableInt32 and assigns it to the MinTLSVersion field.
func (o *PoolAdditional) SetMinTLSVersion(v int32) {
	o.MinTLSVersion.Set(&v)
}
// SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil
func (o *PoolAdditional) SetMinTLSVersionNil() {
	o.MinTLSVersion.Set(nil)
}

// UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
func (o *PoolAdditional) UnsetMinTLSVersion() {
	o.MinTLSVersion.Unset()
}

// GetMaxTLSVersion returns the MaxTLSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolAdditional) GetMaxTLSVersion() int32 {
	if o == nil || o.MaxTLSVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxTLSVersion.Get()
}

// GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolAdditional) GetMaxTLSVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxTLSVersion.Get(), o.MaxTLSVersion.IsSet()
}

// HasMaxTLSVersion returns a boolean if a field has been set.
func (o *PoolAdditional) HasMaxTLSVersion() bool {
	if o != nil && o.MaxTLSVersion.IsSet() {
		return true
	}

	return false
}

// SetMaxTLSVersion gets a reference to the given NullableInt32 and assigns it to the MaxTLSVersion field.
func (o *PoolAdditional) SetMaxTLSVersion(v int32) {
	o.MaxTLSVersion.Set(&v)
}
// SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil
func (o *PoolAdditional) SetMaxTLSVersionNil() {
	o.MaxTLSVersion.Set(nil)
}

// UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
func (o *PoolAdditional) UnsetMaxTLSVersion() {
	o.MaxTLSVersion.Unset()
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
	if o == nil  {
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
	if o == nil  {
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
	if o == nil  {
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
