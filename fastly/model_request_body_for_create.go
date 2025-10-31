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

// RequestBodyForCreate All attributes for creating a domain
type RequestBodyForCreate struct {
	// The fully-qualified domain name for your domain. Can be created, but not updated.
	Fqdn string `json:"fqdn"`
	// The `service_id` associated with your domain or `null` if there is no association.
	ServiceId NullableString `json:"service_id,omitempty"`
	// A freeform descriptive note.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestBodyForCreate RequestBodyForCreate

// NewRequestBodyForCreate instantiates a new RequestBodyForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBodyForCreate(fqdn string) *RequestBodyForCreate {
	this := RequestBodyForCreate{}
	this.Fqdn = fqdn
	return &this
}

// NewRequestBodyForCreateWithDefaults instantiates a new RequestBodyForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBodyForCreateWithDefaults() *RequestBodyForCreate {
	this := RequestBodyForCreate{}
	return &this
}

// GetFqdn returns the Fqdn field value
func (o *RequestBodyForCreate) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *RequestBodyForCreate) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *RequestBodyForCreate) SetFqdn(v string) {
	o.Fqdn = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestBodyForCreate) GetServiceId() string {
	if o == nil || o.ServiceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceId.Get()
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestBodyForCreate) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceId.Get(), o.ServiceId.IsSet()
}

// HasServiceId returns a boolean if a field has been set.
func (o *RequestBodyForCreate) HasServiceId() bool {
	if o != nil && o.ServiceId.IsSet() {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given NullableString and assigns it to the ServiceId field.
func (o *RequestBodyForCreate) SetServiceId(v string) {
	o.ServiceId.Set(&v)
}

// SetServiceIdNil sets the value for ServiceId to be an explicit nil
func (o *RequestBodyForCreate) SetServiceIdNil() {
	o.ServiceId.Set(nil)
}

// UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
func (o *RequestBodyForCreate) UnsetServiceId() {
	o.ServiceId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestBodyForCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBodyForCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestBodyForCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestBodyForCreate) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestBodyForCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.ServiceId.IsSet() {
		toSerialize["service_id"] = o.ServiceId.Get()
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
func (o *RequestBodyForCreate) UnmarshalJSON(bytes []byte) (err error) {
	varRequestBodyForCreate := _RequestBodyForCreate{}

	if err = json.Unmarshal(bytes, &varRequestBodyForCreate); err == nil {
		*o = RequestBodyForCreate(varRequestBodyForCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRequestBodyForCreate is a helper abstraction for handling nullable requestbodyforcreate types.
type NullableRequestBodyForCreate struct {
	value *RequestBodyForCreate
	isSet bool
}

// Get returns the value.
func (v NullableRequestBodyForCreate) Get() *RequestBodyForCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestBodyForCreate) Set(val *RequestBodyForCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestBodyForCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestBodyForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestBodyForCreate returns a pointer to a new instance of NullableRequestBodyForCreate.
func NewNullableRequestBodyForCreate(val *RequestBodyForCreate) *NullableRequestBodyForCreate {
	return &NullableRequestBodyForCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestBodyForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRequestBodyForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
