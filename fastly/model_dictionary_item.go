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

// DictionaryItem struct for DictionaryItem
type DictionaryItem struct {
	// Item key, maximum 256 characters.
	ItemKey *string `json:"item_key,omitempty"`
	// Item value, maximum 8000 characters.
	ItemValue *string `json:"item_value,omitempty"`
	AdditionalProperties map[string]any
}

type _DictionaryItem DictionaryItem

// NewDictionaryItem instantiates a new DictionaryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaryItem() *DictionaryItem {
	this := DictionaryItem{}
	return &this
}

// NewDictionaryItemWithDefaults instantiates a new DictionaryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryItemWithDefaults() *DictionaryItem {
	this := DictionaryItem{}
	return &this
}

// GetItemKey returns the ItemKey field value if set, zero value otherwise.
func (o *DictionaryItem) GetItemKey() string {
	if o == nil || o.ItemKey == nil {
		var ret string
		return ret
	}
	return *o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItem) GetItemKeyOk() (*string, bool) {
	if o == nil || o.ItemKey == nil {
		return nil, false
	}
	return o.ItemKey, true
}

// HasItemKey returns a boolean if a field has been set.
func (o *DictionaryItem) HasItemKey() bool {
	if o != nil && o.ItemKey != nil {
		return true
	}

	return false
}

// SetItemKey gets a reference to the given string and assigns it to the ItemKey field.
func (o *DictionaryItem) SetItemKey(v string) {
	o.ItemKey = &v
}

// GetItemValue returns the ItemValue field value if set, zero value otherwise.
func (o *DictionaryItem) GetItemValue() string {
	if o == nil || o.ItemValue == nil {
		var ret string
		return ret
	}
	return *o.ItemValue
}

// GetItemValueOk returns a tuple with the ItemValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItem) GetItemValueOk() (*string, bool) {
	if o == nil || o.ItemValue == nil {
		return nil, false
	}
	return o.ItemValue, true
}

// HasItemValue returns a boolean if a field has been set.
func (o *DictionaryItem) HasItemValue() bool {
	if o != nil && o.ItemValue != nil {
		return true
	}

	return false
}

// SetItemValue gets a reference to the given string and assigns it to the ItemValue field.
func (o *DictionaryItem) SetItemValue(v string) {
	o.ItemValue = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DictionaryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ItemKey != nil {
		toSerialize["item_key"] = o.ItemKey
	}
	if o.ItemValue != nil {
		toSerialize["item_value"] = o.ItemValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *DictionaryItem) UnmarshalJSON(bytes []byte) (err error) {
	varDictionaryItem := _DictionaryItem{}

	if err = json.Unmarshal(bytes, &varDictionaryItem); err == nil {
		*o = DictionaryItem(varDictionaryItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_key")
		delete(additionalProperties, "item_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDictionaryItem is a helper abstraction for handling nullable dictionaryitem types. 
type NullableDictionaryItem struct {
	value *DictionaryItem
	isSet bool
}

// Get returns the value.
func (v NullableDictionaryItem) Get() *DictionaryItem {
	return v.value
}

// Set modifies the value.
func (v *NullableDictionaryItem) Set(val *DictionaryItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDictionaryItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDictionaryItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDictionaryItem returns a pointer to a new instance of NullableDictionaryItem.
func NewNullableDictionaryItem(val *DictionaryItem) *NullableDictionaryItem {
	return &NullableDictionaryItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDictionaryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDictionaryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
