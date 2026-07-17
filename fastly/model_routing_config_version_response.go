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

// RoutingConfigVersionResponse All attributes for a routing config version response.
type RoutingConfigVersionResponse struct {
	// Alphanumeric string identifying the version.
	Id *string `json:"id,omitempty"`
	// A freeform comment describing the version.
	Comment *string `json:"comment,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Timestamp of when the version was most recently activated. `null` if the version has never been activated.
	ActivatedAt          NullableTime `json:"activated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _RoutingConfigVersionResponse RoutingConfigVersionResponse

// NewRoutingConfigVersionResponse instantiates a new RoutingConfigVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfigVersionResponse() *RoutingConfigVersionResponse {
	this := RoutingConfigVersionResponse{}
	return &this
}

// NewRoutingConfigVersionResponseWithDefaults instantiates a new RoutingConfigVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigVersionResponseWithDefaults() *RoutingConfigVersionResponse {
	this := RoutingConfigVersionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoutingConfigVersionResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigVersionResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoutingConfigVersionResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoutingConfigVersionResponse) SetId(v string) {
	o.Id = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RoutingConfigVersionResponse) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigVersionResponse) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RoutingConfigVersionResponse) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RoutingConfigVersionResponse) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigVersionResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigVersionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RoutingConfigVersionResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *RoutingConfigVersionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RoutingConfigVersionResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RoutingConfigVersionResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoutingConfigVersionResponse) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt.Get()
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingConfigVersionResponse) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivatedAt.Get(), o.ActivatedAt.IsSet()
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *RoutingConfigVersionResponse) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt.IsSet() {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given NullableTime and assigns it to the ActivatedAt field.
func (o *RoutingConfigVersionResponse) SetActivatedAt(v time.Time) {
	o.ActivatedAt.Set(&v)
}

// SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil
func (o *RoutingConfigVersionResponse) SetActivatedAtNil() {
	o.ActivatedAt.Set(nil)
}

// UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
func (o *RoutingConfigVersionResponse) UnsetActivatedAt() {
	o.ActivatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfigVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ActivatedAt.IsSet() {
		toSerialize["activated_at"] = o.ActivatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RoutingConfigVersionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfigVersionResponse := _RoutingConfigVersionResponse{}

	if err = json.Unmarshal(bytes, &varRoutingConfigVersionResponse); err == nil {
		*o = RoutingConfigVersionResponse(varRoutingConfigVersionResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "activated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfigVersionResponse is a helper abstraction for handling nullable routingconfigversionresponse types.
type NullableRoutingConfigVersionResponse struct {
	value *RoutingConfigVersionResponse
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigVersionResponse) Get() *RoutingConfigVersionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigVersionResponse) Set(val *RoutingConfigVersionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigVersionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigVersionResponse returns a pointer to a new instance of NullableRoutingConfigVersionResponse.
func NewNullableRoutingConfigVersionResponse(val *RoutingConfigVersionResponse) *NullableRoutingConfigVersionResponse {
	return &NullableRoutingConfigVersionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
