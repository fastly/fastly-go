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

// OriginInspectorHistoricalMetaFilters Filters that were applied when calculating the results for this query.
type OriginInspectorHistoricalMetaFilters struct {
	Region               *string `json:"region,omitempty"`
	Datacenter           *string `json:"datacenter,omitempty"`
	Host                 *string `json:"host,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorHistoricalMetaFilters OriginInspectorHistoricalMetaFilters

// NewOriginInspectorHistoricalMetaFilters instantiates a new OriginInspectorHistoricalMetaFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorHistoricalMetaFilters() *OriginInspectorHistoricalMetaFilters {
	this := OriginInspectorHistoricalMetaFilters{}
	return &this
}

// NewOriginInspectorHistoricalMetaFiltersWithDefaults instantiates a new OriginInspectorHistoricalMetaFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorHistoricalMetaFiltersWithDefaults() *OriginInspectorHistoricalMetaFilters {
	this := OriginInspectorHistoricalMetaFilters{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMetaFilters) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMetaFilters) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMetaFilters) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *OriginInspectorHistoricalMetaFilters) SetRegion(v string) {
	o.Region = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMetaFilters) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMetaFilters) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMetaFilters) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *OriginInspectorHistoricalMetaFilters) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMetaFilters) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMetaFilters) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMetaFilters) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *OriginInspectorHistoricalMetaFilters) SetHost(v string) {
	o.Host = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorHistoricalMetaFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OriginInspectorHistoricalMetaFilters) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorHistoricalMetaFilters := _OriginInspectorHistoricalMetaFilters{}

	if err = json.Unmarshal(bytes, &varOriginInspectorHistoricalMetaFilters); err == nil {
		*o = OriginInspectorHistoricalMetaFilters(varOriginInspectorHistoricalMetaFilters)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "datacenter")
		delete(additionalProperties, "host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorHistoricalMetaFilters is a helper abstraction for handling nullable origininspectorhistoricalmetafilters types.
type NullableOriginInspectorHistoricalMetaFilters struct {
	value *OriginInspectorHistoricalMetaFilters
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorHistoricalMetaFilters) Get() *OriginInspectorHistoricalMetaFilters {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorHistoricalMetaFilters) Set(val *OriginInspectorHistoricalMetaFilters) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorHistoricalMetaFilters) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorHistoricalMetaFilters) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorHistoricalMetaFilters returns a pointer to a new instance of NullableOriginInspectorHistoricalMetaFilters.
func NewNullableOriginInspectorHistoricalMetaFilters(val *OriginInspectorHistoricalMetaFilters) *NullableOriginInspectorHistoricalMetaFilters {
	return &NullableOriginInspectorHistoricalMetaFilters{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorHistoricalMetaFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorHistoricalMetaFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
