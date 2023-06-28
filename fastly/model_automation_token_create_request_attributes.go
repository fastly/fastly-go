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

// AutomationTokenCreateRequestAttributes struct for AutomationTokenCreateRequestAttributes
type AutomationTokenCreateRequestAttributes struct {
	// name of the token
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
	// List of service ids to limit the token
	Services []string `json:"services,omitempty"`
	Scope *string `json:"scope,omitempty"`
	// A UTC time-stamp of when the token will expire.
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// Indicates whether TLS access is enabled for the token.
	TLSAccess *bool `json:"tls_access,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenCreateRequestAttributes AutomationTokenCreateRequestAttributes

// NewAutomationTokenCreateRequestAttributes instantiates a new AutomationTokenCreateRequestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenCreateRequestAttributes() *AutomationTokenCreateRequestAttributes {
	this := AutomationTokenCreateRequestAttributes{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// NewAutomationTokenCreateRequestAttributesWithDefaults instantiates a new AutomationTokenCreateRequestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenCreateRequestAttributesWithDefaults() *AutomationTokenCreateRequestAttributes {
	this := AutomationTokenCreateRequestAttributes{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutomationTokenCreateRequestAttributes) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequestAttributes) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequestAttributes) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AutomationTokenCreateRequestAttributes) SetRole(v string) {
	o.Role = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequestAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequestAttributes) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *AutomationTokenCreateRequestAttributes) SetServices(v []string) {
	o.Services = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequestAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequestAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AutomationTokenCreateRequestAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationTokenCreateRequestAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenCreateRequestAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *AutomationTokenCreateRequestAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *AutomationTokenCreateRequestAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *AutomationTokenCreateRequestAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetTLSAccess returns the TLSAccess field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequestAttributes) GetTLSAccess() bool {
	if o == nil || o.TLSAccess == nil {
		var ret bool
		return ret
	}
	return *o.TLSAccess
}

// GetTLSAccessOk returns a tuple with the TLSAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequestAttributes) GetTLSAccessOk() (*bool, bool) {
	if o == nil || o.TLSAccess == nil {
		return nil, false
	}
	return o.TLSAccess, true
}

// HasTLSAccess returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequestAttributes) HasTLSAccess() bool {
	if o != nil && o.TLSAccess != nil {
		return true
	}

	return false
}

// SetTLSAccess gets a reference to the given bool and assigns it to the TLSAccess field.
func (o *AutomationTokenCreateRequestAttributes) SetTLSAccess(v bool) {
	o.TLSAccess = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.TLSAccess != nil {
		toSerialize["tls_access"] = o.TLSAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *AutomationTokenCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenCreateRequestAttributes := _AutomationTokenCreateRequestAttributes{}

	if err = json.Unmarshal(bytes, &varAutomationTokenCreateRequestAttributes); err == nil {
		*o = AutomationTokenCreateRequestAttributes(varAutomationTokenCreateRequestAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "services")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "tls_access")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenCreateRequestAttributes is a helper abstraction for handling nullable automationtokencreaterequestattributes types. 
type NullableAutomationTokenCreateRequestAttributes struct {
	value *AutomationTokenCreateRequestAttributes
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenCreateRequestAttributes) Get() *AutomationTokenCreateRequestAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenCreateRequestAttributes) Set(val *AutomationTokenCreateRequestAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenCreateRequestAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenCreateRequestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenCreateRequestAttributes returns a pointer to a new instance of NullableAutomationTokenCreateRequestAttributes.
func NewNullableAutomationTokenCreateRequestAttributes(val *AutomationTokenCreateRequestAttributes) *NullableAutomationTokenCreateRequestAttributes {
	return &NullableAutomationTokenCreateRequestAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableAutomationTokenCreateRequestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
