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

// ScriptUpdate struct for ScriptUpdate
type ScriptUpdate struct {
	// Script authorization status
	AuthorizationStatus *string `json:"authorization_status,omitempty"`
	// Reason for authorization decision
	Justification *string `json:"justification,omitempty"`
	// Hash of authorized script content
	AuthorizedHash       *string `json:"authorized_hash,omitempty"`
	AdditionalProperties map[string]any
}

type _ScriptUpdate ScriptUpdate

// NewScriptUpdate instantiates a new ScriptUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptUpdate() *ScriptUpdate {
	this := ScriptUpdate{}
	return &this
}

// NewScriptUpdateWithDefaults instantiates a new ScriptUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptUpdateWithDefaults() *ScriptUpdate {
	this := ScriptUpdate{}
	return &this
}

// GetAuthorizationStatus returns the AuthorizationStatus field value if set, zero value otherwise.
func (o *ScriptUpdate) GetAuthorizationStatus() string {
	if o == nil || o.AuthorizationStatus == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationStatus
}

// GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptUpdate) GetAuthorizationStatusOk() (*string, bool) {
	if o == nil || o.AuthorizationStatus == nil {
		return nil, false
	}
	return o.AuthorizationStatus, true
}

// HasAuthorizationStatus returns a boolean if a field has been set.
func (o *ScriptUpdate) HasAuthorizationStatus() bool {
	if o != nil && o.AuthorizationStatus != nil {
		return true
	}

	return false
}

// SetAuthorizationStatus gets a reference to the given string and assigns it to the AuthorizationStatus field.
func (o *ScriptUpdate) SetAuthorizationStatus(v string) {
	o.AuthorizationStatus = &v
}

// GetJustification returns the Justification field value if set, zero value otherwise.
func (o *ScriptUpdate) GetJustification() string {
	if o == nil || o.Justification == nil {
		var ret string
		return ret
	}
	return *o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptUpdate) GetJustificationOk() (*string, bool) {
	if o == nil || o.Justification == nil {
		return nil, false
	}
	return o.Justification, true
}

// HasJustification returns a boolean if a field has been set.
func (o *ScriptUpdate) HasJustification() bool {
	if o != nil && o.Justification != nil {
		return true
	}

	return false
}

// SetJustification gets a reference to the given string and assigns it to the Justification field.
func (o *ScriptUpdate) SetJustification(v string) {
	o.Justification = &v
}

// GetAuthorizedHash returns the AuthorizedHash field value if set, zero value otherwise.
func (o *ScriptUpdate) GetAuthorizedHash() string {
	if o == nil || o.AuthorizedHash == nil {
		var ret string
		return ret
	}
	return *o.AuthorizedHash
}

// GetAuthorizedHashOk returns a tuple with the AuthorizedHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptUpdate) GetAuthorizedHashOk() (*string, bool) {
	if o == nil || o.AuthorizedHash == nil {
		return nil, false
	}
	return o.AuthorizedHash, true
}

// HasAuthorizedHash returns a boolean if a field has been set.
func (o *ScriptUpdate) HasAuthorizedHash() bool {
	if o != nil && o.AuthorizedHash != nil {
		return true
	}

	return false
}

// SetAuthorizedHash gets a reference to the given string and assigns it to the AuthorizedHash field.
func (o *ScriptUpdate) SetAuthorizedHash(v string) {
	o.AuthorizedHash = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ScriptUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AuthorizationStatus != nil {
		toSerialize["authorization_status"] = o.AuthorizationStatus
	}
	if o.Justification != nil {
		toSerialize["justification"] = o.Justification
	}
	if o.AuthorizedHash != nil {
		toSerialize["authorized_hash"] = o.AuthorizedHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ScriptUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varScriptUpdate := _ScriptUpdate{}

	if err = json.Unmarshal(bytes, &varScriptUpdate); err == nil {
		*o = ScriptUpdate(varScriptUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization_status")
		delete(additionalProperties, "justification")
		delete(additionalProperties, "authorized_hash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableScriptUpdate is a helper abstraction for handling nullable scriptupdate types.
type NullableScriptUpdate struct {
	value *ScriptUpdate
	isSet bool
}

// Get returns the value.
func (v NullableScriptUpdate) Get() *ScriptUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableScriptUpdate) Set(val *ScriptUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableScriptUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableScriptUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableScriptUpdate returns a pointer to a new instance of NullableScriptUpdate.
func NewNullableScriptUpdate(val *ScriptUpdate) *NullableScriptUpdate {
	return &NullableScriptUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableScriptUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableScriptUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
