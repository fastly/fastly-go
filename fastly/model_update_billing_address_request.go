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

// UpdateBillingAddressRequest struct for UpdateBillingAddressRequest
type UpdateBillingAddressRequest struct {
	// When set to true, the address will be saved without verification
	SkipVerification *bool `json:"skip_verification,omitempty"`
	Data *UpdateBillingAddressRequestData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _UpdateBillingAddressRequest UpdateBillingAddressRequest

// NewUpdateBillingAddressRequest instantiates a new UpdateBillingAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBillingAddressRequest() *UpdateBillingAddressRequest {
	this := UpdateBillingAddressRequest{}
	return &this
}

// NewUpdateBillingAddressRequestWithDefaults instantiates a new UpdateBillingAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBillingAddressRequestWithDefaults() *UpdateBillingAddressRequest {
	this := UpdateBillingAddressRequest{}
	return &this
}

// GetSkipVerification returns the SkipVerification field value if set, zero value otherwise.
func (o *UpdateBillingAddressRequest) GetSkipVerification() bool {
	if o == nil || o.SkipVerification == nil {
		var ret bool
		return ret
	}
	return *o.SkipVerification
}

// GetSkipVerificationOk returns a tuple with the SkipVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBillingAddressRequest) GetSkipVerificationOk() (*bool, bool) {
	if o == nil || o.SkipVerification == nil {
		return nil, false
	}
	return o.SkipVerification, true
}

// HasSkipVerification returns a boolean if a field has been set.
func (o *UpdateBillingAddressRequest) HasSkipVerification() bool {
	if o != nil && o.SkipVerification != nil {
		return true
	}

	return false
}

// SetSkipVerification gets a reference to the given bool and assigns it to the SkipVerification field.
func (o *UpdateBillingAddressRequest) SetSkipVerification(v bool) {
	o.SkipVerification = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateBillingAddressRequest) GetData() UpdateBillingAddressRequestData {
	if o == nil || o.Data == nil {
		var ret UpdateBillingAddressRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBillingAddressRequest) GetDataOk() (*UpdateBillingAddressRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateBillingAddressRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateBillingAddressRequestData and assigns it to the Data field.
func (o *UpdateBillingAddressRequest) SetData(v UpdateBillingAddressRequestData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o UpdateBillingAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.SkipVerification != nil {
		toSerialize["skip_verification"] = o.SkipVerification
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *UpdateBillingAddressRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateBillingAddressRequest := _UpdateBillingAddressRequest{}

	if err = json.Unmarshal(bytes, &varUpdateBillingAddressRequest); err == nil {
		*o = UpdateBillingAddressRequest(varUpdateBillingAddressRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "skip_verification")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableUpdateBillingAddressRequest is a helper abstraction for handling nullable updatebillingaddressrequest types. 
type NullableUpdateBillingAddressRequest struct {
	value *UpdateBillingAddressRequest
	isSet bool
}

// Get returns the value.
func (v NullableUpdateBillingAddressRequest) Get() *UpdateBillingAddressRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableUpdateBillingAddressRequest) Set(val *UpdateBillingAddressRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableUpdateBillingAddressRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableUpdateBillingAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateBillingAddressRequest returns a pointer to a new instance of NullableUpdateBillingAddressRequest.
func NewNullableUpdateBillingAddressRequest(val *UpdateBillingAddressRequest) *NullableUpdateBillingAddressRequest {
	return &NullableUpdateBillingAddressRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableUpdateBillingAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableUpdateBillingAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
