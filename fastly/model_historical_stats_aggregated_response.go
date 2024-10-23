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

// HistoricalStatsAggregatedResponse struct for HistoricalStatsAggregatedResponse
type HistoricalStatsAggregatedResponse struct {
	// Whether or not we were able to successfully execute the query.
	Status *string         `json:"status,omitempty"`
	Meta   *HistoricalMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg                  NullableString `json:"msg,omitempty"`
	Data                 []Results      `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalStatsAggregatedResponse HistoricalStatsAggregatedResponse

// NewHistoricalStatsAggregatedResponse instantiates a new HistoricalStatsAggregatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalStatsAggregatedResponse() *HistoricalStatsAggregatedResponse {
	this := HistoricalStatsAggregatedResponse{}
	return &this
}

// NewHistoricalStatsAggregatedResponseWithDefaults instantiates a new HistoricalStatsAggregatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalStatsAggregatedResponseWithDefaults() *HistoricalStatsAggregatedResponse {
	this := HistoricalStatsAggregatedResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HistoricalStatsAggregatedResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalStatsAggregatedResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HistoricalStatsAggregatedResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HistoricalStatsAggregatedResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HistoricalStatsAggregatedResponse) GetMeta() HistoricalMeta {
	if o == nil || o.Meta == nil {
		var ret HistoricalMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalStatsAggregatedResponse) GetMetaOk() (*HistoricalMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HistoricalStatsAggregatedResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given HistoricalMeta and assigns it to the Meta field.
func (o *HistoricalStatsAggregatedResponse) SetMeta(v HistoricalMeta) {
	o.Meta = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricalStatsAggregatedResponse) GetMsg() string {
	if o == nil || o.Msg.Get() == nil {
		var ret string
		return ret
	}
	return *o.Msg.Get()
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricalStatsAggregatedResponse) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Msg.Get(), o.Msg.IsSet()
}

// HasMsg returns a boolean if a field has been set.
func (o *HistoricalStatsAggregatedResponse) HasMsg() bool {
	if o != nil && o.Msg.IsSet() {
		return true
	}

	return false
}

// SetMsg gets a reference to the given NullableString and assigns it to the Msg field.
func (o *HistoricalStatsAggregatedResponse) SetMsg(v string) {
	o.Msg.Set(&v)
}

// SetMsgNil sets the value for Msg to be an explicit nil
func (o *HistoricalStatsAggregatedResponse) SetMsgNil() {
	o.Msg.Set(nil)
}

// UnsetMsg ensures that no value is present for Msg, not even an explicit nil
func (o *HistoricalStatsAggregatedResponse) UnsetMsg() {
	o.Msg.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalStatsAggregatedResponse) GetData() []Results {
	if o == nil || o.Data == nil {
		var ret []Results
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalStatsAggregatedResponse) GetDataOk() ([]Results, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalStatsAggregatedResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Results and assigns it to the Data field.
func (o *HistoricalStatsAggregatedResponse) SetData(v []Results) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalStatsAggregatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Msg.IsSet() {
		toSerialize["msg"] = o.Msg.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HistoricalStatsAggregatedResponse) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalStatsAggregatedResponse := _HistoricalStatsAggregatedResponse{}

	if err = json.Unmarshal(bytes, &varHistoricalStatsAggregatedResponse); err == nil {
		*o = HistoricalStatsAggregatedResponse(varHistoricalStatsAggregatedResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalStatsAggregatedResponse is a helper abstraction for handling nullable historicalstatsaggregatedresponse types.
type NullableHistoricalStatsAggregatedResponse struct {
	value *HistoricalStatsAggregatedResponse
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalStatsAggregatedResponse) Get() *HistoricalStatsAggregatedResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalStatsAggregatedResponse) Set(val *HistoricalStatsAggregatedResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalStatsAggregatedResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalStatsAggregatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalStatsAggregatedResponse returns a pointer to a new instance of NullableHistoricalStatsAggregatedResponse.
func NewNullableHistoricalStatsAggregatedResponse(val *HistoricalStatsAggregatedResponse) *NullableHistoricalStatsAggregatedResponse {
	return &NullableHistoricalStatsAggregatedResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalStatsAggregatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHistoricalStatsAggregatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
