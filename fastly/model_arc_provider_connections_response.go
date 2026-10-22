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

// ArcProviderConnectionsResponse struct for ArcProviderConnectionsResponse
type ArcProviderConnectionsResponse struct {
	// The provider connections returned by the request.
	Data                 []ArcProviderConnectionListItemResponse `json:"data"`
	Meta                 ArcProviderConnectionsMeta              `json:"meta"`
	AdditionalProperties map[string]any
}

type _ArcProviderConnectionsResponse ArcProviderConnectionsResponse

// NewArcProviderConnectionsResponse instantiates a new ArcProviderConnectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProviderConnectionsResponse(data []ArcProviderConnectionListItemResponse, meta ArcProviderConnectionsMeta) *ArcProviderConnectionsResponse {
	this := ArcProviderConnectionsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewArcProviderConnectionsResponseWithDefaults instantiates a new ArcProviderConnectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderConnectionsResponseWithDefaults() *ArcProviderConnectionsResponse {
	this := ArcProviderConnectionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ArcProviderConnectionsResponse) GetData() []ArcProviderConnectionListItemResponse {
	if o == nil {
		var ret []ArcProviderConnectionListItemResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionsResponse) GetDataOk() ([]ArcProviderConnectionListItemResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ArcProviderConnectionsResponse) SetData(v []ArcProviderConnectionListItemResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ArcProviderConnectionsResponse) GetMeta() ArcProviderConnectionsMeta {
	if o == nil {
		var ret ArcProviderConnectionsMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionsResponse) GetMetaOk() (*ArcProviderConnectionsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ArcProviderConnectionsResponse) SetMeta(v ArcProviderConnectionsMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcProviderConnectionsResponse := _ArcProviderConnectionsResponse{}

	if err = json.Unmarshal(bytes, &varArcProviderConnectionsResponse); err == nil {
		*o = ArcProviderConnectionsResponse(varArcProviderConnectionsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProviderConnectionsResponse is a helper abstraction for handling nullable arcproviderconnectionsresponse types.
type NullableArcProviderConnectionsResponse struct {
	value *ArcProviderConnectionsResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionsResponse) Get() *ArcProviderConnectionsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionsResponse) Set(val *ArcProviderConnectionsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionsResponse returns a pointer to a new instance of NullableArcProviderConnectionsResponse.
func NewNullableArcProviderConnectionsResponse(val *ArcProviderConnectionsResponse) *NullableArcProviderConnectionsResponse {
	return &NullableArcProviderConnectionsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
