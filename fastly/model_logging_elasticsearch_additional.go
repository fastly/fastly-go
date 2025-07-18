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

// LoggingElasticsearchAdditional struct for LoggingElasticsearchAdditional
type LoggingElasticsearchAdditional struct {
	// The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today's date.
	Index *string `json:"index,omitempty"`
	// The URL to stream logs to. Must use HTTPS.
	URL *string `json:"url,omitempty"`
	// The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html).
	Pipeline NullableString `json:"pipeline,omitempty"`
	// Basic Auth username.
	User NullableString `json:"user,omitempty"`
	// Basic Auth password.
	Password NullableString `json:"password,omitempty"`
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest.
	Format               *string `json:"format,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingElasticsearchAdditional LoggingElasticsearchAdditional

// NewLoggingElasticsearchAdditional instantiates a new LoggingElasticsearchAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingElasticsearchAdditional() *LoggingElasticsearchAdditional {
	this := LoggingElasticsearchAdditional{}
	return &this
}

// NewLoggingElasticsearchAdditionalWithDefaults instantiates a new LoggingElasticsearchAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingElasticsearchAdditionalWithDefaults() *LoggingElasticsearchAdditional {
	this := LoggingElasticsearchAdditional{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *LoggingElasticsearchAdditional) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchAdditional) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *LoggingElasticsearchAdditional) SetIndex(v string) {
	o.Index = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingElasticsearchAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingElasticsearchAdditional) SetURL(v string) {
	o.URL = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchAdditional) GetPipeline() string {
	if o == nil || o.Pipeline.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pipeline.Get()
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchAdditional) GetPipelineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pipeline.Get(), o.Pipeline.IsSet()
}

// HasPipeline returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasPipeline() bool {
	if o != nil && o.Pipeline.IsSet() {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given NullableString and assigns it to the Pipeline field.
func (o *LoggingElasticsearchAdditional) SetPipeline(v string) {
	o.Pipeline.Set(&v)
}

// SetPipelineNil sets the value for Pipeline to be an explicit nil
func (o *LoggingElasticsearchAdditional) SetPipelineNil() {
	o.Pipeline.Set(nil)
}

// UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
func (o *LoggingElasticsearchAdditional) UnsetPipeline() {
	o.Pipeline.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchAdditional) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchAdditional) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *LoggingElasticsearchAdditional) SetUser(v string) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *LoggingElasticsearchAdditional) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *LoggingElasticsearchAdditional) UnsetUser() {
	o.User.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchAdditional) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchAdditional) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *LoggingElasticsearchAdditional) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *LoggingElasticsearchAdditional) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *LoggingElasticsearchAdditional) UnsetPassword() {
	o.Password.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingElasticsearchAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingElasticsearchAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingElasticsearchAdditional) SetFormat(v string) {
	o.Format = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingElasticsearchAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}
	if o.Pipeline.IsSet() {
		toSerialize["pipeline"] = o.Pipeline.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingElasticsearchAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingElasticsearchAdditional := _LoggingElasticsearchAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingElasticsearchAdditional); err == nil {
		*o = LoggingElasticsearchAdditional(varLoggingElasticsearchAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "index")
		delete(additionalProperties, "url")
		delete(additionalProperties, "pipeline")
		delete(additionalProperties, "user")
		delete(additionalProperties, "password")
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingElasticsearchAdditional is a helper abstraction for handling nullable loggingelasticsearchadditional types.
type NullableLoggingElasticsearchAdditional struct {
	value *LoggingElasticsearchAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingElasticsearchAdditional) Get() *LoggingElasticsearchAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingElasticsearchAdditional) Set(val *LoggingElasticsearchAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingElasticsearchAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingElasticsearchAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingElasticsearchAdditional returns a pointer to a new instance of NullableLoggingElasticsearchAdditional.
func NewNullableLoggingElasticsearchAdditional(val *LoggingElasticsearchAdditional) *NullableLoggingElasticsearchAdditional {
	return &NullableLoggingElasticsearchAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingElasticsearchAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingElasticsearchAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
