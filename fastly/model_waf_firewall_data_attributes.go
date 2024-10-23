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

// WafFirewallDataAttributes struct for WafFirewallDataAttributes
type WafFirewallDataAttributes struct {
	// The status of the firewall.
	Disabled *bool `json:"disabled,omitempty"`
	// Name of the corresponding condition object.
	PrefetchCondition *string `json:"prefetch_condition,omitempty"`
	// Name of the corresponding response object.
	Response             *string `json:"response,omitempty"`
	ServiceVersionNumber *int32  `json:"service_version_number,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallDataAttributes WafFirewallDataAttributes

// NewWafFirewallDataAttributes instantiates a new WafFirewallDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallDataAttributes() *WafFirewallDataAttributes {
	this := WafFirewallDataAttributes{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewWafFirewallDataAttributesWithDefaults instantiates a new WafFirewallDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallDataAttributesWithDefaults() *WafFirewallDataAttributes {
	this := WafFirewallDataAttributes{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *WafFirewallDataAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallDataAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *WafFirewallDataAttributes) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *WafFirewallDataAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetPrefetchCondition returns the PrefetchCondition field value if set, zero value otherwise.
func (o *WafFirewallDataAttributes) GetPrefetchCondition() string {
	if o == nil || o.PrefetchCondition == nil {
		var ret string
		return ret
	}
	return *o.PrefetchCondition
}

// GetPrefetchConditionOk returns a tuple with the PrefetchCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallDataAttributes) GetPrefetchConditionOk() (*string, bool) {
	if o == nil || o.PrefetchCondition == nil {
		return nil, false
	}
	return o.PrefetchCondition, true
}

// HasPrefetchCondition returns a boolean if a field has been set.
func (o *WafFirewallDataAttributes) HasPrefetchCondition() bool {
	if o != nil && o.PrefetchCondition != nil {
		return true
	}

	return false
}

// SetPrefetchCondition gets a reference to the given string and assigns it to the PrefetchCondition field.
func (o *WafFirewallDataAttributes) SetPrefetchCondition(v string) {
	o.PrefetchCondition = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *WafFirewallDataAttributes) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallDataAttributes) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *WafFirewallDataAttributes) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *WafFirewallDataAttributes) SetResponse(v string) {
	o.Response = &v
}

// GetServiceVersionNumber returns the ServiceVersionNumber field value if set, zero value otherwise.
func (o *WafFirewallDataAttributes) GetServiceVersionNumber() int32 {
	if o == nil || o.ServiceVersionNumber == nil {
		var ret int32
		return ret
	}
	return *o.ServiceVersionNumber
}

// GetServiceVersionNumberOk returns a tuple with the ServiceVersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallDataAttributes) GetServiceVersionNumberOk() (*int32, bool) {
	if o == nil || o.ServiceVersionNumber == nil {
		return nil, false
	}
	return o.ServiceVersionNumber, true
}

// HasServiceVersionNumber returns a boolean if a field has been set.
func (o *WafFirewallDataAttributes) HasServiceVersionNumber() bool {
	if o != nil && o.ServiceVersionNumber != nil {
		return true
	}

	return false
}

// SetServiceVersionNumber gets a reference to the given int32 and assigns it to the ServiceVersionNumber field.
func (o *WafFirewallDataAttributes) SetServiceVersionNumber(v int32) {
	o.ServiceVersionNumber = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.PrefetchCondition != nil {
		toSerialize["prefetch_condition"] = o.PrefetchCondition
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.ServiceVersionNumber != nil {
		toSerialize["service_version_number"] = o.ServiceVersionNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafFirewallDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallDataAttributes := _WafFirewallDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafFirewallDataAttributes); err == nil {
		*o = WafFirewallDataAttributes(varWafFirewallDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "prefetch_condition")
		delete(additionalProperties, "response")
		delete(additionalProperties, "service_version_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallDataAttributes is a helper abstraction for handling nullable waffirewalldataattributes types.
type NullableWafFirewallDataAttributes struct {
	value *WafFirewallDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallDataAttributes) Get() *WafFirewallDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallDataAttributes) Set(val *WafFirewallDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallDataAttributes returns a pointer to a new instance of NullableWafFirewallDataAttributes.
func NewNullableWafFirewallDataAttributes(val *WafFirewallDataAttributes) *NullableWafFirewallDataAttributes {
	return &NullableWafFirewallDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafFirewallDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
