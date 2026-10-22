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

// NetworkProtectNetwork Shared Network Protect network attributes.
type NetworkProtectNetwork struct {
	// Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update.
	Description NullableString `json:"description,omitempty"`
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The CIDR block to protect, masked to its network address. It cannot be changed after creation.
	Prefix string `json:"prefix"`
	// The backend-assigned rule set applied to this network.
	RuleSetId NullableString `json:"rule_set_id,omitempty"`
	// Complete endpoint policy configuration. A network has at most one policy per region and may have no policies.
	EndpointPolicies []NetworkProtectEndpointPolicy `json:"endpoint_policies"`
	// Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does.
	Advertise            bool `json:"advertise"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetwork NetworkProtectNetwork

// NewNetworkProtectNetwork instantiates a new NetworkProtectNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetwork(name string, prefix string, endpointPolicies []NetworkProtectEndpointPolicy, advertise bool) *NetworkProtectNetwork {
	this := NetworkProtectNetwork{}
	this.Name = name
	this.Prefix = prefix
	this.EndpointPolicies = endpointPolicies
	this.Advertise = advertise
	return &this
}

// NewNetworkProtectNetworkWithDefaults instantiates a new NetworkProtectNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworkWithDefaults() *NetworkProtectNetwork {
	this := NetworkProtectNetwork{}
	var advertise bool = false
	this.Advertise = advertise
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetwork) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkProtectNetwork) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkProtectNetwork) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkProtectNetwork) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkProtectNetwork) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *NetworkProtectNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetwork) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectNetwork) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value
func (o *NetworkProtectNetwork) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetwork) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *NetworkProtectNetwork) SetPrefix(v string) {
	o.Prefix = v
}

// GetRuleSetId returns the RuleSetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetwork) GetRuleSetId() string {
	if o == nil || o.RuleSetId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RuleSetId.Get()
}

// GetRuleSetIdOk returns a tuple with the RuleSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetwork) GetRuleSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSetId.Get(), o.RuleSetId.IsSet()
}

// HasRuleSetId returns a boolean if a field has been set.
func (o *NetworkProtectNetwork) HasRuleSetId() bool {
	if o != nil && o.RuleSetId.IsSet() {
		return true
	}

	return false
}

// SetRuleSetId gets a reference to the given NullableString and assigns it to the RuleSetId field.
func (o *NetworkProtectNetwork) SetRuleSetId(v string) {
	o.RuleSetId.Set(&v)
}

// SetRuleSetIdNil sets the value for RuleSetId to be an explicit nil
func (o *NetworkProtectNetwork) SetRuleSetIdNil() {
	o.RuleSetId.Set(nil)
}

// UnsetRuleSetId ensures that no value is present for RuleSetId, not even an explicit nil
func (o *NetworkProtectNetwork) UnsetRuleSetId() {
	o.RuleSetId.Unset()
}

// GetEndpointPolicies returns the EndpointPolicies field value
func (o *NetworkProtectNetwork) GetEndpointPolicies() []NetworkProtectEndpointPolicy {
	if o == nil {
		var ret []NetworkProtectEndpointPolicy
		return ret
	}

	return o.EndpointPolicies
}

// GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetwork) GetEndpointPoliciesOk() ([]NetworkProtectEndpointPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointPolicies, true
}

// SetEndpointPolicies sets field value
func (o *NetworkProtectNetwork) SetEndpointPolicies(v []NetworkProtectEndpointPolicy) {
	o.EndpointPolicies = v
}

// GetAdvertise returns the Advertise field value
func (o *NetworkProtectNetwork) GetAdvertise() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetwork) GetAdvertiseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advertise, true
}

// SetAdvertise sets field value
func (o *NetworkProtectNetwork) SetAdvertise(v bool) {
	o.Advertise = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if o.RuleSetId.IsSet() {
		toSerialize["rule_set_id"] = o.RuleSetId.Get()
	}
	if true {
		toSerialize["endpoint_policies"] = o.EndpointPolicies
	}
	if true {
		toSerialize["advertise"] = o.Advertise
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectNetwork) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetwork := _NetworkProtectNetwork{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetwork); err == nil {
		*o = NetworkProtectNetwork(varNetworkProtectNetwork)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "rule_set_id")
		delete(additionalProperties, "endpoint_policies")
		delete(additionalProperties, "advertise")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetwork is a helper abstraction for handling nullable networkprotectnetwork types.
type NullableNetworkProtectNetwork struct {
	value *NetworkProtectNetwork
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetwork) Get() *NetworkProtectNetwork {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetwork) Set(val *NetworkProtectNetwork) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetwork) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetwork returns a pointer to a new instance of NullableNetworkProtectNetwork.
func NewNullableNetworkProtectNetwork(val *NetworkProtectNetwork) *NullableNetworkProtectNetwork {
	return &NullableNetworkProtectNetwork{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
