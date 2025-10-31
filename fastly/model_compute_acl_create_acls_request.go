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

// ComputeAclCreateAclsRequest struct for ComputeAclCreateAclsRequest
type ComputeAclCreateAclsRequest struct {
	// Human readable name of store
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeAclCreateAclsRequest ComputeAclCreateAclsRequest

// NewComputeAclCreateAclsRequest instantiates a new ComputeAclCreateAclsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAclCreateAclsRequest() *ComputeAclCreateAclsRequest {
	this := ComputeAclCreateAclsRequest{}
	return &this
}

// NewComputeAclCreateAclsRequestWithDefaults instantiates a new ComputeAclCreateAclsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAclCreateAclsRequestWithDefaults() *ComputeAclCreateAclsRequest {
	this := ComputeAclCreateAclsRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeAclCreateAclsRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclCreateAclsRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeAclCreateAclsRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeAclCreateAclsRequest) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeAclCreateAclsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeAclCreateAclsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varComputeAclCreateAclsRequest := _ComputeAclCreateAclsRequest{}

	if err = json.Unmarshal(bytes, &varComputeAclCreateAclsRequest); err == nil {
		*o = ComputeAclCreateAclsRequest(varComputeAclCreateAclsRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeAclCreateAclsRequest is a helper abstraction for handling nullable computeaclcreateaclsrequest types.
type NullableComputeAclCreateAclsRequest struct {
	value *ComputeAclCreateAclsRequest
	isSet bool
}

// Get returns the value.
func (v NullableComputeAclCreateAclsRequest) Get() *ComputeAclCreateAclsRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeAclCreateAclsRequest) Set(val *ComputeAclCreateAclsRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeAclCreateAclsRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeAclCreateAclsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeAclCreateAclsRequest returns a pointer to a new instance of NullableComputeAclCreateAclsRequest.
func NewNullableComputeAclCreateAclsRequest(val *ComputeAclCreateAclsRequest) *NullableComputeAclCreateAclsRequest {
	return &NullableComputeAclCreateAclsRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeAclCreateAclsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeAclCreateAclsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
