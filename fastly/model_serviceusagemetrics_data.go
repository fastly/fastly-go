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

// ServiceusagemetricsData struct for ServiceusagemetricsData
type ServiceusagemetricsData struct {
	CustomerId *string `json:"customer_id,omitempty"`
	// Date and time (in ISO 8601 format) for initiation point of a billing cycle, signifying the start of charges for a service or subscription.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Date and time (in ISO 8601 format) for termination point of a billing cycle, signifying the end of charges for a service or subscription.
	EndTime *time.Time `json:"end_time,omitempty"`
	// The usage type identifier for the usage. This is a single, billable metric for the product.
	UsageType *string `json:"usage_type,omitempty"`
	// The unit for the usage as shown on an invoice. If there is no explicit unit, this field will be \"unit\" (e.g., a request with `product_id` of 'cdn_usage' and `usage_type` of 'North America Requests' has no unit, and will return \"unit\").
	Unit                 *string              `json:"unit,omitempty"`
	Details              []Serviceusagemetric `json:"details,omitempty"`
	Meta                 *Metadata            `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceusagemetricsData ServiceusagemetricsData

// NewServiceusagemetricsData instantiates a new ServiceusagemetricsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceusagemetricsData() *ServiceusagemetricsData {
	this := ServiceusagemetricsData{}
	return &this
}

// NewServiceusagemetricsDataWithDefaults instantiates a new ServiceusagemetricsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceusagemetricsDataWithDefaults() *ServiceusagemetricsData {
	this := ServiceusagemetricsData{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ServiceusagemetricsData) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ServiceusagemetricsData) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ServiceusagemetricsData) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasUsageType() bool {
	if o != nil && o.UsageType != nil {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *ServiceusagemetricsData) SetUsageType(v string) {
	o.UsageType = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ServiceusagemetricsData) SetUnit(v string) {
	o.Unit = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetDetails() []Serviceusagemetric {
	if o == nil || o.Details == nil {
		var ret []Serviceusagemetric
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetDetailsOk() ([]Serviceusagemetric, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []Serviceusagemetric and assigns it to the Details field.
func (o *ServiceusagemetricsData) SetDetails(v []Serviceusagemetric) {
	o.Details = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ServiceusagemetricsData) GetMeta() Metadata {
	if o == nil || o.Meta == nil {
		var ret Metadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceusagemetricsData) GetMetaOk() (*Metadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ServiceusagemetricsData) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Metadata and assigns it to the Meta field.
func (o *ServiceusagemetricsData) SetMeta(v Metadata) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceusagemetricsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServiceusagemetricsData) UnmarshalJSON(bytes []byte) (err error) {
	varServiceusagemetricsData := _ServiceusagemetricsData{}

	if err = json.Unmarshal(bytes, &varServiceusagemetricsData); err == nil {
		*o = ServiceusagemetricsData(varServiceusagemetricsData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "usage_type")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "details")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceusagemetricsData is a helper abstraction for handling nullable serviceusagemetricsdata types.
type NullableServiceusagemetricsData struct {
	value *ServiceusagemetricsData
	isSet bool
}

// Get returns the value.
func (v NullableServiceusagemetricsData) Get() *ServiceusagemetricsData {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceusagemetricsData) Set(val *ServiceusagemetricsData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceusagemetricsData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceusagemetricsData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceusagemetricsData returns a pointer to a new instance of NullableServiceusagemetricsData.
func NewNullableServiceusagemetricsData(val *ServiceusagemetricsData) *NullableServiceusagemetricsData {
	return &NullableServiceusagemetricsData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceusagemetricsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceusagemetricsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
