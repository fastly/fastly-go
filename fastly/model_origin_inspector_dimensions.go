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

// OriginInspectorDimensions The unique combination of dimensions associated with this timeseries.
type OriginInspectorDimensions struct {
	// The geographic region from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across regions.
	Region *string `json:"region,omitempty"`
	// The POP from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across POPs.
	Datacenter *string `json:"datacenter,omitempty"`
	// The origin host from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across origin hosts.
	Host                 *string `json:"host,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorDimensions OriginInspectorDimensions

// NewOriginInspectorDimensions instantiates a new OriginInspectorDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorDimensions() *OriginInspectorDimensions {
	this := OriginInspectorDimensions{}
	return &this
}

// NewOriginInspectorDimensionsWithDefaults instantiates a new OriginInspectorDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorDimensionsWithDefaults() *OriginInspectorDimensions {
	this := OriginInspectorDimensions{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OriginInspectorDimensions) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorDimensions) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OriginInspectorDimensions) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *OriginInspectorDimensions) SetRegion(v string) {
	o.Region = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *OriginInspectorDimensions) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorDimensions) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *OriginInspectorDimensions) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *OriginInspectorDimensions) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OriginInspectorDimensions) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorDimensions) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OriginInspectorDimensions) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *OriginInspectorDimensions) SetHost(v string) {
	o.Host = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorDimensions) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorDimensions) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorDimensions := _OriginInspectorDimensions{}

	if err = json.Unmarshal(bytes, &varOriginInspectorDimensions); err == nil {
		*o = OriginInspectorDimensions(varOriginInspectorDimensions)
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

// NullableOriginInspectorDimensions is a helper abstraction for handling nullable origininspectordimensions types.
type NullableOriginInspectorDimensions struct {
	value *OriginInspectorDimensions
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorDimensions) Get() *OriginInspectorDimensions {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorDimensions) Set(val *OriginInspectorDimensions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorDimensions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorDimensions returns a pointer to a new instance of NullableOriginInspectorDimensions.
func NewNullableOriginInspectorDimensions(val *OriginInspectorDimensions) *NullableOriginInspectorDimensions {
	return &NullableOriginInspectorDimensions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
