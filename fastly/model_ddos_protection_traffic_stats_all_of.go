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

// DdosProtectionTrafficStatsAllOf struct for DdosProtectionTrafficStatsAllOf
type DdosProtectionTrafficStatsAllOf struct {
	// Alphanumeric string identifying the customer.
	CustomerId *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceId  *string                        `json:"service_id,omitempty"`
	Attributes []DdosProtectionAttributeStats `json:"attributes,omitempty"`
	// Rule traffic percentage for the event.
	TrafficPercentage    *float32 `json:"traffic_percentage,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionTrafficStatsAllOf DdosProtectionTrafficStatsAllOf

// NewDdosProtectionTrafficStatsAllOf instantiates a new DdosProtectionTrafficStatsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionTrafficStatsAllOf() *DdosProtectionTrafficStatsAllOf {
	this := DdosProtectionTrafficStatsAllOf{}
	return &this
}

// NewDdosProtectionTrafficStatsAllOfWithDefaults instantiates a new DdosProtectionTrafficStatsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionTrafficStatsAllOfWithDefaults() *DdosProtectionTrafficStatsAllOf {
	this := DdosProtectionTrafficStatsAllOf{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *DdosProtectionTrafficStatsAllOf) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DdosProtectionTrafficStatsAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetAttributes() []DdosProtectionAttributeStats {
	if o == nil || o.Attributes == nil {
		var ret []DdosProtectionAttributeStats
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetAttributesOk() ([]DdosProtectionAttributeStats, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []DdosProtectionAttributeStats and assigns it to the Attributes field.
func (o *DdosProtectionTrafficStatsAllOf) SetAttributes(v []DdosProtectionAttributeStats) {
	o.Attributes = v
}

// GetTrafficPercentage returns the TrafficPercentage field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetTrafficPercentage() float32 {
	if o == nil || o.TrafficPercentage == nil {
		var ret float32
		return ret
	}
	return *o.TrafficPercentage
}

// GetTrafficPercentageOk returns a tuple with the TrafficPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetTrafficPercentageOk() (*float32, bool) {
	if o == nil || o.TrafficPercentage == nil {
		return nil, false
	}
	return o.TrafficPercentage, true
}

// HasTrafficPercentage returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasTrafficPercentage() bool {
	if o != nil && o.TrafficPercentage != nil {
		return true
	}

	return false
}

// SetTrafficPercentage gets a reference to the given float32 and assigns it to the TrafficPercentage field.
func (o *DdosProtectionTrafficStatsAllOf) SetTrafficPercentage(v float32) {
	o.TrafficPercentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionTrafficStatsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.TrafficPercentage != nil {
		toSerialize["traffic_percentage"] = o.TrafficPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionTrafficStatsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionTrafficStatsAllOf := _DdosProtectionTrafficStatsAllOf{}

	if err = json.Unmarshal(bytes, &varDdosProtectionTrafficStatsAllOf); err == nil {
		*o = DdosProtectionTrafficStatsAllOf(varDdosProtectionTrafficStatsAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "traffic_percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionTrafficStatsAllOf is a helper abstraction for handling nullable ddosprotectiontrafficstatsallof types.
type NullableDdosProtectionTrafficStatsAllOf struct {
	value *DdosProtectionTrafficStatsAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionTrafficStatsAllOf) Get() *DdosProtectionTrafficStatsAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionTrafficStatsAllOf) Set(val *DdosProtectionTrafficStatsAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionTrafficStatsAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionTrafficStatsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionTrafficStatsAllOf returns a pointer to a new instance of NullableDdosProtectionTrafficStatsAllOf.
func NewNullableDdosProtectionTrafficStatsAllOf(val *DdosProtectionTrafficStatsAllOf) *NullableDdosProtectionTrafficStatsAllOf {
	return &NullableDdosProtectionTrafficStatsAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionTrafficStatsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionTrafficStatsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
