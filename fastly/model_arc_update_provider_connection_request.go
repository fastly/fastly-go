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

// ArcUpdateProviderConnectionRequest A partial update to a provider connection. Only the fields you include are changed. `models`, `base_url`, and `api_key` apply to `api-key` connections; `customer_role_arn` and `region` apply to `aws-iam` connections.
type ArcUpdateProviderConnectionRequest struct {
	// The AI model identifiers this connection is allowed to use. Must be non-empty if included.
	Models []string `json:"models,omitempty"`
	// The base URL of the provider's API (`api-key` connections only).
	BaseUrl *string `json:"base_url,omitempty"`
	// The provider's secret API key (`api-key` connections only).
	ApiKey *string `json:"api_key,omitempty"`
	// The ARN of the IAM role to assume for Bedrock requests (`aws-iam` connections only).
	CustomerRoleArn *string `json:"customer_role_arn,omitempty"`
	// The AWS region for Bedrock API requests (`aws-iam` connections only).
	Region               *string `json:"region,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcUpdateProviderConnectionRequest ArcUpdateProviderConnectionRequest

// NewArcUpdateProviderConnectionRequest instantiates a new ArcUpdateProviderConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcUpdateProviderConnectionRequest() *ArcUpdateProviderConnectionRequest {
	this := ArcUpdateProviderConnectionRequest{}
	return &this
}

// NewArcUpdateProviderConnectionRequestWithDefaults instantiates a new ArcUpdateProviderConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcUpdateProviderConnectionRequestWithDefaults() *ArcUpdateProviderConnectionRequest {
	this := ArcUpdateProviderConnectionRequest{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *ArcUpdateProviderConnectionRequest) GetModels() []string {
	if o == nil || o.Models == nil {
		var ret []string
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUpdateProviderConnectionRequest) GetModelsOk() ([]string, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *ArcUpdateProviderConnectionRequest) HasModels() bool {
	if o != nil && o.Models != nil {
		return true
	}

	return false
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *ArcUpdateProviderConnectionRequest) SetModels(v []string) {
	o.Models = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ArcUpdateProviderConnectionRequest) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUpdateProviderConnectionRequest) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ArcUpdateProviderConnectionRequest) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ArcUpdateProviderConnectionRequest) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *ArcUpdateProviderConnectionRequest) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUpdateProviderConnectionRequest) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *ArcUpdateProviderConnectionRequest) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *ArcUpdateProviderConnectionRequest) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetCustomerRoleArn returns the CustomerRoleArn field value if set, zero value otherwise.
func (o *ArcUpdateProviderConnectionRequest) GetCustomerRoleArn() string {
	if o == nil || o.CustomerRoleArn == nil {
		var ret string
		return ret
	}
	return *o.CustomerRoleArn
}

// GetCustomerRoleArnOk returns a tuple with the CustomerRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUpdateProviderConnectionRequest) GetCustomerRoleArnOk() (*string, bool) {
	if o == nil || o.CustomerRoleArn == nil {
		return nil, false
	}
	return o.CustomerRoleArn, true
}

// HasCustomerRoleArn returns a boolean if a field has been set.
func (o *ArcUpdateProviderConnectionRequest) HasCustomerRoleArn() bool {
	if o != nil && o.CustomerRoleArn != nil {
		return true
	}

	return false
}

// SetCustomerRoleArn gets a reference to the given string and assigns it to the CustomerRoleArn field.
func (o *ArcUpdateProviderConnectionRequest) SetCustomerRoleArn(v string) {
	o.CustomerRoleArn = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ArcUpdateProviderConnectionRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUpdateProviderConnectionRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ArcUpdateProviderConnectionRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ArcUpdateProviderConnectionRequest) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcUpdateProviderConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Models != nil {
		toSerialize["models"] = o.Models
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.CustomerRoleArn != nil {
		toSerialize["customer_role_arn"] = o.CustomerRoleArn
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcUpdateProviderConnectionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varArcUpdateProviderConnectionRequest := _ArcUpdateProviderConnectionRequest{}

	if err = json.Unmarshal(bytes, &varArcUpdateProviderConnectionRequest); err == nil {
		*o = ArcUpdateProviderConnectionRequest(varArcUpdateProviderConnectionRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "models")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "customer_role_arn")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcUpdateProviderConnectionRequest is a helper abstraction for handling nullable arcupdateproviderconnectionrequest types.
type NullableArcUpdateProviderConnectionRequest struct {
	value *ArcUpdateProviderConnectionRequest
	isSet bool
}

// Get returns the value.
func (v NullableArcUpdateProviderConnectionRequest) Get() *ArcUpdateProviderConnectionRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableArcUpdateProviderConnectionRequest) Set(val *ArcUpdateProviderConnectionRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcUpdateProviderConnectionRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcUpdateProviderConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcUpdateProviderConnectionRequest returns a pointer to a new instance of NullableArcUpdateProviderConnectionRequest.
func NewNullableArcUpdateProviderConnectionRequest(val *ArcUpdateProviderConnectionRequest) *NullableArcUpdateProviderConnectionRequest {
	return &NullableArcUpdateProviderConnectionRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcUpdateProviderConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcUpdateProviderConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
