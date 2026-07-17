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

// RoutingConfigsResponse A paginated list of routing configs.
type RoutingConfigsResponse struct {
	Data                 []RoutingConfigResponse `json:"data,omitempty"`
	Meta                 *PaginationMeta         `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _RoutingConfigsResponse RoutingConfigsResponse

// NewRoutingConfigsResponse instantiates a new RoutingConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfigsResponse() *RoutingConfigsResponse {
	this := RoutingConfigsResponse{}
	return &this
}

// NewRoutingConfigsResponseWithDefaults instantiates a new RoutingConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigsResponseWithDefaults() *RoutingConfigsResponse {
	this := RoutingConfigsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RoutingConfigsResponse) GetData() []RoutingConfigResponse {
	if o == nil || o.Data == nil {
		var ret []RoutingConfigResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigsResponse) GetDataOk() ([]RoutingConfigResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RoutingConfigsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RoutingConfigResponse and assigns it to the Data field.
func (o *RoutingConfigsResponse) SetData(v []RoutingConfigResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RoutingConfigsResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RoutingConfigsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *RoutingConfigsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfigsResponse) MarshalJSON() ([]byte, error) {
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
func (o *RoutingConfigsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfigsResponse := _RoutingConfigsResponse{}

	if err = json.Unmarshal(bytes, &varRoutingConfigsResponse); err == nil {
		*o = RoutingConfigsResponse(varRoutingConfigsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfigsResponse is a helper abstraction for handling nullable routingconfigsresponse types.
type NullableRoutingConfigsResponse struct {
	value *RoutingConfigsResponse
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigsResponse) Get() *RoutingConfigsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigsResponse) Set(val *RoutingConfigsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigsResponse returns a pointer to a new instance of NullableRoutingConfigsResponse.
func NewNullableRoutingConfigsResponse(val *RoutingConfigsResponse) *NullableRoutingConfigsResponse {
	return &NullableRoutingConfigsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
