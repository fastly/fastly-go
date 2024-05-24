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

// BulkUpdateConfigStoreItem struct for BulkUpdateConfigStoreItem
type BulkUpdateConfigStoreItem struct {
	// Item key, maximum 256 characters.
	ItemKey *string `json:"item_key,omitempty"`
	// Item value, maximum 8000 characters.
	ItemValue *string `json:"item_value,omitempty"`
	Op *string `json:"op,omitempty"`
	AdditionalProperties map[string]any
}

type _BulkUpdateConfigStoreItem BulkUpdateConfigStoreItem

// NewBulkUpdateConfigStoreItem instantiates a new BulkUpdateConfigStoreItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateConfigStoreItem() *BulkUpdateConfigStoreItem {
	this := BulkUpdateConfigStoreItem{}
	return &this
}

// NewBulkUpdateConfigStoreItemWithDefaults instantiates a new BulkUpdateConfigStoreItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateConfigStoreItemWithDefaults() *BulkUpdateConfigStoreItem {
	this := BulkUpdateConfigStoreItem{}
	return &this
}

// GetItemKey returns the ItemKey field value if set, zero value otherwise.
func (o *BulkUpdateConfigStoreItem) GetItemKey() string {
	if o == nil || o.ItemKey == nil {
		var ret string
		return ret
	}
	return *o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateConfigStoreItem) GetItemKeyOk() (*string, bool) {
	if o == nil || o.ItemKey == nil {
		return nil, false
	}
	return o.ItemKey, true
}

// HasItemKey returns a boolean if a field has been set.
func (o *BulkUpdateConfigStoreItem) HasItemKey() bool {
	if o != nil && o.ItemKey != nil {
		return true
	}

	return false
}

// SetItemKey gets a reference to the given string and assigns it to the ItemKey field.
func (o *BulkUpdateConfigStoreItem) SetItemKey(v string) {
	o.ItemKey = &v
}

// GetItemValue returns the ItemValue field value if set, zero value otherwise.
func (o *BulkUpdateConfigStoreItem) GetItemValue() string {
	if o == nil || o.ItemValue == nil {
		var ret string
		return ret
	}
	return *o.ItemValue
}

// GetItemValueOk returns a tuple with the ItemValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateConfigStoreItem) GetItemValueOk() (*string, bool) {
	if o == nil || o.ItemValue == nil {
		return nil, false
	}
	return o.ItemValue, true
}

// HasItemValue returns a boolean if a field has been set.
func (o *BulkUpdateConfigStoreItem) HasItemValue() bool {
	if o != nil && o.ItemValue != nil {
		return true
	}

	return false
}

// SetItemValue gets a reference to the given string and assigns it to the ItemValue field.
func (o *BulkUpdateConfigStoreItem) SetItemValue(v string) {
	o.ItemValue = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *BulkUpdateConfigStoreItem) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateConfigStoreItem) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *BulkUpdateConfigStoreItem) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *BulkUpdateConfigStoreItem) SetOp(v string) {
	o.Op = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BulkUpdateConfigStoreItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ItemKey != nil {
		toSerialize["item_key"] = o.ItemKey
	}
	if o.ItemValue != nil {
		toSerialize["item_value"] = o.ItemValue
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BulkUpdateConfigStoreItem) UnmarshalJSON(bytes []byte) (err error) {
	varBulkUpdateConfigStoreItem := _BulkUpdateConfigStoreItem{}

	if err = json.Unmarshal(bytes, &varBulkUpdateConfigStoreItem); err == nil {
		*o = BulkUpdateConfigStoreItem(varBulkUpdateConfigStoreItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_key")
		delete(additionalProperties, "item_value")
		delete(additionalProperties, "op")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBulkUpdateConfigStoreItem is a helper abstraction for handling nullable bulkupdateconfigstoreitem types. 
type NullableBulkUpdateConfigStoreItem struct {
	value *BulkUpdateConfigStoreItem
	isSet bool
}

// Get returns the value.
func (v NullableBulkUpdateConfigStoreItem) Get() *BulkUpdateConfigStoreItem {
	return v.value
}

// Set modifies the value.
func (v *NullableBulkUpdateConfigStoreItem) Set(val *BulkUpdateConfigStoreItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBulkUpdateConfigStoreItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBulkUpdateConfigStoreItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBulkUpdateConfigStoreItem returns a pointer to a new instance of NullableBulkUpdateConfigStoreItem.
func NewNullableBulkUpdateConfigStoreItem(val *BulkUpdateConfigStoreItem) *NullableBulkUpdateConfigStoreItem {
	return &NullableBulkUpdateConfigStoreItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBulkUpdateConfigStoreItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBulkUpdateConfigStoreItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
