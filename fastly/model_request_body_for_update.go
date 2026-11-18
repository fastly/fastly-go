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

// RequestBodyForUpdate All attributes for updating a domain
type RequestBodyForUpdate struct {
	// The `service_id` associated with your domain or `null` if there is no association.
	ServiceID NullableString `json:"service_id,omitempty"`
	// The `routing_configuration_id` associated with your domain or `null` if there is no association.
	RoutingConfigurationID NullableString `json:"routing_configuration_id,omitempty"`
	// A freeform descriptive note.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestBodyForUpdate RequestBodyForUpdate

// NewRequestBodyForUpdate instantiates a new RequestBodyForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBodyForUpdate() *RequestBodyForUpdate {
	this := RequestBodyForUpdate{}
	return &this
}

// NewRequestBodyForUpdateWithDefaults instantiates a new RequestBodyForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBodyForUpdateWithDefaults() *RequestBodyForUpdate {
	this := RequestBodyForUpdate{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestBodyForUpdate) GetServiceID() string {
	if o == nil || o.ServiceID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceID.Get()
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestBodyForUpdate) GetServiceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceID.Get(), o.ServiceID.IsSet()
}

// HasServiceID returns a boolean if a field has been set.
func (o *RequestBodyForUpdate) HasServiceID() bool {
	if o != nil && o.ServiceID.IsSet() {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given NullableString and assigns it to the ServiceID field.
func (o *RequestBodyForUpdate) SetServiceID(v string) {
	o.ServiceID.Set(&v)
}

// SetServiceIDNil sets the value for ServiceID to be an explicit nil
func (o *RequestBodyForUpdate) SetServiceIDNil() {
	o.ServiceID.Set(nil)
}

// UnsetServiceID ensures that no value is present for ServiceID, not even an explicit nil
func (o *RequestBodyForUpdate) UnsetServiceID() {
	o.ServiceID.Unset()
}

// GetRoutingConfigurationID returns the RoutingConfigurationID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestBodyForUpdate) GetRoutingConfigurationID() string {
	if o == nil || o.RoutingConfigurationID.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoutingConfigurationID.Get()
}

// GetRoutingConfigurationIDOk returns a tuple with the RoutingConfigurationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestBodyForUpdate) GetRoutingConfigurationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoutingConfigurationID.Get(), o.RoutingConfigurationID.IsSet()
}

// HasRoutingConfigurationID returns a boolean if a field has been set.
func (o *RequestBodyForUpdate) HasRoutingConfigurationID() bool {
	if o != nil && o.RoutingConfigurationID.IsSet() {
		return true
	}

	return false
}

// SetRoutingConfigurationID gets a reference to the given NullableString and assigns it to the RoutingConfigurationID field.
func (o *RequestBodyForUpdate) SetRoutingConfigurationID(v string) {
	o.RoutingConfigurationID.Set(&v)
}

// SetRoutingConfigurationIDNil sets the value for RoutingConfigurationID to be an explicit nil
func (o *RequestBodyForUpdate) SetRoutingConfigurationIDNil() {
	o.RoutingConfigurationID.Set(nil)
}

// UnsetRoutingConfigurationID ensures that no value is present for RoutingConfigurationID, not even an explicit nil
func (o *RequestBodyForUpdate) UnsetRoutingConfigurationID() {
	o.RoutingConfigurationID.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestBodyForUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBodyForUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestBodyForUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestBodyForUpdate) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestBodyForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID.IsSet() {
		toSerialize["service_id"] = o.ServiceID.Get()
	}
	if o.RoutingConfigurationID.IsSet() {
		toSerialize["routing_configuration_id"] = o.RoutingConfigurationID.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RequestBodyForUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varRequestBodyForUpdate := _RequestBodyForUpdate{}

	if err = json.Unmarshal(bytes, &varRequestBodyForUpdate); err == nil {
		*o = RequestBodyForUpdate(varRequestBodyForUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "routing_configuration_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRequestBodyForUpdate is a helper abstraction for handling nullable requestbodyforupdate types.
type NullableRequestBodyForUpdate struct {
	value *RequestBodyForUpdate
	isSet bool
}

// Get returns the value.
func (v NullableRequestBodyForUpdate) Get() *RequestBodyForUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestBodyForUpdate) Set(val *RequestBodyForUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestBodyForUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestBodyForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestBodyForUpdate returns a pointer to a new instance of NullableRequestBodyForUpdate.
func NewNullableRequestBodyForUpdate(val *RequestBodyForUpdate) *NullableRequestBodyForUpdate {
	return &NullableRequestBodyForUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestBodyForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRequestBodyForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
