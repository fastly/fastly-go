# LoggingRequestCapsCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]

## Methods

### NewLoggingRequestCapsCommon

`func NewLoggingRequestCapsCommon() *LoggingRequestCapsCommon`

NewLoggingRequestCapsCommon instantiates a new LoggingRequestCapsCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingRequestCapsCommonWithDefaults

`func NewLoggingRequestCapsCommonWithDefaults() *LoggingRequestCapsCommon`

NewLoggingRequestCapsCommonWithDefaults instantiates a new LoggingRequestCapsCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestMaxEntries

`func (o *LoggingRequestCapsCommon) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingRequestCapsCommon) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingRequestCapsCommon) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingRequestCapsCommon) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingRequestCapsCommon) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingRequestCapsCommon) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingRequestCapsCommon) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingRequestCapsCommon) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
