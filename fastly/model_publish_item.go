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

// PublishItem An individual message.
type PublishItem struct {
	// The channel to publish the message on.
	Channel string `json:"channel"`
	// The ID of the message.
	Id *string `json:"id,omitempty"`
	// The ID of the previous message published on the same channel.
	PrevId               *string            `json:"prev-id,omitempty"`
	Formats              PublishItemFormats `json:"formats"`
	AdditionalProperties map[string]any
}

type _PublishItem PublishItem

// NewPublishItem instantiates a new PublishItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishItem(channel string, formats PublishItemFormats) *PublishItem {
	this := PublishItem{}
	this.Channel = channel
	this.Formats = formats
	return &this
}

// NewPublishItemWithDefaults instantiates a new PublishItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishItemWithDefaults() *PublishItem {
	this := PublishItem{}
	return &this
}

// GetChannel returns the Channel field value
func (o *PublishItem) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *PublishItem) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *PublishItem) SetChannel(v string) {
	o.Channel = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublishItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublishItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublishItem) SetId(v string) {
	o.Id = &v
}

// GetPrevId returns the PrevId field value if set, zero value otherwise.
func (o *PublishItem) GetPrevId() string {
	if o == nil || o.PrevId == nil {
		var ret string
		return ret
	}
	return *o.PrevId
}

// GetPrevIdOk returns a tuple with the PrevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItem) GetPrevIdOk() (*string, bool) {
	if o == nil || o.PrevId == nil {
		return nil, false
	}
	return o.PrevId, true
}

// HasPrevId returns a boolean if a field has been set.
func (o *PublishItem) HasPrevId() bool {
	if o != nil && o.PrevId != nil {
		return true
	}

	return false
}

// SetPrevId gets a reference to the given string and assigns it to the PrevId field.
func (o *PublishItem) SetPrevId(v string) {
	o.PrevId = &v
}

// GetFormats returns the Formats field value
func (o *PublishItem) GetFormats() PublishItemFormats {
	if o == nil {
		var ret PublishItemFormats
		return ret
	}

	return o.Formats
}

// GetFormatsOk returns a tuple with the Formats field value
// and a boolean to check if the value has been set.
func (o *PublishItem) GetFormatsOk() (*PublishItemFormats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formats, true
}

// SetFormats sets field value
func (o *PublishItem) SetFormats(v PublishItemFormats) {
	o.Formats = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PublishItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PrevId != nil {
		toSerialize["prev-id"] = o.PrevId
	}
	if true {
		toSerialize["formats"] = o.Formats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PublishItem) UnmarshalJSON(bytes []byte) (err error) {
	varPublishItem := _PublishItem{}

	if err = json.Unmarshal(bytes, &varPublishItem); err == nil {
		*o = PublishItem(varPublishItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "id")
		delete(additionalProperties, "prev-id")
		delete(additionalProperties, "formats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePublishItem is a helper abstraction for handling nullable publishitem types.
type NullablePublishItem struct {
	value *PublishItem
	isSet bool
}

// Get returns the value.
func (v NullablePublishItem) Get() *PublishItem {
	return v.value
}

// Set modifies the value.
func (v *NullablePublishItem) Set(val *PublishItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePublishItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePublishItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePublishItem returns a pointer to a new instance of NullablePublishItem.
func NewNullablePublishItem(val *PublishItem) *NullablePublishItem {
	return &NullablePublishItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePublishItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePublishItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
