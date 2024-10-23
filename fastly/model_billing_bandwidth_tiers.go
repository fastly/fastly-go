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

// BillingBandwidthTiers struct for BillingBandwidthTiers
type BillingBandwidthTiers struct {
	Name                 *string  `json:"name,omitempty"`
	Units                *float32 `json:"units,omitempty"`
	Price                *float32 `json:"price,omitempty"`
	DiscountedPrice      *float32 `json:"discounted_price,omitempty"`
	Total                *float32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingBandwidthTiers BillingBandwidthTiers

// NewBillingBandwidthTiers instantiates a new BillingBandwidthTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingBandwidthTiers() *BillingBandwidthTiers {
	this := BillingBandwidthTiers{}
	return &this
}

// NewBillingBandwidthTiersWithDefaults instantiates a new BillingBandwidthTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingBandwidthTiersWithDefaults() *BillingBandwidthTiers {
	this := BillingBandwidthTiers{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingBandwidthTiers) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidthTiers) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingBandwidthTiers) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingBandwidthTiers) SetName(v string) {
	o.Name = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BillingBandwidthTiers) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidthTiers) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BillingBandwidthTiers) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *BillingBandwidthTiers) SetUnits(v float32) {
	o.Units = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingBandwidthTiers) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidthTiers) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingBandwidthTiers) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingBandwidthTiers) SetPrice(v float32) {
	o.Price = &v
}

// GetDiscountedPrice returns the DiscountedPrice field value if set, zero value otherwise.
func (o *BillingBandwidthTiers) GetDiscountedPrice() float32 {
	if o == nil || o.DiscountedPrice == nil {
		var ret float32
		return ret
	}
	return *o.DiscountedPrice
}

// GetDiscountedPriceOk returns a tuple with the DiscountedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidthTiers) GetDiscountedPriceOk() (*float32, bool) {
	if o == nil || o.DiscountedPrice == nil {
		return nil, false
	}
	return o.DiscountedPrice, true
}

// HasDiscountedPrice returns a boolean if a field has been set.
func (o *BillingBandwidthTiers) HasDiscountedPrice() bool {
	if o != nil && o.DiscountedPrice != nil {
		return true
	}

	return false
}

// SetDiscountedPrice gets a reference to the given float32 and assigns it to the DiscountedPrice field.
func (o *BillingBandwidthTiers) SetDiscountedPrice(v float32) {
	o.DiscountedPrice = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BillingBandwidthTiers) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidthTiers) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BillingBandwidthTiers) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *BillingBandwidthTiers) SetTotal(v float32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingBandwidthTiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.DiscountedPrice != nil {
		toSerialize["discounted_price"] = o.DiscountedPrice
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BillingBandwidthTiers) UnmarshalJSON(bytes []byte) (err error) {
	varBillingBandwidthTiers := _BillingBandwidthTiers{}

	if err = json.Unmarshal(bytes, &varBillingBandwidthTiers); err == nil {
		*o = BillingBandwidthTiers(varBillingBandwidthTiers)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "units")
		delete(additionalProperties, "price")
		delete(additionalProperties, "discounted_price")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingBandwidthTiers is a helper abstraction for handling nullable billingbandwidthtiers types.
type NullableBillingBandwidthTiers struct {
	value *BillingBandwidthTiers
	isSet bool
}

// Get returns the value.
func (v NullableBillingBandwidthTiers) Get() *BillingBandwidthTiers {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingBandwidthTiers) Set(val *BillingBandwidthTiers) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingBandwidthTiers) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingBandwidthTiers) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingBandwidthTiers returns a pointer to a new instance of NullableBillingBandwidthTiers.
func NewNullableBillingBandwidthTiers(val *BillingBandwidthTiers) *NullableBillingBandwidthTiers {
	return &NullableBillingBandwidthTiers{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingBandwidthTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBillingBandwidthTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
