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

// Serviceusagemetric struct for Serviceusagemetric
type Serviceusagemetric struct {
	CustomerID *string `json:"customer_id,omitempty"`
	// Service ID associated with the usage.
	ServiceID *string `json:"service_id,omitempty"`
	// Name of the service associated with the usage.
	ServiceName *string `json:"service_name,omitempty"`
	// The quantity of the usage for the billing period. Amount will be in the units provided in the parent object (e.g., a quantity of `1.3` with a unit of `gb` would have a usage amount of 1.3 gigabytes).
	UsageUnits           *float32 `json:"usage_units,omitempty"`
	AdditionalProperties map[string]any
}

type _Serviceusagemetric Serviceusagemetric

// NewServiceusagemetric instantiates a new Serviceusagemetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceusagemetric() *Serviceusagemetric {
	this := Serviceusagemetric{}
	return &this
}

// NewServiceusagemetricWithDefaults instantiates a new Serviceusagemetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceusagemetricWithDefaults() *Serviceusagemetric {
	this := Serviceusagemetric{}
	return &this
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *Serviceusagemetric) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagemetric) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *Serviceusagemetric) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *Serviceusagemetric) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *Serviceusagemetric) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagemetric) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *Serviceusagemetric) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *Serviceusagemetric) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *Serviceusagemetric) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagemetric) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *Serviceusagemetric) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *Serviceusagemetric) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetUsageUnits returns the UsageUnits field value if set, zero value otherwise.
func (o *Serviceusagemetric) GetUsageUnits() float32 {
	if o == nil || o.UsageUnits == nil {
		var ret float32
		return ret
	}
	return *o.UsageUnits
}

// GetUsageUnitsOk returns a tuple with the UsageUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceusagemetric) GetUsageUnitsOk() (*float32, bool) {
	if o == nil || o.UsageUnits == nil {
		return nil, false
	}
	return o.UsageUnits, true
}

// HasUsageUnits returns a boolean if a field has been set.
func (o *Serviceusagemetric) HasUsageUnits() bool {
	if o != nil && o.UsageUnits != nil {
		return true
	}

	return false
}

// SetUsageUnits gets a reference to the given float32 and assigns it to the UsageUnits field.
func (o *Serviceusagemetric) SetUsageUnits(v float32) {
	o.UsageUnits = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Serviceusagemetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.UsageUnits != nil {
		toSerialize["usage_units"] = o.UsageUnits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Serviceusagemetric) UnmarshalJSON(bytes []byte) (err error) {
	varServiceusagemetric := _Serviceusagemetric{}

	if err = json.Unmarshal(bytes, &varServiceusagemetric); err == nil {
		*o = Serviceusagemetric(varServiceusagemetric)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "service_name")
		delete(additionalProperties, "usage_units")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceusagemetric is a helper abstraction for handling nullable serviceusagemetric types.
type NullableServiceusagemetric struct {
	value *Serviceusagemetric
	isSet bool
}

// Get returns the value.
func (v NullableServiceusagemetric) Get() *Serviceusagemetric {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceusagemetric) Set(val *Serviceusagemetric) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceusagemetric) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceusagemetric) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceusagemetric returns a pointer to a new instance of NullableServiceusagemetric.
func NewNullableServiceusagemetric(val *Serviceusagemetric) *NullableServiceusagemetric {
	return &NullableServiceusagemetric{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceusagemetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceusagemetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
