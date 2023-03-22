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

// LoggingAzureblob struct for LoggingAzureblob
type LoggingAzureblob struct {
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
	// How the message should be formatted.
	MessageType *string `json:"message_type,omitempty"`
	// A timestamp format
	TimestampFormat NullableString `json:"timestamp_format,omitempty"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int32 `json:"period,omitempty"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int32 `json:"gzip_level,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *string `json:"compression_codec,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// The unique Azure Blob Storage namespace in which your data objects are stored. Required.
	AccountName *string `json:"account_name,omitempty"`
	// The name of the Azure Blob Storage container in which to store logs. Required.
	Container *string `json:"container,omitempty"`
	// The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
	SasToken *string `json:"sas_token,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
	FileMaxBytes *int32 `json:"file_max_bytes,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingAzureblob LoggingAzureblob

// NewLoggingAzureblob instantiates a new LoggingAzureblob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingAzureblob() *LoggingAzureblob {
	this := LoggingAzureblob{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var messageType string = "classic"
	this.MessageType = &messageType
	var period int32 = 3600
	this.Period = &period
	var gzipLevel int32 = 0
	this.GzipLevel = &gzipLevel
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// NewLoggingAzureblobWithDefaults instantiates a new LoggingAzureblob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingAzureblobWithDefaults() *LoggingAzureblob {
	this := LoggingAzureblob{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var messageType string = "classic"
	this.MessageType = &messageType
	var period int32 = 3600
	this.Period = &period
	var gzipLevel int32 = 0
	this.GzipLevel = &gzipLevel
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingAzureblob) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblob) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblob) GetPlacementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingAzureblob) SetPlacement(v string) {
	o.Placement.Set(&v)
}
// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingAzureblob) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingAzureblob) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingAzureblob) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblob) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblob) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingAzureblob) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingAzureblob) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingAzureblob) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingAzureblob) SetFormat(v string) {
	o.Format = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *LoggingAzureblob) SetMessageType(v string) {
	o.MessageType = &v
}

// GetTimestampFormat returns the TimestampFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblob) GetTimestampFormat() string {
	if o == nil || o.TimestampFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimestampFormat.Get()
}

// GetTimestampFormatOk returns a tuple with the TimestampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblob) GetTimestampFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimestampFormat.Get(), o.TimestampFormat.IsSet()
}

// HasTimestampFormat returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasTimestampFormat() bool {
	if o != nil && o.TimestampFormat.IsSet() {
		return true
	}

	return false
}

// SetTimestampFormat gets a reference to the given NullableString and assigns it to the TimestampFormat field.
func (o *LoggingAzureblob) SetTimestampFormat(v string) {
	o.TimestampFormat.Set(&v)
}
// SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil
func (o *LoggingAzureblob) SetTimestampFormatNil() {
	o.TimestampFormat.Set(nil)
}

// UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
func (o *LoggingAzureblob) UnsetTimestampFormat() {
	o.TimestampFormat.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *LoggingAzureblob) SetPeriod(v int32) {
	o.Period = &v
}

// GetGzipLevel returns the GzipLevel field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetGzipLevel() int32 {
	if o == nil || o.GzipLevel == nil {
		var ret int32
		return ret
	}
	return *o.GzipLevel
}

// GetGzipLevelOk returns a tuple with the GzipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetGzipLevelOk() (*int32, bool) {
	if o == nil || o.GzipLevel == nil {
		return nil, false
	}
	return o.GzipLevel, true
}

// HasGzipLevel returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasGzipLevel() bool {
	if o != nil && o.GzipLevel != nil {
		return true
	}

	return false
}

