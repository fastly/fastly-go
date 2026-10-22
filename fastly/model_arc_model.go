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

// ArcModel An AI model available through AI Runtime Control.
type ArcModel struct {
	// The model identifier.
	Id string `json:"id"`
	// The human-readable name of the model.
	DisplayName string `json:"display_name"`
	// The ID of the provider this model belongs to.
	ProviderId           string `json:"provider_id"`
	AdditionalProperties map[string]any
}

type _ArcModel ArcModel

// NewArcModel instantiates a new ArcModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcModel(id string, displayName string, providerId string) *ArcModel {
	this := ArcModel{}
	this.Id = id
	this.DisplayName = displayName
	this.ProviderId = providerId
	return &this
}

// NewArcModelWithDefaults instantiates a new ArcModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcModelWithDefaults() *ArcModel {
	this := ArcModel{}
	return &this
}

// GetId returns the Id field value
func (o *ArcModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcModel) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *ArcModel) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ArcModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ArcModel) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetProviderId returns the ProviderId field value
func (o *ArcModel) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *ArcModel) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *ArcModel) SetProviderId(v string) {
	o.ProviderId = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["provider_id"] = o.ProviderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcModel) UnmarshalJSON(bytes []byte) (err error) {
	varArcModel := _ArcModel{}

	if err = json.Unmarshal(bytes, &varArcModel); err == nil {
		*o = ArcModel(varArcModel)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "provider_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcModel is a helper abstraction for handling nullable arcmodel types.
type NullableArcModel struct {
	value *ArcModel
	isSet bool
}

// Get returns the value.
func (v NullableArcModel) Get() *ArcModel {
	return v.value
}

// Set modifies the value.
func (v *NullableArcModel) Set(val *ArcModel) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcModel) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcModel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcModel returns a pointer to a new instance of NullableArcModel.
func NewNullableArcModel(val *ArcModel) *NullableArcModel {
	return &NullableArcModel{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
