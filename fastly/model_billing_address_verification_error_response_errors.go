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

// BillingAddressVerificationErrorResponseErrors struct for BillingAddressVerificationErrorResponseErrors
type BillingAddressVerificationErrorResponseErrors struct {
	// The error type.
	Type string `json:"type"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Status float32 `json:"status"`
	Candidates []BillingAddressAttributes `json:"candidates,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressVerificationErrorResponseErrors BillingAddressVerificationErrorResponseErrors

// NewBillingAddressVerificationErrorResponseErrors instantiates a new BillingAddressVerificationErrorResponseErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressVerificationErrorResponseErrors(resourceType string, title string, detail string, status float32) *BillingAddressVerificationErrorResponseErrors {
	this := BillingAddressVerificationErrorResponseErrors{}
	this.Type = resourceType
	this.Title = title
	this.Detail = detail
	this.Status = status
	return &this
}

// NewBillingAddressVerificationErrorResponseErrorsWithDefaults instantiates a new BillingAddressVerificationErrorResponseErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressVerificationErrorResponseErrorsWithDefaults() *BillingAddressVerificationErrorResponseErrors {
	this := BillingAddressVerificationErrorResponseErrors{}
	return &this
}

// GetType returns the Type field value
func (o *BillingAddressVerificationErrorResponseErrors) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponseErrors) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillingAddressVerificationErrorResponseErrors) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *BillingAddressVerificationErrorResponseErrors) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponseErrors) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BillingAddressVerificationErrorResponseErrors) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *BillingAddressVerificationErrorResponseErrors) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponseErrors) GetDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *BillingAddressVerificationErrorResponseErrors) SetDetail(v string) {
	o.Detail = v
}

// GetStatus returns the Status field value
func (o *BillingAddressVerificationErrorResponseErrors) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponseErrors) GetStatusOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BillingAddressVerificationErrorResponseErrors) SetStatus(v float32) {
	o.Status = v
}

// GetCandidates returns the Candidates field value if set, zero value otherwise.
func (o *BillingAddressVerificationErrorResponseErrors) GetCandidates() []BillingAddressAttributes {
	if o == nil || o.Candidates == nil {
		var ret []BillingAddressAttributes
		return ret
	}
	return o.Candidates
}

// GetCandidatesOk returns a tuple with the Candidates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressVerificationErrorResponseErrors) GetCandidatesOk() ([]BillingAddressAttributes, bool) {
	if o == nil || o.Candidates == nil {
		return nil, false
	}
	return o.Candidates, true
}

// HasCandidates returns a boolean if a field has been set.
func (o *BillingAddressVerificationErrorResponseErrors) HasCandidates() bool {
	if o != nil && o.Candidates != nil {
		return true
	}

	return false
}

// SetCandidates gets a reference to the given []BillingAddressAttributes and assigns it to the Candidates field.
func (o *BillingAddressVerificationErrorResponseErrors) SetCandidates(v []BillingAddressAttributes) {
	o.Candidates = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressVerificationErrorResponseErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Candidates != nil {
		toSerialize["candidates"] = o.Candidates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingAddressVerificationErrorResponseErrors) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressVerificationErrorResponseErrors := _BillingAddressVerificationErrorResponseErrors{}

	if err = json.Unmarshal(bytes, &varBillingAddressVerificationErrorResponseErrors); err == nil {
		*o = BillingAddressVerificationErrorResponseErrors(varBillingAddressVerificationErrorResponseErrors)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "status")
		delete(additionalProperties, "candidates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressVerificationErrorResponseErrors is a helper abstraction for handling nullable billingaddressverificationerrorresponseerrors types. 
type NullableBillingAddressVerificationErrorResponseErrors struct {
	value *BillingAddressVerificationErrorResponseErrors
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressVerificationErrorResponseErrors) Get() *BillingAddressVerificationErrorResponseErrors {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressVerificationErrorResponseErrors) Set(val *BillingAddressVerificationErrorResponseErrors) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressVerificationErrorResponseErrors) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressVerificationErrorResponseErrors) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressVerificationErrorResponseErrors returns a pointer to a new instance of NullableBillingAddressVerificationErrorResponseErrors.
func NewNullableBillingAddressVerificationErrorResponseErrors(val *BillingAddressVerificationErrorResponseErrors) *NullableBillingAddressVerificationErrorResponseErrors {
	return &NullableBillingAddressVerificationErrorResponseErrors{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressVerificationErrorResponseErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingAddressVerificationErrorResponseErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
