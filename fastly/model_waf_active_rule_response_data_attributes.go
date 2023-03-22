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
	"time"
)

// WafActiveRuleResponseDataAttributes struct for WafActiveRuleResponseDataAttributes
type WafActiveRuleResponseDataAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// The latest rule revision number that is available for the associated rule revision.
	LatestRevision *int32 `json:"latest_revision,omitempty"`
	// Indicates if the associated rule revision is up to date or not.
	Outdated *bool `json:"outdated,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleResponseDataAttributes WafActiveRuleResponseDataAttributes

// NewWafActiveRuleResponseDataAttributes instantiates a new WafActiveRuleResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleResponseDataAttributes() *WafActiveRuleResponseDataAttributes {
	this := WafActiveRuleResponseDataAttributes{}
	return &this
}

// NewWafActiveRuleResponseDataAttributesWithDefaults instantiates a new WafActiveRuleResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleResponseDataAttributesWithDefaults() *WafActiveRuleResponseDataAttributes {
	this := WafActiveRuleResponseDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafActiveRuleResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafActiveRuleResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *WafActiveRuleResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *WafActiveRuleResponseDataAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *WafActiveRuleResponseDataAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafActiveRuleResponseDataAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafActiveRuleResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *WafActiveRuleResponseDataAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *WafActiveRuleResponseDataAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *WafActiveRuleResponseDataAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafActiveRuleResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafActiveRuleResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *WafActiveRuleResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *WafActiveRuleResponseDataAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *WafActiveRuleResponseDataAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetLatestRevision returns the LatestRevision field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAttributes) GetLatestRevision() int32 {
	if o == nil || o.LatestRevision == nil {
		var ret int32
		return ret
	}
	return *o.LatestRevision
}

// GetLatestRevisionOk returns a tuple with the LatestRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAttributes) GetLatestRevisionOk() (*int32, bool) {
	if o == nil || o.LatestRevision == nil {
		return nil, false
	}
	return o.LatestRevision, true
}

// HasLatestRevision returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributes) HasLatestRevision() bool {
	if o != nil && o.LatestRevision != nil {
		return true
	}

	return false
}

// SetLatestRevision gets a reference to the given int32 and assigns it to the LatestRevision field.
func (o *WafActiveRuleResponseDataAttributes) SetLatestRevision(v int32) {
	o.LatestRevision = &v
}

// GetOutdated returns the Outdated field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAttributes) GetOutdated() bool {
	if o == nil || o.Outdated == nil {
		var ret bool
		return ret
	}
	return *o.Outdated
}

// GetOutdatedOk returns a tuple with the Outdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAttributes) GetOutdatedOk() (*bool, bool) {
	if o == nil || o.Outdated == nil {
		return nil, false
	}
	return o.Outdated, true
}

// HasOutdated returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAttributes) HasOutdated() bool {
	if o != nil && o.Outdated != nil {
		return true
	}

	return false
}

// SetOutdated gets a reference to the given bool and assigns it to the Outdated field.
func (o *WafActiveRuleResponseDataAttributes) SetOutdated(v bool) {
	o.Outdated = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.LatestRevision != nil {
		toSerialize["latest_revision"] = o.LatestRevision
	}
	if o.Outdated != nil {
		toSerialize["outdated"] = o.Outdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafActiveRuleResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleResponseDataAttributes := _WafActiveRuleResponseDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleResponseDataAttributes); err == nil {
		*o = WafActiveRuleResponseDataAttributes(varWafActiveRuleResponseDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "latest_revision")
		delete(additionalProperties, "outdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleResponseDataAttributes is a helper abstraction for handling nullable wafactiveruleresponsedataattributes types. 
type NullableWafActiveRuleResponseDataAttributes struct {
	value *WafActiveRuleResponseDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleResponseDataAttributes) Get() *WafActiveRuleResponseDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleResponseDataAttributes) Set(val *WafActiveRuleResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleResponseDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleResponseDataAttributes returns a pointer to a new instance of NullableWafActiveRuleResponseDataAttributes.
func NewNullableWafActiveRuleResponseDataAttributes(val *WafActiveRuleResponseDataAttributes) *NullableWafActiveRuleResponseDataAttributes {
	return &NullableWafActiveRuleResponseDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafActiveRuleResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
