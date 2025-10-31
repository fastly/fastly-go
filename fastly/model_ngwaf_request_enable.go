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

// NgwafRequestEnable struct for NgwafRequestEnable
type NgwafRequestEnable struct {
	// The workspace to link.
	WorkspaceId          string `json:"workspace_id"`
	AdditionalProperties map[string]any
}

type _NgwafRequestEnable NgwafRequestEnable

// NewNgwafRequestEnable instantiates a new NgwafRequestEnable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgwafRequestEnable(workspaceId string) *NgwafRequestEnable {
	this := NgwafRequestEnable{}
	this.WorkspaceId = workspaceId
	return &this
}

// NewNgwafRequestEnableWithDefaults instantiates a new NgwafRequestEnable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgwafRequestEnableWithDefaults() *NgwafRequestEnable {
	this := NgwafRequestEnable{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *NgwafRequestEnable) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *NgwafRequestEnable) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *NgwafRequestEnable) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NgwafRequestEnable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NgwafRequestEnable) UnmarshalJSON(bytes []byte) (err error) {
	varNgwafRequestEnable := _NgwafRequestEnable{}

	if err = json.Unmarshal(bytes, &varNgwafRequestEnable); err == nil {
		*o = NgwafRequestEnable(varNgwafRequestEnable)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workspace_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNgwafRequestEnable is a helper abstraction for handling nullable ngwafrequestenable types.
type NullableNgwafRequestEnable struct {
	value *NgwafRequestEnable
	isSet bool
}

// Get returns the value.
func (v NullableNgwafRequestEnable) Get() *NgwafRequestEnable {
	return v.value
}

// Set modifies the value.
func (v *NullableNgwafRequestEnable) Set(val *NgwafRequestEnable) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNgwafRequestEnable) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNgwafRequestEnable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNgwafRequestEnable returns a pointer to a new instance of NullableNgwafRequestEnable.
func NewNullableNgwafRequestEnable(val *NgwafRequestEnable) *NullableNgwafRequestEnable {
	return &NullableNgwafRequestEnable{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNgwafRequestEnable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNgwafRequestEnable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
