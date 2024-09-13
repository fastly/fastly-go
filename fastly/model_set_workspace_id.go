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

// SetWorkspaceID struct for SetWorkspaceID
type SetWorkspaceID struct {
	// The workspace to link with the Next-Gen WAF product. Note this body is only supported and required when `product_id` is `ngwaf`
	WorkspaceID *string `json:"workspace_id,omitempty"`
	AdditionalProperties map[string]any
}

type _SetWorkspaceID SetWorkspaceID

// NewSetWorkspaceID instantiates a new SetWorkspaceID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetWorkspaceID() *SetWorkspaceID {
	this := SetWorkspaceID{}
	return &this
}

// NewSetWorkspaceIDWithDefaults instantiates a new SetWorkspaceID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetWorkspaceIDWithDefaults() *SetWorkspaceID {
	this := SetWorkspaceID{}
	return &this
}

// GetWorkspaceID returns the WorkspaceID field value if set, zero value otherwise.
func (o *SetWorkspaceID) GetWorkspaceID() string {
	if o == nil || o.WorkspaceID == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceID
}

// GetWorkspaceIDOk returns a tuple with the WorkspaceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWorkspaceID) GetWorkspaceIDOk() (*string, bool) {
	if o == nil || o.WorkspaceID == nil {
		return nil, false
	}
	return o.WorkspaceID, true
}

// HasWorkspaceID returns a boolean if a field has been set.
func (o *SetWorkspaceID) HasWorkspaceID() bool {
	if o != nil && o.WorkspaceID != nil {
		return true
	}

	return false
}

// SetWorkspaceID gets a reference to the given string and assigns it to the WorkspaceID field.
func (o *SetWorkspaceID) SetWorkspaceID(v string) {
	o.WorkspaceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SetWorkspaceID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WorkspaceID != nil {
		toSerialize["workspace_id"] = o.WorkspaceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SetWorkspaceID) UnmarshalJSON(bytes []byte) (err error) {
	varSetWorkspaceID := _SetWorkspaceID{}

	if err = json.Unmarshal(bytes, &varSetWorkspaceID); err == nil {
		*o = SetWorkspaceID(varSetWorkspaceID)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workspace_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSetWorkspaceID is a helper abstraction for handling nullable setworkspaceid types. 
type NullableSetWorkspaceID struct {
	value *SetWorkspaceID
	isSet bool
}

// Get returns the value.
func (v NullableSetWorkspaceID) Get() *SetWorkspaceID {
	return v.value
}

// Set modifies the value.
func (v *NullableSetWorkspaceID) Set(val *SetWorkspaceID) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSetWorkspaceID) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSetWorkspaceID) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSetWorkspaceID returns a pointer to a new instance of NullableSetWorkspaceID.
func NewNullableSetWorkspaceID(val *SetWorkspaceID) *NullableSetWorkspaceID {
	return &NullableSetWorkspaceID{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSetWorkspaceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSetWorkspaceID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
