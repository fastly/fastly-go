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

// DiscoveredOperationBase struct for DiscoveredOperationBase
type DiscoveredOperationBase struct {
	// The HTTP method for the operation.
	Method *string `json:"method,omitempty"`
	// The domain for the operation.
	Domain *string `json:"domain,omitempty"`
	// The path for the operation, which may include path parameters.
	Path                 *string `json:"path,omitempty"`
	AdditionalProperties map[string]any
}

type _DiscoveredOperationBase DiscoveredOperationBase

// NewDiscoveredOperationBase instantiates a new DiscoveredOperationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredOperationBase() *DiscoveredOperationBase {
	this := DiscoveredOperationBase{}
	return &this
}

// NewDiscoveredOperationBaseWithDefaults instantiates a new DiscoveredOperationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredOperationBaseWithDefaults() *DiscoveredOperationBase {
	this := DiscoveredOperationBase{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *DiscoveredOperationBase) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationBase) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *DiscoveredOperationBase) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *DiscoveredOperationBase) SetMethod(v string) {
	o.Method = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DiscoveredOperationBase) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationBase) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DiscoveredOperationBase) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DiscoveredOperationBase) SetDomain(v string) {
	o.Domain = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DiscoveredOperationBase) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationBase) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DiscoveredOperationBase) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DiscoveredOperationBase) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DiscoveredOperationBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DiscoveredOperationBase) UnmarshalJSON(bytes []byte) (err error) {
	varDiscoveredOperationBase := _DiscoveredOperationBase{}

	if err = json.Unmarshal(bytes, &varDiscoveredOperationBase); err == nil {
		*o = DiscoveredOperationBase(varDiscoveredOperationBase)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDiscoveredOperationBase is a helper abstraction for handling nullable discoveredoperationbase types.
type NullableDiscoveredOperationBase struct {
	value *DiscoveredOperationBase
	isSet bool
}

// Get returns the value.
func (v NullableDiscoveredOperationBase) Get() *DiscoveredOperationBase {
	return v.value
}

// Set modifies the value.
func (v *NullableDiscoveredOperationBase) Set(val *DiscoveredOperationBase) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDiscoveredOperationBase) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDiscoveredOperationBase) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDiscoveredOperationBase returns a pointer to a new instance of NullableDiscoveredOperationBase.
func NewNullableDiscoveredOperationBase(val *DiscoveredOperationBase) *NullableDiscoveredOperationBase {
	return &NullableDiscoveredOperationBase{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDiscoveredOperationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDiscoveredOperationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
