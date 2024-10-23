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
	ID *string `json:"id,omitempty"`
	// The ID of the previous message published on the same channel.
	PrevID               *string            `json:"prev-id,omitempty"`
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

// GetID returns the ID field value if set, zero value otherwise.
func (o *PublishItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *PublishItem) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *PublishItem) SetID(v string) {
	o.ID = &v
}

// GetPrevID returns the PrevID field value if set, zero value otherwise.
func (o *PublishItem) GetPrevID() string {
	if o == nil || o.PrevID == nil {
		var ret string
		return ret
	}
	return *o.PrevID
}

// GetPrevIDOk returns a tuple with the PrevID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItem) GetPrevIDOk() (*string, bool) {
	if o == nil || o.PrevID == nil {
		return nil, false
	}
	return o.PrevID, true
}

// HasPrevID returns a boolean if a field has been set.
func (o *PublishItem) HasPrevID() bool {
	if o != nil && o.PrevID != nil {
		return true
	}

	return false
}

// SetPrevID gets a reference to the given string and assigns it to the PrevID field.
func (o *PublishItem) SetPrevID(v string) {
	o.PrevID = &v
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
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PrevID != nil {
		toSerialize["prev-id"] = o.PrevID
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
