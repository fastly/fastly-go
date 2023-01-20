// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// GetStoresResponseMeta Meta for the pagination.
type GetStoresResponseMeta struct {
	// Cursor for the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Entries returned.
	Limit *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _GetStoresResponseMeta GetStoresResponseMeta

// NewGetStoresResponseMeta instantiates a new GetStoresResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoresResponseMeta() *GetStoresResponseMeta {
	this := GetStoresResponseMeta{}
	return &this
}

// NewGetStoresResponseMetaWithDefaults instantiates a new GetStoresResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoresResponseMetaWithDefaults() *GetStoresResponseMeta {
	this := GetStoresResponseMeta{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *GetStoresResponseMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoresResponseMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *GetStoresResponseMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *GetStoresResponseMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetStoresResponseMeta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoresResponseMeta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetStoresResponseMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetStoresResponseMeta) SetLimit(v int32) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o GetStoresResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *GetStoresResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	varGetStoresResponseMeta := _GetStoresResponseMeta{}

	if err = json.Unmarshal(bytes, &varGetStoresResponseMeta); err == nil {
		*o = GetStoresResponseMeta(varGetStoresResponseMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableGetStoresResponseMeta is a helper abstraction for handling nullable getstoresresponsemeta types. 
type NullableGetStoresResponseMeta struct {
	value *GetStoresResponseMeta
	isSet bool
}

// Get returns the value.
func (v NullableGetStoresResponseMeta) Get() *GetStoresResponseMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableGetStoresResponseMeta) Set(val *GetStoresResponseMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableGetStoresResponseMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableGetStoresResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGetStoresResponseMeta returns a pointer to a new instance of NullableGetStoresResponseMeta.
func NewNullableGetStoresResponseMeta(val *GetStoresResponseMeta) *NullableGetStoresResponseMeta {
	return &NullableGetStoresResponseMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableGetStoresResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableGetStoresResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
