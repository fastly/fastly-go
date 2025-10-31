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

// LoggingGrafanacloudlogsAdditional struct for LoggingGrafanacloudlogsAdditional
type LoggingGrafanacloudlogsAdditional struct {
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
	Format *string `json:"format,omitempty"`
	// The Grafana Cloud Logs Dataset you want to log to.
	User *string `json:"user,omitempty"`
	// The URL of the Loki instance in your Grafana stack.
	Url *string `json:"url,omitempty"`
	// The Grafana Access Policy token with `logs:write` access scoped to your Loki instance.
	Token *string `json:"token,omitempty"`
	// The Stream Labels, a JSON string used to identify the stream.
	Index                *string `json:"index,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGrafanacloudlogsAdditional LoggingGrafanacloudlogsAdditional

// NewLoggingGrafanacloudlogsAdditional instantiates a new LoggingGrafanacloudlogsAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGrafanacloudlogsAdditional() *LoggingGrafanacloudlogsAdditional {
	this := LoggingGrafanacloudlogsAdditional{}
	return &this
}

// NewLoggingGrafanacloudlogsAdditionalWithDefaults instantiates a new LoggingGrafanacloudlogsAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGrafanacloudlogsAdditionalWithDefaults() *LoggingGrafanacloudlogsAdditional {
	this := LoggingGrafanacloudlogsAdditional{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingGrafanacloudlogsAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGrafanacloudlogsAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingGrafanacloudlogsAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingGrafanacloudlogsAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingGrafanacloudlogsAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGrafanacloudlogsAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingGrafanacloudlogsAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingGrafanacloudlogsAdditional) SetUser(v string) {
	o.User = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LoggingGrafanacloudlogsAdditional) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGrafanacloudlogsAdditional) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LoggingGrafanacloudlogsAdditional) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LoggingGrafanacloudlogsAdditional) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingGrafanacloudlogsAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGrafanacloudlogsAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingGrafanacloudlogsAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingGrafanacloudlogsAdditional) SetToken(v string) {
	o.Token = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *LoggingGrafanacloudlogsAdditional) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGrafanacloudlogsAdditional) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *LoggingGrafanacloudlogsAdditional) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *LoggingGrafanacloudlogsAdditional) SetIndex(v string) {
	o.Index = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGrafanacloudlogsAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingGrafanacloudlogsAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGrafanacloudlogsAdditional := _LoggingGrafanacloudlogsAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingGrafanacloudlogsAdditional); err == nil {
		*o = LoggingGrafanacloudlogsAdditional(varLoggingGrafanacloudlogsAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "user")
		delete(additionalProperties, "url")
		delete(additionalProperties, "token")
		delete(additionalProperties, "index")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGrafanacloudlogsAdditional is a helper abstraction for handling nullable logginggrafanacloudlogsadditional types.
type NullableLoggingGrafanacloudlogsAdditional struct {
	value *LoggingGrafanacloudlogsAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGrafanacloudlogsAdditional) Get() *LoggingGrafanacloudlogsAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGrafanacloudlogsAdditional) Set(val *LoggingGrafanacloudlogsAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGrafanacloudlogsAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGrafanacloudlogsAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGrafanacloudlogsAdditional returns a pointer to a new instance of NullableLoggingGrafanacloudlogsAdditional.
func NewNullableLoggingGrafanacloudlogsAdditional(val *LoggingGrafanacloudlogsAdditional) *NullableLoggingGrafanacloudlogsAdditional {
	return &NullableLoggingGrafanacloudlogsAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGrafanacloudlogsAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingGrafanacloudlogsAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
