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

// PathChange Modifications to an existing path between versions.
type PathChange struct {
	// Alphanumeric string identifying the path. Stable across versions of the routing config.
	PathId *string `json:"path_id,omitempty"`
	// The current path pattern.
	Path *string `json:"path,omitempty"`
	// The previous path pattern, if it changed.
	OldPath *string `json:"old_path,omitempty"`
	// Rules that were added to this path.
	RulesAdded []RuleResponse `json:"rules_added,omitempty"`
	// Rules that were modified on this path.
	RulesChanged []RuleChange `json:"rules_changed,omitempty"`
	// Rules that were removed from this path.
	RulesDeleted         []RuleResponse `json:"rules_deleted,omitempty"`
	AdditionalProperties map[string]any
}

type _PathChange PathChange

// NewPathChange instantiates a new PathChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathChange() *PathChange {
	this := PathChange{}
	return &this
}

// NewPathChangeWithDefaults instantiates a new PathChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathChangeWithDefaults() *PathChange {
	this := PathChange{}
	return &this
}

// GetPathId returns the PathId field value if set, zero value otherwise.
func (o *PathChange) GetPathId() string {
	if o == nil || o.PathId == nil {
		var ret string
		return ret
	}
	return *o.PathId
}

// GetPathIdOk returns a tuple with the PathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetPathIdOk() (*string, bool) {
	if o == nil || o.PathId == nil {
		return nil, false
	}
	return o.PathId, true
}

// HasPathId returns a boolean if a field has been set.
func (o *PathChange) HasPathId() bool {
	if o != nil && o.PathId != nil {
		return true
	}

	return false
}

// SetPathId gets a reference to the given string and assigns it to the PathId field.
func (o *PathChange) SetPathId(v string) {
	o.PathId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathChange) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathChange) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathChange) SetPath(v string) {
	o.Path = &v
}

// GetOldPath returns the OldPath field value if set, zero value otherwise.
func (o *PathChange) GetOldPath() string {
	if o == nil || o.OldPath == nil {
		var ret string
		return ret
	}
	return *o.OldPath
}

// GetOldPathOk returns a tuple with the OldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetOldPathOk() (*string, bool) {
	if o == nil || o.OldPath == nil {
		return nil, false
	}
	return o.OldPath, true
}

// HasOldPath returns a boolean if a field has been set.
func (o *PathChange) HasOldPath() bool {
	if o != nil && o.OldPath != nil {
		return true
	}

	return false
}

// SetOldPath gets a reference to the given string and assigns it to the OldPath field.
func (o *PathChange) SetOldPath(v string) {
	o.OldPath = &v
}

// GetRulesAdded returns the RulesAdded field value if set, zero value otherwise.
func (o *PathChange) GetRulesAdded() []RuleResponse {
	if o == nil || o.RulesAdded == nil {
		var ret []RuleResponse
		return ret
	}
	return o.RulesAdded
}

// GetRulesAddedOk returns a tuple with the RulesAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetRulesAddedOk() ([]RuleResponse, bool) {
	if o == nil || o.RulesAdded == nil {
		return nil, false
	}
	return o.RulesAdded, true
}

// HasRulesAdded returns a boolean if a field has been set.
func (o *PathChange) HasRulesAdded() bool {
	if o != nil && o.RulesAdded != nil {
		return true
	}

	return false
}

// SetRulesAdded gets a reference to the given []RuleResponse and assigns it to the RulesAdded field.
func (o *PathChange) SetRulesAdded(v []RuleResponse) {
	o.RulesAdded = v
}

// GetRulesChanged returns the RulesChanged field value if set, zero value otherwise.
func (o *PathChange) GetRulesChanged() []RuleChange {
	if o == nil || o.RulesChanged == nil {
		var ret []RuleChange
		return ret
	}
	return o.RulesChanged
}

// GetRulesChangedOk returns a tuple with the RulesChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetRulesChangedOk() ([]RuleChange, bool) {
	if o == nil || o.RulesChanged == nil {
		return nil, false
	}
	return o.RulesChanged, true
}

// HasRulesChanged returns a boolean if a field has been set.
func (o *PathChange) HasRulesChanged() bool {
	if o != nil && o.RulesChanged != nil {
		return true
	}

	return false
}

// SetRulesChanged gets a reference to the given []RuleChange and assigns it to the RulesChanged field.
func (o *PathChange) SetRulesChanged(v []RuleChange) {
	o.RulesChanged = v
}

// GetRulesDeleted returns the RulesDeleted field value if set, zero value otherwise.
func (o *PathChange) GetRulesDeleted() []RuleResponse {
	if o == nil || o.RulesDeleted == nil {
		var ret []RuleResponse
		return ret
	}
	return o.RulesDeleted
}

// GetRulesDeletedOk returns a tuple with the RulesDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathChange) GetRulesDeletedOk() ([]RuleResponse, bool) {
	if o == nil || o.RulesDeleted == nil {
		return nil, false
	}
	return o.RulesDeleted, true
}

// HasRulesDeleted returns a boolean if a field has been set.
func (o *PathChange) HasRulesDeleted() bool {
	if o != nil && o.RulesDeleted != nil {
		return true
	}

	return false
}

// SetRulesDeleted gets a reference to the given []RuleResponse and assigns it to the RulesDeleted field.
func (o *PathChange) SetRulesDeleted(v []RuleResponse) {
	o.RulesDeleted = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PathChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.PathId != nil {
		toSerialize["path_id"] = o.PathId
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.OldPath != nil {
		toSerialize["old_path"] = o.OldPath
	}
	if o.RulesAdded != nil {
		toSerialize["rules_added"] = o.RulesAdded
	}
	if o.RulesChanged != nil {
		toSerialize["rules_changed"] = o.RulesChanged
	}
	if o.RulesDeleted != nil {
		toSerialize["rules_deleted"] = o.RulesDeleted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PathChange) UnmarshalJSON(bytes []byte) (err error) {
	varPathChange := _PathChange{}

	if err = json.Unmarshal(bytes, &varPathChange); err == nil {
		*o = PathChange(varPathChange)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path_id")
		delete(additionalProperties, "path")
		delete(additionalProperties, "old_path")
		delete(additionalProperties, "rules_added")
		delete(additionalProperties, "rules_changed")
		delete(additionalProperties, "rules_deleted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePathChange is a helper abstraction for handling nullable pathchange types.
type NullablePathChange struct {
	value *PathChange
	isSet bool
}

// Get returns the value.
func (v NullablePathChange) Get() *PathChange {
	return v.value
}

// Set modifies the value.
func (v *NullablePathChange) Set(val *PathChange) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePathChange) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePathChange) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePathChange returns a pointer to a new instance of NullablePathChange.
func NewNullablePathChange(val *PathChange) *NullablePathChange {
	return &NullablePathChange{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePathChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePathChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
