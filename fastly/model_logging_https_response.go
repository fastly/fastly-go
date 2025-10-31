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

// LoggingHttpsResponse struct for LoggingHttpsResponse
type LoggingHttpsResponse struct {
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	Placement NullableString `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
	Format *string `json:"format,omitempty"`
	// The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global.
	LogProcessingRegion *string `json:"log_processing_region,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	FormatVersion *string `json:"format_version,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TlsCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TlsClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TlsClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TlsHostname NullableString `json:"tls_hostname,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` (10k).
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (100MB).
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// The URL to send logs to. Must use HTTPS. Required.
	Url *string `json:"url,omitempty"`
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
	JsonFormat *string `json:"json_format,omitempty"`
	// How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds.
	Period *int32 `json:"period,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	ServiceId            *string      `json:"service_id,omitempty"`
	Version              *string      `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHttpsResponse LoggingHttpsResponse

// NewLoggingHttpsResponse instantiates a new LoggingHttpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHttpsResponse() *LoggingHttpsResponse {
	this := LoggingHttpsResponse{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var logProcessingRegion string = "none"
	this.LogProcessingRegion = &logProcessingRegion
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TlsHostname = *NewNullableString(&tlsHostname)
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
	var period int32 = 5
	this.Period = &period
	return &this
}

// NewLoggingHttpsResponseWithDefaults instantiates a new LoggingHttpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHttpsResponseWithDefaults() *LoggingHttpsResponse {
	this := LoggingHttpsResponse{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var logProcessingRegion string = "none"
	this.LogProcessingRegion = &logProcessingRegion
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TlsHostname = *NewNullableString(&tlsHostname)
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
	var period int32 = 5
	this.Period = &period
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingHttpsResponse) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetPlacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingHttpsResponse) SetPlacement(v string) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingHttpsResponse) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetPlacement() {
	o.Placement.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetResponseConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingHttpsResponse) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}

// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingHttpsResponse) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingHttpsResponse) SetFormat(v string) {
	o.Format = &v
}

// GetLogProcessingRegion returns the LogProcessingRegion field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetLogProcessingRegion() string {
	if o == nil || o.LogProcessingRegion == nil {
		var ret string
		return ret
	}
	return *o.LogProcessingRegion
}

// GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetLogProcessingRegionOk() (*string, bool) {
	if o == nil || o.LogProcessingRegion == nil {
		return nil, false
	}
	return o.LogProcessingRegion, true
}

// HasLogProcessingRegion returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasLogProcessingRegion() bool {
	if o != nil && o.LogProcessingRegion != nil {
		return true
	}

	return false
}

// SetLogProcessingRegion gets a reference to the given string and assigns it to the LogProcessingRegion field.
func (o *LoggingHttpsResponse) SetLogProcessingRegion(v string) {
	o.LogProcessingRegion = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingHttpsResponse) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *LoggingHttpsResponse) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *LoggingHttpsResponse) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *LoggingHttpsResponse) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *LoggingHttpsResponse) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *LoggingHttpsResponse) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *LoggingHttpsResponse) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsHostname returns the TlsHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetTlsHostname() string {
	if o == nil || o.TlsHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsHostname.Get()
}

// GetTlsHostnameOk returns a tuple with the TlsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetTlsHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsHostname.Get(), o.TlsHostname.IsSet()
}

// HasTlsHostname returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasTlsHostname() bool {
	if o != nil && o.TlsHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsHostname gets a reference to the given NullableString and assigns it to the TlsHostname field.
func (o *LoggingHttpsResponse) SetTlsHostname(v string) {
	o.TlsHostname.Set(&v)
}

// SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil
func (o *LoggingHttpsResponse) SetTlsHostnameNil() {
	o.TlsHostname.Set(nil)
}

// UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetTlsHostname() {
	o.TlsHostname.Unset()
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingHttpsResponse) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingHttpsResponse) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LoggingHttpsResponse) SetUrl(v string) {
	o.Url = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *LoggingHttpsResponse) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *LoggingHttpsResponse) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetContentType() {
	o.ContentType.Unset()
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetHeaderName() string {
	if o == nil || o.HeaderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderName.Get()
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName.Get(), o.HeaderName.IsSet()
}

// HasHeaderName returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasHeaderName() bool {
	if o != nil && o.HeaderName.IsSet() {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given NullableString and assigns it to the HeaderName field.
func (o *LoggingHttpsResponse) SetHeaderName(v string) {
	o.HeaderName.Set(&v)
}

// SetHeaderNameNil sets the value for HeaderName to be an explicit nil
func (o *LoggingHttpsResponse) SetHeaderNameNil() {
	o.HeaderName.Set(nil)
}

// UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetHeaderName() {
	o.HeaderName.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingHttpsResponse) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetHeaderValue() string {
	if o == nil || o.HeaderValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderValue.Get()
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderValue.Get(), o.HeaderValue.IsSet()
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasHeaderValue() bool {
	if o != nil && o.HeaderValue.IsSet() {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given NullableString and assigns it to the HeaderValue field.
func (o *LoggingHttpsResponse) SetHeaderValue(v string) {
	o.HeaderValue.Set(&v)
}

// SetHeaderValueNil sets the value for HeaderValue to be an explicit nil
func (o *LoggingHttpsResponse) SetHeaderValueNil() {
	o.HeaderValue.Set(nil)
}

// UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetHeaderValue() {
	o.HeaderValue.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoggingHttpsResponse) SetMethod(v string) {
	o.Method = &v
}

// GetJsonFormat returns the JsonFormat field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetJsonFormat() string {
	if o == nil || o.JsonFormat == nil {
		var ret string
		return ret
	}
	return *o.JsonFormat
}

// GetJsonFormatOk returns a tuple with the JsonFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetJsonFormatOk() (*string, bool) {
	if o == nil || o.JsonFormat == nil {
		return nil, false
	}
	return o.JsonFormat, true
}

// HasJsonFormat returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasJsonFormat() bool {
	if o != nil && o.JsonFormat != nil {
		return true
	}

	return false
}

// SetJsonFormat gets a reference to the given string and assigns it to the JsonFormat field.
func (o *LoggingHttpsResponse) SetJsonFormat(v string) {
	o.JsonFormat = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *LoggingHttpsResponse) SetPeriod(v int32) {
	o.Period = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingHttpsResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingHttpsResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingHttpsResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingHttpsResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingHttpsResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingHttpsResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingHttpsResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *LoggingHttpsResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingHttpsResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingHttpsResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingHttpsResponse) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHttpsResponse) MarshalJSON() ([]byte, error) {
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
	if o.LogProcessingRegion != nil {
		toSerialize["log_processing_region"] = o.LogProcessingRegion
	}
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
	}
	if o.TlsCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TlsCaCert.Get()
	}
	if o.TlsClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TlsClientCert.Get()
	}
	if o.TlsClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TlsClientKey.Get()
	}
	if o.TlsHostname.IsSet() {
		toSerialize["tls_hostname"] = o.TlsHostname.Get()
	}
	if o.RequestMaxEntries != nil {
		toSerialize["request_max_entries"] = o.RequestMaxEntries
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
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
	if o.JsonFormat != nil {
		toSerialize["json_format"] = o.JsonFormat
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
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
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
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
func (o *LoggingHttpsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHttpsResponse := _LoggingHttpsResponse{}

	if err = json.Unmarshal(bytes, &varLoggingHttpsResponse); err == nil {
		*o = LoggingHttpsResponse(varLoggingHttpsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "placement")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "format")
		delete(additionalProperties, "log_processing_region")
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
		delete(additionalProperties, "period")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHttpsResponse is a helper abstraction for handling nullable logginghttpsresponse types.
type NullableLoggingHttpsResponse struct {
	value *LoggingHttpsResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHttpsResponse) Get() *LoggingHttpsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHttpsResponse) Set(val *LoggingHttpsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHttpsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHttpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHttpsResponse returns a pointer to a new instance of NullableLoggingHttpsResponse.
func NewNullableLoggingHttpsResponse(val *LoggingHttpsResponse) *NullableLoggingHttpsResponse {
	return &NullableLoggingHttpsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHttpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingHttpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
