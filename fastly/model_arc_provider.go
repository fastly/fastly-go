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

// ArcProvider An AI provider supported by AI Runtime Control.
type ArcProvider struct {
	// The provider identifier.
	ID string `json:"id"`
	// The human-readable name of the provider.
	DisplayName string `json:"display_name"`
	// The default API base URL for the provider.
	DefaultBaseURL string `json:"default_base_url"`
	// The models available from this provider.
	Models               []ArcModel `json:"models"`
	AdditionalProperties map[string]any
}

type _ArcProvider ArcProvider

// NewArcProvider instantiates a new ArcProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProvider(id string, displayName string, defaultBaseURL string, models []ArcModel) *ArcProvider {
	this := ArcProvider{}
	this.ID = id
	this.DisplayName = displayName
	this.DefaultBaseURL = defaultBaseURL
	this.Models = models
	return &this
}

// NewArcProviderWithDefaults instantiates a new ArcProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderWithDefaults() *ArcProvider {
	this := ArcProvider{}
	return &this
}

// GetID returns the ID field value
func (o *ArcProvider) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *ArcProvider) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *ArcProvider) SetID(v string) {
	o.ID = v
}

// GetDisplayName returns the DisplayName field value
func (o *ArcProvider) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ArcProvider) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ArcProvider) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDefaultBaseURL returns the DefaultBaseURL field value
func (o *ArcProvider) GetDefaultBaseURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBaseURL
}

// GetDefaultBaseURLOk returns a tuple with the DefaultBaseURL field value
// and a boolean to check if the value has been set.
func (o *ArcProvider) GetDefaultBaseURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBaseURL, true
}

// SetDefaultBaseURL sets field value
func (o *ArcProvider) SetDefaultBaseURL(v string) {
	o.DefaultBaseURL = v
}

// GetModels returns the Models field value
func (o *ArcProvider) GetModels() []ArcModel {
	if o == nil {
		var ret []ArcModel
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ArcProvider) GetModelsOk() ([]ArcModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ArcProvider) SetModels(v []ArcModel) {
	o.Models = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["default_base_url"] = o.DefaultBaseURL
	}
	if true {
		toSerialize["models"] = o.Models
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProvider) UnmarshalJSON(bytes []byte) (err error) {
	varArcProvider := _ArcProvider{}

	if err = json.Unmarshal(bytes, &varArcProvider); err == nil {
		*o = ArcProvider(varArcProvider)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "default_base_url")
		delete(additionalProperties, "models")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProvider is a helper abstraction for handling nullable arcprovider types.
type NullableArcProvider struct {
	value *ArcProvider
	isSet bool
}

// Get returns the value.
func (v NullableArcProvider) Get() *ArcProvider {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProvider) Set(val *ArcProvider) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProvider) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProvider) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProvider returns a pointer to a new instance of NullableArcProvider.
func NewNullableArcProvider(val *ArcProvider) *NullableArcProvider {
	return &NullableArcProvider{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
