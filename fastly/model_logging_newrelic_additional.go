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

// LoggingNewrelicAdditional struct for LoggingNewrelicAdditional
type LoggingNewrelicAdditional struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that New Relic Logs can ingest.
	Format *string `json:"format,omitempty"`
	// The Insert API key from the Account page of your New Relic account. Required.
	Token *string `json:"token,omitempty"`
	// The region to which to stream logs.
	Region               *string `json:"region,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingNewrelicAdditional LoggingNewrelicAdditional

// NewLoggingNewrelicAdditional instantiates a new LoggingNewrelicAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingNewrelicAdditional() *LoggingNewrelicAdditional {
	this := LoggingNewrelicAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var region string = "US"
	this.Region = &region
	return &this
}

// NewLoggingNewrelicAdditionalWithDefaults instantiates a new LoggingNewrelicAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingNewrelicAdditionalWithDefaults() *LoggingNewrelicAdditional {
	this := LoggingNewrelicAdditional{}
	var format string = "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"
	this.Format = &format
	var region string = "US"
	this.Region = &region
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingNewrelicAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingNewrelicAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingNewrelicAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingNewrelicAdditional) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicAdditional) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingNewrelicAdditional) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingNewrelicAdditional) SetToken(v string) {
	o.Token = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingNewrelicAdditional) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingNewrelicAdditional) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingNewrelicAdditional) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LoggingNewrelicAdditional) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingNewrelicAdditional) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingNewrelicAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingNewrelicAdditional := _LoggingNewrelicAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingNewrelicAdditional); err == nil {
		*o = LoggingNewrelicAdditional(varLoggingNewrelicAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "token")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingNewrelicAdditional is a helper abstraction for handling nullable loggingnewrelicadditional types.
type NullableLoggingNewrelicAdditional struct {
	value *LoggingNewrelicAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingNewrelicAdditional) Get() *LoggingNewrelicAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingNewrelicAdditional) Set(val *LoggingNewrelicAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingNewrelicAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingNewrelicAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingNewrelicAdditional returns a pointer to a new instance of NullableLoggingNewrelicAdditional.
func NewNullableLoggingNewrelicAdditional(val *LoggingNewrelicAdditional) *NullableLoggingNewrelicAdditional {
	return &NullableLoggingNewrelicAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingNewrelicAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingNewrelicAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
