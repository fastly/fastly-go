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

// LoggingSyslog struct for LoggingSyslog
type LoggingSyslog struct {
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
	// A hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// The port number.
	Port *int32 `json:"port,omitempty"`
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// The hostname used for the syslog endpoint.
	Hostname *string `json:"hostname,omitempty"`
	// The IPv4 address used for the syslog endpoint.
	Ipv4 NullableString `json:"ipv4,omitempty"`
	// Whether to prepend each message with a specific token.
	Token NullableString `json:"token,omitempty"`
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingSyslog LoggingSyslog

// NewLoggingSyslog instantiates a new LoggingSyslog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSyslog() *LoggingSyslog {
	this := LoggingSyslog{}
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
	var port int32 = 514
	this.Port = &port
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var token string = "null"
	this.Token = *NewNullableString(&token)
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingSyslogWithDefaults instantiates a new LoggingSyslog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSyslogWithDefaults() *LoggingSyslog {
	this := LoggingSyslog{}
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
	var port int32 = 514
	this.Port = &port
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var token string = "null"
	this.Token = *NewNullableString(&token)
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingSyslog) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingSyslog) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingSyslog) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetPlacementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingSyslog) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingSyslog) SetPlacement(v string) {
	o.Placement.Set(&v)
}
// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingSyslog) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingSyslog) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingSyslog) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingSyslog) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingSyslog) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingSyslog) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingSyslog) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingSyslog) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingSyslog) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingSyslog) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingSyslog) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingSyslog) SetFormat(v string) {
	o.Format = &v
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *LoggingSyslog) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *LoggingSyslog) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *LoggingSyslog) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *LoggingSyslog) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *LoggingSyslog) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *LoggingSyslog) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *LoggingSyslog) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *LoggingSyslog) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *LoggingSyslog) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *LoggingSyslog) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *LoggingSyslog) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *LoggingSyslog) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSHostname returns the TLSHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetTLSHostname() string {
	if o == nil || o.TLSHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSHostname.Get()
}

// GetTLSHostnameOk returns a tuple with the TLSHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetTLSHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSHostname.Get(), o.TLSHostname.IsSet()
}

// HasTLSHostname returns a boolean if a field has been set.
func (o *LoggingSyslog) HasTLSHostname() bool {
	if o != nil && o.TLSHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSHostname gets a reference to the given NullableString and assigns it to the TLSHostname field.
func (o *LoggingSyslog) SetTLSHostname(v string) {
	o.TLSHostname.Set(&v)
}
// SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil
func (o *LoggingSyslog) SetTLSHostnameNil() {
	o.TLSHostname.Set(nil)
}

// UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
func (o *LoggingSyslog) UnsetTLSHostname() {
	o.TLSHostname.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoggingSyslog) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoggingSyslog) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoggingSyslog) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingSyslog) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingSyslog) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoggingSyslog) SetPort(v int32) {
	o.Port = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingSyslog) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingSyslog) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingSyslog) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LoggingSyslog) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LoggingSyslog) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *LoggingSyslog) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetIpv4() string {
	if o == nil || o.Ipv4.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ipv4.Get()
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetIpv4Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ipv4.Get(), o.Ipv4.IsSet()
}

// HasIpv4 returns a boolean if a field has been set.
func (o *LoggingSyslog) HasIpv4() bool {
	if o != nil && o.Ipv4.IsSet() {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NullableString and assigns it to the Ipv4 field.
func (o *LoggingSyslog) SetIpv4(v string) {
	o.Ipv4.Set(&v)
}
// SetIpv4Nil sets the value for Ipv4 to be an explicit nil
func (o *LoggingSyslog) SetIpv4Nil() {
	o.Ipv4.Set(nil)
}

// UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
func (o *LoggingSyslog) UnsetIpv4() {
	o.Ipv4.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSyslog) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSyslog) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSyslog) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *LoggingSyslog) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *LoggingSyslog) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *LoggingSyslog) UnsetToken() {
	o.Token.Unset()
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingSyslog) GetUseTLS() LoggingUseTLS {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLS
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSyslog) GetUseTLSOk() (*LoggingUseTLS, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingSyslog) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLS and assigns it to the UseTLS field.
func (o *LoggingSyslog) SetUseTLS(v LoggingUseTLS) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSyslog) MarshalJSON() ([]byte, error) {
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
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Ipv4.IsSet() {
		toSerialize["ipv4"] = o.Ipv4.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
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
func (o *LoggingSyslog) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSyslog := _LoggingSyslog{}

	if err = json.Unmarshal(bytes, &varLoggingSyslog); err == nil {
		*o = LoggingSyslog(varLoggingSyslog)
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
		delete(additionalProperties, "address")
		delete(additionalProperties, "port")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSyslog is a helper abstraction for handling nullable loggingsyslog types. 
type NullableLoggingSyslog struct {
	value *LoggingSyslog
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSyslog) Get() *LoggingSyslog {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSyslog) Set(val *LoggingSyslog) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSyslog) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSyslog) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSyslog returns a pointer to a new instance of NullableLoggingSyslog.
func NewNullableLoggingSyslog(val *LoggingSyslog) *NullableLoggingSyslog {
	return &NullableLoggingSyslog{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSyslog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingSyslog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
