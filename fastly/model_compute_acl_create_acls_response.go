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

// ComputeAclCreateAclsResponse struct for ComputeAclCreateAclsResponse
type ComputeAclCreateAclsResponse struct {
	// Human readable name of store
	Name *string `json:"name,omitempty"`
	// An example identifier (UUID).
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeAclCreateAclsResponse ComputeAclCreateAclsResponse

// NewComputeAclCreateAclsResponse instantiates a new ComputeAclCreateAclsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAclCreateAclsResponse() *ComputeAclCreateAclsResponse {
	this := ComputeAclCreateAclsResponse{}
	return &this
}

// NewComputeAclCreateAclsResponseWithDefaults instantiates a new ComputeAclCreateAclsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAclCreateAclsResponseWithDefaults() *ComputeAclCreateAclsResponse {
	this := ComputeAclCreateAclsResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeAclCreateAclsResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclCreateAclsResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeAclCreateAclsResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeAclCreateAclsResponse) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputeAclCreateAclsResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclCreateAclsResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputeAclCreateAclsResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputeAclCreateAclsResponse) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeAclCreateAclsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeAclCreateAclsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varComputeAclCreateAclsResponse := _ComputeAclCreateAclsResponse{}

	if err = json.Unmarshal(bytes, &varComputeAclCreateAclsResponse); err == nil {
		*o = ComputeAclCreateAclsResponse(varComputeAclCreateAclsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeAclCreateAclsResponse is a helper abstraction for handling nullable computeaclcreateaclsresponse types.
type NullableComputeAclCreateAclsResponse struct {
	value *ComputeAclCreateAclsResponse
	isSet bool
}

// Get returns the value.
func (v NullableComputeAclCreateAclsResponse) Get() *ComputeAclCreateAclsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeAclCreateAclsResponse) Set(val *ComputeAclCreateAclsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeAclCreateAclsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeAclCreateAclsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeAclCreateAclsResponse returns a pointer to a new instance of NullableComputeAclCreateAclsResponse.
func NewNullableComputeAclCreateAclsResponse(val *ComputeAclCreateAclsResponse) *NullableComputeAclCreateAclsResponse {
	return &NullableComputeAclCreateAclsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeAclCreateAclsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeAclCreateAclsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
