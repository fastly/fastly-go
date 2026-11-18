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

// NetworkProtectEndpointResponse A Network Protect endpoint response.
type NetworkProtectEndpointResponse struct {
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The GRE destination address, unique within the customer. It cannot be changed after endpoint creation.
	IPAddress string `json:"ip_address"`
	// A read-only unique identifier for a Network Protect resource.
	ID string `json:"id"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpointResponse NetworkProtectEndpointResponse

// NewNetworkProtectEndpointResponse instantiates a new NetworkProtectEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpointResponse(name string, ipAddress string, id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectEndpointResponse {
	this := NetworkProtectEndpointResponse{}
	this.Name = name
	this.IPAddress = ipAddress
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectEndpointResponseWithDefaults instantiates a new NetworkProtectEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointResponseWithDefaults() *NetworkProtectEndpointResponse {
	this := NetworkProtectEndpointResponse{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectEndpointResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectEndpointResponse) SetName(v string) {
	o.Name = v
}

// GetIPAddress returns the IPAddress field value
func (o *NetworkProtectEndpointResponse) GetIPAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointResponse) GetIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPAddress, true
}

// SetIPAddress sets field value
func (o *NetworkProtectEndpointResponse) SetIPAddress(v string) {
	o.IPAddress = v
}

// GetID returns the ID field value
func (o *NetworkProtectEndpointResponse) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointResponse) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectEndpointResponse) SetID(v string) {
	o.ID = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectEndpointResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectEndpointResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectEndpointResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectEndpointResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ip_address"] = o.IPAddress
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
func (o *NetworkProtectEndpointResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpointResponse := _NetworkProtectEndpointResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpointResponse); err == nil {
		*o = NetworkProtectEndpointResponse(varNetworkProtectEndpointResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "ip_address")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpointResponse is a helper abstraction for handling nullable networkprotectendpointresponse types.
type NullableNetworkProtectEndpointResponse struct {
	value *NetworkProtectEndpointResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpointResponse) Get() *NetworkProtectEndpointResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpointResponse) Set(val *NetworkProtectEndpointResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpointResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpointResponse returns a pointer to a new instance of NullableNetworkProtectEndpointResponse.
func NewNullableNetworkProtectEndpointResponse(val *NetworkProtectEndpointResponse) *NullableNetworkProtectEndpointResponse {
	return &NullableNetworkProtectEndpointResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
