# LoggingLogentriesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | The port number. | [optional] [default to 20000]
**Token** | Pointer to **string** | Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)). | [optional] 
**UseTLS** | Pointer to [**LoggingUseTLS**](LoggingUseTLS.md) |  | [optional] [default to LOGGINGUSETLS_no_tls]
**Region** | Pointer to **string** | The region to which to stream logs. | [optional] 

## Methods

### NewLoggingLogentriesAllOf

`func NewLoggingLogentriesAllOf() *LoggingLogentriesAllOf`

NewLoggingLogentriesAllOf instantiates a new LoggingLogentriesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogentriesAllOfWithDefaults

`func NewLoggingLogentriesAllOfWithDefaults() *LoggingLogentriesAllOf`

NewLoggingLogentriesAllOfWithDefaults instantiates a new LoggingLogentriesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *LoggingLogentriesAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingLogentriesAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingLogentriesAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingLogentriesAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *LoggingLogentriesAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingLogentriesAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingLogentriesAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingLogentriesAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUseTLS

`func (o *LoggingLogentriesAllOf) GetUseTLS() LoggingUseTLS`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingLogentriesAllOf) GetUseTLSOk() (*LoggingUseTLS, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingLogentriesAllOf) SetUseTLS(v LoggingUseTLS)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingLogentriesAllOf) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingLogentriesAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingLogentriesAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingLogentriesAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingLogentriesAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
