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

// ArcProviderConnectionListItemResponse struct for ArcProviderConnectionListItemResponse
type ArcProviderConnectionListItemResponse struct {
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
	UpdatedAt time.Time `json:"updated_at"`
	// The date and time the connection was deleted, in RFC 3339 format. Present only for deleted connections, which are returned only when `include_deleted` is `true`.
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcProviderConnectionListItemResponse ArcProviderConnectionListItemResponse

// NewArcProviderConnectionListItemResponse instantiates a new ArcProviderConnectionListItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProviderConnectionListItemResponse(id string, secretId string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time) *ArcProviderConnectionListItemResponse {
	this := ArcProviderConnectionListItemResponse{}
	this.Id = id
	this.SecretId = secretId
	this.Name = name
	this.Models = models
	this.AuthType = authType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewArcProviderConnectionListItemResponseWithDefaults instantiates a new ArcProviderConnectionListItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderConnectionListItemResponseWithDefaults() *ArcProviderConnectionListItemResponse {
	this := ArcProviderConnectionListItemResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ArcProviderConnectionListItemResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcProviderConnectionListItemResponse) SetId(v string) {
	o.Id = v
}

// GetSecretId returns the SecretId field value
func (o *ArcProviderConnectionListItemResponse) GetSecretId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value
func (o *ArcProviderConnectionListItemResponse) SetSecretId(v string) {
	o.SecretId = v
}

// GetName returns the Name field value
func (o *ArcProviderConnectionListItemResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcProviderConnectionListItemResponse) SetName(v string) {
	o.Name = v
}

// GetModels returns the Models field value
func (o *ArcProviderConnectionListItemResponse) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ArcProviderConnectionListItemResponse) SetModels(v []string) {
	o.Models = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ArcProviderConnectionListItemResponse) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ArcProviderConnectionListItemResponse) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ArcProviderConnectionListItemResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetAuthType returns the AuthType field value
func (o *ArcProviderConnectionListItemResponse) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ArcProviderConnectionListItemResponse) SetAuthType(v string) {
	o.AuthType = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ArcProviderConnectionListItemResponse) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ArcProviderConnectionListItemResponse) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ArcProviderConnectionListItemResponse) SetRegion(v string) {
	o.Region = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcProviderConnectionListItemResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcProviderConnectionListItemResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcProviderConnectionListItemResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcProviderConnectionListItemResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ArcProviderConnectionListItemResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionListItemResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArcProviderConnectionListItemResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ArcProviderConnectionListItemResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionListItemResponse) MarshalJSON() ([]byte, error) {
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
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionListItemResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcProviderConnectionListItemResponse := _ArcProviderConnectionListItemResponse{}

	if err = json.Unmarshal(bytes, &varArcProviderConnectionListItemResponse); err == nil {
		*o = ArcProviderConnectionListItemResponse(varArcProviderConnectionListItemResponse)
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
		delete(additionalProperties, "deleted_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProviderConnectionListItemResponse is a helper abstraction for handling nullable arcproviderconnectionlistitemresponse types.
type NullableArcProviderConnectionListItemResponse struct {
	value *ArcProviderConnectionListItemResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionListItemResponse) Get() *ArcProviderConnectionListItemResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionListItemResponse) Set(val *ArcProviderConnectionListItemResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionListItemResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionListItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionListItemResponse returns a pointer to a new instance of NullableArcProviderConnectionListItemResponse.
func NewNullableArcProviderConnectionListItemResponse(val *ArcProviderConnectionListItemResponse) *NullableArcProviderConnectionListItemResponse {
	return &NullableArcProviderConnectionListItemResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionListItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionListItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
