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

// DraftDiff The differences between the draft and active versions of a routing config.
type DraftDiff struct {
	// Paths that exist in the draft but not in the active version.
	Added []PathWithRules `json:"added,omitempty"`
	// Paths that exist in the active version but not in the draft.
	Deleted []PathWithRules `json:"deleted,omitempty"`
	// Paths that exist in both versions but have changed.
	Modified             []PathChange `json:"modified,omitempty"`
	AdditionalProperties map[string]any
}

type _DraftDiff DraftDiff

// NewDraftDiff instantiates a new DraftDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftDiff() *DraftDiff {
	this := DraftDiff{}
	return &this
}

// NewDraftDiffWithDefaults instantiates a new DraftDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftDiffWithDefaults() *DraftDiff {
	this := DraftDiff{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *DraftDiff) GetAdded() []PathWithRules {
	if o == nil || o.Added == nil {
		var ret []PathWithRules
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftDiff) GetAddedOk() ([]PathWithRules, bool) {
	if o == nil || o.Added == nil {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *DraftDiff) HasAdded() bool {
	if o != nil && o.Added != nil {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []PathWithRules and assigns it to the Added field.
func (o *DraftDiff) SetAdded(v []PathWithRules) {
	o.Added = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DraftDiff) GetDeleted() []PathWithRules {
	if o == nil || o.Deleted == nil {
		var ret []PathWithRules
		return ret
	}
	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftDiff) GetDeletedOk() ([]PathWithRules, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DraftDiff) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given []PathWithRules and assigns it to the Deleted field.
func (o *DraftDiff) SetDeleted(v []PathWithRules) {
	o.Deleted = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *DraftDiff) GetModified() []PathChange {
	if o == nil || o.Modified == nil {
		var ret []PathChange
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftDiff) GetModifiedOk() ([]PathChange, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *DraftDiff) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given []PathChange and assigns it to the Modified field.
func (o *DraftDiff) SetModified(v []PathChange) {
	o.Modified = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DraftDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Added != nil {
		toSerialize["added"] = o.Added
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DraftDiff) UnmarshalJSON(bytes []byte) (err error) {
	varDraftDiff := _DraftDiff{}

	if err = json.Unmarshal(bytes, &varDraftDiff); err == nil {
		*o = DraftDiff(varDraftDiff)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "added")
		delete(additionalProperties, "deleted")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDraftDiff is a helper abstraction for handling nullable draftdiff types.
type NullableDraftDiff struct {
	value *DraftDiff
	isSet bool
}

// Get returns the value.
func (v NullableDraftDiff) Get() *DraftDiff {
	return v.value
}

// Set modifies the value.
func (v *NullableDraftDiff) Set(val *DraftDiff) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDraftDiff) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDraftDiff) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDraftDiff returns a pointer to a new instance of NullableDraftDiff.
func NewNullableDraftDiff(val *DraftDiff) *NullableDraftDiff {
	return &NullableDraftDiff{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDraftDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDraftDiff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
