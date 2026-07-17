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

// RoutingConfigResponseAllOf struct for RoutingConfigResponseAllOf
type RoutingConfigResponseAllOf struct {
	// Alphanumeric string identifying the routing config.
	Id *string `json:"id,omitempty"`
	// The user-defined name for the routing config.
	Name  *string             `json:"name,omitempty"`
	State *RoutingConfigState `json:"state,omitempty"`
	// Timestamp of when the version was most recently activated. `null` if the version has never been activated.
	ActivatedAt NullableTime `json:"activated_at,omitempty"`
	// HATEOAS links to related resources.
	Links                *map[string]string `json:"links,omitempty"`
	AdditionalProperties map[string]any
}

type _RoutingConfigResponseAllOf RoutingConfigResponseAllOf

// NewRoutingConfigResponseAllOf instantiates a new RoutingConfigResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfigResponseAllOf() *RoutingConfigResponseAllOf {
	this := RoutingConfigResponseAllOf{}
	return &this
}

// NewRoutingConfigResponseAllOfWithDefaults instantiates a new RoutingConfigResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigResponseAllOfWithDefaults() *RoutingConfigResponseAllOf {
	this := RoutingConfigResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoutingConfigResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoutingConfigResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoutingConfigResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingConfigResponseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingConfigResponseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingConfigResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RoutingConfigResponseAllOf) GetState() RoutingConfigState {
	if o == nil || o.State == nil {
		var ret RoutingConfigState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponseAllOf) GetStateOk() (*RoutingConfigState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RoutingConfigResponseAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given RoutingConfigState and assigns it to the State field.
func (o *RoutingConfigResponseAllOf) SetState(v RoutingConfigState) {
	o.State = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigResponseAllOf) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt.Get()
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigResponseAllOf) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivatedAt.Get(), o.ActivatedAt.IsSet()
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *RoutingConfigResponseAllOf) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt.IsSet() {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given NullableTime and assigns it to the ActivatedAt field.
func (o *RoutingConfigResponseAllOf) SetActivatedAt(v time.Time) {
	o.ActivatedAt.Set(&v)
}

// SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil
func (o *RoutingConfigResponseAllOf) SetActivatedAtNil() {
	o.ActivatedAt.Set(nil)
}

// UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
func (o *RoutingConfigResponseAllOf) UnsetActivatedAt() {
	o.ActivatedAt.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoutingConfigResponseAllOf) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponseAllOf) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoutingConfigResponseAllOf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *RoutingConfigResponseAllOf) SetLinks(v map[string]string) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfigResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ActivatedAt.IsSet() {
		toSerialize["activated_at"] = o.ActivatedAt.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RoutingConfigResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfigResponseAllOf := _RoutingConfigResponseAllOf{}

	if err = json.Unmarshal(bytes, &varRoutingConfigResponseAllOf); err == nil {
		*o = RoutingConfigResponseAllOf(varRoutingConfigResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "activated_at")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfigResponseAllOf is a helper abstraction for handling nullable routingconfigresponseallof types.
type NullableRoutingConfigResponseAllOf struct {
	value *RoutingConfigResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigResponseAllOf) Get() *RoutingConfigResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigResponseAllOf) Set(val *RoutingConfigResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigResponseAllOf returns a pointer to a new instance of NullableRoutingConfigResponseAllOf.
func NewNullableRoutingConfigResponseAllOf(val *RoutingConfigResponseAllOf) *NullableRoutingConfigResponseAllOf {
	return &NullableRoutingConfigResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