// SetGzipLevel gets a reference to the given int32 and assigns it to the GzipLevel field.
func (o *LoggingAzureblob) SetGzipLevel(v int32) {
	o.GzipLevel = &v
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetCompressionCodecOk() (*string, bool) {
	if o == nil || o.CompressionCodec == nil {
		return nil, false
	}
	return o.CompressionCodec, true
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec != nil {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given string and assigns it to the CompressionCodec field.
func (o *LoggingAzureblob) SetCompressionCodec(v string) {
	o.CompressionCodec = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblob) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblob) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingAzureblob) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingAzureblob) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingAzureblob) UnsetPath() {
	o.Path.Unset()
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *LoggingAzureblob) SetAccountName(v string) {
	o.AccountName = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *LoggingAzureblob) SetContainer(v string) {
	o.Container = &v
}

// GetSasToken returns the SasToken field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetSasToken() string {
	if o == nil || o.SasToken == nil {
		var ret string
		return ret
	}
	return *o.SasToken
}

// GetSasTokenOk returns a tuple with the SasToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetSasTokenOk() (*string, bool) {
	if o == nil || o.SasToken == nil {
		return nil, false
	}
	return o.SasToken, true
}

// HasSasToken returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasSasToken() bool {
	if o != nil && o.SasToken != nil {
		return true
	}

	return false
}

// SetSasToken gets a reference to the given string and assigns it to the SasToken field.
func (o *LoggingAzureblob) SetSasToken(v string) {
	o.SasToken = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingAzureblob) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingAzureblob) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingAzureblob) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingAzureblob) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingAzureblob) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetFileMaxBytes returns the FileMaxBytes field value if set, zero value otherwise.
func (o *LoggingAzureblob) GetFileMaxBytes() int32 {
	if o == nil || o.FileMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.FileMaxBytes
}

// GetFileMaxBytesOk returns a tuple with the FileMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingAzureblob) GetFileMaxBytesOk() (*int32, bool) {
	if o == nil || o.FileMaxBytes == nil {
		return nil, false
	}
	return o.FileMaxBytes, true
}

// HasFileMaxBytes returns a boolean if a field has been set.
func (o *LoggingAzureblob) HasFileMaxBytes() bool {
	if o != nil && o.FileMaxBytes != nil {
		return true
	}

	return false
}

// SetFileMaxBytes gets a reference to the given int32 and assigns it to the FileMaxBytes field.
func (o *LoggingAzureblob) SetFileMaxBytes(v int32) {
	o.FileMaxBytes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingAzureblob) MarshalJSON() ([]byte, error) {
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
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.TimestampFormat.IsSet() {
		toSerialize["timestamp_format"] = o.TimestampFormat.Get()
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.GzipLevel != nil {
		toSerialize["gzip_level"] = o.GzipLevel
	}
	if o.CompressionCodec != nil {
		toSerialize["compression_codec"] = o.CompressionCodec
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.SasToken != nil {
		toSerialize["sas_token"] = o.SasToken
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.FileMaxBytes != nil {
		toSerialize["file_max_bytes"] = o.FileMaxBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingAzureblob) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingAzureblob := _LoggingAzureblob{}

	if err = json.Unmarshal(bytes, &varLoggingAzureblob); err == nil {
		*o = LoggingAzureblob(varLoggingAzureblob)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "placement")
		delete(additionalProperties, "format_version")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "format")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "timestamp_format")
		delete(additionalProperties, "period")
		delete(additionalProperties, "gzip_level")
		delete(additionalProperties, "compression_codec")
		delete(additionalProperties, "path")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "container")
		delete(additionalProperties, "sas_token")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "file_max_bytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingAzureblob is a helper abstraction for handling nullable loggingazureblob types. 
type NullableLoggingAzureblob struct {
	value *LoggingAzureblob
	isSet bool
}

// Get returns the value.
func (v NullableLoggingAzureblob) Get() *LoggingAzureblob {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingAzureblob) Set(val *LoggingAzureblob) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingAzureblob) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingAzureblob) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingAzureblob returns a pointer to a new instance of NullableLoggingAzureblob.
func NewNullableLoggingAzureblob(val *LoggingAzureblob) *NullableLoggingAzureblob {
	return &NullableLoggingAzureblob{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingAzureblob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingAzureblob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
