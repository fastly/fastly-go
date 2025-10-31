# LoggingLogentriesAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | The port number. | [optional] [default to 20000]
**Token** | Pointer to **string** | Use token based authentication. | [optional] 
**UseTls** | Pointer to [**LoggingUseTlsString**](LoggingUseTlsString.md) |  | [optional] [default to LOGGINGUSETLSSTRING_no_tls]
**Region** | Pointer to **string** | The region to which to stream logs. | [optional] 

## Methods

### NewLoggingLogentriesAdditional

`func NewLoggingLogentriesAdditional() *LoggingLogentriesAdditional`

NewLoggingLogentriesAdditional instantiates a new LoggingLogentriesAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogentriesAdditionalWithDefaults

`func NewLoggingLogentriesAdditionalWithDefaults() *LoggingLogentriesAdditional`

NewLoggingLogentriesAdditionalWithDefaults instantiates a new LoggingLogentriesAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *LoggingLogentriesAdditional) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingLogentriesAdditional) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingLogentriesAdditional) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingLogentriesAdditional) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *LoggingLogentriesAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingLogentriesAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingLogentriesAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingLogentriesAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUseTls

`func (o *LoggingLogentriesAdditional) GetUseTls() LoggingUseTlsString`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *LoggingLogentriesAdditional) GetUseTlsOk() (*LoggingUseTlsString, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *LoggingLogentriesAdditional) SetUseTls(v LoggingUseTlsString)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *LoggingLogentriesAdditional) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingLogentriesAdditional) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingLogentriesAdditional) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingLogentriesAdditional) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingLogentriesAdditional) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


