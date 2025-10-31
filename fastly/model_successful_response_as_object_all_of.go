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

// SuccessfulResponseAsObjectAllOf struct for SuccessfulResponseAsObjectAllOf
type SuccessfulResponseAsObjectAllOf struct {
	// Domain Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// The fully-qualified domain name for your domain. Can be created, but not updated.
	Fqdn *string `json:"fqdn,omitempty"`
	// The `service_id` associated with your domain or `null` if there is no association.
	ServiceId NullableString `json:"service_id,omitempty"`
	// A freeform descriptive note.
	Description *string `json:"description,omitempty"`
	// Denotes if the domain has at least one TLS activation or not.
	Activated *bool `json:"activated,omitempty"`
	// Denotes that the customer has proven ownership over the domain by obtaining a Fastly-managed TLS certificate for it or by providing a their own TLS certificate from a publicly-trusted CA that includes the domain or matching wildcard.
	Verified             *bool `json:"verified,omitempty"`
	AdditionalProperties map[string]any
}

type _SuccessfulResponseAsObjectAllOf SuccessfulResponseAsObjectAllOf

// NewSuccessfulResponseAsObjectAllOf instantiates a new SuccessfulResponseAsObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulResponseAsObjectAllOf() *SuccessfulResponseAsObjectAllOf {
	this := SuccessfulResponseAsObjectAllOf{}
	return &this
}

// NewSuccessfulResponseAsObjectAllOfWithDefaults instantiates a new SuccessfulResponseAsObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulResponseAsObjectAllOfWithDefaults() *SuccessfulResponseAsObjectAllOf {
	this := SuccessfulResponseAsObjectAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SuccessfulResponseAsObjectAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulResponseAsObjectAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SuccessfulResponseAsObjectAllOf) SetId(v string) {
	o.Id = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *SuccessfulResponseAsObjectAllOf) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulResponseAsObjectAllOf) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *SuccessfulResponseAsObjectAllOf) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SuccessfulResponseAsObjectAllOf) GetServiceId() string {
	if o == nil || o.ServiceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceId.Get()
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SuccessfulResponseAsObjectAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceId.Get(), o.ServiceId.IsSet()
}

// HasServiceId returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId.IsSet() {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given NullableString and assigns it to the ServiceId field.
func (o *SuccessfulResponseAsObjectAllOf) SetServiceId(v string) {
	o.ServiceId.Set(&v)
}

// SetServiceIdNil sets the value for ServiceId to be an explicit nil
func (o *SuccessfulResponseAsObjectAllOf) SetServiceIdNil() {
	o.ServiceId.Set(nil)
}

// UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
func (o *SuccessfulResponseAsObjectAllOf) UnsetServiceId() {
	o.ServiceId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SuccessfulResponseAsObjectAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulResponseAsObjectAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SuccessfulResponseAsObjectAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *SuccessfulResponseAsObjectAllOf) GetActivated() bool {
	if o == nil || o.Activated == nil {
		var ret bool
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulResponseAsObjectAllOf) GetActivatedOk() (*bool, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given bool and assigns it to the Activated field.
func (o *SuccessfulResponseAsObjectAllOf) SetActivated(v bool) {
	o.Activated = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *SuccessfulResponseAsObjectAllOf) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulResponseAsObjectAllOf) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *SuccessfulResponseAsObjectAllOf) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *SuccessfulResponseAsObjectAllOf) SetVerified(v bool) {
	o.Verified = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SuccessfulResponseAsObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.ServiceId.IsSet() {
		toSerialize["service_id"] = o.ServiceId.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SuccessfulResponseAsObjectAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSuccessfulResponseAsObjectAllOf := _SuccessfulResponseAsObjectAllOf{}

	if err = json.Unmarshal(bytes, &varSuccessfulResponseAsObjectAllOf); err == nil {
		*o = SuccessfulResponseAsObjectAllOf(varSuccessfulResponseAsObjectAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "activated")
		delete(additionalProperties, "verified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSuccessfulResponseAsObjectAllOf is a helper abstraction for handling nullable successfulresponseasobjectallof types.
type NullableSuccessfulResponseAsObjectAllOf struct {
	value *SuccessfulResponseAsObjectAllOf
	isSet bool
}

// Get returns the value.
func (v NullableSuccessfulResponseAsObjectAllOf) Get() *SuccessfulResponseAsObjectAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableSuccessfulResponseAsObjectAllOf) Set(val *SuccessfulResponseAsObjectAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSuccessfulResponseAsObjectAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSuccessfulResponseAsObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSuccessfulResponseAsObjectAllOf returns a pointer to a new instance of NullableSuccessfulResponseAsObjectAllOf.
func NewNullableSuccessfulResponseAsObjectAllOf(val *SuccessfulResponseAsObjectAllOf) *NullableSuccessfulResponseAsObjectAllOf {
	return &NullableSuccessfulResponseAsObjectAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSuccessfulResponseAsObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSuccessfulResponseAsObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
