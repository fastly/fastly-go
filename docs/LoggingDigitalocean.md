# LoggingDigitalocean

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
**BucketName** | Pointer to **string** | The name of the DigitalOcean Space. | [optional] 
**AccessKey** | Pointer to **string** | Your DigitalOcean Spaces account access key. | [optional] 
**SecretKey** | Pointer to **string** | Your DigitalOcean Spaces account secret key. | [optional] 
**Domain** | Pointer to **string** | The domain of the DigitalOcean Spaces endpoint. | [optional] [default to "nyc3.digitaloceanspaces.com"]
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]

## Methods

### NewLoggingDigitalocean

`func NewLoggingDigitalocean() *LoggingDigitalocean`

NewLoggingDigitalocean instantiates a new LoggingDigitalocean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDigitaloceanWithDefaults

`func NewLoggingDigitaloceanWithDefaults() *LoggingDigitalocean`

NewLoggingDigitaloceanWithDefaults instantiates a new LoggingDigitalocean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingDigitalocean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingDigitalocean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingDigitalocean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingDigitalocean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingDigitalocean) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingDigitalocean) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingDigitalocean) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingDigitalocean) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingDigitalocean) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingDigitalocean) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingDigitalocean) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingDigitalocean) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingDigitalocean) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingDigitalocean) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetResponseCondition

`func (o *LoggingDigitalocean) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingDigitalocean) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingDigitalocean) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingDigitalocean) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingDigitalocean) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingDigitalocean) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingDigitalocean) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingDigitalocean) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingDigitalocean) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingDigitalocean) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMessageType

`func (o *LoggingDigitalocean) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingDigitalocean) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingDigitalocean) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingDigitalocean) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *LoggingDigitalocean) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *LoggingDigitalocean) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *LoggingDigitalocean) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *LoggingDigitalocean) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### SetTimestampFormatNil

`func (o *LoggingDigitalocean) SetTimestampFormatNil(b bool)`

 SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil

### UnsetTimestampFormat
`func (o *LoggingDigitalocean) UnsetTimestampFormat()`

UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
### GetPeriod

`func (o *LoggingDigitalocean) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingDigitalocean) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingDigitalocean) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingDigitalocean) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingDigitalocean) GetGzipLevel() int32`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingDigitalocean) GetGzipLevelOk() (*int32, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingDigitalocean) SetGzipLevel(v int32)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingDigitalocean) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *LoggingDigitalocean) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingDigitalocean) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingDigitalocean) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingDigitalocean) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingDigitalocean) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingDigitalocean) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingDigitalocean) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingDigitalocean) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetAccessKey

`func (o *LoggingDigitalocean) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingDigitalocean) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingDigitalocean) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingDigitalocean) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingDigitalocean) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingDigitalocean) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingDigitalocean) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingDigitalocean) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetDomain

`func (o *LoggingDigitalocean) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LoggingDigitalocean) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LoggingDigitalocean) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LoggingDigitalocean) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *LoggingDigitalocean) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingDigitalocean) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingDigitalocean) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingDigitalocean) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingDigitalocean) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingDigitalocean) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingDigitalocean) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingDigitalocean) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingDigitalocean) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingDigitalocean) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingDigitalocean) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingDigitalocean) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
