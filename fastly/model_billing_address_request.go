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

// BillingAddressRequest struct for BillingAddressRequest
type BillingAddressRequest struct {
	// When set to true, the address will be saved without verification
	SkipVerification *bool `json:"skip_verification,omitempty"`
	Data *BillingAddressRequestData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressRequest BillingAddressRequest

// NewBillingAddressRequest instantiates a new BillingAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressRequest() *BillingAddressRequest {
	this := BillingAddressRequest{}
	return &this
}

// NewBillingAddressRequestWithDefaults instantiates a new BillingAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressRequestWithDefaults() *BillingAddressRequest {
	this := BillingAddressRequest{}
	return &this
}

// GetSkipVerification returns the SkipVerification field value if set, zero value otherwise.
func (o *BillingAddressRequest) GetSkipVerification() bool {
	if o == nil || o.SkipVerification == nil {
		var ret bool
		return ret
	}
	return *o.SkipVerification
}

// GetSkipVerificationOk returns a tuple with the SkipVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressRequest) GetSkipVerificationOk() (*bool, bool) {
	if o == nil || o.SkipVerification == nil {
		return nil, false
	}
	return o.SkipVerification, true
}

// HasSkipVerification returns a boolean if a field has been set.
func (o *BillingAddressRequest) HasSkipVerification() bool {
	if o != nil && o.SkipVerification != nil {
		return true
	}

	return false
}

// SetSkipVerification gets a reference to the given bool and assigns it to the SkipVerification field.
func (o *BillingAddressRequest) SetSkipVerification(v bool) {
	o.SkipVerification = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BillingAddressRequest) GetData() BillingAddressRequestData {
	if o == nil || o.Data == nil {
		var ret BillingAddressRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressRequest) GetDataOk() (*BillingAddressRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingAddressRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BillingAddressRequestData and assigns it to the Data field.
func (o *BillingAddressRequest) SetData(v BillingAddressRequestData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressRequest) MarshalJSON() ([]byte, error) {
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
func (o *BillingAddressRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressRequest := _BillingAddressRequest{}

	if err = json.Unmarshal(bytes, &varBillingAddressRequest); err == nil {
		*o = BillingAddressRequest(varBillingAddressRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "skip_verification")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressRequest is a helper abstraction for handling nullable billingaddressrequest types. 
type NullableBillingAddressRequest struct {
	value *BillingAddressRequest
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressRequest) Get() *BillingAddressRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressRequest) Set(val *BillingAddressRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressRequest returns a pointer to a new instance of NullableBillingAddressRequest.
func NewNullableBillingAddressRequest(val *BillingAddressRequest) *NullableBillingAddressRequest {
	return &NullableBillingAddressRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
