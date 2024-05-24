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

// LoggingKafkaResponsePost struct for LoggingKafkaResponsePost
type LoggingKafkaResponsePost struct {
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`. 
	Placement NullableString `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`. 
	FormatVersion *int32 `json:"format_version,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname NullableString `json:"tls_hostname,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version *string `json:"version,omitempty"`
	// The Kafka topic to send logs to. Required.
	Topic *string `json:"topic,omitempty"`
	// A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.
	Brokers *string `json:"brokers,omitempty"`
	// The codec used for compression of your logs.
	CompressionCodec NullableString `json:"compression_codec,omitempty"`
	// The number of acknowledgements a leader must receive before a write is considered successful.
	RequiredAcks *int32 `json:"required_acks,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (no limit).
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).
	ParseLogKeyvals *bool `json:"parse_log_keyvals,omitempty"`
	// SASL authentication method.
	AuthMethod *string `json:"auth_method,omitempty"`
	// SASL user.
	User *string `json:"user,omitempty"`
	// SASL password.
	Password *string `json:"password,omitempty"`
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingKafkaResponsePost LoggingKafkaResponsePost

// NewLoggingKafkaResponsePost instantiates a new LoggingKafkaResponsePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingKafkaResponsePost() *LoggingKafkaResponsePost {
	this := LoggingKafkaResponsePost{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TLSHostname = *NewNullableString(&tlsHostname)
	var requiredAcks int32 = 1
	this.RequiredAcks = &requiredAcks
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingKafkaResponsePostWithDefaults instantiates a new LoggingKafkaResponsePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingKafkaResponsePostWithDefaults() *LoggingKafkaResponsePost {
	this := LoggingKafkaResponsePost{}
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TLSHostname = *NewNullableString(&tlsHostname)
	var requiredAcks int32 = 1
	this.RequiredAcks = &requiredAcks
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLS = LOGGINGUSETLS_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingKafkaResponsePost) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetPlacementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingKafkaResponsePost) SetPlacement(v string) {
	o.Placement.Set(&v)
}
// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingKafkaResponsePost) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetPlacement() {
	o.Placement.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingKafkaResponsePost) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingKafkaResponsePost) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingKafkaResponsePost) SetFormat(v string) {
	o.Format = &v
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingKafkaResponsePost) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *LoggingKafkaResponsePost) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *LoggingKafkaResponsePost) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *LoggingKafkaResponsePost) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *LoggingKafkaResponsePost) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *LoggingKafkaResponsePost) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *LoggingKafkaResponsePost) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSHostname returns the TLSHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetTLSHostname() string {
	if o == nil || o.TLSHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSHostname.Get()
}

// GetTLSHostnameOk returns a tuple with the TLSHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetTLSHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSHostname.Get(), o.TLSHostname.IsSet()
}

// HasTLSHostname returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasTLSHostname() bool {
	if o != nil && o.TLSHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSHostname gets a reference to the given NullableString and assigns it to the TLSHostname field.
func (o *LoggingKafkaResponsePost) SetTLSHostname(v string) {
	o.TLSHostname.Set(&v)
}
// SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil
func (o *LoggingKafkaResponsePost) SetTLSHostnameNil() {
	o.TLSHostname.Set(nil)
}

// UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetTLSHostname() {
	o.TLSHostname.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *LoggingKafkaResponsePost) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *LoggingKafkaResponsePost) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *LoggingKafkaResponsePost) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *LoggingKafkaResponsePost) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *LoggingKafkaResponsePost) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *LoggingKafkaResponsePost) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *LoggingKafkaResponsePost) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoggingKafkaResponsePost) SetVersion(v string) {
	o.Version = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingKafkaResponsePost) SetTopic(v string) {
	o.Topic = &v
}

// GetBrokers returns the Brokers field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetBrokers() string {
	if o == nil || o.Brokers == nil {
		var ret string
		return ret
	}
	return *o.Brokers
}

// GetBrokersOk returns a tuple with the Brokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetBrokersOk() (*string, bool) {
	if o == nil || o.Brokers == nil {
		return nil, false
	}
	return o.Brokers, true
}

// HasBrokers returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasBrokers() bool {
	if o != nil && o.Brokers != nil {
		return true
	}

	return false
}

// SetBrokers gets a reference to the given string and assigns it to the Brokers field.
func (o *LoggingKafkaResponsePost) SetBrokers(v string) {
	o.Brokers = &v
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaResponsePost) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec.Get() == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec.Get()
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaResponsePost) GetCompressionCodecOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompressionCodec.Get(), o.CompressionCodec.IsSet()
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec.IsSet() {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given NullableString and assigns it to the CompressionCodec field.
func (o *LoggingKafkaResponsePost) SetCompressionCodec(v string) {
	o.CompressionCodec.Set(&v)
}
// SetCompressionCodecNil sets the value for CompressionCodec to be an explicit nil
func (o *LoggingKafkaResponsePost) SetCompressionCodecNil() {
	o.CompressionCodec.Set(nil)
}

// UnsetCompressionCodec ensures that no value is present for CompressionCodec, not even an explicit nil
func (o *LoggingKafkaResponsePost) UnsetCompressionCodec() {
	o.CompressionCodec.Unset()
}

// GetRequiredAcks returns the RequiredAcks field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetRequiredAcks() int32 {
	if o == nil || o.RequiredAcks == nil {
		var ret int32
		return ret
	}
	return *o.RequiredAcks
}

// GetRequiredAcksOk returns a tuple with the RequiredAcks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetRequiredAcksOk() (*int32, bool) {
	if o == nil || o.RequiredAcks == nil {
		return nil, false
	}
	return o.RequiredAcks, true
}

// HasRequiredAcks returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasRequiredAcks() bool {
	if o != nil && o.RequiredAcks != nil {
		return true
	}

	return false
}

// SetRequiredAcks gets a reference to the given int32 and assigns it to the RequiredAcks field.
func (o *LoggingKafkaResponsePost) SetRequiredAcks(v int32) {
	o.RequiredAcks = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingKafkaResponsePost) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetParseLogKeyvals returns the ParseLogKeyvals field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetParseLogKeyvals() bool {
	if o == nil || o.ParseLogKeyvals == nil {
		var ret bool
		return ret
	}
	return *o.ParseLogKeyvals
}

// GetParseLogKeyvalsOk returns a tuple with the ParseLogKeyvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetParseLogKeyvalsOk() (*bool, bool) {
	if o == nil || o.ParseLogKeyvals == nil {
		return nil, false
	}
	return o.ParseLogKeyvals, true
}

// HasParseLogKeyvals returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasParseLogKeyvals() bool {
	if o != nil && o.ParseLogKeyvals != nil {
		return true
	}

	return false
}

// SetParseLogKeyvals gets a reference to the given bool and assigns it to the ParseLogKeyvals field.
func (o *LoggingKafkaResponsePost) SetParseLogKeyvals(v bool) {
	o.ParseLogKeyvals = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetAuthMethod() string {
	if o == nil || o.AuthMethod == nil {
		var ret string
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetAuthMethodOk() (*string, bool) {
	if o == nil || o.AuthMethod == nil {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasAuthMethod() bool {
	if o != nil && o.AuthMethod != nil {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.
func (o *LoggingKafkaResponsePost) SetAuthMethod(v string) {
	o.AuthMethod = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingKafkaResponsePost) SetUser(v string) {
	o.User = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoggingKafkaResponsePost) SetPassword(v string) {
	o.Password = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingKafkaResponsePost) GetUseTLS() LoggingUseTLS {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLS
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaResponsePost) GetUseTLSOk() (*LoggingUseTLS, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingKafkaResponsePost) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLS and assigns it to the UseTLS field.
func (o *LoggingKafkaResponsePost) SetUseTLS(v LoggingUseTLS) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingKafkaResponsePost) MarshalJSON() ([]byte, error) {
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
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Brokers != nil {
		toSerialize["brokers"] = o.Brokers
	}
	if o.CompressionCodec.IsSet() {
		toSerialize["compression_codec"] = o.CompressionCodec.Get()
	}
	if o.RequiredAcks != nil {
		toSerialize["required_acks"] = o.RequiredAcks
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}
	if o.ParseLogKeyvals != nil {
		toSerialize["parse_log_keyvals"] = o.ParseLogKeyvals
	}
	if o.AuthMethod != nil {
		toSerialize["auth_method"] = o.AuthMethod
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
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
func (o *LoggingKafkaResponsePost) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingKafkaResponsePost := _LoggingKafkaResponsePost{}

	if err = json.Unmarshal(bytes, &varLoggingKafkaResponsePost); err == nil {
		*o = LoggingKafkaResponsePost(varLoggingKafkaResponsePost)
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
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "topic")
		delete(additionalProperties, "brokers")
		delete(additionalProperties, "compression_codec")
		delete(additionalProperties, "required_acks")
		delete(additionalProperties, "request_max_bytes")
		delete(additionalProperties, "parse_log_keyvals")
		delete(additionalProperties, "auth_method")
		delete(additionalProperties, "user")
		delete(additionalProperties, "password")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingKafkaResponsePost is a helper abstraction for handling nullable loggingkafkaresponsepost types. 
type NullableLoggingKafkaResponsePost struct {
	value *LoggingKafkaResponsePost
	isSet bool
}

// Get returns the value.
func (v NullableLoggingKafkaResponsePost) Get() *LoggingKafkaResponsePost {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingKafkaResponsePost) Set(val *LoggingKafkaResponsePost) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingKafkaResponsePost) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingKafkaResponsePost) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingKafkaResponsePost returns a pointer to a new instance of NullableLoggingKafkaResponsePost.
func NewNullableLoggingKafkaResponsePost(val *LoggingKafkaResponsePost) *NullableLoggingKafkaResponsePost {
	return &NullableLoggingKafkaResponsePost{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingKafkaResponsePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingKafkaResponsePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
