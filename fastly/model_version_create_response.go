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
)

// VersionCreateResponse struct for VersionCreateResponse
type VersionCreateResponse struct {
	Number *int32 `json:"number,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _VersionCreateResponse VersionCreateResponse

// NewVersionCreateResponse instantiates a new VersionCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionCreateResponse() *VersionCreateResponse {
	this := VersionCreateResponse{}
	return &this
}

// NewVersionCreateResponseWithDefaults instantiates a new VersionCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionCreateResponseWithDefaults() *VersionCreateResponse {
	this := VersionCreateResponse{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *VersionCreateResponse) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionCreateResponse) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *VersionCreateResponse) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *VersionCreateResponse) SetNumber(v int32) {
	o.Number = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *VersionCreateResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionCreateResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *VersionCreateResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *VersionCreateResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VersionCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *VersionCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varVersionCreateResponse := _VersionCreateResponse{}

	if err = json.Unmarshal(bytes, &varVersionCreateResponse); err == nil {
		*o = VersionCreateResponse(varVersionCreateResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "number")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVersionCreateResponse is a helper abstraction for handling nullable versioncreateresponse types. 
type NullableVersionCreateResponse struct {
	value *VersionCreateResponse
	isSet bool
}

// Get returns the value.
func (v NullableVersionCreateResponse) Get() *VersionCreateResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableVersionCreateResponse) Set(val *VersionCreateResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVersionCreateResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVersionCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVersionCreateResponse returns a pointer to a new instance of NullableVersionCreateResponse.
func NewNullableVersionCreateResponse(val *VersionCreateResponse) *NullableVersionCreateResponse {
	return &NullableVersionCreateResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVersionCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableVersionCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
