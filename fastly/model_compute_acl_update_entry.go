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

// ComputeACLUpdateEntry Defines the structure of an ACL update request entry.
type ComputeACLUpdateEntry struct {
	// One of \"create\", \"update\", or \"delete\" indicating the operation to perform on the update.
	Op *string `json:"op,omitempty"`
	// An IP prefix defined in Classless Inter-Domain Routing (CIDR) format, i.e. a valid IP address (v4 or v6) followed by a forward slash (/) and a prefix length (0-32 or 0-128, depending on address family).
	Prefix *string `json:"prefix,omitempty"`
	// The action taken on the IP address, one of \"BLOCK\" or \"ALLOW\". If using the \"delete\" operation, no action should be specified.
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeACLUpdateEntry ComputeACLUpdateEntry

// NewComputeACLUpdateEntry instantiates a new ComputeACLUpdateEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeACLUpdateEntry() *ComputeACLUpdateEntry {
	this := ComputeACLUpdateEntry{}
	return &this
}

// NewComputeACLUpdateEntryWithDefaults instantiates a new ComputeACLUpdateEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeACLUpdateEntryWithDefaults() *ComputeACLUpdateEntry {
	this := ComputeACLUpdateEntry{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *ComputeACLUpdateEntry) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLUpdateEntry) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *ComputeACLUpdateEntry) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *ComputeACLUpdateEntry) SetOp(v string) {
	o.Op = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ComputeACLUpdateEntry) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLUpdateEntry) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ComputeACLUpdateEntry) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ComputeACLUpdateEntry) SetPrefix(v string) {
	o.Prefix = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ComputeACLUpdateEntry) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLUpdateEntry) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ComputeACLUpdateEntry) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ComputeACLUpdateEntry) SetAction(v string) {
	o.Action = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeACLUpdateEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeACLUpdateEntry) UnmarshalJSON(bytes []byte) (err error) {
	varComputeACLUpdateEntry := _ComputeACLUpdateEntry{}

	if err = json.Unmarshal(bytes, &varComputeACLUpdateEntry); err == nil {
		*o = ComputeACLUpdateEntry(varComputeACLUpdateEntry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeACLUpdateEntry is a helper abstraction for handling nullable computeaclupdateentry types.
type NullableComputeACLUpdateEntry struct {
	value *ComputeACLUpdateEntry
	isSet bool
}

// Get returns the value.
func (v NullableComputeACLUpdateEntry) Get() *ComputeACLUpdateEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeACLUpdateEntry) Set(val *ComputeACLUpdateEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeACLUpdateEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeACLUpdateEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeACLUpdateEntry returns a pointer to a new instance of NullableComputeACLUpdateEntry.
func NewNullableComputeACLUpdateEntry(val *ComputeACLUpdateEntry) *NullableComputeACLUpdateEntry {
	return &NullableComputeACLUpdateEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeACLUpdateEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeACLUpdateEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
