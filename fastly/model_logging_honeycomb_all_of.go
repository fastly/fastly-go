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

// LoggingHoneycombAllOf struct for LoggingHoneycombAllOf
type LoggingHoneycombAllOf struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest.
	Format *string `json:"format,omitempty"`
	// The Honeycomb Dataset you want to log to.
	Dataset *string `json:"dataset,omitempty"`
	// The Write Key from the Account page of your Honeycomb account.
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHoneycombAllOf LoggingHoneycombAllOf

// NewLoggingHoneycombAllOf instantiates a new LoggingHoneycombAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHoneycombAllOf() *LoggingHoneycombAllOf {
	this := LoggingHoneycombAllOf{}
	return &this
}

// NewLoggingHoneycombAllOfWithDefaults instantiates a new LoggingHoneycombAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHoneycombAllOfWithDefaults() *LoggingHoneycombAllOf {
	this := LoggingHoneycombAllOf{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingHoneycombAllOf) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHoneycombAllOf) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingHoneycombAllOf) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingHoneycombAllOf) SetFormat(v string) {
	o.Format = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *LoggingHoneycombAllOf) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHoneycombAllOf) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *LoggingHoneycombAllOf) HasDataset() bool {
	if o != nil && o.Dataset != nil {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *LoggingHoneycombAllOf) SetDataset(v string) {
	o.Dataset = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingHoneycombAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHoneycombAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingHoneycombAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingHoneycombAllOf) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHoneycombAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingHoneycombAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHoneycombAllOf := _LoggingHoneycombAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingHoneycombAllOf); err == nil {
		*o = LoggingHoneycombAllOf(varLoggingHoneycombAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "dataset")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHoneycombAllOf is a helper abstraction for handling nullable logginghoneycomballof types. 
type NullableLoggingHoneycombAllOf struct {
	value *LoggingHoneycombAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHoneycombAllOf) Get() *LoggingHoneycombAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHoneycombAllOf) Set(val *LoggingHoneycombAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHoneycombAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHoneycombAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHoneycombAllOf returns a pointer to a new instance of NullableLoggingHoneycombAllOf.
func NewNullableLoggingHoneycombAllOf(val *LoggingHoneycombAllOf) *NullableLoggingHoneycombAllOf {
	return &NullableLoggingHoneycombAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHoneycombAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingHoneycombAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
