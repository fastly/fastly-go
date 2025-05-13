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

// DdosProtectionTrafficStatsAllOf struct for DdosProtectionTrafficStatsAllOf
type DdosProtectionTrafficStatsAllOf struct {
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceID            *string                        `json:"service_id,omitempty"`
	Attributes           []DdosProtectionAttributeStats `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionTrafficStatsAllOf DdosProtectionTrafficStatsAllOf

// NewDdosProtectionTrafficStatsAllOf instantiates a new DdosProtectionTrafficStatsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionTrafficStatsAllOf() *DdosProtectionTrafficStatsAllOf {
	this := DdosProtectionTrafficStatsAllOf{}
	return &this
}

// NewDdosProtectionTrafficStatsAllOfWithDefaults instantiates a new DdosProtectionTrafficStatsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionTrafficStatsAllOfWithDefaults() *DdosProtectionTrafficStatsAllOf {
	this := DdosProtectionTrafficStatsAllOf{}
	return &this
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *DdosProtectionTrafficStatsAllOf) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DdosProtectionTrafficStatsAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DdosProtectionTrafficStatsAllOf) GetAttributes() []DdosProtectionAttributeStats {
	if o == nil || o.Attributes == nil {
		var ret []DdosProtectionAttributeStats
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionTrafficStatsAllOf) GetAttributesOk() ([]DdosProtectionAttributeStats, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DdosProtectionTrafficStatsAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []DdosProtectionAttributeStats and assigns it to the Attributes field.
func (o *DdosProtectionTrafficStatsAllOf) SetAttributes(v []DdosProtectionAttributeStats) {
	o.Attributes = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionTrafficStatsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionTrafficStatsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionTrafficStatsAllOf := _DdosProtectionTrafficStatsAllOf{}

	if err = json.Unmarshal(bytes, &varDdosProtectionTrafficStatsAllOf); err == nil {
		*o = DdosProtectionTrafficStatsAllOf(varDdosProtectionTrafficStatsAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionTrafficStatsAllOf is a helper abstraction for handling nullable ddosprotectiontrafficstatsallof types.
type NullableDdosProtectionTrafficStatsAllOf struct {
	value *DdosProtectionTrafficStatsAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionTrafficStatsAllOf) Get() *DdosProtectionTrafficStatsAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionTrafficStatsAllOf) Set(val *DdosProtectionTrafficStatsAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionTrafficStatsAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionTrafficStatsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionTrafficStatsAllOf returns a pointer to a new instance of NullableDdosProtectionTrafficStatsAllOf.
func NewNullableDdosProtectionTrafficStatsAllOf(val *DdosProtectionTrafficStatsAllOf) *NullableDdosProtectionTrafficStatsAllOf {
	return &NullableDdosProtectionTrafficStatsAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionTrafficStatsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionTrafficStatsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
