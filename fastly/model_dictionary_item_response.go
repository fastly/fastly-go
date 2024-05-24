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

// DictionaryItemResponse struct for DictionaryItemResponse
type DictionaryItemResponse struct {
	// Item key, maximum 256 characters.
	ItemKey *string `json:"item_key,omitempty"`
	// Item value, maximum 8000 characters.
	ItemValue *string `json:"item_value,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	DictionaryID *string `json:"dictionary_id,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _DictionaryItemResponse DictionaryItemResponse

// NewDictionaryItemResponse instantiates a new DictionaryItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaryItemResponse() *DictionaryItemResponse {
	this := DictionaryItemResponse{}
	return &this
}

// NewDictionaryItemResponseWithDefaults instantiates a new DictionaryItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryItemResponseWithDefaults() *DictionaryItemResponse {
	this := DictionaryItemResponse{}
	return &this
}

// GetItemKey returns the ItemKey field value if set, zero value otherwise.
func (o *DictionaryItemResponse) GetItemKey() string {
	if o == nil || o.ItemKey == nil {
		var ret string
		return ret
	}
	return *o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponse) GetItemKeyOk() (*string, bool) {
	if o == nil || o.ItemKey == nil {
		return nil, false
	}
	return o.ItemKey, true
}

// HasItemKey returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasItemKey() bool {
	if o != nil && o.ItemKey != nil {
		return true
	}

	return false
}

// SetItemKey gets a reference to the given string and assigns it to the ItemKey field.
func (o *DictionaryItemResponse) SetItemKey(v string) {
	o.ItemKey = &v
}

// GetItemValue returns the ItemValue field value if set, zero value otherwise.
func (o *DictionaryItemResponse) GetItemValue() string {
	if o == nil || o.ItemValue == nil {
		var ret string
		return ret
	}
	return *o.ItemValue
}

// GetItemValueOk returns a tuple with the ItemValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponse) GetItemValueOk() (*string, bool) {
	if o == nil || o.ItemValue == nil {
		return nil, false
	}
	return o.ItemValue, true
}

// HasItemValue returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasItemValue() bool {
	if o != nil && o.ItemValue != nil {
		return true
	}

	return false
}

// SetItemValue gets a reference to the given string and assigns it to the ItemValue field.
func (o *DictionaryItemResponse) SetItemValue(v string) {
	o.ItemValue = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DictionaryItemResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DictionaryItemResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *DictionaryItemResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *DictionaryItemResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *DictionaryItemResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DictionaryItemResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DictionaryItemResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *DictionaryItemResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *DictionaryItemResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *DictionaryItemResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DictionaryItemResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DictionaryItemResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *DictionaryItemResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *DictionaryItemResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *DictionaryItemResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetDictionaryID returns the DictionaryID field value if set, zero value otherwise.
func (o *DictionaryItemResponse) GetDictionaryID() string {
	if o == nil || o.DictionaryID == nil {
		var ret string
		return ret
	}
	return *o.DictionaryID
}

// GetDictionaryIDOk returns a tuple with the DictionaryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponse) GetDictionaryIDOk() (*string, bool) {
	if o == nil || o.DictionaryID == nil {
		return nil, false
	}
	return o.DictionaryID, true
}

// HasDictionaryID returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasDictionaryID() bool {
	if o != nil && o.DictionaryID != nil {
		return true
	}

	return false
}

// SetDictionaryID gets a reference to the given string and assigns it to the DictionaryID field.
func (o *DictionaryItemResponse) SetDictionaryID(v string) {
	o.DictionaryID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DictionaryItemResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DictionaryItemResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DictionaryItemResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DictionaryItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ItemKey != nil {
		toSerialize["item_key"] = o.ItemKey
	}
	if o.ItemValue != nil {
		toSerialize["item_value"] = o.ItemValue
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
	if o.DictionaryID != nil {
		toSerialize["dictionary_id"] = o.DictionaryID
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
func (o *DictionaryItemResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDictionaryItemResponse := _DictionaryItemResponse{}

	if err = json.Unmarshal(bytes, &varDictionaryItemResponse); err == nil {
		*o = DictionaryItemResponse(varDictionaryItemResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_key")
		delete(additionalProperties, "item_value")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "dictionary_id")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDictionaryItemResponse is a helper abstraction for handling nullable dictionaryitemresponse types. 
type NullableDictionaryItemResponse struct {
	value *DictionaryItemResponse
	isSet bool
}

// Get returns the value.
func (v NullableDictionaryItemResponse) Get() *DictionaryItemResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDictionaryItemResponse) Set(val *DictionaryItemResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDictionaryItemResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDictionaryItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDictionaryItemResponse returns a pointer to a new instance of NullableDictionaryItemResponse.
func NewNullableDictionaryItemResponse(val *DictionaryItemResponse) *NullableDictionaryItemResponse {
	return &NullableDictionaryItemResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDictionaryItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDictionaryItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
