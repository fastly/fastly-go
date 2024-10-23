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

// ComputeACLCreateAclsResponse struct for ComputeACLCreateAclsResponse
type ComputeACLCreateAclsResponse struct {
	// Human readable name of store
	Name *string `json:"name,omitempty"`
	// An example identifier (UUID).
	ID                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeACLCreateAclsResponse ComputeACLCreateAclsResponse

// NewComputeACLCreateAclsResponse instantiates a new ComputeACLCreateAclsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeACLCreateAclsResponse() *ComputeACLCreateAclsResponse {
	this := ComputeACLCreateAclsResponse{}
	return &this
}

// NewComputeACLCreateAclsResponseWithDefaults instantiates a new ComputeACLCreateAclsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeACLCreateAclsResponseWithDefaults() *ComputeACLCreateAclsResponse {
	this := ComputeACLCreateAclsResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeACLCreateAclsResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLCreateAclsResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeACLCreateAclsResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeACLCreateAclsResponse) SetName(v string) {
	o.Name = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ComputeACLCreateAclsResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLCreateAclsResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ComputeACLCreateAclsResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ComputeACLCreateAclsResponse) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeACLCreateAclsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeACLCreateAclsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varComputeACLCreateAclsResponse := _ComputeACLCreateAclsResponse{}

	if err = json.Unmarshal(bytes, &varComputeACLCreateAclsResponse); err == nil {
		*o = ComputeACLCreateAclsResponse(varComputeACLCreateAclsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeACLCreateAclsResponse is a helper abstraction for handling nullable computeaclcreateaclsresponse types.
type NullableComputeACLCreateAclsResponse struct {
	value *ComputeACLCreateAclsResponse
	isSet bool
}

// Get returns the value.
func (v NullableComputeACLCreateAclsResponse) Get() *ComputeACLCreateAclsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeACLCreateAclsResponse) Set(val *ComputeACLCreateAclsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeACLCreateAclsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeACLCreateAclsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeACLCreateAclsResponse returns a pointer to a new instance of NullableComputeACLCreateAclsResponse.
func NewNullableComputeACLCreateAclsResponse(val *ComputeACLCreateAclsResponse) *NullableComputeACLCreateAclsResponse {
	return &NullableComputeACLCreateAclsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeACLCreateAclsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeACLCreateAclsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
