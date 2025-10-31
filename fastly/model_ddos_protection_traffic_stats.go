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
	"time"
)

// DdosProtectionTrafficStats struct for DdosProtectionTrafficStats
type DdosProtectionTrafficStats struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerId *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceId  *string                        `json:"service_id,omitempty"`
	Attributes []DdosProtectionAttributeStats `json:"attributes,omitempty"`
	// Rule traffic percentage for the event.
	TrafficPercentage    *float32 `json:"traffic_percentage,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionTrafficStats DdosProtectionTrafficStats

// NewDdosProtectionTrafficStats instantiates a new DdosProtectionTrafficStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionTrafficStats() *DdosProtectionTrafficStats {
	this := DdosProtectionTrafficStats{}
	return &this
}

// NewDdosProtectionTrafficStatsWithDefaults instantiates a new DdosProtectionTrafficStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionTrafficStatsWithDefaults() *DdosProtectionTrafficStats {
	this := DdosProtectionTrafficStats{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionTrafficStats) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionTrafficStats) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *DdosProtectionTrafficStats) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *DdosProtectionTrafficStats) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *DdosProtectionTrafficStats) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionTrafficStats) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionTrafficStats) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *DdosProtectionTrafficStats) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *DdosProtectionTrafficStats) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *DdosProtectionTrafficStats) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStats) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStats) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *DdosProtectionTrafficStats) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStats) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStats) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DdosProtectionTrafficStats) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStats) GetAttributes() []DdosProtectionAttributeStats {
	if o == nil || o.Attributes == nil {
		var ret []DdosProtectionAttributeStats
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStats) GetAttributesOk() ([]DdosProtectionAttributeStats, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []DdosProtectionAttributeStats and assigns it to the Attributes field.
func (o *DdosProtectionTrafficStats) SetAttributes(v []DdosProtectionAttributeStats) {
	o.Attributes = v
}

// GetTrafficPercentage returns the TrafficPercentage field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStats) GetTrafficPercentage() float32 {
	if o == nil || o.TrafficPercentage == nil {
		var ret float32
		return ret
	}
	return *o.TrafficPercentage
}

// GetTrafficPercentageOk returns a tuple with the TrafficPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStats) GetTrafficPercentageOk() (*float32, bool) {
	if o == nil || o.TrafficPercentage == nil {
		return nil, false
	}
	return o.TrafficPercentage, true
}

// HasTrafficPercentage returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStats) HasTrafficPercentage() bool {
	if o != nil && o.TrafficPercentage != nil {
		return true
	}

	return false
}

// SetTrafficPercentage gets a reference to the given float32 and assigns it to the TrafficPercentage field.
func (o *DdosProtectionTrafficStats) SetTrafficPercentage(v float32) {
	o.TrafficPercentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionTrafficStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
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
func (o *DdosProtectionTrafficStats) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionTrafficStats := _DdosProtectionTrafficStats{}

	if err = json.Unmarshal(bytes, &varDdosProtectionTrafficStats); err == nil {
		*o = DdosProtectionTrafficStats(varDdosProtectionTrafficStats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "traffic_percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionTrafficStats is a helper abstraction for handling nullable ddosprotectiontrafficstats types.
type NullableDdosProtectionTrafficStats struct {
	value *DdosProtectionTrafficStats
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionTrafficStats) Get() *DdosProtectionTrafficStats {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionTrafficStats) Set(val *DdosProtectionTrafficStats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionTrafficStats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionTrafficStats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionTrafficStats returns a pointer to a new instance of NullableDdosProtectionTrafficStats.
func NewNullableDdosProtectionTrafficStats(val *DdosProtectionTrafficStats) *NullableDdosProtectionTrafficStats {
	return &NullableDdosProtectionTrafficStats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionTrafficStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionTrafficStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
