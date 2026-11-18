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

// ArcBedrockSetupResponse struct for ArcBedrockSetupResponse
type ArcBedrockSetupResponse struct {
	// The ARN of the AI Runtime Control IAM role your AWS account's trust policy must trust.
	ArcRoleArn string `json:"arc_role_arn"`
	// The external ID to require in your trust policy. This is your Fastly customer ID.
	ExternalID           string `json:"external_id"`
	AdditionalProperties map[string]any
}

type _ArcBedrockSetupResponse ArcBedrockSetupResponse

// NewArcBedrockSetupResponse instantiates a new ArcBedrockSetupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcBedrockSetupResponse(arcRoleArn string, externalID string) *ArcBedrockSetupResponse {
	this := ArcBedrockSetupResponse{}
	this.ArcRoleArn = arcRoleArn
	this.ExternalID = externalID
	return &this
}

// NewArcBedrockSetupResponseWithDefaults instantiates a new ArcBedrockSetupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcBedrockSetupResponseWithDefaults() *ArcBedrockSetupResponse {
	this := ArcBedrockSetupResponse{}
	return &this
}

// GetArcRoleArn returns the ArcRoleArn field value
func (o *ArcBedrockSetupResponse) GetArcRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArcRoleArn
}

// GetArcRoleArnOk returns a tuple with the ArcRoleArn field value
// and a boolean to check if the value has been set.
func (o *ArcBedrockSetupResponse) GetArcRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArcRoleArn, true
}

// SetArcRoleArn sets field value
func (o *ArcBedrockSetupResponse) SetArcRoleArn(v string) {
	o.ArcRoleArn = v
}

// GetExternalID returns the ExternalID field value
func (o *ArcBedrockSetupResponse) GetExternalID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value
// and a boolean to check if the value has been set.
func (o *ArcBedrockSetupResponse) GetExternalIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalID, true
}

// SetExternalID sets field value
func (o *ArcBedrockSetupResponse) SetExternalID(v string) {
	o.ExternalID = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcBedrockSetupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["arc_role_arn"] = o.ArcRoleArn
	}
	if true {
		toSerialize["external_id"] = o.ExternalID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcBedrockSetupResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcBedrockSetupResponse := _ArcBedrockSetupResponse{}

	if err = json.Unmarshal(bytes, &varArcBedrockSetupResponse); err == nil {
		*o = ArcBedrockSetupResponse(varArcBedrockSetupResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "arc_role_arn")
		delete(additionalProperties, "external_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcBedrockSetupResponse is a helper abstraction for handling nullable arcbedrocksetupresponse types.
type NullableArcBedrockSetupResponse struct {
	value *ArcBedrockSetupResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcBedrockSetupResponse) Get() *ArcBedrockSetupResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcBedrockSetupResponse) Set(val *ArcBedrockSetupResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcBedrockSetupResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcBedrockSetupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcBedrockSetupResponse returns a pointer to a new instance of NullableArcBedrockSetupResponse.
func NewNullableArcBedrockSetupResponse(val *ArcBedrockSetupResponse) *NullableArcBedrockSetupResponse {
	return &NullableArcBedrockSetupResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcBedrockSetupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcBedrockSetupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
