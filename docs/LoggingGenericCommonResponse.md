# LoggingGenericCommonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to **string** | How the message should be formatted. | [optional] [default to "classic"]
**TimestampFormat** | Pointer to **NullableString** | A timestamp format | [optional] [readonly] 
**CompressionCodec** | Pointer to **string** | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] 
**Period** | Pointer to **string** | How frequently log files are finalized so they can be available for reading (in seconds). | [optional] [default to "3600"]
**GzipLevel** | Pointer to **string** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] [default to "0"]

## Methods

### NewLoggingGenericCommonResponse

`func NewLoggingGenericCommonResponse() *LoggingGenericCommonResponse`

NewLoggingGenericCommonResponse instantiates a new LoggingGenericCommonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingGenericCommonResponseWithDefaults

`func NewLoggingGenericCommonResponseWithDefaults() *LoggingGenericCommonResponse`

NewLoggingGenericCommonResponseWithDefaults instantiates a new LoggingGenericCommonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *LoggingGenericCommonResponse) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingGenericCommonResponse) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingGenericCommonResponse) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingGenericCommonResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *LoggingGenericCommonResponse) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *LoggingGenericCommonResponse) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *LoggingGenericCommonResponse) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *LoggingGenericCommonResponse) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### SetTimestampFormatNil

`func (o *LoggingGenericCommonResponse) SetTimestampFormatNil(b bool)`

 SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil

### UnsetTimestampFormat
`func (o *LoggingGenericCommonResponse) UnsetTimestampFormat()`

UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
### GetCompressionCodec

`func (o *LoggingGenericCommonResponse) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingGenericCommonResponse) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingGenericCommonResponse) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingGenericCommonResponse) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetPeriod

`func (o *LoggingGenericCommonResponse) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingGenericCommonResponse) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingGenericCommonResponse) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingGenericCommonResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingGenericCommonResponse) GetGzipLevel() string`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingGenericCommonResponse) GetGzipLevelOk() (*string, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingGenericCommonResponse) SetGzipLevel(v string)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingGenericCommonResponse) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
