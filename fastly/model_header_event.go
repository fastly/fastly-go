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
	"time"
)

// HeaderEvent struct for HeaderEvent
type HeaderEvent struct {
	Id                   *string    `json:"id,omitempty"`
	HeaderName           *string    `json:"header_name,omitempty"`
	OldValue             *string    `json:"old_value,omitempty"`
	NewValue             *string    `json:"new_value,omitempty"`
	ChangedAt            *time.Time `json:"changed_at,omitempty"`
	AdditionalProperties map[string]any
}

type _HeaderEvent HeaderEvent

// NewHeaderEvent instantiates a new HeaderEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderEvent() *HeaderEvent {
	this := HeaderEvent{}
	return &this
}

// NewHeaderEventWithDefaults instantiates a new HeaderEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderEventWithDefaults() *HeaderEvent {
	this := HeaderEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HeaderEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HeaderEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HeaderEvent) SetId(v string) {
	o.Id = &v
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise.
func (o *HeaderEvent) GetHeaderName() string {
	if o == nil || o.HeaderName == nil {
		var ret string
		return ret
	}
	return *o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderEvent) GetHeaderNameOk() (*string, bool) {
	if o == nil || o.HeaderName == nil {
		return nil, false
	}
	return o.HeaderName, true
}

// HasHeaderName returns a boolean if a field has been set.
func (o *HeaderEvent) HasHeaderName() bool {
	if o != nil && o.HeaderName != nil {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given string and assigns it to the HeaderName field.
func (o *HeaderEvent) SetHeaderName(v string) {
	o.HeaderName = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *HeaderEvent) GetOldValue() string {
	if o == nil || o.OldValue == nil {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderEvent) GetOldValueOk() (*string, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *HeaderEvent) HasOldValue() bool {
	if o != nil && o.OldValue != nil {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *HeaderEvent) SetOldValue(v string) {
	o.OldValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *HeaderEvent) GetNewValue() string {
	if o == nil || o.NewValue == nil {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderEvent) GetNewValueOk() (*string, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *HeaderEvent) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *HeaderEvent) SetNewValue(v string) {
	o.NewValue = &v
}

// GetChangedAt returns the ChangedAt field value if set, zero value otherwise.
func (o *HeaderEvent) GetChangedAt() time.Time {
	if o == nil || o.ChangedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ChangedAt
}

// GetChangedAtOk returns a tuple with the ChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderEvent) GetChangedAtOk() (*time.Time, bool) {
	if o == nil || o.ChangedAt == nil {
		return nil, false
	}
	return o.ChangedAt, true
}

// HasChangedAt returns a boolean if a field has been set.
func (o *HeaderEvent) HasChangedAt() bool {
	if o != nil && o.ChangedAt != nil {
		return true
	}

	return false
}

// SetChangedAt gets a reference to the given time.Time and assigns it to the ChangedAt field.
func (o *HeaderEvent) SetChangedAt(v time.Time) {
	o.ChangedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HeaderEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.HeaderName != nil {
		toSerialize["header_name"] = o.HeaderName
	}
	if o.OldValue != nil {
		toSerialize["old_value"] = o.OldValue
	}
	if o.NewValue != nil {
		toSerialize["new_value"] = o.NewValue
	}
	if o.ChangedAt != nil {
		toSerialize["changed_at"] = o.ChangedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HeaderEvent) UnmarshalJSON(bytes []byte) (err error) {
	varHeaderEvent := _HeaderEvent{}

	if err = json.Unmarshal(bytes, &varHeaderEvent); err == nil {
		*o = HeaderEvent(varHeaderEvent)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "header_name")
		delete(additionalProperties, "old_value")
		delete(additionalProperties, "new_value")
		delete(additionalProperties, "changed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHeaderEvent is a helper abstraction for handling nullable headerevent types.
type NullableHeaderEvent struct {
	value *HeaderEvent
	isSet bool
}

// Get returns the value.
func (v NullableHeaderEvent) Get() *HeaderEvent {
	return v.value
}

// Set modifies the value.
func (v *NullableHeaderEvent) Set(val *HeaderEvent) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHeaderEvent) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHeaderEvent) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHeaderEvent returns a pointer to a new instance of NullableHeaderEvent.
func NewNullableHeaderEvent(val *HeaderEvent) *NullableHeaderEvent {
	return &NullableHeaderEvent{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHeaderEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHeaderEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
