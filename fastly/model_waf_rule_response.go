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

// WafRuleResponse struct for WafRuleResponse
type WafRuleResponse struct {
	Data *WafRuleResponseData `json:"data,omitempty"`
	Included []IncludedWithWafRuleItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleResponse WafRuleResponse

// NewWafRuleResponse instantiates a new WafRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleResponse() *WafRuleResponse {
	this := WafRuleResponse{}
	return &this
}

// NewWafRuleResponseWithDefaults instantiates a new WafRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleResponseWithDefaults() *WafRuleResponse {
	this := WafRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafRuleResponse) GetData() WafRuleResponseData {
	if o == nil || o.Data == nil {
		var ret WafRuleResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponse) GetDataOk() (*WafRuleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WafRuleResponseData and assigns it to the Data field.
func (o *WafRuleResponse) SetData(v WafRuleResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafRuleResponse) GetIncluded() []IncludedWithWafRuleItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithWafRuleItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponse) GetIncludedOk() ([]IncludedWithWafRuleItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafRuleResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithWafRuleItem and assigns it to the Included field.
func (o *WafRuleResponse) SetIncluded(v []IncludedWithWafRuleItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleResponse) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleResponse := _WafRuleResponse{}

	if err = json.Unmarshal(bytes, &varWafRuleResponse); err == nil {
		*o = WafRuleResponse(varWafRuleResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleResponse is a helper abstraction for handling nullable wafruleresponse types. 
type NullableWafRuleResponse struct {
	value *WafRuleResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleResponse) Get() *WafRuleResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleResponse) Set(val *WafRuleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleResponse returns a pointer to a new instance of NullableWafRuleResponse.
func NewNullableWafRuleResponse(val *WafRuleResponse) *NullableWafRuleResponse {
	return &NullableWafRuleResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
