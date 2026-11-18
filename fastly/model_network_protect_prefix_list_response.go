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

// NetworkProtectPrefixListResponse A Network Protect prefix-list response.
type NetworkProtectPrefixListResponse struct {
	// A human-readable name, unique within the customer.
	Name   string                         `json:"name"`
	Family NetworkProtectPrefixListFamily `json:"family"`
	// Complete canonical list contents. Bare addresses are canonicalized, CIDRs are masked, equivalent entries are deduplicated, and overlaps are retained.
	Addresses []string `json:"addresses"`
	// A read-only unique identifier for a Network Protect resource.
	ID string `json:"id"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPrefixListResponse NetworkProtectPrefixListResponse

// NewNetworkProtectPrefixListResponse instantiates a new NetworkProtectPrefixListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPrefixListResponse(name string, family NetworkProtectPrefixListFamily, addresses []string, id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectPrefixListResponse {
	this := NetworkProtectPrefixListResponse{}
	this.Name = name
	this.Family = family
	this.Addresses = addresses
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectPrefixListResponseWithDefaults instantiates a new NetworkProtectPrefixListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPrefixListResponseWithDefaults() *NetworkProtectPrefixListResponse {
	this := NetworkProtectPrefixListResponse{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectPrefixListResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectPrefixListResponse) SetName(v string) {
	o.Name = v
}

// GetFamily returns the Family field value
func (o *NetworkProtectPrefixListResponse) GetFamily() NetworkProtectPrefixListFamily {
	if o == nil {
		var ret NetworkProtectPrefixListFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetFamilyOk() (*NetworkProtectPrefixListFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *NetworkProtectPrefixListResponse) SetFamily(v NetworkProtectPrefixListFamily) {
	o.Family = v
}

// GetAddresses returns the Addresses field value
func (o *NetworkProtectPrefixListResponse) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *NetworkProtectPrefixListResponse) SetAddresses(v []string) {
	o.Addresses = v
}

// GetID returns the ID field value
func (o *NetworkProtectPrefixListResponse) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectPrefixListResponse) SetID(v string) {
	o.ID = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectPrefixListResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectPrefixListResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectPrefixListResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectPrefixListResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPrefixListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["id"] = o.ID
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
func (o *NetworkProtectPrefixListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPrefixListResponse := _NetworkProtectPrefixListResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPrefixListResponse); err == nil {
		*o = NetworkProtectPrefixListResponse(varNetworkProtectPrefixListResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "family")
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPrefixListResponse is a helper abstraction for handling nullable networkprotectprefixlistresponse types.
type NullableNetworkProtectPrefixListResponse struct {
	value *NetworkProtectPrefixListResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixListResponse) Get() *NetworkProtectPrefixListResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixListResponse) Set(val *NetworkProtectPrefixListResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixListResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixListResponse returns a pointer to a new instance of NullableNetworkProtectPrefixListResponse.
func NewNullableNetworkProtectPrefixListResponse(val *NetworkProtectPrefixListResponse) *NullableNetworkProtectPrefixListResponse {
	return &NullableNetworkProtectPrefixListResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
