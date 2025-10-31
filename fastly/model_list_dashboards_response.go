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

// ListDashboardsResponse struct for ListDashboardsResponse
type ListDashboardsResponse struct {
	Data []Dashboard `json:"data,omitempty"`
	// Meta for the pagination.
	Meta                 interface{} `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _ListDashboardsResponse ListDashboardsResponse

// NewListDashboardsResponse instantiates a new ListDashboardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDashboardsResponse() *ListDashboardsResponse {
	this := ListDashboardsResponse{}
	return &this
}

// NewListDashboardsResponseWithDefaults instantiates a new ListDashboardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDashboardsResponseWithDefaults() *ListDashboardsResponse {
	this := ListDashboardsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDashboardsResponse) GetData() []Dashboard {
	if o == nil || o.Data == nil {
		var ret []Dashboard
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDashboardsResponse) GetDataOk() ([]Dashboard, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDashboardsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Dashboard and assigns it to the Data field.
func (o *ListDashboardsResponse) SetData(v []Dashboard) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListDashboardsResponse) GetMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListDashboardsResponse) GetMetaOk() (*interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListDashboardsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given interface{} and assigns it to the Meta field.
func (o *ListDashboardsResponse) SetMeta(v interface{}) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ListDashboardsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ListDashboardsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListDashboardsResponse := _ListDashboardsResponse{}

	if err = json.Unmarshal(bytes, &varListDashboardsResponse); err == nil {
		*o = ListDashboardsResponse(varListDashboardsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableListDashboardsResponse is a helper abstraction for handling nullable listdashboardsresponse types.
type NullableListDashboardsResponse struct {
	value *ListDashboardsResponse
	isSet bool
}

// Get returns the value.
func (v NullableListDashboardsResponse) Get() *ListDashboardsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableListDashboardsResponse) Set(val *ListDashboardsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableListDashboardsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableListDashboardsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListDashboardsResponse returns a pointer to a new instance of NullableListDashboardsResponse.
func NewNullableListDashboardsResponse(val *ListDashboardsResponse) *NullableListDashboardsResponse {
	return &NullableListDashboardsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableListDashboardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableListDashboardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
