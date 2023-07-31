// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// LoggingGooglePubsubAdditional struct for LoggingGooglePubsubAdditional
type LoggingGooglePubsubAdditional struct {
	// The Google Cloud Pub/Sub topic to which logs will be published. Required.
	Topic *string `json:"topic,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGooglePubsubAdditional LoggingGooglePubsubAdditional

// NewLoggingGooglePubsubAdditional instantiates a new LoggingGooglePubsubAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGooglePubsubAdditional() *LoggingGooglePubsubAdditional {
	this := LoggingGooglePubsubAdditional{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	return &this
}

// NewLoggingGooglePubsubAdditionalWithDefaults instantiates a new LoggingGooglePubsubAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGooglePubsubAdditionalWithDefaults() *LoggingGooglePubsubAdditional {
	this := LoggingGooglePubsubAdditional{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAdditional) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAdditional) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAdditional) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingGooglePubsubAdditional) SetTopic(v string) {
	o.Topic = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingGooglePubsubAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetProjectID returns the ProjectID field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAdditional) GetProjectID() string {
	if o == nil || o.ProjectID == nil {
		var ret string
		return ret
	}
	return *o.ProjectID
}

// GetProjectIDOk returns a tuple with the ProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAdditional) GetProjectIDOk() (*string, bool) {
	if o == nil || o.ProjectID == nil {
		return nil, false
	}
	return o.ProjectID, true
}

// HasProjectID returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAdditional) HasProjectID() bool {
	if o != nil && o.ProjectID != nil {
		return true
	}

	return false
}

// SetProjectID gets a reference to the given string and assigns it to the ProjectID field.
func (o *LoggingGooglePubsubAdditional) SetProjectID(v string) {
	o.ProjectID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGooglePubsubAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.ProjectID != nil {
		toSerialize["project_id"] = o.ProjectID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingGooglePubsubAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGooglePubsubAdditional := _LoggingGooglePubsubAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingGooglePubsubAdditional); err == nil {
		*o = LoggingGooglePubsubAdditional(varLoggingGooglePubsubAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "topic")
		delete(additionalProperties, "format")
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGooglePubsubAdditional is a helper abstraction for handling nullable logginggooglepubsubadditional types. 
type NullableLoggingGooglePubsubAdditional struct {
	value *LoggingGooglePubsubAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGooglePubsubAdditional) Get() *LoggingGooglePubsubAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGooglePubsubAdditional) Set(val *LoggingGooglePubsubAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGooglePubsubAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGooglePubsubAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGooglePubsubAdditional returns a pointer to a new instance of NullableLoggingGooglePubsubAdditional.
func NewNullableLoggingGooglePubsubAdditional(val *LoggingGooglePubsubAdditional) *NullableLoggingGooglePubsubAdditional {
	return &NullableLoggingGooglePubsubAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGooglePubsubAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingGooglePubsubAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
