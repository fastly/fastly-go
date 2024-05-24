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

// WafRuleRevisionResponse struct for WafRuleRevisionResponse
type WafRuleRevisionResponse struct {
	Data *WafRuleRevisionResponseData `json:"data,omitempty"`
	Included []WafRule `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleRevisionResponse WafRuleRevisionResponse

// NewWafRuleRevisionResponse instantiates a new WafRuleRevisionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleRevisionResponse() *WafRuleRevisionResponse {
	this := WafRuleRevisionResponse{}
	return &this
}

// NewWafRuleRevisionResponseWithDefaults instantiates a new WafRuleRevisionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleRevisionResponseWithDefaults() *WafRuleRevisionResponse {
	this := WafRuleRevisionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafRuleRevisionResponse) GetData() WafRuleRevisionResponseData {
	if o == nil || o.Data == nil {
		var ret WafRuleRevisionResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponse) GetDataOk() (*WafRuleRevisionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafRuleRevisionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WafRuleRevisionResponseData and assigns it to the Data field.
func (o *WafRuleRevisionResponse) SetData(v WafRuleRevisionResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafRuleRevisionResponse) GetIncluded() []WafRule {
	if o == nil || o.Included == nil {
		var ret []WafRule
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponse) GetIncludedOk() ([]WafRule, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafRuleRevisionResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []WafRule and assigns it to the Included field.
func (o *WafRuleRevisionResponse) SetIncluded(v []WafRule) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionResponse) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleRevisionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleRevisionResponse := _WafRuleRevisionResponse{}

	if err = json.Unmarshal(bytes, &varWafRuleRevisionResponse); err == nil {
		*o = WafRuleRevisionResponse(varWafRuleRevisionResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleRevisionResponse is a helper abstraction for handling nullable wafrulerevisionresponse types. 
type NullableWafRuleRevisionResponse struct {
	value *WafRuleRevisionResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionResponse) Get() *WafRuleRevisionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionResponse) Set(val *WafRuleRevisionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionResponse returns a pointer to a new instance of NullableWafRuleRevisionResponse.
func NewNullableWafRuleRevisionResponse(val *WafRuleRevisionResponse) *NullableWafRuleRevisionResponse {
	return &NullableWafRuleRevisionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleRevisionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
