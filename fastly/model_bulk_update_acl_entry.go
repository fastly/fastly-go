// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// BulkUpdateACLEntry struct for BulkUpdateACLEntry
type BulkUpdateACLEntry struct {
	// Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.
	Negated *int32 `json:"negated,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// An IP address.
	IP *string `json:"ip,omitempty"`
	// Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied.
	Subnet NullableInt32 `json:"subnet,omitempty"`
	Op *string `json:"op,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateACLEntry BulkUpdateACLEntry

// NewBulkUpdateACLEntry instantiates a new BulkUpdateACLEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateACLEntry() *BulkUpdateACLEntry {
	this := BulkUpdateACLEntry{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// NewBulkUpdateACLEntryWithDefaults instantiates a new BulkUpdateACLEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateACLEntryWithDefaults() *BulkUpdateACLEntry {
	this := BulkUpdateACLEntry{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *BulkUpdateACLEntry) GetNegated() int32 {
	if o == nil || o.Negated == nil {
		var ret int32
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntry) GetNegatedOk() (*int32, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasNegated() bool {
	if o != nil && o.Negated != nil {
		return true
	}

	return false
}

// SetNegated gets a reference to the given int32 and assigns it to the Negated field.
func (o *BulkUpdateACLEntry) SetNegated(v int32) {
	o.Negated = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkUpdateACLEntry) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateACLEntry) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *BulkUpdateACLEntry) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *BulkUpdateACLEntry) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *BulkUpdateACLEntry) UnsetComment() {
	o.Comment.Unset()
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *BulkUpdateACLEntry) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntry) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *BulkUpdateACLEntry) SetIP(v string) {
	o.IP = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkUpdateACLEntry) GetSubnet() int32 {
	if o == nil || o.Subnet.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateACLEntry) GetSubnetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableInt32 and assigns it to the Subnet field.
func (o *BulkUpdateACLEntry) SetSubnet(v int32) {
	o.Subnet.Set(&v)
}
// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *BulkUpdateACLEntry) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *BulkUpdateACLEntry) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateACLEntry) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntry) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateACLEntry) SetOp(v string) {
	o.Op = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BulkUpdateACLEntry) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateACLEntry) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BulkUpdateACLEntry) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BulkUpdateACLEntry) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateACLEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Negated != nil {
		toSerialize["negated"] = o.Negated
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.IP != nil {
		toSerialize["ip"] = o.IP
	}
	if o.Subnet.IsSet() {
		toSerialize["subnet"] = o.Subnet.Get()
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BulkUpdateACLEntry) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateACLEntry := _BulkUpdateACLEntry{}

	if err = json.Unmarshal(bytes, &varBulkUpdateACLEntry); err == nil {
		*o = BulkUpdateACLEntry(varBulkUpdateACLEntry)
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

// NullableBulkUpdateACLEntry is a helper abstraction for handling nullable bulkupdateaclentry types. 
type NullableBulkUpdateACLEntry struct {
	value *BulkUpdateACLEntry
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateACLEntry) Get() *BulkUpdateACLEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateACLEntry) Set(val *BulkUpdateACLEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateACLEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateACLEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateACLEntry returns a pointer to a new instance of NullableBulkUpdateACLEntry.
func NewNullableBulkUpdateACLEntry(val *BulkUpdateACLEntry) *NullableBulkUpdateACLEntry {
	return &NullableBulkUpdateACLEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateACLEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateACLEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
