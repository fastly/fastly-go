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

// LoggingElasticsearchResponse struct for LoggingElasticsearchResponse
type LoggingElasticsearchResponse struct {
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	Placement NullableString `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest.
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
	// The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today's date.
	Index *string `json:"index,omitempty"`
	// The URL to stream logs to. Must use HTTPS.
	Url *string `json:"url,omitempty"`
	// The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html).
	Pipeline NullableString `json:"pipeline,omitempty"`
	// Basic Auth username.
	User NullableString `json:"user,omitempty"`
	// Basic Auth password.
	Password NullableString `json:"password,omitempty"`
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

type _LoggingElasticsearchResponse LoggingElasticsearchResponse

// NewLoggingElasticsearchResponse instantiates a new LoggingElasticsearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingElasticsearchResponse() *LoggingElasticsearchResponse {
	this := LoggingElasticsearchResponse{}
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
	return &this
}

// NewLoggingElasticsearchResponseWithDefaults instantiates a new LoggingElasticsearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingElasticsearchResponseWithDefaults() *LoggingElasticsearchResponse {
	this := LoggingElasticsearchResponse{}
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
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingElasticsearchResponse) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetPlacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingElasticsearchResponse) SetPlacement(v string) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingElasticsearchResponse) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetPlacement() {
	o.Placement.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetResponseConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingElasticsearchResponse) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}

// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingElasticsearchResponse) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingElasticsearchResponse) SetFormat(v string) {
	o.Format = &v
}

// GetLogProcessingRegion returns the LogProcessingRegion field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetLogProcessingRegion() string {
	if o == nil || o.LogProcessingRegion == nil {
		var ret string
		return ret
	}
	return *o.LogProcessingRegion
}

// GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetLogProcessingRegionOk() (*string, bool) {
	if o == nil || o.LogProcessingRegion == nil {
		return nil, false
	}
	return o.LogProcessingRegion, true
}

// HasLogProcessingRegion returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasLogProcessingRegion() bool {
	if o != nil && o.LogProcessingRegion != nil {
		return true
	}

	return false
}

// SetLogProcessingRegion gets a reference to the given string and assigns it to the LogProcessingRegion field.
func (o *LoggingElasticsearchResponse) SetLogProcessingRegion(v string) {
	o.LogProcessingRegion = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingElasticsearchResponse) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *LoggingElasticsearchResponse) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *LoggingElasticsearchResponse) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *LoggingElasticsearchResponse) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *LoggingElasticsearchResponse) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *LoggingElasticsearchResponse) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *LoggingElasticsearchResponse) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsHostname returns the TlsHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetTlsHostname() string {
	if o == nil || o.TlsHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsHostname.Get()
}

// GetTlsHostnameOk returns a tuple with the TlsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetTlsHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsHostname.Get(), o.TlsHostname.IsSet()
}

// HasTlsHostname returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasTlsHostname() bool {
	if o != nil && o.TlsHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsHostname gets a reference to the given NullableString and assigns it to the TlsHostname field.
func (o *LoggingElasticsearchResponse) SetTlsHostname(v string) {
	o.TlsHostname.Set(&v)
}

// SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil
func (o *LoggingElasticsearchResponse) SetTlsHostnameNil() {
	o.TlsHostname.Set(nil)
}

// UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetTlsHostname() {
	o.TlsHostname.Unset()
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingElasticsearchResponse) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingElasticsearchResponse) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *LoggingElasticsearchResponse) SetIndex(v string) {
	o.Index = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LoggingElasticsearchResponse) SetUrl(v string) {
	o.Url = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetPipeline() string {
	if o == nil || o.Pipeline.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pipeline.Get()
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetPipelineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pipeline.Get(), o.Pipeline.IsSet()
}

// HasPipeline returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasPipeline() bool {
	if o != nil && o.Pipeline.IsSet() {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given NullableString and assigns it to the Pipeline field.
func (o *LoggingElasticsearchResponse) SetPipeline(v string) {
	o.Pipeline.Set(&v)
}

// SetPipelineNil sets the value for Pipeline to be an explicit nil
func (o *LoggingElasticsearchResponse) SetPipelineNil() {
	o.Pipeline.Set(nil)
}

// UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetPipeline() {
	o.Pipeline.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *LoggingElasticsearchResponse) SetUser(v string) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *LoggingElasticsearchResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetUser() {
	o.User.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *LoggingElasticsearchResponse) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *LoggingElasticsearchResponse) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetPassword() {
	o.Password.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingElasticsearchResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingElasticsearchResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingElasticsearchResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingElasticsearchResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingElasticsearchResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingElasticsearchResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingElasticsearchResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingElasticsearchResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingElasticsearchResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *LoggingElasticsearchResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingElasticsearchResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingElasticsearchResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingElasticsearchResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingElasticsearchResponse) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingElasticsearchResponse) MarshalJSON() ([]byte, error) {
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
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
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
func (o *LoggingElasticsearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingElasticsearchResponse := _LoggingElasticsearchResponse{}

	if err = json.Unmarshal(bytes, &varLoggingElasticsearchResponse); err == nil {
		*o = LoggingElasticsearchResponse(varLoggingElasticsearchResponse)
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
		delete(additionalProperties, "index")
		delete(additionalProperties, "url")
		delete(additionalProperties, "pipeline")
		delete(additionalProperties, "user")
		delete(additionalProperties, "password")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingElasticsearchResponse is a helper abstraction for handling nullable loggingelasticsearchresponse types.
type NullableLoggingElasticsearchResponse struct {
	value *LoggingElasticsearchResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingElasticsearchResponse) Get() *LoggingElasticsearchResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingElasticsearchResponse) Set(val *LoggingElasticsearchResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingElasticsearchResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingElasticsearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingElasticsearchResponse returns a pointer to a new instance of NullableLoggingElasticsearchResponse.
func NewNullableLoggingElasticsearchResponse(val *LoggingElasticsearchResponse) *NullableLoggingElasticsearchResponse {
	return &NullableLoggingElasticsearchResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingElasticsearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingElasticsearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
