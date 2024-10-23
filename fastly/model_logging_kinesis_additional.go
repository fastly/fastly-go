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

// LoggingKinesisAdditional struct for LoggingKinesisAdditional
type LoggingKinesisAdditional struct {
	// The name for the real-time logging configuration.
	Name      *string                  `json:"name,omitempty"`
	Placement NullableLoggingPlacement `json:"placement,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The Amazon Kinesis stream to send logs to. Required.
	Topic  *string    `json:"topic,omitempty"`
	Region *AwsRegion `json:"region,omitempty"`
	// The secret key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified.
	SecretKey NullableString `json:"secret_key,omitempty"`
	// The access key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified.
	AccessKey NullableString `json:"access_key,omitempty"`
	// The ARN for an IAM role granting Fastly access to the target Amazon Kinesis stream. Not required if `access_key` and `secret_key` are provided.
	IamRole              NullableString `json:"iam_role,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingKinesisAdditional LoggingKinesisAdditional

// NewLoggingKinesisAdditional instantiates a new LoggingKinesisAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingKinesisAdditional() *LoggingKinesisAdditional {
	this := LoggingKinesisAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	return &this
}

// NewLoggingKinesisAdditionalWithDefaults instantiates a new LoggingKinesisAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingKinesisAdditionalWithDefaults() *LoggingKinesisAdditional {
	this := LoggingKinesisAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingKinesisAdditional) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisAdditional) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingKinesisAdditional) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisAdditional) GetPlacement() LoggingPlacement {
	if o == nil || o.Placement.Get() == nil {
		var ret LoggingPlacement
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisAdditional) GetPlacementOk() (*LoggingPlacement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableLoggingPlacement and assigns it to the Placement field.
func (o *LoggingKinesisAdditional) SetPlacement(v LoggingPlacement) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingKinesisAdditional) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingKinesisAdditional) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingKinesisAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingKinesisAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingKinesisAdditional) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisAdditional) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingKinesisAdditional) SetTopic(v string) {
	o.Topic = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingKinesisAdditional) GetRegion() AwsRegion {
	if o == nil || o.Region == nil {
		var ret AwsRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisAdditional) GetRegionOk() (*AwsRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegion and assigns it to the Region field.
func (o *LoggingKinesisAdditional) SetRegion(v AwsRegion) {
	o.Region = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisAdditional) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisAdditional) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingKinesisAdditional) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}

// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingKinesisAdditional) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingKinesisAdditional) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisAdditional) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisAdditional) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *LoggingKinesisAdditional) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}

// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *LoggingKinesisAdditional) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *LoggingKinesisAdditional) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetIamRole returns the IamRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisAdditional) GetIamRole() string {
	if o == nil || o.IamRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.IamRole.Get()
}

// GetIamRoleOk returns a tuple with the IamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisAdditional) GetIamRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IamRole.Get(), o.IamRole.IsSet()
}

// HasIamRole returns a boolean if a field has been set.
func (o *LoggingKinesisAdditional) HasIamRole() bool {
	if o != nil && o.IamRole.IsSet() {
		return true
	}

	return false
}

// SetIamRole gets a reference to the given NullableString and assigns it to the IamRole field.
func (o *LoggingKinesisAdditional) SetIamRole(v string) {
	o.IamRole.Set(&v)
}

// SetIamRoleNil sets the value for IamRole to be an explicit nil
func (o *LoggingKinesisAdditional) SetIamRoleNil() {
	o.IamRole.Set(nil)
}

// UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
func (o *LoggingKinesisAdditional) UnsetIamRole() {
	o.IamRole.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingKinesisAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Placement.IsSet() {
		toSerialize["placement"] = o.Placement.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.SecretKey.IsSet() {
		toSerialize["secret_key"] = o.SecretKey.Get()
	}
	if o.AccessKey.IsSet() {
		toSerialize["access_key"] = o.AccessKey.Get()
	}
	if o.IamRole.IsSet() {
		toSerialize["iam_role"] = o.IamRole.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingKinesisAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingKinesisAdditional := _LoggingKinesisAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingKinesisAdditional); err == nil {
		*o = LoggingKinesisAdditional(varLoggingKinesisAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "placement")
		delete(additionalProperties, "format")
		delete(additionalProperties, "topic")
		delete(additionalProperties, "region")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "iam_role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingKinesisAdditional is a helper abstraction for handling nullable loggingkinesisadditional types.
type NullableLoggingKinesisAdditional struct {
	value *LoggingKinesisAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingKinesisAdditional) Get() *LoggingKinesisAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingKinesisAdditional) Set(val *LoggingKinesisAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingKinesisAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingKinesisAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingKinesisAdditional returns a pointer to a new instance of NullableLoggingKinesisAdditional.
func NewNullableLoggingKinesisAdditional(val *LoggingKinesisAdditional) *NullableLoggingKinesisAdditional {
	return &NullableLoggingKinesisAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingKinesisAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingKinesisAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
