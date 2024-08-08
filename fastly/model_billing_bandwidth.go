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

// BillingBandwidth struct for BillingBandwidth
type BillingBandwidth struct {
	Total *float32 `json:"total,omitempty"`
	Tiers []BillingBandwidthTiers `json:"tiers,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingBandwidth BillingBandwidth

// NewBillingBandwidth instantiates a new BillingBandwidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingBandwidth() *BillingBandwidth {
	this := BillingBandwidth{}
	return &this
}

// NewBillingBandwidthWithDefaults instantiates a new BillingBandwidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingBandwidthWithDefaults() *BillingBandwidth {
	this := BillingBandwidth{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BillingBandwidth) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidth) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BillingBandwidth) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *BillingBandwidth) SetTotal(v float32) {
	o.Total = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *BillingBandwidth) GetTiers() []BillingBandwidthTiers {
	if o == nil || o.Tiers == nil {
		var ret []BillingBandwidthTiers
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingBandwidth) GetTiersOk() ([]BillingBandwidthTiers, bool) {
	if o == nil || o.Tiers == nil {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *BillingBandwidth) HasTiers() bool {
	if o != nil && o.Tiers != nil {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []BillingBandwidthTiers and assigns it to the Tiers field.
func (o *BillingBandwidth) SetTiers(v []BillingBandwidthTiers) {
	o.Tiers = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingBandwidth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingBandwidth) UnmarshalJSON(bytes []byte) (err error) {
	varBillingBandwidth := _BillingBandwidth{}

	if err = json.Unmarshal(bytes, &varBillingBandwidth); err == nil {
		*o = BillingBandwidth(varBillingBandwidth)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "tiers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingBandwidth is a helper abstraction for handling nullable billingbandwidth types. 
type NullableBillingBandwidth struct {
	value *BillingBandwidth
	isSet bool
}

// Get returns the value.
func (v NullableBillingBandwidth) Get() *BillingBandwidth {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingBandwidth) Set(val *BillingBandwidth) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingBandwidth) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingBandwidth returns a pointer to a new instance of NullableBillingBandwidth.
func NewNullableBillingBandwidth(val *BillingBandwidth) *NullableBillingBandwidth {
	return &NullableBillingBandwidth{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
