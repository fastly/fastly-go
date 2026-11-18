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

// ServiceResponseAllOf struct for ServiceResponseAllOf
type ServiceResponseAllOf struct {
	ID *string `json:"id,omitempty"`
	// Unused at this time.
	PublishKey *string `json:"publish_key,omitempty"`
	// Whether the service is paused. Services are paused due to a lack of traffic for an extended period of time. Services are resumed either when a draft version is activated or a locked version is cloned and reactivated.
	Paused *bool `json:"paused,omitempty"`
	// A list of [versions](https://www.fastly.com/documentation/reference/api/services/version/) associated with the service.
	Versions []SchemasVersionResponse `json:"versions,omitempty"`
	// A list of environments where the service has been deployed.
	Environments         []Environment `json:"environments,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceResponseAllOf ServiceResponseAllOf

// NewServiceResponseAllOf instantiates a new ServiceResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResponseAllOf() *ServiceResponseAllOf {
	this := ServiceResponseAllOf{}
	return &this
}

// NewServiceResponseAllOfWithDefaults instantiates a new ServiceResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResponseAllOfWithDefaults() *ServiceResponseAllOf {
	this := ServiceResponseAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceResponseAllOf) SetID(v string) {
	o.ID = &v
}

// GetPublishKey returns the PublishKey field value if set, zero value otherwise.
func (o *ServiceResponseAllOf) GetPublishKey() string {
	if o == nil || o.PublishKey == nil {
		var ret string
		return ret
	}
	return *o.PublishKey
}

// GetPublishKeyOk returns a tuple with the PublishKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseAllOf) GetPublishKeyOk() (*string, bool) {
	if o == nil || o.PublishKey == nil {
		return nil, false
	}
	return o.PublishKey, true
}

// HasPublishKey returns a boolean if a field has been set.
func (o *ServiceResponseAllOf) HasPublishKey() bool {
	if o != nil && o.PublishKey != nil {
		return true
	}

	return false
}

// SetPublishKey gets a reference to the given string and assigns it to the PublishKey field.
func (o *ServiceResponseAllOf) SetPublishKey(v string) {
	o.PublishKey = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *ServiceResponseAllOf) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseAllOf) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *ServiceResponseAllOf) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *ServiceResponseAllOf) SetPaused(v bool) {
	o.Paused = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ServiceResponseAllOf) GetVersions() []SchemasVersionResponse {
	if o == nil || o.Versions == nil {
		var ret []SchemasVersionResponse
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseAllOf) GetVersionsOk() ([]SchemasVersionResponse, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ServiceResponseAllOf) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []SchemasVersionResponse and assigns it to the Versions field.
func (o *ServiceResponseAllOf) SetVersions(v []SchemasVersionResponse) {
	o.Versions = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ServiceResponseAllOf) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseAllOf) GetEnvironmentsOk() ([]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ServiceResponseAllOf) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *ServiceResponseAllOf) SetEnvironments(v []Environment) {
	o.Environments = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PublishKey != nil {
		toSerialize["publish_key"] = o.PublishKey
	}
	if o.Paused != nil {
		toSerialize["paused"] = o.Paused
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServiceResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceResponseAllOf := _ServiceResponseAllOf{}

	if err = json.Unmarshal(bytes, &varServiceResponseAllOf); err == nil {
		*o = ServiceResponseAllOf(varServiceResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "publish_key")
		delete(additionalProperties, "paused")
		delete(additionalProperties, "versions")
		delete(additionalProperties, "environments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceResponseAllOf is a helper abstraction for handling nullable serviceresponseallof types.
type NullableServiceResponseAllOf struct {
	value *ServiceResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableServiceResponseAllOf) Get() *ServiceResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceResponseAllOf) Set(val *ServiceResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceResponseAllOf returns a pointer to a new instance of NullableServiceResponseAllOf.
func NewNullableServiceResponseAllOf(val *ServiceResponseAllOf) *NullableServiceResponseAllOf {
	return &NullableServiceResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
