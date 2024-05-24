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

// BillingAddressVerificationErrorResponse struct for BillingAddressVerificationErrorResponse
type BillingAddressVerificationErrorResponse struct {
	Errors []BillingAddressVerificationErrorResponseErrors `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressVerificationErrorResponse BillingAddressVerificationErrorResponse

// NewBillingAddressVerificationErrorResponse instantiates a new BillingAddressVerificationErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressVerificationErrorResponse() *BillingAddressVerificationErrorResponse {
	this := BillingAddressVerificationErrorResponse{}
	return &this
}

// NewBillingAddressVerificationErrorResponseWithDefaults instantiates a new BillingAddressVerificationErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressVerificationErrorResponseWithDefaults() *BillingAddressVerificationErrorResponse {
	this := BillingAddressVerificationErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BillingAddressVerificationErrorResponse) GetErrors() []BillingAddressVerificationErrorResponseErrors {
	if o == nil || o.Errors == nil {
		var ret []BillingAddressVerificationErrorResponseErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponse) GetErrorsOk() ([]BillingAddressVerificationErrorResponseErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BillingAddressVerificationErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BillingAddressVerificationErrorResponseErrors and assigns it to the Errors field.
func (o *BillingAddressVerificationErrorResponse) SetErrors(v []BillingAddressVerificationErrorResponseErrors) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressVerificationErrorResponse) MarshalJSON() ([]byte, error) {
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
func (o *BillingAddressVerificationErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressVerificationErrorResponse := _BillingAddressVerificationErrorResponse{}

	if err = json.Unmarshal(bytes, &varBillingAddressVerificationErrorResponse); err == nil {
		*o = BillingAddressVerificationErrorResponse(varBillingAddressVerificationErrorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressVerificationErrorResponse is a helper abstraction for handling nullable billingaddressverificationerrorresponse types. 
type NullableBillingAddressVerificationErrorResponse struct {
	value *BillingAddressVerificationErrorResponse
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressVerificationErrorResponse) Get() *BillingAddressVerificationErrorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressVerificationErrorResponse) Set(val *BillingAddressVerificationErrorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressVerificationErrorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressVerificationErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressVerificationErrorResponse returns a pointer to a new instance of NullableBillingAddressVerificationErrorResponse.
func NewNullableBillingAddressVerificationErrorResponse(val *BillingAddressVerificationErrorResponse) *NullableBillingAddressVerificationErrorResponse {
	return &NullableBillingAddressVerificationErrorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressVerificationErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingAddressVerificationErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
