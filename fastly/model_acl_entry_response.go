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

// AclEntryResponse struct for AclEntryResponse
type AclEntryResponse struct {
	// Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.
	Negated *int32 `json:"negated,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// An IP address.
	Ip *string `json:"ip,omitempty"`
	// Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied.
	Subnet NullableInt32 `json:"subnet,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	AclId                *string      `json:"acl_id,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	ServiceId            *string      `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _AclEntryResponse AclEntryResponse

// NewAclEntryResponse instantiates a new AclEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclEntryResponse() *AclEntryResponse {
	this := AclEntryResponse{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// NewAclEntryResponseWithDefaults instantiates a new AclEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclEntryResponseWithDefaults() *AclEntryResponse {
	this := AclEntryResponse{}
	var negated int32 = 0
	this.Negated = &negated
	return &this
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *AclEntryResponse) GetNegated() int32 {
	if o == nil || o.Negated == nil {
		var ret int32
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponse) GetNegatedOk() (*int32, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *AclEntryResponse) HasNegated() bool {
	if o != nil && o.Negated != nil {
		return true
	}

	return false
}

// SetNegated gets a reference to the given int32 and assigns it to the Negated field.
func (o *AclEntryResponse) SetNegated(v int32) {
	o.Negated = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AclEntryResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclEntryResponse) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *AclEntryResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *AclEntryResponse) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *AclEntryResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *AclEntryResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AclEntryResponse) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponse) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *AclEntryResponse) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AclEntryResponse) SetIp(v string) {
	o.Ip = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AclEntryResponse) GetSubnet() int32 {
	if o == nil || o.Subnet.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclEntryResponse) GetSubnetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *AclEntryResponse) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableInt32 and assigns it to the Subnet field.
func (o *AclEntryResponse) SetSubnet(v int32) {
	o.Subnet.Set(&v)
}

// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *AclEntryResponse) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *AclEntryResponse) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AclEntryResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclEntryResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AclEntryResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *AclEntryResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *AclEntryResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *AclEntryResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AclEntryResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclEntryResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AclEntryResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *AclEntryResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *AclEntryResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *AclEntryResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AclEntryResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclEntryResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AclEntryResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *AclEntryResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *AclEntryResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *AclEntryResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *AclEntryResponse) GetAclId() string {
	if o == nil || o.AclId == nil {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponse) GetAclIdOk() (*string, bool) {
	if o == nil || o.AclId == nil {
		return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *AclEntryResponse) HasAclId() bool {
	if o != nil && o.AclId != nil {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *AclEntryResponse) SetAclId(v string) {
	o.AclId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AclEntryResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AclEntryResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AclEntryResponse) SetId(v string) {
	o.Id = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *AclEntryResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclEntryResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *AclEntryResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *AclEntryResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AclEntryResponse) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.AclId != nil {
		toSerialize["acl_id"] = o.AclId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AclEntryResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAclEntryResponse := _AclEntryResponse{}

	if err = json.Unmarshal(bytes, &varAclEntryResponse); err == nil {
		*o = AclEntryResponse(varAclEntryResponse)
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

// NullableAclEntryResponse is a helper abstraction for handling nullable aclentryresponse types.
type NullableAclEntryResponse struct {
	value *AclEntryResponse
	isSet bool
}

// Get returns the value.
func (v NullableAclEntryResponse) Get() *AclEntryResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAclEntryResponse) Set(val *AclEntryResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAclEntryResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAclEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAclEntryResponse returns a pointer to a new instance of NullableAclEntryResponse.
func NewNullableAclEntryResponse(val *AclEntryResponse) *NullableAclEntryResponse {
	return &NullableAclEntryResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAclEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAclEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
