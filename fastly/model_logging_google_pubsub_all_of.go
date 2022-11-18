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

// LoggingGooglePubsubAllOf struct for LoggingGooglePubsubAllOf
type LoggingGooglePubsubAllOf struct {
	// The Google Cloud Pub/Sub topic to which logs will be published. Required.
	Topic *string `json:"topic,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGooglePubsubAllOf LoggingGooglePubsubAllOf

// NewLoggingGooglePubsubAllOf instantiates a new LoggingGooglePubsubAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGooglePubsubAllOf() *LoggingGooglePubsubAllOf {
	this := LoggingGooglePubsubAllOf{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	return &this
}

// NewLoggingGooglePubsubAllOfWithDefaults instantiates a new LoggingGooglePubsubAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGooglePubsubAllOfWithDefaults() *LoggingGooglePubsubAllOf {
	this := LoggingGooglePubsubAllOf{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAllOf) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAllOf) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAllOf) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingGooglePubsubAllOf) SetTopic(v string) {
	o.Topic = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAllOf) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAllOf) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAllOf) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingGooglePubsubAllOf) SetFormat(v string) {
	o.Format = &v
}

// GetProjectID returns the ProjectID field value if set, zero value otherwise.
func (o *LoggingGooglePubsubAllOf) GetProjectID() string {
	if o == nil || o.ProjectID == nil {
		var ret string
		return ret
	}
	return *o.ProjectID
}

// GetProjectIDOk returns a tuple with the ProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGooglePubsubAllOf) GetProjectIDOk() (*string, bool) {
	if o == nil || o.ProjectID == nil {
		return nil, false
	}
	return o.ProjectID, true
}

// HasProjectID returns a boolean if a field has been set.
func (o *LoggingGooglePubsubAllOf) HasProjectID() bool {
	if o != nil && o.ProjectID != nil {
		return true
	}

	return false
}

// SetProjectID gets a reference to the given string and assigns it to the ProjectID field.
func (o *LoggingGooglePubsubAllOf) SetProjectID(v string) {
	o.ProjectID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGooglePubsubAllOf) MarshalJSON() ([]byte, error) {
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
func (o *LoggingGooglePubsubAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGooglePubsubAllOf := _LoggingGooglePubsubAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingGooglePubsubAllOf); err == nil {
		*o = LoggingGooglePubsubAllOf(varLoggingGooglePubsubAllOf)
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

// NullableLoggingGooglePubsubAllOf is a helper abstraction for handling nullable logginggooglepubsuballof types. 
type NullableLoggingGooglePubsubAllOf struct {
	value *LoggingGooglePubsubAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGooglePubsubAllOf) Get() *LoggingGooglePubsubAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGooglePubsubAllOf) Set(val *LoggingGooglePubsubAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGooglePubsubAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGooglePubsubAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGooglePubsubAllOf returns a pointer to a new instance of NullableLoggingGooglePubsubAllOf.
func NewNullableLoggingGooglePubsubAllOf(val *LoggingGooglePubsubAllOf) *NullableLoggingGooglePubsubAllOf {
	return &NullableLoggingGooglePubsubAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGooglePubsubAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingGooglePubsubAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
