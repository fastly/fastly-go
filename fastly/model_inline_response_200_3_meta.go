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

// InlineResponse2003Meta Meta for the pagination.
type InlineResponse2003Meta struct {
	// Cursor for the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Entries returned.
	Limit                *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2003Meta InlineResponse2003Meta

// NewInlineResponse2003Meta instantiates a new InlineResponse2003Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003Meta() *InlineResponse2003Meta {
	this := InlineResponse2003Meta{}
	return &this
}

// NewInlineResponse2003MetaWithDefaults instantiates a new InlineResponse2003Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003MetaWithDefaults() *InlineResponse2003Meta {
	this := InlineResponse2003Meta{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *InlineResponse2003Meta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Meta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *InlineResponse2003Meta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *InlineResponse2003Meta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *InlineResponse2003Meta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Meta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *InlineResponse2003Meta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *InlineResponse2003Meta) SetLimit(v int32) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2003Meta) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2003Meta) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2003Meta := _InlineResponse2003Meta{}

	if err = json.Unmarshal(bytes, &varInlineResponse2003Meta); err == nil {
		*o = InlineResponse2003Meta(varInlineResponse2003Meta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2003Meta is a helper abstraction for handling nullable inlineresponse2003meta types.
type NullableInlineResponse2003Meta struct {
	value *InlineResponse2003Meta
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2003Meta) Get() *InlineResponse2003Meta {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2003Meta) Set(val *InlineResponse2003Meta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2003Meta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2003Meta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2003Meta returns a pointer to a new instance of NullableInlineResponse2003Meta.
func NewNullableInlineResponse2003Meta(val *InlineResponse2003Meta) *NullableInlineResponse2003Meta {
	return &NullableInlineResponse2003Meta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2003Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2003Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
