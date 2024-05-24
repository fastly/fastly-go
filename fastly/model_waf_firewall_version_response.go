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

// WafFirewallVersionResponse struct for WafFirewallVersionResponse
type WafFirewallVersionResponse struct {
	Data *WafFirewallVersionResponseData `json:"data,omitempty"`
	Included []IncludedWithWafFirewallVersionItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionResponse WafFirewallVersionResponse

// NewWafFirewallVersionResponse instantiates a new WafFirewallVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionResponse() *WafFirewallVersionResponse {
	this := WafFirewallVersionResponse{}
	return &this
}

// NewWafFirewallVersionResponseWithDefaults instantiates a new WafFirewallVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionResponseWithDefaults() *WafFirewallVersionResponse {
	this := WafFirewallVersionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafFirewallVersionResponse) GetData() WafFirewallVersionResponseData {
	if o == nil || o.Data == nil {
		var ret WafFirewallVersionResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponse) GetDataOk() (*WafFirewallVersionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafFirewallVersionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WafFirewallVersionResponseData and assigns it to the Data field.
func (o *WafFirewallVersionResponse) SetData(v WafFirewallVersionResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafFirewallVersionResponse) GetIncluded() []IncludedWithWafFirewallVersionItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithWafFirewallVersionItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponse) GetIncludedOk() ([]IncludedWithWafFirewallVersionItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafFirewallVersionResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithWafFirewallVersionItem and assigns it to the Included field.
func (o *WafFirewallVersionResponse) SetIncluded(v []IncludedWithWafFirewallVersionItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionResponse) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionResponse := _WafFirewallVersionResponse{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionResponse); err == nil {
		*o = WafFirewallVersionResponse(varWafFirewallVersionResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallVersionResponse is a helper abstraction for handling nullable waffirewallversionresponse types. 
type NullableWafFirewallVersionResponse struct {
	value *WafFirewallVersionResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionResponse) Get() *WafFirewallVersionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionResponse) Set(val *WafFirewallVersionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionResponse returns a pointer to a new instance of NullableWafFirewallVersionResponse.
func NewNullableWafFirewallVersionResponse(val *WafFirewallVersionResponse) *NullableWafFirewallVersionResponse {
	return &NullableWafFirewallVersionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
