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

// ArcProviderConnectionResponse A provider connection.
type ArcProviderConnectionResponse struct {
	// The unique identifier of the provider connection.
	Id string `json:"id"`
	// An opaque identifier for the secret material backing this connection. Not the secret itself.
	SecretId string `json:"secret_id"`
	// The display name of the provider connection.
	Name string `json:"name"`
	// The AI model identifiers this connection is allowed to use.
	Models []string `json:"models"`
	// The base URL of the provider's API. Present only for `api-key` connections.
	BaseUrl *string `json:"base_url,omitempty"`
	// The authentication type for this connection.
	AuthType string `json:"auth_type"`
	// The AWS region for Bedrock API requests. Present only for `aws-iam` connections.
	Region *string `json:"region,omitempty"`
	// The date and time the connection was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the connection was last updated, in RFC 3339 format.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _ArcProviderConnectionResponse ArcProviderConnectionResponse

// NewArcProviderConnectionResponse instantiates a new ArcProviderConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProviderConnectionResponse(id string, secretId string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time) *ArcProviderConnectionResponse {
	this := ArcProviderConnectionResponse{}
	this.Id = id
	this.SecretId = secretId
	this.Name = name
	this.Models = models
	this.AuthType = authType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewArcProviderConnectionResponseWithDefaults instantiates a new ArcProviderConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderConnectionResponseWithDefaults() *ArcProviderConnectionResponse {
	this := ArcProviderConnectionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ArcProviderConnectionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcProviderConnectionResponse) SetId(v string) {
	o.Id = v
}

// GetSecretId returns the SecretId field value
func (o *ArcProviderConnectionResponse) GetSecretId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value
func (o *ArcProviderConnectionResponse) SetSecretId(v string) {
	o.SecretId = v
}

// GetName returns the Name field value
func (o *ArcProviderConnectionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcProviderConnectionResponse) SetName(v string) {
	o.Name = v
}

// GetModels returns the Models field value
func (o *ArcProviderConnectionResponse) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ArcProviderConnectionResponse) SetModels(v []string) {
	o.Models = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ArcProviderConnectionResponse) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ArcProviderConnectionResponse) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ArcProviderConnectionResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetAuthType returns the AuthType field value
func (o *ArcProviderConnectionResponse) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ArcProviderConnectionResponse) SetAuthType(v string) {
	o.AuthType = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ArcProviderConnectionResponse) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ArcProviderConnectionResponse) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ArcProviderConnectionResponse) SetRegion(v string) {
	o.Region = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcProviderConnectionResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcProviderConnectionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcProviderConnectionResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcProviderConnectionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["secret_id"] = o.SecretId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["models"] = o.Models
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
	if true {
		toSerialize["auth_type"] = o.AuthType
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcProviderConnectionResponse := _ArcProviderConnectionResponse{}

	if err = json.Unmarshal(bytes, &varArcProviderConnectionResponse); err == nil {
		*o = ArcProviderConnectionResponse(varArcProviderConnectionResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "secret_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "models")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "region")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProviderConnectionResponse is a helper abstraction for handling nullable arcproviderconnectionresponse types.
type NullableArcProviderConnectionResponse struct {
	value *ArcProviderConnectionResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionResponse) Get() *ArcProviderConnectionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionResponse) Set(val *ArcProviderConnectionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionResponse returns a pointer to a new instance of NullableArcProviderConnectionResponse.
func NewNullableArcProviderConnectionResponse(val *ArcProviderConnectionResponse) *NullableArcProviderConnectionResponse {
	return &NullableArcProviderConnectionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
