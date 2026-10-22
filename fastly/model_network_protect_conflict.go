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

// NetworkProtectConflict Conflict returned when resource state prevents the request.
type NetworkProtectConflict struct {
	// Short summary of the conflict.
	Title string `json:"title"`
	// HTTP status code for this error.
	Status int32 `json:"status"`
	// Explanation of the conflicting resource state.
	Detail string `json:"detail"`
	// Optional conflict error details.
	Errors               []NetworkProtectConflictError `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectConflict NetworkProtectConflict

// NewNetworkProtectConflict instantiates a new NetworkProtectConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectConflict(title string, status int32, detail string) *NetworkProtectConflict {
	this := NetworkProtectConflict{}
	this.Title = title
	this.Status = status
	this.Detail = detail
	return &this
}

// NewNetworkProtectConflictWithDefaults instantiates a new NetworkProtectConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectConflictWithDefaults() *NetworkProtectConflict {
	this := NetworkProtectConflict{}
	return &this
}

// GetTitle returns the Title field value
func (o *NetworkProtectConflict) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflict) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NetworkProtectConflict) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *NetworkProtectConflict) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflict) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NetworkProtectConflict) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *NetworkProtectConflict) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflict) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *NetworkProtectConflict) SetDetail(v string) {
	o.Detail = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *NetworkProtectConflict) GetErrors() []NetworkProtectConflictError {
	if o == nil || o.Errors == nil {
		var ret []NetworkProtectConflictError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectConflict) GetErrorsOk() ([]NetworkProtectConflictError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *NetworkProtectConflict) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []NetworkProtectConflictError and assigns it to the Errors field.
func (o *NetworkProtectConflict) SetErrors(v []NetworkProtectConflictError) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectConflict) MarshalJSON() ([]byte, error) {
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
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectConflict) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectConflict := _NetworkProtectConflict{}

	if err = json.Unmarshal(bytes, &varNetworkProtectConflict); err == nil {
		*o = NetworkProtectConflict(varNetworkProtectConflict)
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

// NullableNetworkProtectConflict is a helper abstraction for handling nullable networkprotectconflict types.
type NullableNetworkProtectConflict struct {
	value *NetworkProtectConflict
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectConflict) Get() *NetworkProtectConflict {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectConflict) Set(val *NetworkProtectConflict) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectConflict) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectConflict) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectConflict returns a pointer to a new instance of NullableNetworkProtectConflict.
func NewNullableNetworkProtectConflict(val *NetworkProtectConflict) *NullableNetworkProtectConflict {
	return &NullableNetworkProtectConflict{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
