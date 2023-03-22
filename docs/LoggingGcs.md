# LoggingGcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**FormatVersion** | Pointer to **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to 2]
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**MessageType** | Pointer to **string** | How the message should be formatted. | [optional] [default to "classic"]
**TimestampFormat** | Pointer to **NullableString** | A timestamp format | [optional] [readonly] 
**Period** | Pointer to **int32** | How frequently log files are finalized so they can be available for reading (in seconds). | [optional] [default to 3600]
**GzipLevel** | Pointer to **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] [default to 0]
**CompressionCodec** | Pointer to **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] 
**User** | Pointer to **string** | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. | [optional] 
**SecretKey** | Pointer to **string** | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. | [optional] 
**AccountName** | Pointer to **string** | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. | [optional] 
**BucketName** | Pointer to **string** | The name of the GCS bucket. | [optional] 
**Path** | Pointer to **string** | The path to upload logs to. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**ProjectID** | Pointer to **string** | Your Google Cloud Platform project ID. Required | [optional] 

## Methods

### NewLoggingGcs

`func NewLoggingGcs() *LoggingGcs`

NewLoggingGcs instantiates a new LoggingGcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingGcsWithDefaults

`func NewLoggingGcsWithDefaults() *LoggingGcs`

NewLoggingGcsWithDefaults instantiates a new LoggingGcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingGcs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingGcs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingGcs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingGcs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingGcs) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingGcs) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingGcs) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingGcs) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingGcs) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingGcs) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingGcs) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingGcs) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingGcs) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingGcs) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetResponseCondition

`func (o *LoggingGcs) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingGcs) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingGcs) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingGcs) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingGcs) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingGcs) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingGcs) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingGcs) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingGcs) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingGcs) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMessageType

`func (o *LoggingGcs) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingGcs) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingGcs) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingGcs) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *LoggingGcs) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *LoggingGcs) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *LoggingGcs) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *LoggingGcs) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### SetTimestampFormatNil

`func (o *LoggingGcs) SetTimestampFormatNil(b bool)`

 SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil

### UnsetTimestampFormat
`func (o *LoggingGcs) UnsetTimestampFormat()`

UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
### GetPeriod

`func (o *LoggingGcs) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingGcs) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingGcs) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingGcs) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingGcs) GetGzipLevel() int32`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingGcs) GetGzipLevelOk() (*int32, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingGcs) SetGzipLevel(v int32)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingGcs) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *LoggingGcs) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingGcs) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingGcs) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingGcs) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetUser

`func (o *LoggingGcs) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingGcs) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingGcs) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingGcs) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingGcs) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingGcs) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingGcs) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingGcs) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetAccountName

`func (o *LoggingGcs) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *LoggingGcs) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *LoggingGcs) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *LoggingGcs) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingGcs) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingGcs) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingGcs) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingGcs) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPath

`func (o *LoggingGcs) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingGcs) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingGcs) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingGcs) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingGcs) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingGcs) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingGcs) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingGcs) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingGcs) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingGcs) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetProjectID

`func (o *LoggingGcs) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingGcs) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingGcs) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingGcs) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
