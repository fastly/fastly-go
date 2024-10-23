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

// LoggingNewrelicotlpAdditional struct for LoggingNewrelicotlpAdditional
type LoggingNewrelicotlpAdditional struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The Insert API key from the Account page of your New Relic account. Required.
	Token *string `json:"token,omitempty"`
	// The region to which to stream logs.
	Region *string `json:"region,omitempty"`
	// (Optional) URL of the New Relic Trace Observer, if you are using New Relic Infinite Tracing.
	URL                  NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingNewrelicotlpAdditional LoggingNewrelicotlpAdditional

// NewLoggingNewrelicotlpAdditional instantiates a new LoggingNewrelicotlpAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingNewrelicotlpAdditional() *LoggingNewrelicotlpAdditional {
	this := LoggingNewrelicotlpAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var region string = "US"
	this.Region = &region
	var url string = "null"
	this.URL = *NewNullableString(&url)
	return &this
}

// NewLoggingNewrelicotlpAdditionalWithDefaults instantiates a new LoggingNewrelicotlpAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingNewrelicotlpAdditionalWithDefaults() *LoggingNewrelicotlpAdditional {
	this := LoggingNewrelicotlpAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var region string = "US"
	this.Region = &region
	var url string = "null"
	this.URL = *NewNullableString(&url)
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingNewrelicotlpAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicotlpAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingNewrelicotlpAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingNewrelicotlpAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingNewrelicotlpAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicotlpAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingNewrelicotlpAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingNewrelicotlpAdditional) SetToken(v string) {
	o.Token = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingNewrelicotlpAdditional) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicotlpAdditional) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingNewrelicotlpAdditional) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LoggingNewrelicotlpAdditional) SetRegion(v string) {
	o.Region = &v
}

// GetURL returns the URL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingNewrelicotlpAdditional) GetURL() string {
	if o == nil || o.URL.Get() == nil {
		var ret string
		return ret
	}
	return *o.URL.Get()
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingNewrelicotlpAdditional) GetURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.URL.Get(), o.URL.IsSet()
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingNewrelicotlpAdditional) HasURL() bool {
	if o != nil && o.URL.IsSet() {
		return true
	}

	return false
}

// SetURL gets a reference to the given NullableString and assigns it to the URL field.
func (o *LoggingNewrelicotlpAdditional) SetURL(v string) {
	o.URL.Set(&v)
}

// SetURLNil sets the value for URL to be an explicit nil
func (o *LoggingNewrelicotlpAdditional) SetURLNil() {
	o.URL.Set(nil)
}

// UnsetURL ensures that no value is present for URL, not even an explicit nil
func (o *LoggingNewrelicotlpAdditional) UnsetURL() {
	o.URL.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingNewrelicotlpAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.URL.IsSet() {
		toSerialize["url"] = o.URL.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingNewrelicotlpAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingNewrelicotlpAdditional := _LoggingNewrelicotlpAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingNewrelicotlpAdditional); err == nil {
		*o = LoggingNewrelicotlpAdditional(varLoggingNewrelicotlpAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "token")
		delete(additionalProperties, "region")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingNewrelicotlpAdditional is a helper abstraction for handling nullable loggingnewrelicotlpadditional types.
type NullableLoggingNewrelicotlpAdditional struct {
	value *LoggingNewrelicotlpAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingNewrelicotlpAdditional) Get() *LoggingNewrelicotlpAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingNewrelicotlpAdditional) Set(val *LoggingNewrelicotlpAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingNewrelicotlpAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingNewrelicotlpAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingNewrelicotlpAdditional returns a pointer to a new instance of NullableLoggingNewrelicotlpAdditional.
func NewNullableLoggingNewrelicotlpAdditional(val *LoggingNewrelicotlpAdditional) *NullableLoggingNewrelicotlpAdditional {
	return &NullableLoggingNewrelicotlpAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingNewrelicotlpAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingNewrelicotlpAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
