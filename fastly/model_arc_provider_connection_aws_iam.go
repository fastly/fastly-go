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

// ArcProviderConnectionAwsIam A provider connection to Amazon Bedrock authenticated with an AWS IAM role. Requires AWS IAM authentication to be enabled on your account.
type ArcProviderConnectionAwsIam struct {
	// The provider name. Must be `bedrock` for `aws-iam` connections.
	Name string `json:"name"`
	// The Bedrock model identifiers this connection is allowed to use.
	Models []string `json:"models"`
	// The authentication type for this connection.
	AuthType string `json:"auth_type"`
	// The ARN of the IAM role in your AWS account that AI Runtime Control assumes to access Bedrock. See [get Bedrock setup information](https://www.fastly.com/documentation/reference/api/ai-runtime-control/providers#arc-get-bedrock-setup) for the role ARN and external ID your trust policy must reference.
	CustomerRoleArn string `json:"customer_role_arn"`
	// The AWS region for Bedrock API requests.
	Region               string `json:"region"`
	AdditionalProperties map[string]any
}

type _ArcProviderConnectionAwsIam ArcProviderConnectionAwsIam

// NewArcProviderConnectionAwsIam instantiates a new ArcProviderConnectionAwsIam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProviderConnectionAwsIam(name string, models []string, authType string, customerRoleArn string, region string) *ArcProviderConnectionAwsIam {
	this := ArcProviderConnectionAwsIam{}
	this.Name = name
	this.Models = models
	this.AuthType = authType
	this.CustomerRoleArn = customerRoleArn
	this.Region = region
	return &this
}

// NewArcProviderConnectionAwsIamWithDefaults instantiates a new ArcProviderConnectionAwsIam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProviderConnectionAwsIamWithDefaults() *ArcProviderConnectionAwsIam {
	this := ArcProviderConnectionAwsIam{}
	return &this
}

// GetName returns the Name field value
func (o *ArcProviderConnectionAwsIam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionAwsIam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcProviderConnectionAwsIam) SetName(v string) {
	o.Name = v
}

// GetModels returns the Models field value
func (o *ArcProviderConnectionAwsIam) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionAwsIam) GetModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ArcProviderConnectionAwsIam) SetModels(v []string) {
	o.Models = v
}

// GetAuthType returns the AuthType field value
func (o *ArcProviderConnectionAwsIam) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionAwsIam) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ArcProviderConnectionAwsIam) SetAuthType(v string) {
	o.AuthType = v
}

// GetCustomerRoleArn returns the CustomerRoleArn field value
func (o *ArcProviderConnectionAwsIam) GetCustomerRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerRoleArn
}

// GetCustomerRoleArnOk returns a tuple with the CustomerRoleArn field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionAwsIam) GetCustomerRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerRoleArn, true
}

// SetCustomerRoleArn sets field value
func (o *ArcProviderConnectionAwsIam) SetCustomerRoleArn(v string) {
	o.CustomerRoleArn = v
}

// GetRegion returns the Region field value
func (o *ArcProviderConnectionAwsIam) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ArcProviderConnectionAwsIam) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ArcProviderConnectionAwsIam) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionAwsIam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["models"] = o.Models
	}
	if true {
		toSerialize["auth_type"] = o.AuthType
	}
	if true {
		toSerialize["customer_role_arn"] = o.CustomerRoleArn
	}
	if true {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionAwsIam) UnmarshalJSON(bytes []byte) (err error) {
	varArcProviderConnectionAwsIam := _ArcProviderConnectionAwsIam{}

	if err = json.Unmarshal(bytes, &varArcProviderConnectionAwsIam); err == nil {
		*o = ArcProviderConnectionAwsIam(varArcProviderConnectionAwsIam)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "models")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "customer_role_arn")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProviderConnectionAwsIam is a helper abstraction for handling nullable arcproviderconnectionawsiam types.
type NullableArcProviderConnectionAwsIam struct {
	value *ArcProviderConnectionAwsIam
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionAwsIam) Get() *ArcProviderConnectionAwsIam {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionAwsIam) Set(val *ArcProviderConnectionAwsIam) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionAwsIam) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionAwsIam) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionAwsIam returns a pointer to a new instance of NullableArcProviderConnectionAwsIam.
func NewNullableArcProviderConnectionAwsIam(val *ArcProviderConnectionAwsIam) *NullableArcProviderConnectionAwsIam {
	return &NullableArcProviderConnectionAwsIam{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionAwsIam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionAwsIam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
