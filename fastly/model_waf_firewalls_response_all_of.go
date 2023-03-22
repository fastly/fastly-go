// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// WafFirewallsResponseAllOf struct for WafFirewallsResponseAllOf
type WafFirewallsResponseAllOf struct {
	Data []WafFirewallResponseData `json:"data,omitempty"`
	Included []SchemasWafFirewallVersion `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallsResponseAllOf WafFirewallsResponseAllOf

// NewWafFirewallsResponseAllOf instantiates a new WafFirewallsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallsResponseAllOf() *WafFirewallsResponseAllOf {
	this := WafFirewallsResponseAllOf{}
	return &this
}

// NewWafFirewallsResponseAllOfWithDefaults instantiates a new WafFirewallsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallsResponseAllOfWithDefaults() *WafFirewallsResponseAllOf {
	this := WafFirewallsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafFirewallsResponseAllOf) GetData() []WafFirewallResponseData {
	if o == nil || o.Data == nil {
		var ret []WafFirewallResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponseAllOf) GetDataOk() ([]WafFirewallResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafFirewallsResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafFirewallResponseData and assigns it to the Data field.
func (o *WafFirewallsResponseAllOf) SetData(v []WafFirewallResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafFirewallsResponseAllOf) GetIncluded() []SchemasWafFirewallVersion {
	if o == nil || o.Included == nil {
		var ret []SchemasWafFirewallVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponseAllOf) GetIncludedOk() ([]SchemasWafFirewallVersion, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafFirewallsResponseAllOf) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SchemasWafFirewallVersion and assigns it to the Included field.
func (o *WafFirewallsResponseAllOf) SetIncluded(v []SchemasWafFirewallVersion) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafFirewallsResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallsResponseAllOf := _WafFirewallsResponseAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallsResponseAllOf); err == nil {
		*o = WafFirewallsResponseAllOf(varWafFirewallsResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallsResponseAllOf is a helper abstraction for handling nullable waffirewallsresponseallof types. 
type NullableWafFirewallsResponseAllOf struct {
	value *WafFirewallsResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallsResponseAllOf) Get() *WafFirewallsResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallsResponseAllOf) Set(val *WafFirewallsResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallsResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallsResponseAllOf returns a pointer to a new instance of NullableWafFirewallsResponseAllOf.
func NewNullableWafFirewallsResponseAllOf(val *WafFirewallsResponseAllOf) *NullableWafFirewallsResponseAllOf {
	return &NullableWafFirewallsResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
