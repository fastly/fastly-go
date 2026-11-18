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

// NetworkProtectNetworkCreate Attributes accepted when creating a network.
type NetworkProtectNetworkCreate struct {
	// Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update.
	Description NullableString `json:"description,omitempty"`
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The CIDR block to protect, masked to its network address. It cannot be changed after creation.
	Prefix string `json:"prefix"`
	// Custom rule-set attachment is unavailable. Omit this value or send an empty string to use the backend-assigned default rule set.
	RuleSetID *string `json:"rule_set_id,omitempty"`
	// Complete endpoint policy configuration. A network has at most one policy per region and may have no policies.
	EndpointPolicies     []NetworkProtectEndpointPolicy `json:"endpoint_policies,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetworkCreate NetworkProtectNetworkCreate

// NewNetworkProtectNetworkCreate instantiates a new NetworkProtectNetworkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetworkCreate(name string, prefix string) *NetworkProtectNetworkCreate {
	this := NetworkProtectNetworkCreate{}
	this.Name = name
	this.Prefix = prefix
	return &this
}

// NewNetworkProtectNetworkCreateWithDefaults instantiates a new NetworkProtectNetworkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworkCreateWithDefaults() *NetworkProtectNetworkCreate {
	this := NetworkProtectNetworkCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkCreate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkProtectNetworkCreate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkProtectNetworkCreate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkProtectNetworkCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkProtectNetworkCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *NetworkProtectNetworkCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectNetworkCreate) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value
func (o *NetworkProtectNetworkCreate) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkCreate) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *NetworkProtectNetworkCreate) SetPrefix(v string) {
	o.Prefix = v
}

// GetRuleSetID returns the RuleSetID field value if set, zero value otherwise.
func (o *NetworkProtectNetworkCreate) GetRuleSetID() string {
	if o == nil || o.RuleSetID == nil {
		var ret string
		return ret
	}
	return *o.RuleSetID
}

// GetRuleSetIDOk returns a tuple with the RuleSetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkCreate) GetRuleSetIDOk() (*string, bool) {
	if o == nil || o.RuleSetID == nil {
		return nil, false
	}
	return o.RuleSetID, true
}

// HasRuleSetID returns a boolean if a field has been set.
func (o *NetworkProtectNetworkCreate) HasRuleSetID() bool {
	if o != nil && o.RuleSetID != nil {
		return true
	}

	return false
}

// SetRuleSetID gets a reference to the given string and assigns it to the RuleSetID field.
func (o *NetworkProtectNetworkCreate) SetRuleSetID(v string) {
	o.RuleSetID = &v
}

// GetEndpointPolicies returns the EndpointPolicies field value if set, zero value otherwise.
func (o *NetworkProtectNetworkCreate) GetEndpointPolicies() []NetworkProtectEndpointPolicy {
	if o == nil || o.EndpointPolicies == nil {
		var ret []NetworkProtectEndpointPolicy
		return ret
	}
	return o.EndpointPolicies
}

// GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkCreate) GetEndpointPoliciesOk() ([]NetworkProtectEndpointPolicy, bool) {
	if o == nil || o.EndpointPolicies == nil {
		return nil, false
	}
	return o.EndpointPolicies, true
}

// HasEndpointPolicies returns a boolean if a field has been set.
func (o *NetworkProtectNetworkCreate) HasEndpointPolicies() bool {
	if o != nil && o.EndpointPolicies != nil {
		return true
	}

	return false
}

// SetEndpointPolicies gets a reference to the given []NetworkProtectEndpointPolicy and assigns it to the EndpointPolicies field.
func (o *NetworkProtectNetworkCreate) SetEndpointPolicies(v []NetworkProtectEndpointPolicy) {
	o.EndpointPolicies = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetworkCreate) MarshalJSON() ([]byte, error) {
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
	if o.RuleSetID != nil {
		toSerialize["rule_set_id"] = o.RuleSetID
	}
	if o.EndpointPolicies != nil {
		toSerialize["endpoint_policies"] = o.EndpointPolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectNetworkCreate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetworkCreate := _NetworkProtectNetworkCreate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetworkCreate); err == nil {
		*o = NetworkProtectNetworkCreate(varNetworkProtectNetworkCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "rule_set_id")
		delete(additionalProperties, "endpoint_policies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetworkCreate is a helper abstraction for handling nullable networkprotectnetworkcreate types.
type NullableNetworkProtectNetworkCreate struct {
	value *NetworkProtectNetworkCreate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetworkCreate) Get() *NetworkProtectNetworkCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetworkCreate) Set(val *NetworkProtectNetworkCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetworkCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetworkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetworkCreate returns a pointer to a new instance of NullableNetworkProtectNetworkCreate.
func NewNullableNetworkProtectNetworkCreate(val *NetworkProtectNetworkCreate) *NullableNetworkProtectNetworkCreate {
	return &NullableNetworkProtectNetworkCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetworkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetworkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
