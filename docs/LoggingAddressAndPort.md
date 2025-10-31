# LoggingAddressAndPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A hostname or IPv4 address. | [optional] 
**Port** | Pointer to **int32** | The port number. | [optional] [default to 514]

## Methods

### NewLoggingAddressAndPort

`func NewLoggingAddressAndPort() *LoggingAddressAndPort`

NewLoggingAddressAndPort instantiates a new LoggingAddressAndPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingAddressAndPortWithDefaults

`func NewLoggingAddressAndPortWithDefaults() *LoggingAddressAndPort`

NewLoggingAddressAndPortWithDefaults instantiates a new LoggingAddressAndPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LoggingAddressAndPort) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoggingAddressAndPort) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoggingAddressAndPort) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoggingAddressAndPort) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *LoggingAddressAndPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingAddressAndPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingAddressAndPort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingAddressAndPort) HasPort() bool`

HasPort returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


