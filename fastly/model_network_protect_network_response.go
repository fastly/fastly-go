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

// NetworkProtectNetworkResponse A Network Protect network response.
type NetworkProtectNetworkResponse struct {
	// Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update.
	Description NullableString `json:"description,omitempty"`
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The CIDR block to protect, masked to its network address. It cannot be changed after creation.
	Prefix string `json:"prefix"`
	// The backend-assigned rule set applied to this network.
	RuleSetID NullableString `json:"rule_set_id,omitempty"`
	// Complete endpoint policy configuration. A network has at most one policy per region and may have no policies.
	EndpointPolicies []NetworkProtectEndpointPolicy `json:"endpoint_policies"`
	// Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does.
	Advertise bool `json:"advertise"`
	// A read-only unique identifier for a Network Protect resource.
	ID        string                   `json:"id"`
	RoaStatus *NetworkProtectRoaStatus `json:"roa_status,omitempty"`
	// When Fastly last completed a ROA check. Null while the status is pending. A failed check leaves this unchanged; an old value can indicate delayed or failing checks. Null also means the controller did not report a check time.
	RoaCheckedAt NullableTime `json:"roa_checked_at,omitempty"`
	// When the RPKI dataset used by the last check was built. This is not the same as `roa_checked_at`; a recent check can use an older dataset. Null while the status is pending or when no dataset build time was reported.
	RoaDataAsOf NullableTime             `json:"roa_data_as_of,omitempty"`
	LoaStatus   *NetworkProtectLoaStatus `json:"loa_status,omitempty"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetworkResponse NetworkProtectNetworkResponse

// NewNetworkProtectNetworkResponse instantiates a new NetworkProtectNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetworkResponse(name string, prefix string, endpointPolicies []NetworkProtectEndpointPolicy, advertise bool, id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectNetworkResponse {
	this := NetworkProtectNetworkResponse{}
	this.Name = name
	this.Prefix = prefix
	this.EndpointPolicies = endpointPolicies
	this.Advertise = advertise
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectNetworkResponseWithDefaults instantiates a new NetworkProtectNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworkResponseWithDefaults() *NetworkProtectNetworkResponse {
	this := NetworkProtectNetworkResponse{}
	var advertise bool = false
	this.Advertise = advertise
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkProtectNetworkResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkProtectNetworkResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkProtectNetworkResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *NetworkProtectNetworkResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectNetworkResponse) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value
func (o *NetworkProtectNetworkResponse) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *NetworkProtectNetworkResponse) SetPrefix(v string) {
	o.Prefix = v
}

// GetRuleSetID returns the RuleSetID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponse) GetRuleSetID() string {
	if o == nil || o.RuleSetID.Get() == nil {
		var ret string
		return ret
	}
	return *o.RuleSetID.Get()
}

// GetRuleSetIDOk returns a tuple with the RuleSetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponse) GetRuleSetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSetID.Get(), o.RuleSetID.IsSet()
}

// HasRuleSetID returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasRuleSetID() bool {
	if o != nil && o.RuleSetID.IsSet() {
		return true
	}

	return false
}

// SetRuleSetID gets a reference to the given NullableString and assigns it to the RuleSetID field.
func (o *NetworkProtectNetworkResponse) SetRuleSetID(v string) {
	o.RuleSetID.Set(&v)
}

// SetRuleSetIDNil sets the value for RuleSetID to be an explicit nil
func (o *NetworkProtectNetworkResponse) SetRuleSetIDNil() {
	o.RuleSetID.Set(nil)
}

// UnsetRuleSetID ensures that no value is present for RuleSetID, not even an explicit nil
func (o *NetworkProtectNetworkResponse) UnsetRuleSetID() {
	o.RuleSetID.Unset()
}

// GetEndpointPolicies returns the EndpointPolicies field value
func (o *NetworkProtectNetworkResponse) GetEndpointPolicies() []NetworkProtectEndpointPolicy {
	if o == nil {
		var ret []NetworkProtectEndpointPolicy
		return ret
	}

	return o.EndpointPolicies
}

// GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetEndpointPoliciesOk() ([]NetworkProtectEndpointPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointPolicies, true
}

// SetEndpointPolicies sets field value
func (o *NetworkProtectNetworkResponse) SetEndpointPolicies(v []NetworkProtectEndpointPolicy) {
	o.EndpointPolicies = v
}

// GetAdvertise returns the Advertise field value
func (o *NetworkProtectNetworkResponse) GetAdvertise() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetAdvertiseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advertise, true
}

// SetAdvertise sets field value
func (o *NetworkProtectNetworkResponse) SetAdvertise(v bool) {
	o.Advertise = v
}

// GetID returns the ID field value
func (o *NetworkProtectNetworkResponse) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectNetworkResponse) SetID(v string) {
	o.ID = v
}

// GetRoaStatus returns the RoaStatus field value if set, zero value otherwise.
func (o *NetworkProtectNetworkResponse) GetRoaStatus() NetworkProtectRoaStatus {
	if o == nil || o.RoaStatus == nil {
		var ret NetworkProtectRoaStatus
		return ret
	}
	return *o.RoaStatus
}

// GetRoaStatusOk returns a tuple with the RoaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetRoaStatusOk() (*NetworkProtectRoaStatus, bool) {
	if o == nil || o.RoaStatus == nil {
		return nil, false
	}
	return o.RoaStatus, true
}

// HasRoaStatus returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasRoaStatus() bool {
	if o != nil && o.RoaStatus != nil {
		return true
	}

	return false
}

// SetRoaStatus gets a reference to the given NetworkProtectRoaStatus and assigns it to the RoaStatus field.
func (o *NetworkProtectNetworkResponse) SetRoaStatus(v NetworkProtectRoaStatus) {
	o.RoaStatus = &v
}

// GetRoaCheckedAt returns the RoaCheckedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponse) GetRoaCheckedAt() time.Time {
	if o == nil || o.RoaCheckedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RoaCheckedAt.Get()
}

// GetRoaCheckedAtOk returns a tuple with the RoaCheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponse) GetRoaCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoaCheckedAt.Get(), o.RoaCheckedAt.IsSet()
}

// HasRoaCheckedAt returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasRoaCheckedAt() bool {
	if o != nil && o.RoaCheckedAt.IsSet() {
		return true
	}

	return false
}

// SetRoaCheckedAt gets a reference to the given NullableTime and assigns it to the RoaCheckedAt field.
func (o *NetworkProtectNetworkResponse) SetRoaCheckedAt(v time.Time) {
	o.RoaCheckedAt.Set(&v)
}

// SetRoaCheckedAtNil sets the value for RoaCheckedAt to be an explicit nil
func (o *NetworkProtectNetworkResponse) SetRoaCheckedAtNil() {
	o.RoaCheckedAt.Set(nil)
}

// UnsetRoaCheckedAt ensures that no value is present for RoaCheckedAt, not even an explicit nil
func (o *NetworkProtectNetworkResponse) UnsetRoaCheckedAt() {
	o.RoaCheckedAt.Unset()
}

// GetRoaDataAsOf returns the RoaDataAsOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponse) GetRoaDataAsOf() time.Time {
	if o == nil || o.RoaDataAsOf.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RoaDataAsOf.Get()
}

// GetRoaDataAsOfOk returns a tuple with the RoaDataAsOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponse) GetRoaDataAsOfOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoaDataAsOf.Get(), o.RoaDataAsOf.IsSet()
}

// HasRoaDataAsOf returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasRoaDataAsOf() bool {
	if o != nil && o.RoaDataAsOf.IsSet() {
		return true
	}

	return false
}

// SetRoaDataAsOf gets a reference to the given NullableTime and assigns it to the RoaDataAsOf field.
func (o *NetworkProtectNetworkResponse) SetRoaDataAsOf(v time.Time) {
	o.RoaDataAsOf.Set(&v)
}

// SetRoaDataAsOfNil sets the value for RoaDataAsOf to be an explicit nil
func (o *NetworkProtectNetworkResponse) SetRoaDataAsOfNil() {
	o.RoaDataAsOf.Set(nil)
}

// UnsetRoaDataAsOf ensures that no value is present for RoaDataAsOf, not even an explicit nil
func (o *NetworkProtectNetworkResponse) UnsetRoaDataAsOf() {
	o.RoaDataAsOf.Unset()
}

// GetLoaStatus returns the LoaStatus field value if set, zero value otherwise.
func (o *NetworkProtectNetworkResponse) GetLoaStatus() NetworkProtectLoaStatus {
	if o == nil || o.LoaStatus == nil {
		var ret NetworkProtectLoaStatus
		return ret
	}
	return *o.LoaStatus
}

// GetLoaStatusOk returns a tuple with the LoaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetLoaStatusOk() (*NetworkProtectLoaStatus, bool) {
	if o == nil || o.LoaStatus == nil {
		return nil, false
	}
	return o.LoaStatus, true
}

// HasLoaStatus returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponse) HasLoaStatus() bool {
	if o != nil && o.LoaStatus != nil {
		return true
	}

	return false
}

// SetLoaStatus gets a reference to the given NetworkProtectLoaStatus and assigns it to the LoaStatus field.
func (o *NetworkProtectNetworkResponse) SetLoaStatus(v NetworkProtectLoaStatus) {
	o.LoaStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectNetworkResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectNetworkResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectNetworkResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectNetworkResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetworkResponse) MarshalJSON() ([]byte, error) {
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
	if o.RuleSetID.IsSet() {
		toSerialize["rule_set_id"] = o.RuleSetID.Get()
	}
	if true {
		toSerialize["endpoint_policies"] = o.EndpointPolicies
	}
	if true {
		toSerialize["advertise"] = o.Advertise
	}
	if true {
		toSerialize["id"] = o.ID
	}
	if o.RoaStatus != nil {
		toSerialize["roa_status"] = o.RoaStatus
	}
	if o.RoaCheckedAt.IsSet() {
		toSerialize["roa_checked_at"] = o.RoaCheckedAt.Get()
	}
	if o.RoaDataAsOf.IsSet() {
		toSerialize["roa_data_as_of"] = o.RoaDataAsOf.Get()
	}
	if o.LoaStatus != nil {
		toSerialize["loa_status"] = o.LoaStatus
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
func (o *NetworkProtectNetworkResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetworkResponse := _NetworkProtectNetworkResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetworkResponse); err == nil {
		*o = NetworkProtectNetworkResponse(varNetworkProtectNetworkResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "rule_set_id")
		delete(additionalProperties, "endpoint_policies")
		delete(additionalProperties, "advertise")
		delete(additionalProperties, "id")
		delete(additionalProperties, "roa_status")
		delete(additionalProperties, "roa_checked_at")
		delete(additionalProperties, "roa_data_as_of")
		delete(additionalProperties, "loa_status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetworkResponse is a helper abstraction for handling nullable networkprotectnetworkresponse types.
type NullableNetworkProtectNetworkResponse struct {
	value *NetworkProtectNetworkResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetworkResponse) Get() *NetworkProtectNetworkResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetworkResponse) Set(val *NetworkProtectNetworkResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetworkResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetworkResponse returns a pointer to a new instance of NullableNetworkProtectNetworkResponse.
func NewNullableNetworkProtectNetworkResponse(val *NetworkProtectNetworkResponse) *NullableNetworkProtectNetworkResponse {
	return &NullableNetworkProtectNetworkResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
