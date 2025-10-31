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

// TlsDnsRecord struct for TlsDnsRecord
type TlsDnsRecord struct {
	// Specifies the regions that will be used to route traffic. Select DNS records with a `global` region to route traffic to the most performant point of presence (POP) worldwide (global pricing will apply). Select DNS records with a `na/eu` region to exclusively land traffic on North American and European POPs.
	Region *string `json:"region,omitempty"`
	// The type of the DNS record. `A` specifies an IPv4 address to be used for an A record to be used for apex domains (e.g., `example.com`). `AAAA` specifies an IPv6 address for use in an A record for apex domains. `CNAME` specifies the hostname to be used for a CNAME record for subdomains or wildcard domains (e.g., `www.example.com` or `*.example.com`).
	RecordType           *string `json:"record_type,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsDnsRecord TlsDnsRecord

// NewTlsDnsRecord instantiates a new TlsDnsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsDnsRecord() *TlsDnsRecord {
	this := TlsDnsRecord{}
	return &this
}

// NewTlsDnsRecordWithDefaults instantiates a new TlsDnsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsDnsRecordWithDefaults() *TlsDnsRecord {
	this := TlsDnsRecord{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *TlsDnsRecord) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDnsRecord) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *TlsDnsRecord) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *TlsDnsRecord) SetRegion(v string) {
	o.Region = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *TlsDnsRecord) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDnsRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *TlsDnsRecord) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *TlsDnsRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsDnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RecordType != nil {
		toSerialize["record_type"] = o.RecordType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsDnsRecord) UnmarshalJSON(bytes []byte) (err error) {
	varTlsDnsRecord := _TlsDnsRecord{}

	if err = json.Unmarshal(bytes, &varTlsDnsRecord); err == nil {
		*o = TlsDnsRecord(varTlsDnsRecord)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "record_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsDnsRecord is a helper abstraction for handling nullable tlsdnsrecord types.
type NullableTlsDnsRecord struct {
	value *TlsDnsRecord
	isSet bool
}

// Get returns the value.
func (v NullableTlsDnsRecord) Get() *TlsDnsRecord {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsDnsRecord) Set(val *TlsDnsRecord) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsDnsRecord) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsDnsRecord returns a pointer to a new instance of NullableTlsDnsRecord.
func NewNullableTlsDnsRecord(val *TlsDnsRecord) *NullableTlsDnsRecord {
	return &NullableTlsDnsRecord{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
