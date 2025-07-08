# LoggingCloudfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**LogProcessingRegion** | Pointer to **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [optional] [default to "none"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**MessageType** | Pointer to **string** | How the message should be formatted. | [optional] [default to "classic"]
**TimestampFormat** | Pointer to **NullableString** | A timestamp format | [optional] [readonly] 
**CompressionCodec** | Pointer to **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] 
**Period** | Pointer to **string** | How frequently log files are finalized so they can be available for reading (in seconds). | [optional] [default to "3600"]
**GzipLevel** | Pointer to **string** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] [default to "0"]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**AccessKey** | Pointer to **string** | Your Cloud Files account access key. | [optional] 
**BucketName** | Pointer to **string** | The name of your Cloud Files container. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**Region** | Pointer to **NullableString** | The region to stream logs to. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**User** | Pointer to **string** | The username for your Cloud Files account. | [optional] 

## Methods

### NewLoggingCloudfilesResponse

`func NewLoggingCloudfilesResponse() *LoggingCloudfilesResponse`

NewLoggingCloudfilesResponse instantiates a new LoggingCloudfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingCloudfilesResponseWithDefaults

`func NewLoggingCloudfilesResponseWithDefaults() *LoggingCloudfilesResponse`

NewLoggingCloudfilesResponseWithDefaults instantiates a new LoggingCloudfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingCloudfilesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingCloudfilesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingCloudfilesResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingCloudfilesResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingCloudfilesResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingCloudfilesResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingCloudfilesResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingCloudfilesResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingCloudfilesResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingCloudfilesResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingCloudfilesResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingCloudfilesResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingCloudfilesResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingCloudfilesResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingCloudfilesResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingCloudfilesResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingCloudfilesResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingCloudfilesResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingCloudfilesResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingCloudfilesResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLogProcessingRegion

`func (o *LoggingCloudfilesResponse) GetLogProcessingRegion() string`

GetLogProcessingRegion returns the LogProcessingRegion field if non-nil, zero value otherwise.

### GetLogProcessingRegionOk

`func (o *LoggingCloudfilesResponse) GetLogProcessingRegionOk() (*string, bool)`

GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogProcessingRegion

`func (o *LoggingCloudfilesResponse) SetLogProcessingRegion(v string)`

SetLogProcessingRegion sets LogProcessingRegion field to given value.

### HasLogProcessingRegion

`func (o *LoggingCloudfilesResponse) HasLogProcessingRegion() bool`

HasLogProcessingRegion returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingCloudfilesResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingCloudfilesResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingCloudfilesResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingCloudfilesResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetMessageType

`func (o *LoggingCloudfilesResponse) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingCloudfilesResponse) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingCloudfilesResponse) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingCloudfilesResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *LoggingCloudfilesResponse) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *LoggingCloudfilesResponse) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *LoggingCloudfilesResponse) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *LoggingCloudfilesResponse) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### SetTimestampFormatNil

`func (o *LoggingCloudfilesResponse) SetTimestampFormatNil(b bool)`

 SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil

### UnsetTimestampFormat
`func (o *LoggingCloudfilesResponse) UnsetTimestampFormat()`

UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
### GetCompressionCodec

`func (o *LoggingCloudfilesResponse) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingCloudfilesResponse) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingCloudfilesResponse) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingCloudfilesResponse) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetPeriod

`func (o *LoggingCloudfilesResponse) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingCloudfilesResponse) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingCloudfilesResponse) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingCloudfilesResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingCloudfilesResponse) GetGzipLevel() string`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingCloudfilesResponse) GetGzipLevelOk() (*string, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingCloudfilesResponse) SetGzipLevel(v string)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingCloudfilesResponse) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingCloudfilesResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingCloudfilesResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingCloudfilesResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingCloudfilesResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingCloudfilesResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingCloudfilesResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingCloudfilesResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingCloudfilesResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingCloudfilesResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingCloudfilesResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingCloudfilesResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingCloudfilesResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingCloudfilesResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingCloudfilesResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingCloudfilesResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingCloudfilesResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingCloudfilesResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingCloudfilesResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingCloudfilesResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingCloudfilesResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingCloudfilesResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingCloudfilesResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingCloudfilesResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingCloudfilesResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingCloudfilesResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingCloudfilesResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccessKey

`func (o *LoggingCloudfilesResponse) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingCloudfilesResponse) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingCloudfilesResponse) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingCloudfilesResponse) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingCloudfilesResponse) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingCloudfilesResponse) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingCloudfilesResponse) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingCloudfilesResponse) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPath

`func (o *LoggingCloudfilesResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingCloudfilesResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingCloudfilesResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingCloudfilesResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingCloudfilesResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingCloudfilesResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRegion

`func (o *LoggingCloudfilesResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingCloudfilesResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingCloudfilesResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingCloudfilesResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *LoggingCloudfilesResponse) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *LoggingCloudfilesResponse) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPublicKey

`func (o *LoggingCloudfilesResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingCloudfilesResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingCloudfilesResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingCloudfilesResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingCloudfilesResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingCloudfilesResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetUser

`func (o *LoggingCloudfilesResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingCloudfilesResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingCloudfilesResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingCloudfilesResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
