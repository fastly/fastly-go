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

// ACLEntryResponse struct for ACLEntryResponse
type ACLEntryResponse struct {
	// Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.
	Negated *int32 `json:"negated,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// An IP address.
	IP *string `json:"ip,omitempty"`
	// Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied.
	Subnet NullableInt32 `json:"subnet,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	ACLID                *string      `json:"acl_id,omitempty"`
	ID                   *string      `json:"id,omitempty"`
	ServiceID            *string      `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ACLEntryResponse ACLEntryResponse

// NewACLEntryResponse instantiates a new ACLEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLEntryResponse() *ACLEntryResponse {
	this := ACLEntryResponse{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// NewACLEntryResponseWithDefaults instantiates a new ACLEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLEntryResponseWithDefaults() *ACLEntryResponse {
	this := ACLEntryResponse{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *ACLEntryResponse) GetNegated() int32 {
	if o == nil || o.Negated == nil {
		var ret int32
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponse) GetNegatedOk() (*int32, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasNegated() bool {
	if o != nil && o.Negated != nil {
		return true
	}

	return false
}

// SetNegated gets a reference to the given int32 and assigns it to the Negated field.
func (o *ACLEntryResponse) SetNegated(v int32) {
	o.Negated = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ACLEntryResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ACLEntryResponse) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ACLEntryResponse) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ACLEntryResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ACLEntryResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *ACLEntryResponse) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponse) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *ACLEntryResponse) SetIP(v string) {
	o.IP = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ACLEntryResponse) GetSubnet() int32 {
	if o == nil || o.Subnet.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ACLEntryResponse) GetSubnetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableInt32 and assigns it to the Subnet field.
func (o *ACLEntryResponse) SetSubnet(v int32) {
	o.Subnet.Set(&v)
}

// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *ACLEntryResponse) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *ACLEntryResponse) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ACLEntryResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ACLEntryResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ACLEntryResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ACLEntryResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ACLEntryResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ACLEntryResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ACLEntryResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ACLEntryResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ACLEntryResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ACLEntryResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ACLEntryResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ACLEntryResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ACLEntryResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ACLEntryResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ACLEntryResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetACLID returns the ACLID field value if set, zero value otherwise.
func (o *ACLEntryResponse) GetACLID() string {
	if o == nil || o.ACLID == nil {
		var ret string
		return ret
	}
	return *o.ACLID
}

// GetACLIDOk returns a tuple with the ACLID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponse) GetACLIDOk() (*string, bool) {
	if o == nil || o.ACLID == nil {
		return nil, false
	}
	return o.ACLID, true
}

// HasACLID returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasACLID() bool {
	if o != nil && o.ACLID != nil {
		return true
	}

	return false
}

// SetACLID gets a reference to the given string and assigns it to the ACLID field.
func (o *ACLEntryResponse) SetACLID(v string) {
	o.ACLID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ACLEntryResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ACLEntryResponse) SetID(v string) {
	o.ID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ACLEntryResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLEntryResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ACLEntryResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ACLEntryResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ACLEntryResponse) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ACLID != nil {
		toSerialize["acl_id"] = o.ACLID
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ACLEntryResponse) UnmarshalJSON(bytes []byte) (err error) {
	varACLEntryResponse := _ACLEntryResponse{}

	if err = json.Unmarshal(bytes, &varACLEntryResponse); err == nil {
		*o = ACLEntryResponse(varACLEntryResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "negated")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "acl_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableACLEntryResponse is a helper abstraction for handling nullable aclentryresponse types.
type NullableACLEntryResponse struct {
	value *ACLEntryResponse
	isSet bool
}

// Get returns the value.
func (v NullableACLEntryResponse) Get() *ACLEntryResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableACLEntryResponse) Set(val *ACLEntryResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableACLEntryResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableACLEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableACLEntryResponse returns a pointer to a new instance of NullableACLEntryResponse.
func NewNullableACLEntryResponse(val *ACLEntryResponse) *NullableACLEntryResponse {
	return &NullableACLEntryResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableACLEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableACLEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
