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

// AgentKey An agent key used for configuring a Next-Gen WAF agent.
type AgentKey struct {
	// Agent configuration access key value.
	AccessKey string `json:"access_key"`
	// Agent configuration secret key value.
	SecretKey string `json:"secret_key"`
	// Whether the agent key is the primary key that should be used to configure the agent.
	IsPrimary bool `json:"is_primary"`
	// Date and time the agent key was created.
	CreatedAt time.Time `json:"created_at"`
	// Date and time the agent key was last updated.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _AgentKey AgentKey

// NewAgentKey instantiates a new AgentKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentKey(accessKey string, secretKey string, isPrimary bool, createdAt time.Time, updatedAt time.Time) *AgentKey {
	this := AgentKey{}
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	this.IsPrimary = isPrimary
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAgentKeyWithDefaults instantiates a new AgentKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentKeyWithDefaults() *AgentKey {
	this := AgentKey{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AgentKey) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AgentKey) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *AgentKey) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *AgentKey) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *AgentKey) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *AgentKey) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetIsPrimary returns the IsPrimary field value
func (o *AgentKey) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *AgentKey) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *AgentKey) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AgentKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AgentKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AgentKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AgentKey) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AgentKey) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AgentKey) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AgentKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["access_key"] = o.AccessKey
	}
	if true {
		toSerialize["secret_key"] = o.SecretKey
	}
	if true {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AgentKey) UnmarshalJSON(bytes []byte) (err error) {
	varAgentKey := _AgentKey{}

	if err = json.Unmarshal(bytes, &varAgentKey); err == nil {
		*o = AgentKey(varAgentKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "is_primary")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAgentKey is a helper abstraction for handling nullable agentkey types.
type NullableAgentKey struct {
	value *AgentKey
	isSet bool
}

// Get returns the value.
func (v NullableAgentKey) Get() *AgentKey {
	return v.value
}

// Set modifies the value.
func (v *NullableAgentKey) Set(val *AgentKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAgentKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAgentKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAgentKey returns a pointer to a new instance of NullableAgentKey.
func NewNullableAgentKey(val *AgentKey) *NullableAgentKey {
	return &NullableAgentKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAgentKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAgentKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
