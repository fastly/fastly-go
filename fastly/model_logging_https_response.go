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

// LoggingHTTPSResponse struct for LoggingHTTPSResponse
type LoggingHTTPSResponse struct {
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	Placement NullableString `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	FormatVersion *string `json:"format_version,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname NullableString `json:"tls_hostname,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` (10k).
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (100MB).
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// The URL to send logs to. Must use HTTPS. Required.
	URL *string `json:"url,omitempty"`
	// Content type of the header sent with the request.
	ContentType NullableString `json:"content_type,omitempty"`
	// Name of the custom header sent with the request.
	HeaderName  NullableString      `json:"header_name,omitempty"`
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// Value of the custom header sent with the request.
	HeaderValue NullableString `json:"header_value,omitempty"`
	// HTTP method used for request.
	Method *string `json:"method,omitempty"`
	// Enforces valid JSON formatting for log entries.
	JSONFormat *string `json:"json_format,omitempty"`
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

type _LoggingHTTPSResponse LoggingHTTPSResponse

// NewLoggingHTTPSResponse instantiates a new LoggingHTTPSResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHTTPSResponse() *LoggingHTTPSResponse {
	this := LoggingHTTPSResponse{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
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
	var contentType string = "null"
	this.ContentType = *NewNullableString(&contentType)
	var headerName string = "null"
	this.HeaderName = *NewNullableString(&headerName)
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var headerValue string = "null"
	this.HeaderValue = *NewNullableString(&headerValue)
	var method string = "POST"
	this.Method = &method
	return &this
}

// NewLoggingHTTPSResponseWithDefaults instantiates a new LoggingHTTPSResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHTTPSResponseWithDefaults() *LoggingHTTPSResponse {
	this := LoggingHTTPSResponse{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
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
	var contentType string = "null"
	this.ContentType = *NewNullableString(&contentType)
	var headerName string = "null"
	this.HeaderName = *NewNullableString(&headerName)
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var headerValue string = "null"
	this.HeaderValue = *NewNullableString(&headerValue)
	var method string = "POST"
	this.Method = &method
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingHTTPSResponse) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetPlacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingHTTPSResponse) SetPlacement(v string) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingHTTPSResponse) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetPlacement() {
	o.Placement.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetResponseConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingHTTPSResponse) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}

// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingHTTPSResponse) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingHTTPSResponse) SetFormat(v string) {
	o.Format = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingHTTPSResponse) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetTLSCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *LoggingHTTPSResponse) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}

// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *LoggingHTTPSResponse) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetTLSClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *LoggingHTTPSResponse) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}

// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *LoggingHTTPSResponse) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetTLSClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *LoggingHTTPSResponse) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}

// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *LoggingHTTPSResponse) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSHostname returns the TLSHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetTLSHostname() string {
	if o == nil || o.TLSHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSHostname.Get()
}

// GetTLSHostnameOk returns a tuple with the TLSHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetTLSHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSHostname.Get(), o.TLSHostname.IsSet()
}

// HasTLSHostname returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasTLSHostname() bool {
	if o != nil && o.TLSHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSHostname gets a reference to the given NullableString and assigns it to the TLSHostname field.
func (o *LoggingHTTPSResponse) SetTLSHostname(v string) {
	o.TLSHostname.Set(&v)
}

// SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil
func (o *LoggingHTTPSResponse) SetTLSHostnameNil() {
	o.TLSHostname.Set(nil)
}

// UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetTLSHostname() {
	o.TLSHostname.Unset()
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingHTTPSResponse) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingHTTPSResponse) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingHTTPSResponse) SetURL(v string) {
	o.URL = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *LoggingHTTPSResponse) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *LoggingHTTPSResponse) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetContentType() {
	o.ContentType.Unset()
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetHeaderName() string {
	if o == nil || o.HeaderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderName.Get()
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName.Get(), o.HeaderName.IsSet()
}

// HasHeaderName returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasHeaderName() bool {
	if o != nil && o.HeaderName.IsSet() {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given NullableString and assigns it to the HeaderName field.
func (o *LoggingHTTPSResponse) SetHeaderName(v string) {
	o.HeaderName.Set(&v)
}

// SetHeaderNameNil sets the value for HeaderName to be an explicit nil
func (o *LoggingHTTPSResponse) SetHeaderNameNil() {
	o.HeaderName.Set(nil)
}

// UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetHeaderName() {
	o.HeaderName.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingHTTPSResponse) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetHeaderValue() string {
	if o == nil || o.HeaderValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderValue.Get()
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderValue.Get(), o.HeaderValue.IsSet()
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasHeaderValue() bool {
	if o != nil && o.HeaderValue.IsSet() {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given NullableString and assigns it to the HeaderValue field.
func (o *LoggingHTTPSResponse) SetHeaderValue(v string) {
	o.HeaderValue.Set(&v)
}

// SetHeaderValueNil sets the value for HeaderValue to be an explicit nil
func (o *LoggingHTTPSResponse) SetHeaderValueNil() {
	o.HeaderValue.Set(nil)
}

// UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetHeaderValue() {
	o.HeaderValue.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoggingHTTPSResponse) SetMethod(v string) {
	o.Method = &v
}

// GetJSONFormat returns the JSONFormat field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetJSONFormat() string {
	if o == nil || o.JSONFormat == nil {
		var ret string
		return ret
	}
	return *o.JSONFormat
}

// GetJSONFormatOk returns a tuple with the JSONFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetJSONFormatOk() (*string, bool) {
	if o == nil || o.JSONFormat == nil {
		return nil, false
	}
	return o.JSONFormat, true
}

// HasJSONFormat returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasJSONFormat() bool {
	if o != nil && o.JSONFormat != nil {
		return true
	}

	return false
}

// SetJSONFormat gets a reference to the given string and assigns it to the JSONFormat field.
func (o *LoggingHTTPSResponse) SetJSONFormat(v string) {
	o.JSONFormat = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingHTTPSResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingHTTPSResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingHTTPSResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingHTTPSResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingHTTPSResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingHTTPSResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingHTTPSResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *LoggingHTTPSResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingHTTPSResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingHTTPSResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingHTTPSResponse) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHTTPSResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Placement.IsSet() {
		toSerialize["placement"] = o.Placement.Get()
	}
	if o.ResponseCondition.IsSet() {
		toSerialize["response_condition"] = o.ResponseCondition.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
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
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.HeaderName.IsSet() {
		toSerialize["header_name"] = o.HeaderName.Get()
	}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.HeaderValue.IsSet() {
		toSerialize["header_value"] = o.HeaderValue.Get()
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.JSONFormat != nil {
		toSerialize["json_format"] = o.JSONFormat
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
func (o *LoggingHTTPSResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHTTPSResponse := _LoggingHTTPSResponse{}

	if err = json.Unmarshal(bytes, &varLoggingHTTPSResponse); err == nil {
		*o = LoggingHTTPSResponse(varLoggingHTTPSResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "placement")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "format")
		delete(additionalProperties, "format_version")
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_hostname")
		delete(additionalProperties, "request_max_entries")
		delete(additionalProperties, "request_max_bytes")
		delete(additionalProperties, "url")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "header_name")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "header_value")
		delete(additionalProperties, "method")
		delete(additionalProperties, "json_format")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHTTPSResponse is a helper abstraction for handling nullable logginghttpsresponse types.
type NullableLoggingHTTPSResponse struct {
	value *LoggingHTTPSResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHTTPSResponse) Get() *LoggingHTTPSResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHTTPSResponse) Set(val *LoggingHTTPSResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHTTPSResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHTTPSResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHTTPSResponse returns a pointer to a new instance of NullableLoggingHTTPSResponse.
func NewNullableLoggingHTTPSResponse(val *LoggingHTTPSResponse) *NullableLoggingHTTPSResponse {
	return &NullableLoggingHTTPSResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHTTPSResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingHTTPSResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
