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

// IncludedWithTlsConfigurationItem struct for IncludedWithTlsConfigurationItem
type IncludedWithTlsConfigurationItem struct {
	// The IP address or hostname of the DNS record.
	Id                   *string           `json:"id,omitempty"`
	Type                 *TypeTlsDnsRecord `json:"type,omitempty"`
	Attributes           *TlsDnsRecord     `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _IncludedWithTlsConfigurationItem IncludedWithTlsConfigurationItem

// NewIncludedWithTlsConfigurationItem instantiates a new IncludedWithTlsConfigurationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludedWithTlsConfigurationItem() *IncludedWithTlsConfigurationItem {
	this := IncludedWithTlsConfigurationItem{}
	var type_ TypeTlsDnsRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &type_
	return &this
}

// NewIncludedWithTlsConfigurationItemWithDefaults instantiates a new IncludedWithTlsConfigurationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludedWithTlsConfigurationItemWithDefaults() *IncludedWithTlsConfigurationItem {
	this := IncludedWithTlsConfigurationItem{}
	var type_ TypeTlsDnsRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IncludedWithTlsConfigurationItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTlsConfigurationItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IncludedWithTlsConfigurationItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IncludedWithTlsConfigurationItem) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncludedWithTlsConfigurationItem) GetType() TypeTlsDnsRecord {
	if o == nil || o.Type == nil {
		var ret TypeTlsDnsRecord
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTlsConfigurationItem) GetTypeOk() (*TypeTlsDnsRecord, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncludedWithTlsConfigurationItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsDnsRecord and assigns it to the Type field.
func (o *IncludedWithTlsConfigurationItem) SetType(v TypeTlsDnsRecord) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncludedWithTlsConfigurationItem) GetAttributes() TlsDnsRecord {
	if o == nil || o.Attributes == nil {
		var ret TlsDnsRecord
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTlsConfigurationItem) GetAttributesOk() (*TlsDnsRecord, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncludedWithTlsConfigurationItem) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsDnsRecord and assigns it to the Attributes field.
func (o *IncludedWithTlsConfigurationItem) SetAttributes(v TlsDnsRecord) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IncludedWithTlsConfigurationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *IncludedWithTlsConfigurationItem) UnmarshalJSON(bytes []byte) (err error) {
	varIncludedWithTlsConfigurationItem := _IncludedWithTlsConfigurationItem{}

	if err = json.Unmarshal(bytes, &varIncludedWithTlsConfigurationItem); err == nil {
		*o = IncludedWithTlsConfigurationItem(varIncludedWithTlsConfigurationItem)
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

// NullableIncludedWithTlsConfigurationItem is a helper abstraction for handling nullable includedwithtlsconfigurationitem types.
type NullableIncludedWithTlsConfigurationItem struct {
	value *IncludedWithTlsConfigurationItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithTlsConfigurationItem) Get() *IncludedWithTlsConfigurationItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithTlsConfigurationItem) Set(val *IncludedWithTlsConfigurationItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithTlsConfigurationItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithTlsConfigurationItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithTlsConfigurationItem returns a pointer to a new instance of NullableIncludedWithTlsConfigurationItem.
func NewNullableIncludedWithTlsConfigurationItem(val *IncludedWithTlsConfigurationItem) *NullableIncludedWithTlsConfigurationItem {
	return &NullableIncludedWithTlsConfigurationItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithTlsConfigurationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIncludedWithTlsConfigurationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
