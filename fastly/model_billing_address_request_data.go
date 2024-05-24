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

// BillingAddressRequestData struct for BillingAddressRequestData
type BillingAddressRequestData struct {
	Type *TypeBillingAddress `json:"type,omitempty"`
	Attributes *BillingAddressAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressRequestData BillingAddressRequestData

// NewBillingAddressRequestData instantiates a new BillingAddressRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressRequestData() *BillingAddressRequestData {
	this := BillingAddressRequestData{}
	var resourceType TypeBillingAddress = TYPEBILLINGADDRESS_BILLING_ADDRESS
	this.Type = &resourceType
	return &this
}

// NewBillingAddressRequestDataWithDefaults instantiates a new BillingAddressRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressRequestDataWithDefaults() *BillingAddressRequestData {
	this := BillingAddressRequestData{}
	var resourceType TypeBillingAddress = TYPEBILLINGADDRESS_BILLING_ADDRESS
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingAddressRequestData) GetType() TypeBillingAddress {
	if o == nil || o.Type == nil {
		var ret TypeBillingAddress
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressRequestData) GetTypeOk() (*TypeBillingAddress, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingAddressRequestData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeBillingAddress and assigns it to the Type field.
func (o *BillingAddressRequestData) SetType(v TypeBillingAddress) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BillingAddressRequestData) GetAttributes() BillingAddressAttributes {
	if o == nil || o.Attributes == nil {
		var ret BillingAddressAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressRequestData) GetAttributesOk() (*BillingAddressAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BillingAddressRequestData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BillingAddressAttributes and assigns it to the Attributes field.
func (o *BillingAddressRequestData) SetAttributes(v BillingAddressAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingAddressRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressRequestData := _BillingAddressRequestData{}

	if err = json.Unmarshal(bytes, &varBillingAddressRequestData); err == nil {
		*o = BillingAddressRequestData(varBillingAddressRequestData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressRequestData is a helper abstraction for handling nullable billingaddressrequestdata types. 
type NullableBillingAddressRequestData struct {
	value *BillingAddressRequestData
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressRequestData) Get() *BillingAddressRequestData {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressRequestData) Set(val *BillingAddressRequestData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressRequestData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressRequestData returns a pointer to a new instance of NullableBillingAddressRequestData.
func NewNullableBillingAddressRequestData(val *BillingAddressRequestData) *NullableBillingAddressRequestData {
	return &NullableBillingAddressRequestData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingAddressRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
