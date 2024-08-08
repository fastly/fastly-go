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

// BillingRegions struct for BillingRegions
type BillingRegions struct {
	Cost *float32 `json:"cost,omitempty"`
	Bandwidth *BillingBandwidth `json:"bandwidth,omitempty"`
	Percentile *BillingBandwidth `json:"percentile,omitempty"`
	Requests *BillingBandwidth `json:"requests,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingRegions BillingRegions

// NewBillingRegions instantiates a new BillingRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingRegions() *BillingRegions {
	this := BillingRegions{}
	return &this
}

// NewBillingRegionsWithDefaults instantiates a new BillingRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingRegionsWithDefaults() *BillingRegions {
	this := BillingRegions{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingRegions) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingRegions) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingRegions) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingRegions) SetCost(v float32) {
	o.Cost = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *BillingRegions) GetBandwidth() BillingBandwidth {
	if o == nil || o.Bandwidth == nil {
		var ret BillingBandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingRegions) GetBandwidthOk() (*BillingBandwidth, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *BillingRegions) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given BillingBandwidth and assigns it to the Bandwidth field.
func (o *BillingRegions) SetBandwidth(v BillingBandwidth) {
	o.Bandwidth = &v
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *BillingRegions) GetPercentile() BillingBandwidth {
	if o == nil || o.Percentile == nil {
		var ret BillingBandwidth
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingRegions) GetPercentileOk() (*BillingBandwidth, bool) {
	if o == nil || o.Percentile == nil {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *BillingRegions) HasPercentile() bool {
	if o != nil && o.Percentile != nil {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given BillingBandwidth and assigns it to the Percentile field.
func (o *BillingRegions) SetPercentile(v BillingBandwidth) {
	o.Percentile = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *BillingRegions) GetRequests() BillingBandwidth {
	if o == nil || o.Requests == nil {
		var ret BillingBandwidth
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingRegions) GetRequestsOk() (*BillingBandwidth, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BillingRegions) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given BillingBandwidth and assigns it to the Requests field.
func (o *BillingRegions) SetRequests(v BillingBandwidth) {
	o.Requests = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingRegions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.Percentile != nil {
		toSerialize["percentile"] = o.Percentile
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingRegions) UnmarshalJSON(bytes []byte) (err error) {
	varBillingRegions := _BillingRegions{}

	if err = json.Unmarshal(bytes, &varBillingRegions); err == nil {
		*o = BillingRegions(varBillingRegions)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cost")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "percentile")
		delete(additionalProperties, "requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingRegions is a helper abstraction for handling nullable billingregions types. 
type NullableBillingRegions struct {
	value *BillingRegions
	isSet bool
}

// Get returns the value.
func (v NullableBillingRegions) Get() *BillingRegions {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingRegions) Set(val *BillingRegions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingRegions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingRegions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingRegions returns a pointer to a new instance of NullableBillingRegions.
func NewNullableBillingRegions(val *BillingRegions) *NullableBillingRegions {
	return &NullableBillingRegions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
