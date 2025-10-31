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

// AutomationTokenErrorResponse struct for AutomationTokenErrorResponse
type AutomationTokenErrorResponse struct {
	Detail               *string       `json:"detail,omitempty"`
	Errors               []interface{} `json:"errors,omitempty"`
	Status               *int32        `json:"status,omitempty"`
	Title                *string       `json:"title,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenErrorResponse AutomationTokenErrorResponse

// NewAutomationTokenErrorResponse instantiates a new AutomationTokenErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenErrorResponse() *AutomationTokenErrorResponse {
	this := AutomationTokenErrorResponse{}
	return &this
}

// NewAutomationTokenErrorResponseWithDefaults instantiates a new AutomationTokenErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenErrorResponseWithDefaults() *AutomationTokenErrorResponse {
	this := AutomationTokenErrorResponse{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AutomationTokenErrorResponse) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenErrorResponse) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AutomationTokenErrorResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AutomationTokenErrorResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AutomationTokenErrorResponse) GetErrors() []interface{} {
	if o == nil || o.Errors == nil {
		var ret []interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenErrorResponse) GetErrorsOk() ([]interface{}, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AutomationTokenErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []interface{} and assigns it to the Errors field.
func (o *AutomationTokenErrorResponse) SetErrors(v []interface{}) {
	o.Errors = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AutomationTokenErrorResponse) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenErrorResponse) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AutomationTokenErrorResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AutomationTokenErrorResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AutomationTokenErrorResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenErrorResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AutomationTokenErrorResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AutomationTokenErrorResponse) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenErrorResponse := _AutomationTokenErrorResponse{}

	if err = json.Unmarshal(bytes, &varAutomationTokenErrorResponse); err == nil {
		*o = AutomationTokenErrorResponse(varAutomationTokenErrorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "status")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenErrorResponse is a helper abstraction for handling nullable automationtokenerrorresponse types.
type NullableAutomationTokenErrorResponse struct {
	value *AutomationTokenErrorResponse
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenErrorResponse) Get() *AutomationTokenErrorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenErrorResponse) Set(val *AutomationTokenErrorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenErrorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenErrorResponse returns a pointer to a new instance of NullableAutomationTokenErrorResponse.
func NewNullableAutomationTokenErrorResponse(val *AutomationTokenErrorResponse) *NullableAutomationTokenErrorResponse {
	return &NullableAutomationTokenErrorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
