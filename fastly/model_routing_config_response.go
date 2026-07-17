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

// RoutingConfigResponse All attributes for a routing config response.
type RoutingConfigResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
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

type _RoutingConfigResponse RoutingConfigResponse

// NewRoutingConfigResponse instantiates a new RoutingConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfigResponse() *RoutingConfigResponse {
	this := RoutingConfigResponse{}
	return &this
}

// NewRoutingConfigResponseWithDefaults instantiates a new RoutingConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigResponseWithDefaults() *RoutingConfigResponse {
	this := RoutingConfigResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *RoutingConfigResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RoutingConfigResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RoutingConfigResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *RoutingConfigResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *RoutingConfigResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *RoutingConfigResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoutingConfigResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoutingConfigResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingConfigResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingConfigResponse) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RoutingConfigResponse) GetState() RoutingConfigState {
	if o == nil || o.State == nil {
		var ret RoutingConfigState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponse) GetStateOk() (*RoutingConfigState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given RoutingConfigState and assigns it to the State field.
func (o *RoutingConfigResponse) SetState(v RoutingConfigState) {
	o.State = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigResponse) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt.Get()
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigResponse) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivatedAt.Get(), o.ActivatedAt.IsSet()
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt.IsSet() {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given NullableTime and assigns it to the ActivatedAt field.
func (o *RoutingConfigResponse) SetActivatedAt(v time.Time) {
	o.ActivatedAt.Set(&v)
}

// SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil
func (o *RoutingConfigResponse) SetActivatedAtNil() {
	o.ActivatedAt.Set(nil)
}

// UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
func (o *RoutingConfigResponse) UnsetActivatedAt() {
	o.ActivatedAt.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoutingConfigResponse) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigResponse) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoutingConfigResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *RoutingConfigResponse) SetLinks(v map[string]string) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
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
func (o *RoutingConfigResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfigResponse := _RoutingConfigResponse{}

	if err = json.Unmarshal(bytes, &varRoutingConfigResponse); err == nil {
		*o = RoutingConfigResponse(varRoutingConfigResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "activated_at")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfigResponse is a helper abstraction for handling nullable routingconfigresponse types.
type NullableRoutingConfigResponse struct {
	value *RoutingConfigResponse
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigResponse) Get() *RoutingConfigResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigResponse) Set(val *RoutingConfigResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigResponse returns a pointer to a new instance of NullableRoutingConfigResponse.
func NewNullableRoutingConfigResponse(val *RoutingConfigResponse) *NullableRoutingConfigResponse {
	return &NullableRoutingConfigResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
