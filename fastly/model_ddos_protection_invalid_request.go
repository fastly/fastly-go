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

// DdosProtectionInvalidRequest struct for DdosProtectionInvalidRequest
type DdosProtectionInvalidRequest struct {
	Title                string                      `json:"title"`
	Status               int32                       `json:"status"`
	Detail               string                      `json:"detail"`
	Errors               []DdosProtectionErrorErrors `json:"errors"`
	AdditionalProperties map[string]any
}

type _DdosProtectionInvalidRequest DdosProtectionInvalidRequest

// NewDdosProtectionInvalidRequest instantiates a new DdosProtectionInvalidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionInvalidRequest(title string, status int32, detail string, errors []DdosProtectionErrorErrors) *DdosProtectionInvalidRequest {
	this := DdosProtectionInvalidRequest{}
	this.Title = title
	this.Status = status
	this.Detail = detail
	this.Errors = errors
	return &this
}

// NewDdosProtectionInvalidRequestWithDefaults instantiates a new DdosProtectionInvalidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionInvalidRequestWithDefaults() *DdosProtectionInvalidRequest {
	this := DdosProtectionInvalidRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *DdosProtectionInvalidRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DdosProtectionInvalidRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DdosProtectionInvalidRequest) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *DdosProtectionInvalidRequest) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DdosProtectionInvalidRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DdosProtectionInvalidRequest) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *DdosProtectionInvalidRequest) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *DdosProtectionInvalidRequest) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *DdosProtectionInvalidRequest) SetDetail(v string) {
	o.Detail = v
}

// GetErrors returns the Errors field value
func (o *DdosProtectionInvalidRequest) GetErrors() []DdosProtectionErrorErrors {
	if o == nil {
		var ret []DdosProtectionErrorErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *DdosProtectionInvalidRequest) GetErrorsOk() ([]DdosProtectionErrorErrors, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *DdosProtectionInvalidRequest) SetErrors(v []DdosProtectionErrorErrors) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionInvalidRequest) MarshalJSON() ([]byte, error) {
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
func (o *DdosProtectionInvalidRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionInvalidRequest := _DdosProtectionInvalidRequest{}

	if err = json.Unmarshal(bytes, &varDdosProtectionInvalidRequest); err == nil {
		*o = DdosProtectionInvalidRequest(varDdosProtectionInvalidRequest)
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

// NullableDdosProtectionInvalidRequest is a helper abstraction for handling nullable ddosprotectioninvalidrequest types.
type NullableDdosProtectionInvalidRequest struct {
	value *DdosProtectionInvalidRequest
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionInvalidRequest) Get() *DdosProtectionInvalidRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionInvalidRequest) Set(val *DdosProtectionInvalidRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionInvalidRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionInvalidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionInvalidRequest returns a pointer to a new instance of NullableDdosProtectionInvalidRequest.
func NewNullableDdosProtectionInvalidRequest(val *DdosProtectionInvalidRequest) *NullableDdosProtectionInvalidRequest {
	return &NullableDdosProtectionInvalidRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionInvalidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionInvalidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
