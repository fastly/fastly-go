# LoggingAzureblob

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
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**AccountName** | Pointer to **string** | The unique Azure Blob Storage namespace in which your data objects are stored. Required. | [optional] 
**Container** | Pointer to **string** | The name of the Azure Blob Storage container in which to store logs. Required. | [optional] 
**SasToken** | Pointer to **string** | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**FileMaxBytes** | Pointer to **int32** | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.) | [optional] 

## Methods

### NewLoggingAzureblob

`func NewLoggingAzureblob() *LoggingAzureblob`

NewLoggingAzureblob instantiates a new LoggingAzureblob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingAzureblobWithDefaults

`func NewLoggingAzureblobWithDefaults() *LoggingAzureblob`

NewLoggingAzureblobWithDefaults instantiates a new LoggingAzureblob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingAzureblob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingAzureblob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingAzureblob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingAzureblob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingAzureblob) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingAzureblob) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingAzureblob) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingAzureblob) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingAzureblob) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingAzureblob) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingAzureblob) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingAzureblob) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingAzureblob) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingAzureblob) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetResponseCondition

`func (o *LoggingAzureblob) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingAzureblob) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingAzureblob) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingAzureblob) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingAzureblob) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingAzureblob) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingAzureblob) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingAzureblob) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingAzureblob) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingAzureblob) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMessageType

`func (o *LoggingAzureblob) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingAzureblob) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingAzureblob) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingAzureblob) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *LoggingAzureblob) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *LoggingAzureblob) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *LoggingAzureblob) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *LoggingAzureblob) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### SetTimestampFormatNil

`func (o *LoggingAzureblob) SetTimestampFormatNil(b bool)`

 SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil

### UnsetTimestampFormat
`func (o *LoggingAzureblob) UnsetTimestampFormat()`

UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
### GetPeriod

`func (o *LoggingAzureblob) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingAzureblob) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingAzureblob) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingAzureblob) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingAzureblob) GetGzipLevel() int32`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingAzureblob) GetGzipLevelOk() (*int32, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingAzureblob) SetGzipLevel(v int32)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingAzureblob) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *LoggingAzureblob) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingAzureblob) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingAzureblob) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingAzureblob) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetPath

`func (o *LoggingAzureblob) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingAzureblob) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingAzureblob) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingAzureblob) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingAzureblob) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingAzureblob) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetAccountName

`func (o *LoggingAzureblob) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *LoggingAzureblob) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *LoggingAzureblob) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *LoggingAzureblob) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetContainer

`func (o *LoggingAzureblob) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LoggingAzureblob) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LoggingAzureblob) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *LoggingAzureblob) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetSasToken

`func (o *LoggingAzureblob) GetSasToken() string`

GetSasToken returns the SasToken field if non-nil, zero value otherwise.

### GetSasTokenOk

`func (o *LoggingAzureblob) GetSasTokenOk() (*string, bool)`

GetSasTokenOk returns a tuple with the SasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasToken

`func (o *LoggingAzureblob) SetSasToken(v string)`

SetSasToken sets SasToken field to given value.

### HasSasToken

`func (o *LoggingAzureblob) HasSasToken() bool`

HasSasToken returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingAzureblob) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingAzureblob) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingAzureblob) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingAzureblob) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingAzureblob) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingAzureblob) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetFileMaxBytes

`func (o *LoggingAzureblob) GetFileMaxBytes() int32`

GetFileMaxBytes returns the FileMaxBytes field if non-nil, zero value otherwise.

### GetFileMaxBytesOk

`func (o *LoggingAzureblob) GetFileMaxBytesOk() (*int32, bool)`

GetFileMaxBytesOk returns a tuple with the FileMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMaxBytes

`func (o *LoggingAzureblob) SetFileMaxBytes(v int32)`

SetFileMaxBytes sets FileMaxBytes field to given value.

### HasFileMaxBytes

`func (o *LoggingAzureblob) HasFileMaxBytes() bool`

HasFileMaxBytes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
