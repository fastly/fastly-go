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

// DraftUpdate All attributes for updating the draft version.
type DraftUpdate struct {
	// A freeform comment for the draft version.
	Comment              *string `json:"comment,omitempty"`
	AdditionalProperties map[string]any
}

type _DraftUpdate DraftUpdate

// NewDraftUpdate instantiates a new DraftUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftUpdate() *DraftUpdate {
	this := DraftUpdate{}
	return &this
}

// NewDraftUpdateWithDefaults instantiates a new DraftUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftUpdateWithDefaults() *DraftUpdate {
	this := DraftUpdate{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DraftUpdate) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftUpdate) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DraftUpdate) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DraftUpdate) SetComment(v string) {
	o.Comment = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DraftUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DraftUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varDraftUpdate := _DraftUpdate{}

	if err = json.Unmarshal(bytes, &varDraftUpdate); err == nil {
		*o = DraftUpdate(varDraftUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDraftUpdate is a helper abstraction for handling nullable draftupdate types.
type NullableDraftUpdate struct {
	value *DraftUpdate
	isSet bool
}

// Get returns the value.
func (v NullableDraftUpdate) Get() *DraftUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableDraftUpdate) Set(val *DraftUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDraftUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDraftUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDraftUpdate returns a pointer to a new instance of NullableDraftUpdate.
func NewNullableDraftUpdate(val *DraftUpdate) *NullableDraftUpdate {
	return &NullableDraftUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDraftUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDraftUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
