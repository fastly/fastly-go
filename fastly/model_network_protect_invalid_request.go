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

// NetworkProtectInvalidRequest Validation error returned for an invalid request.
type NetworkProtectInvalidRequest struct {
	// Short summary of the validation problem.
	Title string `json:"title"`
	// HTTP status code for this error.
	Status int32 `json:"status"`
	// Corrective explanation for the invalid request.
	Detail string `json:"detail"`
	// Field-level validation errors, when available.
	Errors               []NetworkProtectErrorDetail `json:"errors"`
	AdditionalProperties map[string]any
}

type _NetworkProtectInvalidRequest NetworkProtectInvalidRequest

// NewNetworkProtectInvalidRequest instantiates a new NetworkProtectInvalidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectInvalidRequest(title string, status int32, detail string, errors []NetworkProtectErrorDetail) *NetworkProtectInvalidRequest {
	this := NetworkProtectInvalidRequest{}
	this.Title = title
	this.Status = status
	this.Detail = detail
	this.Errors = errors
	return &this
}

// NewNetworkProtectInvalidRequestWithDefaults instantiates a new NetworkProtectInvalidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectInvalidRequestWithDefaults() *NetworkProtectInvalidRequest {
	this := NetworkProtectInvalidRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *NetworkProtectInvalidRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectInvalidRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NetworkProtectInvalidRequest) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *NetworkProtectInvalidRequest) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectInvalidRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NetworkProtectInvalidRequest) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *NetworkProtectInvalidRequest) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectInvalidRequest) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *NetworkProtectInvalidRequest) SetDetail(v string) {
	o.Detail = v
}

// GetErrors returns the Errors field value
func (o *NetworkProtectInvalidRequest) GetErrors() []NetworkProtectErrorDetail {
	if o == nil {
		var ret []NetworkProtectErrorDetail
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectInvalidRequest) GetErrorsOk() ([]NetworkProtectErrorDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *NetworkProtectInvalidRequest) SetErrors(v []NetworkProtectErrorDetail) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectInvalidRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if true {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectInvalidRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectInvalidRequest := _NetworkProtectInvalidRequest{}

	if err = json.Unmarshal(bytes, &varNetworkProtectInvalidRequest); err == nil {
		*o = NetworkProtectInvalidRequest(varNetworkProtectInvalidRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectInvalidRequest is a helper abstraction for handling nullable networkprotectinvalidrequest types.
type NullableNetworkProtectInvalidRequest struct {
	value *NetworkProtectInvalidRequest
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectInvalidRequest) Get() *NetworkProtectInvalidRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectInvalidRequest) Set(val *NetworkProtectInvalidRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectInvalidRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectInvalidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectInvalidRequest returns a pointer to a new instance of NullableNetworkProtectInvalidRequest.
func NewNullableNetworkProtectInvalidRequest(val *NetworkProtectInvalidRequest) *NullableNetworkProtectInvalidRequest {
	return &NullableNetworkProtectInvalidRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectInvalidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectInvalidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
