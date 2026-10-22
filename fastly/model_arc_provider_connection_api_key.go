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

// ArcProviderConnectionApiKey A provider connection authenticated with the provider's own API key.
type ArcProviderConnectionApiKey struct {
	// A display name for the provider connection. Cannot match a name reserved by a built-in provider.
	Name string `json:"name"`
	// The AI model identifiers this connection is allowed to use.
	Models []string `json:"models"`
	// The base URL of the provider's API. Must be an absolute `http` or `https` URL and must not end in `/v1`.
	BaseUrl string `json:"base_url"`
	// The provider's secret API key.
	ApiKey string `json:"api_key"`
	// The authentication type for this connection. Omit for the default `api-key` behavior.
	AuthType             *string `json:"auth_type,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcProviderConnectionApiKey ArcProviderConnectionApiKey

// NewArcProviderConnectionApiKey instantiates a new ArcProviderConnectionApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProviderConnectionApiKey(name string, models []string, baseUrl string, apiKey string) *ArcProviderConnectionApiKey {
	this := ArcProviderConnectionApiKey{}
	this.Name = name
	this.Models = models
	this.BaseUrl = baseUrl
	this.ApiKey = apiKey
	var authType string = "api-key"
	this.AuthType = &authType
	return &this
}

// NewArcProviderConnectionApiKeyWithDefaults instantiates a new ArcProviderConnectionApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderConnectionApiKeyWithDefaults() *ArcProviderConnectionApiKey {
	this := ArcProviderConnectionApiKey{}
	var authType string = "api-key"
	this.AuthType = &authType
	return &this
}

// GetName returns the Name field value
func (o *ArcProviderConnectionApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcProviderConnectionApiKey) SetName(v string) {
	o.Name = v
}

// GetModels returns the Models field value
func (o *ArcProviderConnectionApiKey) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionApiKey) GetModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ArcProviderConnectionApiKey) SetModels(v []string) {
	o.Models = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *ArcProviderConnectionApiKey) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionApiKey) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *ArcProviderConnectionApiKey) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetApiKey returns the ApiKey field value
func (o *ArcProviderConnectionApiKey) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionApiKey) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *ArcProviderConnectionApiKey) SetApiKey(v string) {
	o.ApiKey = v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *ArcProviderConnectionApiKey) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionApiKey) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *ArcProviderConnectionApiKey) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *ArcProviderConnectionApiKey) SetAuthType(v string) {
	o.AuthType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["models"] = o.Models
	}
	if true {
		toSerialize["base_url"] = o.BaseUrl
	}
	if true {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionApiKey) UnmarshalJSON(bytes []byte) (err error) {
	varArcProviderConnectionApiKey := _ArcProviderConnectionApiKey{}

	if err = json.Unmarshal(bytes, &varArcProviderConnectionApiKey); err == nil {
		*o = ArcProviderConnectionApiKey(varArcProviderConnectionApiKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "models")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "auth_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProviderConnectionApiKey is a helper abstraction for handling nullable arcproviderconnectionapikey types.
type NullableArcProviderConnectionApiKey struct {
	value *ArcProviderConnectionApiKey
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionApiKey) Get() *ArcProviderConnectionApiKey {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionApiKey) Set(val *ArcProviderConnectionApiKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionApiKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionApiKey returns a pointer to a new instance of NullableArcProviderConnectionApiKey.
func NewNullableArcProviderConnectionApiKey(val *ArcProviderConnectionApiKey) *NullableArcProviderConnectionApiKey {
	return &NullableArcProviderConnectionApiKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
