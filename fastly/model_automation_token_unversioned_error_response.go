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

// AutomationTokenUnversionedErrorResponse Error response format used by the workspaces endpoint.
type AutomationTokenUnversionedErrorResponse struct {
	Errors               []AutomationTokenUnversionedError `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenUnversionedErrorResponse AutomationTokenUnversionedErrorResponse

// NewAutomationTokenUnversionedErrorResponse instantiates a new AutomationTokenUnversionedErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenUnversionedErrorResponse() *AutomationTokenUnversionedErrorResponse {
	this := AutomationTokenUnversionedErrorResponse{}
	return &this
}

// NewAutomationTokenUnversionedErrorResponseWithDefaults instantiates a new AutomationTokenUnversionedErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenUnversionedErrorResponseWithDefaults() *AutomationTokenUnversionedErrorResponse {
	this := AutomationTokenUnversionedErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AutomationTokenUnversionedErrorResponse) GetErrors() []AutomationTokenUnversionedError {
	if o == nil || o.Errors == nil {
		var ret []AutomationTokenUnversionedError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenUnversionedErrorResponse) GetErrorsOk() ([]AutomationTokenUnversionedError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AutomationTokenUnversionedErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AutomationTokenUnversionedError and assigns it to the Errors field.
func (o *AutomationTokenUnversionedErrorResponse) SetErrors(v []AutomationTokenUnversionedError) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenUnversionedErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *AutomationTokenUnversionedErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenUnversionedErrorResponse := _AutomationTokenUnversionedErrorResponse{}

	if err = json.Unmarshal(bytes, &varAutomationTokenUnversionedErrorResponse); err == nil {
		*o = AutomationTokenUnversionedErrorResponse(varAutomationTokenUnversionedErrorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenUnversionedErrorResponse is a helper abstraction for handling nullable automationtokenunversionederrorresponse types.
type NullableAutomationTokenUnversionedErrorResponse struct {
	value *AutomationTokenUnversionedErrorResponse
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenUnversionedErrorResponse) Get() *AutomationTokenUnversionedErrorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenUnversionedErrorResponse) Set(val *AutomationTokenUnversionedErrorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenUnversionedErrorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenUnversionedErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenUnversionedErrorResponse returns a pointer to a new instance of NullableAutomationTokenUnversionedErrorResponse.
func NewNullableAutomationTokenUnversionedErrorResponse(val *AutomationTokenUnversionedErrorResponse) *NullableAutomationTokenUnversionedErrorResponse {
	return &NullableAutomationTokenUnversionedErrorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenUnversionedErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenUnversionedErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
