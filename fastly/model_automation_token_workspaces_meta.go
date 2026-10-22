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

// AutomationTokenWorkspacesMeta struct for AutomationTokenWorkspacesMeta
type AutomationTokenWorkspacesMeta struct {
	// Number of results requested.
	Limit int32 `json:"limit"`
	// Cursor for the next page. Empty string when there are no more results.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Cursor for the previous page. Empty string when on the first page.
	PreviousCursor       *string `json:"previous_cursor,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenWorkspacesMeta AutomationTokenWorkspacesMeta

// NewAutomationTokenWorkspacesMeta instantiates a new AutomationTokenWorkspacesMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenWorkspacesMeta(limit int32) *AutomationTokenWorkspacesMeta {
	this := AutomationTokenWorkspacesMeta{}
	this.Limit = limit
	return &this
}

// NewAutomationTokenWorkspacesMetaWithDefaults instantiates a new AutomationTokenWorkspacesMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenWorkspacesMetaWithDefaults() *AutomationTokenWorkspacesMeta {
	this := AutomationTokenWorkspacesMeta{}
	return &this
}

// GetLimit returns the Limit field value
func (o *AutomationTokenWorkspacesMeta) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesMeta) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *AutomationTokenWorkspacesMeta) SetLimit(v int32) {
	o.Limit = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *AutomationTokenWorkspacesMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *AutomationTokenWorkspacesMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *AutomationTokenWorkspacesMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise.
func (o *AutomationTokenWorkspacesMeta) GetPreviousCursor() string {
	if o == nil || o.PreviousCursor == nil {
		var ret string
		return ret
	}
	return *o.PreviousCursor
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesMeta) GetPreviousCursorOk() (*string, bool) {
	if o == nil || o.PreviousCursor == nil {
		return nil, false
	}
	return o.PreviousCursor, true
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *AutomationTokenWorkspacesMeta) HasPreviousCursor() bool {
	if o != nil && o.PreviousCursor != nil {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given string and assigns it to the PreviousCursor field.
func (o *AutomationTokenWorkspacesMeta) SetPreviousCursor(v string) {
	o.PreviousCursor = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenWorkspacesMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.PreviousCursor != nil {
		toSerialize["previous_cursor"] = o.PreviousCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenWorkspacesMeta) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenWorkspacesMeta := _AutomationTokenWorkspacesMeta{}

	if err = json.Unmarshal(bytes, &varAutomationTokenWorkspacesMeta); err == nil {
		*o = AutomationTokenWorkspacesMeta(varAutomationTokenWorkspacesMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "previous_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenWorkspacesMeta is a helper abstraction for handling nullable automationtokenworkspacesmeta types.
type NullableAutomationTokenWorkspacesMeta struct {
	value *AutomationTokenWorkspacesMeta
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenWorkspacesMeta) Get() *AutomationTokenWorkspacesMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenWorkspacesMeta) Set(val *AutomationTokenWorkspacesMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenWorkspacesMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenWorkspacesMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenWorkspacesMeta returns a pointer to a new instance of NullableAutomationTokenWorkspacesMeta.
func NewNullableAutomationTokenWorkspacesMeta(val *AutomationTokenWorkspacesMeta) *NullableAutomationTokenWorkspacesMeta {
	return &NullableAutomationTokenWorkspacesMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenWorkspacesMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenWorkspacesMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
