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

// LoggingKinesisResponse struct for LoggingKinesisResponse
type LoggingKinesisResponse struct {
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
	IamRole NullableString `json:"iam_role,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	FormatVersion *string `json:"format_version,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	ServiceID            *string      `json:"service_id,omitempty"`
	Version              *string      `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingKinesisResponse LoggingKinesisResponse

// NewLoggingKinesisResponse instantiates a new LoggingKinesisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingKinesisResponse() *LoggingKinesisResponse {
	this := LoggingKinesisResponse{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// NewLoggingKinesisResponseWithDefaults instantiates a new LoggingKinesisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingKinesisResponseWithDefaults() *LoggingKinesisResponse {
	this := LoggingKinesisResponse{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingKinesisResponse) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetPlacement() LoggingPlacement {
	if o == nil || o.Placement.Get() == nil {
		var ret LoggingPlacement
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetPlacementOk() (*LoggingPlacement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableLoggingPlacement and assigns it to the Placement field.
func (o *LoggingKinesisResponse) SetPlacement(v LoggingPlacement) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingKinesisResponse) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingKinesisResponse) SetFormat(v string) {
	o.Format = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingKinesisResponse) SetTopic(v string) {
	o.Topic = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetRegion() AwsRegion {
	if o == nil || o.Region == nil {
		var ret AwsRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetRegionOk() (*AwsRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegion and assigns it to the Region field.
func (o *LoggingKinesisResponse) SetRegion(v AwsRegion) {
	o.Region = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingKinesisResponse) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}

// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingKinesisResponse) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *LoggingKinesisResponse) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}

// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *LoggingKinesisResponse) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetIamRole returns the IamRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetIamRole() string {
	if o == nil || o.IamRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.IamRole.Get()
}

// GetIamRoleOk returns a tuple with the IamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetIamRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IamRole.Get(), o.IamRole.IsSet()
}

// HasIamRole returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasIamRole() bool {
	if o != nil && o.IamRole.IsSet() {
		return true
	}

	return false
}

// SetIamRole gets a reference to the given NullableString and assigns it to the IamRole field.
func (o *LoggingKinesisResponse) SetIamRole(v string) {
	o.IamRole.Set(&v)
}

// SetIamRoleNil sets the value for IamRole to be an explicit nil
func (o *LoggingKinesisResponse) SetIamRoleNil() {
	o.IamRole.Set(nil)
}

// UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetIamRole() {
	o.IamRole.Unset()
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingKinesisResponse) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingKinesisResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingKinesisResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingKinesisResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingKinesisResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKinesisResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKinesisResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingKinesisResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingKinesisResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingKinesisResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *LoggingKinesisResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingKinesisResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKinesisResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingKinesisResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingKinesisResponse) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingKinesisResponse) MarshalJSON() ([]byte, error) {
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
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingKinesisResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingKinesisResponse := _LoggingKinesisResponse{}

	if err = json.Unmarshal(bytes, &varLoggingKinesisResponse); err == nil {
		*o = LoggingKinesisResponse(varLoggingKinesisResponse)
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
		delete(additionalProperties, "format_version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingKinesisResponse is a helper abstraction for handling nullable loggingkinesisresponse types.
type NullableLoggingKinesisResponse struct {
	value *LoggingKinesisResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingKinesisResponse) Get() *LoggingKinesisResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingKinesisResponse) Set(val *LoggingKinesisResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingKinesisResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingKinesisResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingKinesisResponse returns a pointer to a new instance of NullableLoggingKinesisResponse.
func NewNullableLoggingKinesisResponse(val *LoggingKinesisResponse) *NullableLoggingKinesisResponse {
	return &NullableLoggingKinesisResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingKinesisResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingKinesisResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
