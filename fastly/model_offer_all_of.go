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

// OfferAllOf struct for OfferAllOf
type OfferAllOf struct {
	// The aftermarket vendor.
	Vendor *string `json:"vendor,omitempty"`
	// The price for the domain from the aftermarket vendor.
	Price *string `json:"price,omitempty"`
	// The currency for the aftermarket offer.
	Currency             *string `json:"currency,omitempty"`
	AdditionalProperties map[string]any
}

type _OfferAllOf OfferAllOf

// NewOfferAllOf instantiates a new OfferAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferAllOf() *OfferAllOf {
	this := OfferAllOf{}
	return &this
}

// NewOfferAllOfWithDefaults instantiates a new OfferAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferAllOfWithDefaults() *OfferAllOf {
	this := OfferAllOf{}
	return &this
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *OfferAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *OfferAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *OfferAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OfferAllOf) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferAllOf) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OfferAllOf) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OfferAllOf) SetPrice(v string) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OfferAllOf) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferAllOf) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OfferAllOf) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OfferAllOf) SetCurrency(v string) {
	o.Currency = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OfferAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OfferAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOfferAllOf := _OfferAllOf{}

	if err = json.Unmarshal(bytes, &varOfferAllOf); err == nil {
		*o = OfferAllOf(varOfferAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "price")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOfferAllOf is a helper abstraction for handling nullable offerallof types.
type NullableOfferAllOf struct {
	value *OfferAllOf
	isSet bool
}

// Get returns the value.
func (v NullableOfferAllOf) Get() *OfferAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableOfferAllOf) Set(val *OfferAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOfferAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOfferAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOfferAllOf returns a pointer to a new instance of NullableOfferAllOf.
func NewNullableOfferAllOf(val *OfferAllOf) *NullableOfferAllOf {
	return &NullableOfferAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOfferAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOfferAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
