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

// LoggingSplunk struct for LoggingSplunk
type LoggingSplunk struct {
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`. 
	Placement NullableString `json:"placement,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`. 
	FormatVersion *int32 `json:"format_version,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname NullableString `json:"tls_hostname,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token *string `json:"token,omitempty"`
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSplunk LoggingSplunk

// NewLoggingSplunk instantiates a new LoggingSplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSplunk() *LoggingSplunk {
	this := LoggingSplunk{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TLSHostname = *NewNullableString(&tlsHostname)
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingSplunkWithDefaults instantiates a new LoggingSplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSplunkWithDefaults() *LoggingSplunk {
	this := LoggingSplunk{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TLSHostname = *NewNullableString(&tlsHostname)
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingSplunk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingSplunk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingSplunk) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetPlacementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingSplunk) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingSplunk) SetPlacement(v string) {
	o.Placement.Set(&v)
}
// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingSplunk) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingSplunk) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingSplunk) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingSplunk) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingSplunk) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingSplunk) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingSplunk) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingSplunk) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingSplunk) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingSplunk) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingSplunk) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingSplunk) SetFormat(v string) {
	o.Format = &v
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *LoggingSplunk) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *LoggingSplunk) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *LoggingSplunk) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *LoggingSplunk) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *LoggingSplunk) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *LoggingSplunk) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *LoggingSplunk) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *LoggingSplunk) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *LoggingSplunk) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *LoggingSplunk) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *LoggingSplunk) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *LoggingSplunk) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSHostname returns the TLSHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunk) GetTLSHostname() string {
	if o == nil || o.TLSHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSHostname.Get()
}

// GetTLSHostnameOk returns a tuple with the TLSHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunk) GetTLSHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSHostname.Get(), o.TLSHostname.IsSet()
}

// HasTLSHostname returns a boolean if a field has been set.
func (o *LoggingSplunk) HasTLSHostname() bool {
	if o != nil && o.TLSHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSHostname gets a reference to the given NullableString and assigns it to the TLSHostname field.
func (o *LoggingSplunk) SetTLSHostname(v string) {
	o.TLSHostname.Set(&v)
}
// SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil
func (o *LoggingSplunk) SetTLSHostnameNil() {
	o.TLSHostname.Set(nil)
}

// UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
func (o *LoggingSplunk) UnsetTLSHostname() {
	o.TLSHostname.Unset()
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingSplunk) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingSplunk) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingSplunk) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingSplunk) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingSplunk) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingSplunk) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingSplunk) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingSplunk) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingSplunk) SetURL(v string) {
	o.URL = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingSplunk) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSplunk) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingSplunk) SetToken(v string) {
	o.Token = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingSplunk) GetUseTLS() LoggingUseTLS {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLS
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunk) GetUseTLSOk() (*LoggingUseTLS, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingSplunk) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLS and assigns it to the UseTLS field.
func (o *LoggingSplunk) SetUseTLS(v LoggingUseTLS) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Placement.IsSet() {
		toSerialize["placement"] = o.Placement.Get()
	}
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
	}
	if o.ResponseCondition.IsSet() {
		toSerialize["response_condition"] = o.ResponseCondition.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.TLSCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TLSCaCert.Get()
	}
	if o.TLSClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TLSClientCert.Get()
	}
	if o.TLSClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TLSClientKey.Get()
	}
	if o.TLSHostname.IsSet() {
		toSerialize["tls_hostname"] = o.TLSHostname.Get()
	}
	if o.RequestMaxEntries != nil {
		toSerialize["request_max_entries"] = o.RequestMaxEntries
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingSplunk) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSplunk := _LoggingSplunk{}

	if err = json.Unmarshal(bytes, &varLoggingSplunk); err == nil {
		*o = LoggingSplunk(varLoggingSplunk)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "placement")
		delete(additionalProperties, "format_version")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "format")
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_hostname")
		delete(additionalProperties, "request_max_entries")
		delete(additionalProperties, "request_max_bytes")
		delete(additionalProperties, "url")
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSplunk is a helper abstraction for handling nullable loggingsplunk types. 
type NullableLoggingSplunk struct {
	value *LoggingSplunk
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSplunk) Get() *LoggingSplunk {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSplunk) Set(val *LoggingSplunk) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSplunk) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSplunk) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSplunk returns a pointer to a new instance of NullableLoggingSplunk.
func NewNullableLoggingSplunk(val *LoggingSplunk) *NullableLoggingSplunk {
	return &NullableLoggingSplunk{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
