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

// BotRequestMetricsResponse struct for BotRequestMetricsResponse
type BotRequestMetricsResponse struct {
	// An array of time series data items containing metric counts. Items are grouped by `timestamp` by default, or by a combination of `timestamp` and `bot` or `datacenter` or `region` when the `group_by` parameter is set. This dataset is sparse: only keys with non-zero values are included.
	Data                 []BotRequestMetricsDataItem `json:"data,omitempty"`
	Meta                 *BotRequestMetricsMetadata  `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _BotRequestMetricsResponse BotRequestMetricsResponse

// NewBotRequestMetricsResponse instantiates a new BotRequestMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotRequestMetricsResponse() *BotRequestMetricsResponse {
	this := BotRequestMetricsResponse{}
	return &this
}

// NewBotRequestMetricsResponseWithDefaults instantiates a new BotRequestMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotRequestMetricsResponseWithDefaults() *BotRequestMetricsResponse {
	this := BotRequestMetricsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BotRequestMetricsResponse) GetData() []BotRequestMetricsDataItem {
	if o == nil || o.Data == nil {
		var ret []BotRequestMetricsDataItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsResponse) GetDataOk() ([]BotRequestMetricsDataItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BotRequestMetricsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BotRequestMetricsDataItem and assigns it to the Data field.
func (o *BotRequestMetricsResponse) SetData(v []BotRequestMetricsDataItem) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BotRequestMetricsResponse) GetMeta() BotRequestMetricsMetadata {
	if o == nil || o.Meta == nil {
		var ret BotRequestMetricsMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsResponse) GetMetaOk() (*BotRequestMetricsMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BotRequestMetricsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given BotRequestMetricsMetadata and assigns it to the Meta field.
func (o *BotRequestMetricsResponse) SetMeta(v BotRequestMetricsMetadata) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotRequestMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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
func (o *BotRequestMetricsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotRequestMetricsResponse := _BotRequestMetricsResponse{}

	if err = json.Unmarshal(bytes, &varBotRequestMetricsResponse); err == nil {
		*o = BotRequestMetricsResponse(varBotRequestMetricsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotRequestMetricsResponse is a helper abstraction for handling nullable botrequestmetricsresponse types.
type NullableBotRequestMetricsResponse struct {
	value *BotRequestMetricsResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotRequestMetricsResponse) Get() *BotRequestMetricsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotRequestMetricsResponse) Set(val *BotRequestMetricsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotRequestMetricsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotRequestMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotRequestMetricsResponse returns a pointer to a new instance of NullableBotRequestMetricsResponse.
func NewNullableBotRequestMetricsResponse(val *BotRequestMetricsResponse) *NullableBotRequestMetricsResponse {
	return &NullableBotRequestMetricsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotRequestMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotRequestMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
