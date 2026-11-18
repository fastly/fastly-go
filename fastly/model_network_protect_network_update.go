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

// NetworkProtectNetworkUpdate Attributes accepted when updating a network.
type NetworkProtectNetworkUpdate struct {
	// Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update.
	Description NullableString `json:"description,omitempty"`
	// A human-readable name, unique within the customer.
	Name *string `json:"name,omitempty"`
	// Custom rule-set attachment is unavailable. Omission, null, and an empty string preserve the current backend-assigned rule set.
	RuleSetID NullableString `json:"rule_set_id,omitempty"`
	// Complete endpoint policy configuration. A network has at most one policy per region and may have no policies.
	EndpointPolicies []NetworkProtectEndpointPolicy `json:"endpoint_policies,omitempty"`
	// Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does.
	Advertise            *bool `json:"advertise,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetworkUpdate NetworkProtectNetworkUpdate

// NewNetworkProtectNetworkUpdate instantiates a new NetworkProtectNetworkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetworkUpdate() *NetworkProtectNetworkUpdate {
	this := NetworkProtectNetworkUpdate{}
	var advertise bool = false
	this.Advertise = &advertise
	return &this
}

// NewNetworkProtectNetworkUpdateWithDefaults instantiates a new NetworkProtectNetworkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworkUpdateWithDefaults() *NetworkProtectNetworkUpdate {
	this := NetworkProtectNetworkUpdate{}
	var advertise bool = false
	this.Advertise = &advertise
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkProtectNetworkUpdate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkProtectNetworkUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkProtectNetworkUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkProtectNetworkUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkProtectNetworkUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkProtectNetworkUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkProtectNetworkUpdate) SetName(v string) {
	o.Name = &v
}

// GetRuleSetID returns the RuleSetID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkUpdate) GetRuleSetID() string {
	if o == nil || o.RuleSetID.Get() == nil {
		var ret string
		return ret
	}
	return *o.RuleSetID.Get()
}

// GetRuleSetIDOk returns a tuple with the RuleSetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkUpdate) GetRuleSetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSetID.Get(), o.RuleSetID.IsSet()
}

// HasRuleSetID returns a boolean if a field has been set.
func (o *NetworkProtectNetworkUpdate) HasRuleSetID() bool {
	if o != nil && o.RuleSetID.IsSet() {
		return true
	}

	return false
}

// SetRuleSetID gets a reference to the given NullableString and assigns it to the RuleSetID field.
func (o *NetworkProtectNetworkUpdate) SetRuleSetID(v string) {
	o.RuleSetID.Set(&v)
}

// SetRuleSetIDNil sets the value for RuleSetID to be an explicit nil
func (o *NetworkProtectNetworkUpdate) SetRuleSetIDNil() {
	o.RuleSetID.Set(nil)
}

// UnsetRuleSetID ensures that no value is present for RuleSetID, not even an explicit nil
func (o *NetworkProtectNetworkUpdate) UnsetRuleSetID() {
	o.RuleSetID.Unset()
}

// GetEndpointPolicies returns the EndpointPolicies field value if set, zero value otherwise.
func (o *NetworkProtectNetworkUpdate) GetEndpointPolicies() []NetworkProtectEndpointPolicy {
	if o == nil || o.EndpointPolicies == nil {
		var ret []NetworkProtectEndpointPolicy
		return ret
	}
	return o.EndpointPolicies
}

// GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkUpdate) GetEndpointPoliciesOk() ([]NetworkProtectEndpointPolicy, bool) {
	if o == nil || o.EndpointPolicies == nil {
		return nil, false
	}
	return o.EndpointPolicies, true
}

// HasEndpointPolicies returns a boolean if a field has been set.
func (o *NetworkProtectNetworkUpdate) HasEndpointPolicies() bool {
	if o != nil && o.EndpointPolicies != nil {
		return true
	}

	return false
}

// SetEndpointPolicies gets a reference to the given []NetworkProtectEndpointPolicy and assigns it to the EndpointPolicies field.
func (o *NetworkProtectNetworkUpdate) SetEndpointPolicies(v []NetworkProtectEndpointPolicy) {
	o.EndpointPolicies = v
}

// GetAdvertise returns the Advertise field value if set, zero value otherwise.
func (o *NetworkProtectNetworkUpdate) GetAdvertise() bool {
	if o == nil || o.Advertise == nil {
		var ret bool
		return ret
	}
	return *o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkUpdate) GetAdvertiseOk() (*bool, bool) {
	if o == nil || o.Advertise == nil {
		return nil, false
	}
	return o.Advertise, true
}

// HasAdvertise returns a boolean if a field has been set.
func (o *NetworkProtectNetworkUpdate) HasAdvertise() bool {
	if o != nil && o.Advertise != nil {
		return true
	}

	return false
}

// SetAdvertise gets a reference to the given bool and assigns it to the Advertise field.
func (o *NetworkProtectNetworkUpdate) SetAdvertise(v bool) {
	o.Advertise = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetworkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RuleSetID.IsSet() {
		toSerialize["rule_set_id"] = o.RuleSetID.Get()
	}
	if o.EndpointPolicies != nil {
		toSerialize["endpoint_policies"] = o.EndpointPolicies
	}
	if o.Advertise != nil {
		toSerialize["advertise"] = o.Advertise
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectNetworkUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetworkUpdate := _NetworkProtectNetworkUpdate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetworkUpdate); err == nil {
		*o = NetworkProtectNetworkUpdate(varNetworkProtectNetworkUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rule_set_id")
		delete(additionalProperties, "endpoint_policies")
		delete(additionalProperties, "advertise")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetworkUpdate is a helper abstraction for handling nullable networkprotectnetworkupdate types.
type NullableNetworkProtectNetworkUpdate struct {
	value *NetworkProtectNetworkUpdate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetworkUpdate) Get() *NetworkProtectNetworkUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetworkUpdate) Set(val *NetworkProtectNetworkUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetworkUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetworkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetworkUpdate returns a pointer to a new instance of NullableNetworkProtectNetworkUpdate.
func NewNullableNetworkProtectNetworkUpdate(val *NetworkProtectNetworkUpdate) *NullableNetworkProtectNetworkUpdate {
	return &NullableNetworkProtectNetworkUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetworkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetworkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
