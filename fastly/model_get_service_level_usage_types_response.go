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

// GetServiceLevelUsageTypesResponse struct for GetServiceLevelUsageTypesResponse
type GetServiceLevelUsageTypesResponse struct {
	Data                 []Serviceusagetype `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _GetServiceLevelUsageTypesResponse GetServiceLevelUsageTypesResponse

// NewGetServiceLevelUsageTypesResponse instantiates a new GetServiceLevelUsageTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceLevelUsageTypesResponse() *GetServiceLevelUsageTypesResponse {
	this := GetServiceLevelUsageTypesResponse{}
	return &this
}

// NewGetServiceLevelUsageTypesResponseWithDefaults instantiates a new GetServiceLevelUsageTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceLevelUsageTypesResponseWithDefaults() *GetServiceLevelUsageTypesResponse {
	this := GetServiceLevelUsageTypesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetServiceLevelUsageTypesResponse) GetData() []Serviceusagetype {
	if o == nil || o.Data == nil {
		var ret []Serviceusagetype
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceLevelUsageTypesResponse) GetDataOk() ([]Serviceusagetype, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetServiceLevelUsageTypesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Serviceusagetype and assigns it to the Data field.
func (o *GetServiceLevelUsageTypesResponse) SetData(v []Serviceusagetype) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o GetServiceLevelUsageTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *GetServiceLevelUsageTypesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetServiceLevelUsageTypesResponse := _GetServiceLevelUsageTypesResponse{}

	if err = json.Unmarshal(bytes, &varGetServiceLevelUsageTypesResponse); err == nil {
		*o = GetServiceLevelUsageTypesResponse(varGetServiceLevelUsageTypesResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableGetServiceLevelUsageTypesResponse is a helper abstraction for handling nullable getservicelevelusagetypesresponse types.
type NullableGetServiceLevelUsageTypesResponse struct {
	value *GetServiceLevelUsageTypesResponse
	isSet bool
}

// Get returns the value.
func (v NullableGetServiceLevelUsageTypesResponse) Get() *GetServiceLevelUsageTypesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableGetServiceLevelUsageTypesResponse) Set(val *GetServiceLevelUsageTypesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableGetServiceLevelUsageTypesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableGetServiceLevelUsageTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGetServiceLevelUsageTypesResponse returns a pointer to a new instance of NullableGetServiceLevelUsageTypesResponse.
func NewNullableGetServiceLevelUsageTypesResponse(val *GetServiceLevelUsageTypesResponse) *NullableGetServiceLevelUsageTypesResponse {
	return &NullableGetServiceLevelUsageTypesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableGetServiceLevelUsageTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableGetServiceLevelUsageTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
