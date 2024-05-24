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

// AutomationTokenCreateRequest struct for AutomationTokenCreateRequest
type AutomationTokenCreateRequest struct {
	Attributes *AutomationTokenCreateRequestAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenCreateRequest AutomationTokenCreateRequest

// NewAutomationTokenCreateRequest instantiates a new AutomationTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenCreateRequest() *AutomationTokenCreateRequest {
	this := AutomationTokenCreateRequest{}
	return &this
}

// NewAutomationTokenCreateRequestWithDefaults instantiates a new AutomationTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenCreateRequestWithDefaults() *AutomationTokenCreateRequest {
	this := AutomationTokenCreateRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetAttributes() AutomationTokenCreateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret AutomationTokenCreateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetAttributesOk() (*AutomationTokenCreateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AutomationTokenCreateRequestAttributes and assigns it to the Attributes field.
func (o *AutomationTokenCreateRequest) SetAttributes(v AutomationTokenCreateRequestAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *AutomationTokenCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenCreateRequest := _AutomationTokenCreateRequest{}

	if err = json.Unmarshal(bytes, &varAutomationTokenCreateRequest); err == nil {
		*o = AutomationTokenCreateRequest(varAutomationTokenCreateRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenCreateRequest is a helper abstraction for handling nullable automationtokencreaterequest types. 
type NullableAutomationTokenCreateRequest struct {
	value *AutomationTokenCreateRequest
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenCreateRequest) Get() *AutomationTokenCreateRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenCreateRequest) Set(val *AutomationTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenCreateRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenCreateRequest returns a pointer to a new instance of NullableAutomationTokenCreateRequest.
func NewNullableAutomationTokenCreateRequest(val *AutomationTokenCreateRequest) *NullableAutomationTokenCreateRequest {
	return &NullableAutomationTokenCreateRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableAutomationTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
