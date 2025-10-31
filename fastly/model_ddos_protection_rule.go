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

// DdosProtectionRule struct for DdosProtectionRule
type DdosProtectionRule struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Unique ID of the rule.
	Id *string `json:"id,omitempty"`
	// A human-readable name for the rule.
	Name *string `json:"name,omitempty"`
	// Action types for a rule. Supported action values are default, block, log, off. The default action value follows the current protection mode of the associated service.
	Action *string `json:"action,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerId *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceId *string `json:"service_id,omitempty"`
	// Source IP address attribute.
	SourceIp NullableString `json:"source_ip,omitempty"`
	// Country code attribute.
	CountryCode NullableString `json:"country_code,omitempty"`
	// Host attribute.
	Host NullableString `json:"host,omitempty"`
	// ASN attribute.
	Asn NullableString `json:"asn,omitempty"`
	// Source IP prefix attribute.
	SourceIpPrefix NullableString `json:"source_ip_prefix,omitempty"`
	// Attribute category for additional, unlisted attributes used in a rule.
	AdditionalAttributes []string `json:"additional_attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRule DdosProtectionRule

// NewDdosProtectionRule instantiates a new DdosProtectionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRule() *DdosProtectionRule {
	this := DdosProtectionRule{}
	var action string = "default"
	this.Action = &action
	return &this
}

// NewDdosProtectionRuleWithDefaults instantiates a new DdosProtectionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRuleWithDefaults() *DdosProtectionRule {
	this := DdosProtectionRule{}
	var action string = "default"
	this.Action = &action
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *DdosProtectionRule) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *DdosProtectionRule) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *DdosProtectionRule) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *DdosProtectionRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *DdosProtectionRule) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *DdosProtectionRule) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DdosProtectionRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DdosProtectionRule) SetName(v string) {
	o.Name = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DdosProtectionRule) SetAction(v string) {
	o.Action = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *DdosProtectionRule) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DdosProtectionRule) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetSourceIp() string {
	if o == nil || o.SourceIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceIp.Get()
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetSourceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIp.Get(), o.SourceIp.IsSet()
}

// HasSourceIp returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasSourceIp() bool {
	if o != nil && o.SourceIp.IsSet() {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given NullableString and assigns it to the SourceIp field.
func (o *DdosProtectionRule) SetSourceIp(v string) {
	o.SourceIp.Set(&v)
}

// SetSourceIpNil sets the value for SourceIp to be an explicit nil
func (o *DdosProtectionRule) SetSourceIpNil() {
	o.SourceIp.Set(nil)
}

// UnsetSourceIp ensures that no value is present for SourceIp, not even an explicit nil
func (o *DdosProtectionRule) UnsetSourceIp() {
	o.SourceIp.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *DdosProtectionRule) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *DdosProtectionRule) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *DdosProtectionRule) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *DdosProtectionRule) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *DdosProtectionRule) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *DdosProtectionRule) UnsetHost() {
	o.Host.Unset()
}

// GetAsn returns the Asn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetAsn() string {
	if o == nil || o.Asn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetAsnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// HasAsn returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasAsn() bool {
	if o != nil && o.Asn.IsSet() {
		return true
	}

	return false
}

// SetAsn gets a reference to the given NullableString and assigns it to the Asn field.
func (o *DdosProtectionRule) SetAsn(v string) {
	o.Asn.Set(&v)
}

// SetAsnNil sets the value for Asn to be an explicit nil
func (o *DdosProtectionRule) SetAsnNil() {
	o.Asn.Set(nil)
}

// UnsetAsn ensures that no value is present for Asn, not even an explicit nil
func (o *DdosProtectionRule) UnsetAsn() {
	o.Asn.Unset()
}

// GetSourceIpPrefix returns the SourceIpPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRule) GetSourceIpPrefix() string {
	if o == nil || o.SourceIpPrefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceIpPrefix.Get()
}

// GetSourceIpPrefixOk returns a tuple with the SourceIpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRule) GetSourceIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIpPrefix.Get(), o.SourceIpPrefix.IsSet()
}

// HasSourceIpPrefix returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasSourceIpPrefix() bool {
	if o != nil && o.SourceIpPrefix.IsSet() {
		return true
	}

	return false
}

// SetSourceIpPrefix gets a reference to the given NullableString and assigns it to the SourceIpPrefix field.
func (o *DdosProtectionRule) SetSourceIpPrefix(v string) {
	o.SourceIpPrefix.Set(&v)
}

// SetSourceIpPrefixNil sets the value for SourceIpPrefix to be an explicit nil
func (o *DdosProtectionRule) SetSourceIpPrefixNil() {
	o.SourceIpPrefix.Set(nil)
}

// UnsetSourceIpPrefix ensures that no value is present for SourceIpPrefix, not even an explicit nil
func (o *DdosProtectionRule) UnsetSourceIpPrefix() {
	o.SourceIpPrefix.Unset()
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *DdosProtectionRule) GetAdditionalAttributes() []string {
	if o == nil || o.AdditionalAttributes == nil {
		var ret []string
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRule) GetAdditionalAttributesOk() ([]string, bool) {
	if o == nil || o.AdditionalAttributes == nil {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *DdosProtectionRule) HasAdditionalAttributes() bool {
	if o != nil && o.AdditionalAttributes != nil {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []string and assigns it to the AdditionalAttributes field.
func (o *DdosProtectionRule) SetAdditionalAttributes(v []string) {
	o.AdditionalAttributes = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.SourceIp.IsSet() {
		toSerialize["source_ip"] = o.SourceIp.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Asn.IsSet() {
		toSerialize["asn"] = o.Asn.Get()
	}
	if o.SourceIpPrefix.IsSet() {
		toSerialize["source_ip_prefix"] = o.SourceIpPrefix.Get()
	}
	if o.AdditionalAttributes != nil {
		toSerialize["additional_attributes"] = o.AdditionalAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionRule) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRule := _DdosProtectionRule{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRule); err == nil {
		*o = DdosProtectionRule(varDdosProtectionRule)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "action")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "source_ip")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "host")
		delete(additionalProperties, "asn")
		delete(additionalProperties, "source_ip_prefix")
		delete(additionalProperties, "additional_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionRule is a helper abstraction for handling nullable ddosprotectionrule types.
type NullableDdosProtectionRule struct {
	value *DdosProtectionRule
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRule) Get() *DdosProtectionRule {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRule) Set(val *DdosProtectionRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRule returns a pointer to a new instance of NullableDdosProtectionRule.
func NewNullableDdosProtectionRule(val *DdosProtectionRule) *NullableDdosProtectionRule {
	return &NullableDdosProtectionRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
