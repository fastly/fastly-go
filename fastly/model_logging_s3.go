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

// LoggingS3 struct for LoggingS3
type LoggingS3 struct {
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
	// The access key for your S3 account. Not required if `iam_role` is provided.
	AccessKey NullableString `json:"access_key,omitempty"`
	// The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information.
	ACL *string `json:"acl,omitempty"`
	// The bucket name for S3 account.
	BucketName *string `json:"bucket_name,omitempty"`
	// The domain of the Amazon S3 endpoint.
	Domain *string `json:"domain,omitempty"`
	// The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided.
	IamRole NullableString `json:"iam_role,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The S3 redundancy level.
	Redundancy NullableString `json:"redundancy,omitempty"`
	// The secret key for your S3 account. Not required if `iam_role` is provided.
	SecretKey NullableString `json:"secret_key,omitempty"`
	// Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`.
	ServerSideEncryptionKmsKeyID NullableString `json:"server_side_encryption_kms_key_id,omitempty"`
	// Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption.
	ServerSideEncryption NullableString `json:"server_side_encryption,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingS3 LoggingS3

// NewLoggingS3 instantiates a new LoggingS3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingS3() *LoggingS3 {
	this := LoggingS3{}
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
	var redundancy string = "null"
	this.Redundancy = *NewNullableString(&redundancy)
	var serverSideEncryptionKmsKeyID string = "null"
	this.ServerSideEncryptionKmsKeyID = *NewNullableString(&serverSideEncryptionKmsKeyID)
	var serverSideEncryption string = "null"
	this.ServerSideEncryption = *NewNullableString(&serverSideEncryption)
	return &this
}

// NewLoggingS3WithDefaults instantiates a new LoggingS3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingS3WithDefaults() *LoggingS3 {
	this := LoggingS3{}
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
	var redundancy string = "null"
	this.Redundancy = *NewNullableString(&redundancy)
	var serverSideEncryptionKmsKeyID string = "null"
	this.ServerSideEncryptionKmsKeyID = *NewNullableString(&serverSideEncryptionKmsKeyID)
	var serverSideEncryption string = "null"
	this.ServerSideEncryption = *NewNullableString(&serverSideEncryption)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingS3) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingS3) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingS3) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetPlacement() string {
	if o == nil || o.Placement.Get() == nil {
		var ret string
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetPlacementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *LoggingS3) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableString and assigns it to the Placement field.
func (o *LoggingS3) SetPlacement(v string) {
	o.Placement.Set(&v)
}
// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *LoggingS3) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *LoggingS3) UnsetPlacement() {
	o.Placement.Unset()
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingS3) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingS3) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingS3) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *LoggingS3) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *LoggingS3) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *LoggingS3) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *LoggingS3) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingS3) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingS3) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingS3) SetFormat(v string) {
	o.Format = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingS3) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingS3) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *LoggingS3) SetMessageType(v string) {
	o.MessageType = &v
}

// GetTimestampFormat returns the TimestampFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetTimestampFormat() string {
	if o == nil || o.TimestampFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimestampFormat.Get()
}

// GetTimestampFormatOk returns a tuple with the TimestampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetTimestampFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimestampFormat.Get(), o.TimestampFormat.IsSet()
}

// HasTimestampFormat returns a boolean if a field has been set.
func (o *LoggingS3) HasTimestampFormat() bool {
	if o != nil && o.TimestampFormat.IsSet() {
		return true
	}

	return false
}

// SetTimestampFormat gets a reference to the given NullableString and assigns it to the TimestampFormat field.
func (o *LoggingS3) SetTimestampFormat(v string) {
	o.TimestampFormat.Set(&v)
}
// SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil
func (o *LoggingS3) SetTimestampFormatNil() {
	o.TimestampFormat.Set(nil)
}

// UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
func (o *LoggingS3) UnsetTimestampFormat() {
	o.TimestampFormat.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingS3) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingS3) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *LoggingS3) SetPeriod(v int32) {
	o.Period = &v
}

// GetGzipLevel returns the GzipLevel field value if set, zero value otherwise.
func (o *LoggingS3) GetGzipLevel() int32 {
	if o == nil || o.GzipLevel == nil {
		var ret int32
		return ret
	}
	return *o.GzipLevel
}

// GetGzipLevelOk returns a tuple with the GzipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetGzipLevelOk() (*int32, bool) {
	if o == nil || o.GzipLevel == nil {
		return nil, false
	}
	return o.GzipLevel, true
}

// HasGzipLevel returns a boolean if a field has been set.
func (o *LoggingS3) HasGzipLevel() bool {
	if o != nil && o.GzipLevel != nil {
		return true
	}

	return false
}

// SetGzipLevel gets a reference to the given int32 and assigns it to the GzipLevel field.
func (o *LoggingS3) SetGzipLevel(v int32) {
	o.GzipLevel = &v
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise.
func (o *LoggingS3) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetCompressionCodecOk() (*string, bool) {
	if o == nil || o.CompressionCodec == nil {
		return nil, false
	}
	return o.CompressionCodec, true
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingS3) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec != nil {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given string and assigns it to the CompressionCodec field.
func (o *LoggingS3) SetCompressionCodec(v string) {
	o.CompressionCodec = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingS3) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *LoggingS3) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}
// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *LoggingS3) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *LoggingS3) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetACL returns the ACL field value if set, zero value otherwise.
func (o *LoggingS3) GetACL() string {
	if o == nil || o.ACL == nil {
		var ret string
		return ret
	}
	return *o.ACL
}

// GetACLOk returns a tuple with the ACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetACLOk() (*string, bool) {
	if o == nil || o.ACL == nil {
		return nil, false
	}
	return o.ACL, true
}

// HasACL returns a boolean if a field has been set.
func (o *LoggingS3) HasACL() bool {
	if o != nil && o.ACL != nil {
		return true
	}

	return false
}

// SetACL gets a reference to the given string and assigns it to the ACL field.
func (o *LoggingS3) SetACL(v string) {
	o.ACL = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingS3) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingS3) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingS3) SetBucketName(v string) {
	o.BucketName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LoggingS3) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LoggingS3) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LoggingS3) SetDomain(v string) {
	o.Domain = &v
}

// GetIamRole returns the IamRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetIamRole() string {
	if o == nil || o.IamRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.IamRole.Get()
}

// GetIamRoleOk returns a tuple with the IamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetIamRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IamRole.Get(), o.IamRole.IsSet()
}

// HasIamRole returns a boolean if a field has been set.
func (o *LoggingS3) HasIamRole() bool {
	if o != nil && o.IamRole.IsSet() {
		return true
	}

	return false
}

// SetIamRole gets a reference to the given NullableString and assigns it to the IamRole field.
func (o *LoggingS3) SetIamRole(v string) {
	o.IamRole.Set(&v)
}
// SetIamRoleNil sets the value for IamRole to be an explicit nil
func (o *LoggingS3) SetIamRoleNil() {
	o.IamRole.Set(nil)
}

// UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
func (o *LoggingS3) UnsetIamRole() {
	o.IamRole.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingS3) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingS3) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingS3) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingS3) UnsetPath() {
	o.Path.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingS3) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingS3) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingS3) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingS3) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetRedundancy() string {
	if o == nil || o.Redundancy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Redundancy.Get()
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetRedundancyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Redundancy.Get(), o.Redundancy.IsSet()
}

// HasRedundancy returns a boolean if a field has been set.
func (o *LoggingS3) HasRedundancy() bool {
	if o != nil && o.Redundancy.IsSet() {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given NullableString and assigns it to the Redundancy field.
func (o *LoggingS3) SetRedundancy(v string) {
	o.Redundancy.Set(&v)
}
// SetRedundancyNil sets the value for Redundancy to be an explicit nil
func (o *LoggingS3) SetRedundancyNil() {
	o.Redundancy.Set(nil)
}

// UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
func (o *LoggingS3) UnsetRedundancy() {
	o.Redundancy.Unset()
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingS3) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingS3) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingS3) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingS3) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetServerSideEncryptionKmsKeyID returns the ServerSideEncryptionKmsKeyID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetServerSideEncryptionKmsKeyID() string {
	if o == nil || o.ServerSideEncryptionKmsKeyID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerSideEncryptionKmsKeyID.Get()
}

// GetServerSideEncryptionKmsKeyIDOk returns a tuple with the ServerSideEncryptionKmsKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetServerSideEncryptionKmsKeyIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerSideEncryptionKmsKeyID.Get(), o.ServerSideEncryptionKmsKeyID.IsSet()
}

// HasServerSideEncryptionKmsKeyID returns a boolean if a field has been set.
func (o *LoggingS3) HasServerSideEncryptionKmsKeyID() bool {
	if o != nil && o.ServerSideEncryptionKmsKeyID.IsSet() {
		return true
	}

	return false
}

// SetServerSideEncryptionKmsKeyID gets a reference to the given NullableString and assigns it to the ServerSideEncryptionKmsKeyID field.
func (o *LoggingS3) SetServerSideEncryptionKmsKeyID(v string) {
	o.ServerSideEncryptionKmsKeyID.Set(&v)
}
// SetServerSideEncryptionKmsKeyIDNil sets the value for ServerSideEncryptionKmsKeyID to be an explicit nil
func (o *LoggingS3) SetServerSideEncryptionKmsKeyIDNil() {
	o.ServerSideEncryptionKmsKeyID.Set(nil)
}

// UnsetServerSideEncryptionKmsKeyID ensures that no value is present for ServerSideEncryptionKmsKeyID, not even an explicit nil
func (o *LoggingS3) UnsetServerSideEncryptionKmsKeyID() {
	o.ServerSideEncryptionKmsKeyID.Unset()
}

// GetServerSideEncryption returns the ServerSideEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3) GetServerSideEncryption() string {
	if o == nil || o.ServerSideEncryption.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerSideEncryption.Get()
}

// GetServerSideEncryptionOk returns a tuple with the ServerSideEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3) GetServerSideEncryptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerSideEncryption.Get(), o.ServerSideEncryption.IsSet()
}

// HasServerSideEncryption returns a boolean if a field has been set.
func (o *LoggingS3) HasServerSideEncryption() bool {
	if o != nil && o.ServerSideEncryption.IsSet() {
		return true
	}

	return false
}

// SetServerSideEncryption gets a reference to the given NullableString and assigns it to the ServerSideEncryption field.
func (o *LoggingS3) SetServerSideEncryption(v string) {
	o.ServerSideEncryption.Set(&v)
}
// SetServerSideEncryptionNil sets the value for ServerSideEncryption to be an explicit nil
func (o *LoggingS3) SetServerSideEncryptionNil() {
	o.ServerSideEncryption.Set(nil)
}

// UnsetServerSideEncryption ensures that no value is present for ServerSideEncryption, not even an explicit nil
func (o *LoggingS3) UnsetServerSideEncryption() {
	o.ServerSideEncryption.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingS3) MarshalJSON() ([]byte, error) {
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
	if o.AccessKey.IsSet() {
		toSerialize["access_key"] = o.AccessKey.Get()
	}
	if o.ACL != nil {
		toSerialize["acl"] = o.ACL
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.IamRole.IsSet() {
		toSerialize["iam_role"] = o.IamRole.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.Redundancy.IsSet() {
		toSerialize["redundancy"] = o.Redundancy.Get()
	}
	if o.SecretKey.IsSet() {
		toSerialize["secret_key"] = o.SecretKey.Get()
	}
	if o.ServerSideEncryptionKmsKeyID.IsSet() {
		toSerialize["server_side_encryption_kms_key_id"] = o.ServerSideEncryptionKmsKeyID.Get()
	}
	if o.ServerSideEncryption.IsSet() {
		toSerialize["server_side_encryption"] = o.ServerSideEncryption.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingS3) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingS3 := _LoggingS3{}

	if err = json.Unmarshal(bytes, &varLoggingS3); err == nil {
		*o = LoggingS3(varLoggingS3)
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
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "acl")
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "iam_role")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "server_side_encryption_kms_key_id")
		delete(additionalProperties, "server_side_encryption")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingS3 is a helper abstraction for handling nullable loggings3 types. 
type NullableLoggingS3 struct {
	value *LoggingS3
	isSet bool
}

// Get returns the value.
func (v NullableLoggingS3) Get() *LoggingS3 {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingS3) Set(val *LoggingS3) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingS3) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingS3) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingS3 returns a pointer to a new instance of NullableLoggingS3.
func NewNullableLoggingS3(val *LoggingS3) *NullableLoggingS3 {
	return &NullableLoggingS3{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingS3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingS3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
