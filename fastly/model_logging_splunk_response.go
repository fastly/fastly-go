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

// LoggingSplunkResponse struct for LoggingSplunkResponse
type LoggingSplunkResponse struct {
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
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// The URL to post logs to.
	Url *string `json:"url,omitempty"`
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token  *string              `json:"token,omitempty"`
	UseTls *LoggingUseTlsString `json:"use_tls,omitempty"`
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

type _LoggingSplunkResponse LoggingSplunkResponse

// NewLoggingSplunkResponse instantiates a new LoggingSplunkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingSplunkResponse() *LoggingSplunkResponse {
	this := LoggingSplunkResponse{}
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
	var useTls LoggingUseTlsString = LOGGINGUSETLSSTRING_no_tls
	this.UseTls = &useTls
	return &this
}

// NewLoggingSplunkResponseWithDefaults instantiates a new LoggingSplunkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingSplunkResponseWithDefaults() *LoggingSplunkResponse {
	this := LoggingSplunkResponse{}
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
	var useTls LoggingUseTlsString = LOGGINGUSETLSSTRING_no_tls
	this.UseTls = &useTls
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingSplunkResponse) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetPlacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingSplunkResponse) SetPlacement(v string) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingSplunkResponse) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetPlacement() {
	o.Placement.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetResponseConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingSplunkResponse) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}

// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingSplunkResponse) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingSplunkResponse) SetFormat(v string) {
	o.Format = &v
}

// GetLogProcessingRegion returns the LogProcessingRegion field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetLogProcessingRegion() string {
	if o == nil || o.LogProcessingRegion == nil {
		var ret string
		return ret
	}
	return *o.LogProcessingRegion
}

// GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetLogProcessingRegionOk() (*string, bool) {
	if o == nil || o.LogProcessingRegion == nil {
		return nil, false
	}
	return o.LogProcessingRegion, true
}

// HasLogProcessingRegion returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasLogProcessingRegion() bool {
	if o != nil && o.LogProcessingRegion != nil {
		return true
	}

	return false
}

// SetLogProcessingRegion gets a reference to the given string and assigns it to the LogProcessingRegion field.
func (o *LoggingSplunkResponse) SetLogProcessingRegion(v string) {
	o.LogProcessingRegion = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingSplunkResponse) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *LoggingSplunkResponse) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *LoggingSplunkResponse) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *LoggingSplunkResponse) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *LoggingSplunkResponse) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *LoggingSplunkResponse) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *LoggingSplunkResponse) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsHostname returns the TlsHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetTlsHostname() string {
	if o == nil || o.TlsHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsHostname.Get()
}

// GetTlsHostnameOk returns a tuple with the TlsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetTlsHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsHostname.Get(), o.TlsHostname.IsSet()
}

// HasTlsHostname returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasTlsHostname() bool {
	if o != nil && o.TlsHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsHostname gets a reference to the given NullableString and assigns it to the TlsHostname field.
func (o *LoggingSplunkResponse) SetTlsHostname(v string) {
	o.TlsHostname.Set(&v)
}

// SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil
func (o *LoggingSplunkResponse) SetTlsHostnameNil() {
	o.TlsHostname.Set(nil)
}

// UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetTlsHostname() {
	o.TlsHostname.Unset()
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingSplunkResponse) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingSplunkResponse) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LoggingSplunkResponse) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingSplunkResponse) SetToken(v string) {
	o.Token = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetUseTls() LoggingUseTlsString {
	if o == nil || o.UseTls == nil {
		var ret LoggingUseTlsString
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetUseTlsOk() (*LoggingUseTlsString, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given LoggingUseTlsString and assigns it to the UseTls field.
func (o *LoggingSplunkResponse) SetUseTls(v LoggingUseTlsString) {
	o.UseTls = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingSplunkResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingSplunkResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingSplunkResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingSplunkResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingSplunkResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingSplunkResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingSplunkResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingSplunkResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingSplunkResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *LoggingSplunkResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingSplunkResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingSplunkResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingSplunkResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingSplunkResponse) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingSplunkResponse) MarshalJSON() ([]byte, error) {
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
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
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
func (o *LoggingSplunkResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingSplunkResponse := _LoggingSplunkResponse{}

	if err = json.Unmarshal(bytes, &varLoggingSplunkResponse); err == nil {
		*o = LoggingSplunkResponse(varLoggingSplunkResponse)
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
		delete(additionalProperties, "token")
		delete(additionalProperties, "use_tls")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingSplunkResponse is a helper abstraction for handling nullable loggingsplunkresponse types.
type NullableLoggingSplunkResponse struct {
	value *LoggingSplunkResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingSplunkResponse) Get() *LoggingSplunkResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingSplunkResponse) Set(val *LoggingSplunkResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingSplunkResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingSplunkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingSplunkResponse returns a pointer to a new instance of NullableLoggingSplunkResponse.
func NewNullableLoggingSplunkResponse(val *LoggingSplunkResponse) *NullableLoggingSplunkResponse {
	return &NullableLoggingSplunkResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingSplunkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingSplunkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
