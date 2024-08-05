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

// IncludedWithTLSConfigurationItem struct for IncludedWithTLSConfigurationItem
type IncludedWithTLSConfigurationItem struct {
	// The IP address or hostname of the DNS record.
	ID *string `json:"id,omitempty"`
	Type *TypeTLSDNSRecord `json:"type,omitempty"`
	Attributes *TLSDNSRecord `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _IncludedWithTLSConfigurationItem IncludedWithTLSConfigurationItem

// NewIncludedWithTLSConfigurationItem instantiates a new IncludedWithTLSConfigurationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludedWithTLSConfigurationItem() *IncludedWithTLSConfigurationItem {
	this := IncludedWithTLSConfigurationItem{}
	var resourceType TypeTLSDNSRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &resourceType
	return &this
}

// NewIncludedWithTLSConfigurationItemWithDefaults instantiates a new IncludedWithTLSConfigurationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludedWithTLSConfigurationItemWithDefaults() *IncludedWithTLSConfigurationItem {
	this := IncludedWithTLSConfigurationItem{}
	var resourceType TypeTLSDNSRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &resourceType
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *IncludedWithTLSConfigurationItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTLSConfigurationItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *IncludedWithTLSConfigurationItem) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *IncludedWithTLSConfigurationItem) SetID(v string) {
	o.ID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncludedWithTLSConfigurationItem) GetType() TypeTLSDNSRecord {
	if o == nil || o.Type == nil {
		var ret TypeTLSDNSRecord
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTLSConfigurationItem) GetTypeOk() (*TypeTLSDNSRecord, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncludedWithTLSConfigurationItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSDNSRecord and assigns it to the Type field.
func (o *IncludedWithTLSConfigurationItem) SetType(v TypeTLSDNSRecord) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncludedWithTLSConfigurationItem) GetAttributes() TLSDNSRecord {
	if o == nil || o.Attributes == nil {
		var ret TLSDNSRecord
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTLSConfigurationItem) GetAttributesOk() (*TLSDNSRecord, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncludedWithTLSConfigurationItem) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSDNSRecord and assigns it to the Attributes field.
func (o *IncludedWithTLSConfigurationItem) SetAttributes(v TLSDNSRecord) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IncludedWithTLSConfigurationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *IncludedWithTLSConfigurationItem) UnmarshalJSON(bytes []byte) (err error) {
	varIncludedWithTLSConfigurationItem := _IncludedWithTLSConfigurationItem{}

	if err = json.Unmarshal(bytes, &varIncludedWithTLSConfigurationItem); err == nil {
		*o = IncludedWithTLSConfigurationItem(varIncludedWithTLSConfigurationItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIncludedWithTLSConfigurationItem is a helper abstraction for handling nullable includedwithtlsconfigurationitem types. 
type NullableIncludedWithTLSConfigurationItem struct {
	value *IncludedWithTLSConfigurationItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithTLSConfigurationItem) Get() *IncludedWithTLSConfigurationItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithTLSConfigurationItem) Set(val *IncludedWithTLSConfigurationItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithTLSConfigurationItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithTLSConfigurationItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithTLSConfigurationItem returns a pointer to a new instance of NullableIncludedWithTLSConfigurationItem.
func NewNullableIncludedWithTLSConfigurationItem(val *IncludedWithTLSConfigurationItem) *NullableIncludedWithTLSConfigurationItem {
	return &NullableIncludedWithTLSConfigurationItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithTLSConfigurationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIncludedWithTLSConfigurationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
