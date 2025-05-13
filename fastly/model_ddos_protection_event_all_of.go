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

// DdosProtectionEventAllOf struct for DdosProtectionEventAllOf
type DdosProtectionEventAllOf struct {
	// Unique ID of the event.
	ID *string `json:"id,omitempty"`
	// A human-readable name for the event.
	Name *string `json:"name,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceID *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	StartedAt NullableTime `json:"started_at,omitempty"`
	// Date and time in ISO 8601 format.
	EndedAt              NullableTime `json:"ended_at,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionEventAllOf DdosProtectionEventAllOf

// NewDdosProtectionEventAllOf instantiates a new DdosProtectionEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionEventAllOf() *DdosProtectionEventAllOf {
	this := DdosProtectionEventAllOf{}
	return &this
}

// NewDdosProtectionEventAllOfWithDefaults instantiates a new DdosProtectionEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionEventAllOfWithDefaults() *DdosProtectionEventAllOf {
	this := DdosProtectionEventAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *DdosProtectionEventAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionEventAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *DdosProtectionEventAllOf) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DdosProtectionEventAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionEventAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DdosProtectionEventAllOf) SetName(v string) {
	o.Name = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *DdosProtectionEventAllOf) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionEventAllOf) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *DdosProtectionEventAllOf) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DdosProtectionEventAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionEventAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DdosProtectionEventAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionEventAllOf) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionEventAllOf) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *DdosProtectionEventAllOf) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *DdosProtectionEventAllOf) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *DdosProtectionEventAllOf) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionEventAllOf) GetEndedAt() time.Time {
	if o == nil || o.EndedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionEventAllOf) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// HasEndedAt returns a boolean if a field has been set.
func (o *DdosProtectionEventAllOf) HasEndedAt() bool {
	if o != nil && o.EndedAt.IsSet() {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given NullableTime and assigns it to the EndedAt field.
func (o *DdosProtectionEventAllOf) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}

// SetEndedAtNil sets the value for EndedAt to be an explicit nil
func (o *DdosProtectionEventAllOf) SetEndedAtNil() {
	o.EndedAt.Set(nil)
}

// UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
func (o *DdosProtectionEventAllOf) UnsetEndedAt() {
	o.EndedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	if o.EndedAt.IsSet() {
		toSerialize["ended_at"] = o.EndedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionEventAllOf := _DdosProtectionEventAllOf{}

	if err = json.Unmarshal(bytes, &varDdosProtectionEventAllOf); err == nil {
		*o = DdosProtectionEventAllOf(varDdosProtectionEventAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "ended_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionEventAllOf is a helper abstraction for handling nullable ddosprotectioneventallof types.
type NullableDdosProtectionEventAllOf struct {
	value *DdosProtectionEventAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionEventAllOf) Get() *DdosProtectionEventAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionEventAllOf) Set(val *DdosProtectionEventAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionEventAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionEventAllOf returns a pointer to a new instance of NullableDdosProtectionEventAllOf.
func NewNullableDdosProtectionEventAllOf(val *DdosProtectionEventAllOf) *NullableDdosProtectionEventAllOf {
	return &NullableDdosProtectionEventAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
