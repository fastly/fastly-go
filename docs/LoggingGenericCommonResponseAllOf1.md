# LoggingGenericCommonResponseAllOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** | How frequently log files are finalized so they can be available for reading (in seconds). | [optional] [default to "3600"]
**GzipLevel** | Pointer to **string** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] [default to "0"]

## Methods

### NewLoggingGenericCommonResponseAllOf1

`func NewLoggingGenericCommonResponseAllOf1() *LoggingGenericCommonResponseAllOf1`

NewLoggingGenericCommonResponseAllOf1 instantiates a new LoggingGenericCommonResponseAllOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingGenericCommonResponseAllOf1WithDefaults

`func NewLoggingGenericCommonResponseAllOf1WithDefaults() *LoggingGenericCommonResponseAllOf1`

NewLoggingGenericCommonResponseAllOf1WithDefaults instantiates a new LoggingGenericCommonResponseAllOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *LoggingGenericCommonResponseAllOf1) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingGenericCommonResponseAllOf1) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingGenericCommonResponseAllOf1) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingGenericCommonResponseAllOf1) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingGenericCommonResponseAllOf1) GetGzipLevel() string`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingGenericCommonResponseAllOf1) GetGzipLevelOk() (*string, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingGenericCommonResponseAllOf1) SetGzipLevel(v string)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingGenericCommonResponseAllOf1) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
