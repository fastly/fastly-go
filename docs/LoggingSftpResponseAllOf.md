# LoggingSftpResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A hostname or IPv4 address. | [optional] 
**Port** | Pointer to **string** | The port number. | [optional] [default to "22"]
**Period** | Pointer to **string** | How frequently log files are finalized so they can be available for reading (in seconds). | [optional] [default to "3600"]
**GzipLevel** | Pointer to **int32** | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error. | [optional] [default to 0]

## Methods

### NewLoggingSftpResponseAllOf

`func NewLoggingSftpResponseAllOf() *LoggingSftpResponseAllOf`

NewLoggingSftpResponseAllOf instantiates a new LoggingSftpResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSftpResponseAllOfWithDefaults

`func NewLoggingSftpResponseAllOfWithDefaults() *LoggingSftpResponseAllOf`

NewLoggingSftpResponseAllOfWithDefaults instantiates a new LoggingSftpResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LoggingSftpResponseAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoggingSftpResponseAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoggingSftpResponseAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoggingSftpResponseAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *LoggingSftpResponseAllOf) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingSftpResponseAllOf) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingSftpResponseAllOf) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingSftpResponseAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPeriod

`func (o *LoggingSftpResponseAllOf) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingSftpResponseAllOf) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingSftpResponseAllOf) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingSftpResponseAllOf) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetGzipLevel

`func (o *LoggingSftpResponseAllOf) GetGzipLevel() int32`

GetGzipLevel returns the GzipLevel field if non-nil, zero value otherwise.

### GetGzipLevelOk

`func (o *LoggingSftpResponseAllOf) GetGzipLevelOk() (*int32, bool)`

GetGzipLevelOk returns a tuple with the GzipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipLevel

`func (o *LoggingSftpResponseAllOf) SetGzipLevel(v int32)`

SetGzipLevel sets GzipLevel field to given value.

### HasGzipLevel

`func (o *LoggingSftpResponseAllOf) HasGzipLevel() bool`

HasGzipLevel returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


