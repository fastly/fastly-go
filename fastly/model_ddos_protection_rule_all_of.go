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

// DdosProtectionRuleAllOf struct for DdosProtectionRuleAllOf
type DdosProtectionRuleAllOf struct {
	// Unique ID of the rule.
	ID *string `json:"id,omitempty"`
	// A human-readable name for the rule.
	Name   *string               `json:"name,omitempty"`
	Action *DdosProtectionAction `json:"action,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceID *string `json:"service_id,omitempty"`
	// Source IP address attribute.
	SourceIP NullableString `json:"source_ip,omitempty"`
	// Country code attribute.
	CountryCode NullableString `json:"country_code,omitempty"`
	// Host attribute.
	Host NullableString `json:"host,omitempty"`
	// ASN attribute.
	Asn NullableString `json:"asn,omitempty"`
	// Source IP prefix attribute.
	SourceIPPrefix NullableString `json:"source_ip_prefix,omitempty"`
	// Attribute category for additional, unlisted attributes used in a rule.
	AdditionalAttributes []string `json:"additional_attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRuleAllOf DdosProtectionRuleAllOf

// NewDdosProtectionRuleAllOf instantiates a new DdosProtectionRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRuleAllOf() *DdosProtectionRuleAllOf {
	this := DdosProtectionRuleAllOf{}
	var action DdosProtectionAction = DDOSPROTECTIONACTION_DEFAULT
	this.Action = &action
	return &this
}

// NewDdosProtectionRuleAllOfWithDefaults instantiates a new DdosProtectionRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRuleAllOfWithDefaults() *DdosProtectionRuleAllOf {
	this := DdosProtectionRuleAllOf{}
	var action DdosProtectionAction = DDOSPROTECTIONACTION_DEFAULT
	this.Action = &action
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *DdosProtectionRuleAllOf) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DdosProtectionRuleAllOf) SetName(v string) {
	o.Name = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetAction() DdosProtectionAction {
	if o == nil || o.Action == nil {
		var ret DdosProtectionAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetActionOk() (*DdosProtectionAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given DdosProtectionAction and assigns it to the Action field.
func (o *DdosProtectionRuleAllOf) SetAction(v DdosProtectionAction) {
	o.Action = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *DdosProtectionRuleAllOf) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DdosProtectionRuleAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetSourceIP returns the SourceIP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRuleAllOf) GetSourceIP() string {
	if o == nil || o.SourceIP.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceIP.Get()
}

// GetSourceIPOk returns a tuple with the SourceIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRuleAllOf) GetSourceIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIP.Get(), o.SourceIP.IsSet()
}

// HasSourceIP returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasSourceIP() bool {
	if o != nil && o.SourceIP.IsSet() {
		return true
	}

	return false
}

// SetSourceIP gets a reference to the given NullableString and assigns it to the SourceIP field.
func (o *DdosProtectionRuleAllOf) SetSourceIP(v string) {
	o.SourceIP.Set(&v)
}

// SetSourceIPNil sets the value for SourceIP to be an explicit nil
func (o *DdosProtectionRuleAllOf) SetSourceIPNil() {
	o.SourceIP.Set(nil)
}

// UnsetSourceIP ensures that no value is present for SourceIP, not even an explicit nil
func (o *DdosProtectionRuleAllOf) UnsetSourceIP() {
	o.SourceIP.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRuleAllOf) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRuleAllOf) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *DdosProtectionRuleAllOf) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *DdosProtectionRuleAllOf) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *DdosProtectionRuleAllOf) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRuleAllOf) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRuleAllOf) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *DdosProtectionRuleAllOf) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *DdosProtectionRuleAllOf) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *DdosProtectionRuleAllOf) UnsetHost() {
	o.Host.Unset()
}

// GetAsn returns the Asn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRuleAllOf) GetAsn() string {
	if o == nil || o.Asn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRuleAllOf) GetAsnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// HasAsn returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasAsn() bool {
	if o != nil && o.Asn.IsSet() {
		return true
	}

	return false
}

// SetAsn gets a reference to the given NullableString and assigns it to the Asn field.
func (o *DdosProtectionRuleAllOf) SetAsn(v string) {
	o.Asn.Set(&v)
}

// SetAsnNil sets the value for Asn to be an explicit nil
func (o *DdosProtectionRuleAllOf) SetAsnNil() {
	o.Asn.Set(nil)
}

// UnsetAsn ensures that no value is present for Asn, not even an explicit nil
func (o *DdosProtectionRuleAllOf) UnsetAsn() {
	o.Asn.Unset()
}

// GetSourceIPPrefix returns the SourceIPPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosProtectionRuleAllOf) GetSourceIPPrefix() string {
	if o == nil || o.SourceIPPrefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceIPPrefix.Get()
}

// GetSourceIPPrefixOk returns a tuple with the SourceIPPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosProtectionRuleAllOf) GetSourceIPPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIPPrefix.Get(), o.SourceIPPrefix.IsSet()
}

// HasSourceIPPrefix returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasSourceIPPrefix() bool {
	if o != nil && o.SourceIPPrefix.IsSet() {
		return true
	}

	return false
}

// SetSourceIPPrefix gets a reference to the given NullableString and assigns it to the SourceIPPrefix field.
func (o *DdosProtectionRuleAllOf) SetSourceIPPrefix(v string) {
	o.SourceIPPrefix.Set(&v)
}

// SetSourceIPPrefixNil sets the value for SourceIPPrefix to be an explicit nil
func (o *DdosProtectionRuleAllOf) SetSourceIPPrefixNil() {
	o.SourceIPPrefix.Set(nil)
}

// UnsetSourceIPPrefix ensures that no value is present for SourceIPPrefix, not even an explicit nil
func (o *DdosProtectionRuleAllOf) UnsetSourceIPPrefix() {
	o.SourceIPPrefix.Unset()
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *DdosProtectionRuleAllOf) GetAdditionalAttributes() []string {
	if o == nil || o.AdditionalAttributes == nil {
		var ret []string
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRuleAllOf) GetAdditionalAttributesOk() ([]string, bool) {
	if o == nil || o.AdditionalAttributes == nil {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *DdosProtectionRuleAllOf) HasAdditionalAttributes() bool {
	if o != nil && o.AdditionalAttributes != nil {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []string and assigns it to the AdditionalAttributes field.
func (o *DdosProtectionRuleAllOf) SetAdditionalAttributes(v []string) {
	o.AdditionalAttributes = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.SourceIP.IsSet() {
		toSerialize["source_ip"] = o.SourceIP.Get()
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
	if o.SourceIPPrefix.IsSet() {
		toSerialize["source_ip_prefix"] = o.SourceIPPrefix.Get()
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
func (o *DdosProtectionRuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRuleAllOf := _DdosProtectionRuleAllOf{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRuleAllOf); err == nil {
		*o = DdosProtectionRuleAllOf(varDdosProtectionRuleAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

// NullableDdosProtectionRuleAllOf is a helper abstraction for handling nullable ddosprotectionruleallof types.
type NullableDdosProtectionRuleAllOf struct {
	value *DdosProtectionRuleAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRuleAllOf) Get() *DdosProtectionRuleAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRuleAllOf) Set(val *DdosProtectionRuleAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRuleAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRuleAllOf returns a pointer to a new instance of NullableDdosProtectionRuleAllOf.
func NewNullableDdosProtectionRuleAllOf(val *DdosProtectionRuleAllOf) *NullableDdosProtectionRuleAllOf {
	return &NullableDdosProtectionRuleAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
