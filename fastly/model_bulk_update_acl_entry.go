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

// BulkUpdateAclEntry struct for BulkUpdateAclEntry
type BulkUpdateAclEntry struct {
	// Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.
	Negated *int32 `json:"negated,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// An IP address.
	Ip *string `json:"ip,omitempty"`
	// Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied.
	Subnet               NullableInt32 `json:"subnet,omitempty"`
	Op                   *string       `json:"op,omitempty"`
	Id                   *string       `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateAclEntry BulkUpdateAclEntry

// NewBulkUpdateAclEntry instantiates a new BulkUpdateAclEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateAclEntry() *BulkUpdateAclEntry {
	this := BulkUpdateAclEntry{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// NewBulkUpdateAclEntryWithDefaults instantiates a new BulkUpdateAclEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateAclEntryWithDefaults() *BulkUpdateAclEntry {
	this := BulkUpdateAclEntry{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *BulkUpdateAclEntry) GetNegated() int32 {
	if o == nil || o.Negated == nil {
		var ret int32
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntry) GetNegatedOk() (*int32, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasNegated() bool {
	if o != nil && o.Negated != nil {
		return true
	}

	return false
}

// SetNegated gets a reference to the given int32 and assigns it to the Negated field.
func (o *BulkUpdateAclEntry) SetNegated(v int32) {
	o.Negated = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkUpdateAclEntry) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateAclEntry) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *BulkUpdateAclEntry) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *BulkUpdateAclEntry) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *BulkUpdateAclEntry) UnsetComment() {
	o.Comment.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *BulkUpdateAclEntry) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntry) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *BulkUpdateAclEntry) SetIp(v string) {
	o.Ip = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkUpdateAclEntry) GetSubnet() int32 {
	if o == nil || o.Subnet.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateAclEntry) GetSubnetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableInt32 and assigns it to the Subnet field.
func (o *BulkUpdateAclEntry) SetSubnet(v int32) {
	o.Subnet.Set(&v)
}

// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *BulkUpdateAclEntry) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *BulkUpdateAclEntry) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateAclEntry) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntry) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateAclEntry) SetOp(v string) {
	o.Op = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkUpdateAclEntry) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateAclEntry) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkUpdateAclEntry) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkUpdateAclEntry) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateAclEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Negated != nil {
		toSerialize["negated"] = o.Negated
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Subnet.IsSet() {
		toSerialize["subnet"] = o.Subnet.Get()
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BulkUpdateAclEntry) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateAclEntry := _BulkUpdateAclEntry{}

	if err = json.Unmarshal(bytes, &varBulkUpdateAclEntry); err == nil {
		*o = BulkUpdateAclEntry(varBulkUpdateAclEntry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "negated")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "op")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateAclEntry is a helper abstraction for handling nullable bulkupdateaclentry types.
type NullableBulkUpdateAclEntry struct {
	value *BulkUpdateAclEntry
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateAclEntry) Get() *BulkUpdateAclEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateAclEntry) Set(val *BulkUpdateAclEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateAclEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateAclEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateAclEntry returns a pointer to a new instance of NullableBulkUpdateAclEntry.
func NewNullableBulkUpdateAclEntry(val *BulkUpdateAclEntry) *NullableBulkUpdateAclEntry {
	return &NullableBulkUpdateAclEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateAclEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBulkUpdateAclEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
