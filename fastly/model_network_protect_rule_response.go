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

// NetworkProtectRuleResponse A rule as returned by the API.
type NetworkProtectRuleResponse struct {
	IpProtocol NetworkProtectRuleIpProtocol `json:"ip_protocol"`
	// Match traffic whose source appears in this prefix list.
	SourcePrefixListId string `json:"source_prefix_list_id"`
	// Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family.
	DestinationIpAddress string                             `json:"destination_ip_address"`
	SourcePorts          NetworkProtectRuleSourcePorts      `json:"source_ports"`
	DestinationPorts     NetworkProtectRuleDestinationPorts `json:"destination_ports"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleResponse NetworkProtectRuleResponse

// NewNetworkProtectRuleResponse instantiates a new NetworkProtectRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleResponse(ipProtocol NetworkProtectRuleIpProtocol, sourcePrefixListId string, destinationIpAddress string, sourcePorts NetworkProtectRuleSourcePorts, destinationPorts NetworkProtectRuleDestinationPorts) *NetworkProtectRuleResponse {
	this := NetworkProtectRuleResponse{}
	this.IpProtocol = ipProtocol
	this.SourcePrefixListId = sourcePrefixListId
	this.DestinationIpAddress = destinationIpAddress
	this.SourcePorts = sourcePorts
	this.DestinationPorts = destinationPorts
	return &this
}

// NewNetworkProtectRuleResponseWithDefaults instantiates a new NetworkProtectRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleResponseWithDefaults() *NetworkProtectRuleResponse {
	this := NetworkProtectRuleResponse{}
	return &this
}

// GetIpProtocol returns the IpProtocol field value
func (o *NetworkProtectRuleResponse) GetIpProtocol() NetworkProtectRuleIpProtocol {
	if o == nil {
		var ret NetworkProtectRuleIpProtocol
		return ret
	}

	return o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetIpProtocolOk() (*NetworkProtectRuleIpProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpProtocol, true
}

// SetIpProtocol sets field value
func (o *NetworkProtectRuleResponse) SetIpProtocol(v NetworkProtectRuleIpProtocol) {
	o.IpProtocol = v
}

// GetSourcePrefixListId returns the SourcePrefixListId field value
func (o *NetworkProtectRuleResponse) GetSourcePrefixListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourcePrefixListId
}

// GetSourcePrefixListIdOk returns a tuple with the SourcePrefixListId field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetSourcePrefixListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePrefixListId, true
}

// SetSourcePrefixListId sets field value
func (o *NetworkProtectRuleResponse) SetSourcePrefixListId(v string) {
	o.SourcePrefixListId = v
}

// GetDestinationIpAddress returns the DestinationIpAddress field value
func (o *NetworkProtectRuleResponse) GetDestinationIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIpAddress
}

// GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetDestinationIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIpAddress, true
}

// SetDestinationIpAddress sets field value
func (o *NetworkProtectRuleResponse) SetDestinationIpAddress(v string) {
	o.DestinationIpAddress = v
}

// GetSourcePorts returns the SourcePorts field value
func (o *NetworkProtectRuleResponse) GetSourcePorts() NetworkProtectRuleSourcePorts {
	if o == nil {
		var ret NetworkProtectRuleSourcePorts
		return ret
	}

	return o.SourcePorts
}

// GetSourcePortsOk returns a tuple with the SourcePorts field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetSourcePortsOk() (*NetworkProtectRuleSourcePorts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePorts, true
}

// SetSourcePorts sets field value
func (o *NetworkProtectRuleResponse) SetSourcePorts(v NetworkProtectRuleSourcePorts) {
	o.SourcePorts = v
}

// GetDestinationPorts returns the DestinationPorts field value
func (o *NetworkProtectRuleResponse) GetDestinationPorts() NetworkProtectRuleDestinationPorts {
	if o == nil {
		var ret NetworkProtectRuleDestinationPorts
		return ret
	}

	return o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetDestinationPortsOk() (*NetworkProtectRuleDestinationPorts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationPorts, true
}

// SetDestinationPorts sets field value
func (o *NetworkProtectRuleResponse) SetDestinationPorts(v NetworkProtectRuleDestinationPorts) {
	o.DestinationPorts = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkProtectRuleResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkProtectRuleResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkProtectRuleResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkProtectRuleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkProtectRuleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkProtectRuleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["ip_protocol"] = o.IpProtocol
	}
	if true {
		toSerialize["source_prefix_list_id"] = o.SourcePrefixListId
	}
	if true {
		toSerialize["destination_ip_address"] = o.DestinationIpAddress
	}
	if true {
		toSerialize["source_ports"] = o.SourcePorts
	}
	if true {
		toSerialize["destination_ports"] = o.DestinationPorts
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleResponse := _NetworkProtectRuleResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleResponse); err == nil {
		*o = NetworkProtectRuleResponse(varNetworkProtectRuleResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_protocol")
		delete(additionalProperties, "source_prefix_list_id")
		delete(additionalProperties, "destination_ip_address")
		delete(additionalProperties, "source_ports")
		delete(additionalProperties, "destination_ports")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleResponse is a helper abstraction for handling nullable networkprotectruleresponse types.
type NullableNetworkProtectRuleResponse struct {
	value *NetworkProtectRuleResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleResponse) Get() *NetworkProtectRuleResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleResponse) Set(val *NetworkProtectRuleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleResponse returns a pointer to a new instance of NullableNetworkProtectRuleResponse.
func NewNullableNetworkProtectRuleResponse(val *NetworkProtectRuleResponse) *NullableNetworkProtectRuleResponse {
	return &NullableNetworkProtectRuleResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
