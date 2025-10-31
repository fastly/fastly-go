# LoggingSyslogAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**Hostname** | Pointer to **string** | The hostname used for the syslog endpoint. | [optional] 
**Ipv4** | Pointer to **NullableString** | The IPv4 address used for the syslog endpoint. | [optional] 
**Token** | Pointer to **NullableString** | Whether to prepend each message with a specific token. | [optional] [default to "null"]
**UseTls** | Pointer to [**LoggingUseTlsString**](LoggingUseTlsString.md) |  | [optional] [default to LOGGINGUSETLSSTRING_no_tls]

## Methods

### NewLoggingSyslogAdditional

`func NewLoggingSyslogAdditional() *LoggingSyslogAdditional`

NewLoggingSyslogAdditional instantiates a new LoggingSyslogAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSyslogAdditionalWithDefaults

`func NewLoggingSyslogAdditionalWithDefaults() *LoggingSyslogAdditional`

NewLoggingSyslogAdditionalWithDefaults instantiates a new LoggingSyslogAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *LoggingSyslogAdditional) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingSyslogAdditional) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingSyslogAdditional) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingSyslogAdditional) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHostname

`func (o *LoggingSyslogAdditional) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoggingSyslogAdditional) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoggingSyslogAdditional) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LoggingSyslogAdditional) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpv4

`func (o *LoggingSyslogAdditional) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LoggingSyslogAdditional) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LoggingSyslogAdditional) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LoggingSyslogAdditional) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### SetIpv4Nil

`func (o *LoggingSyslogAdditional) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *LoggingSyslogAdditional) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetToken

`func (o *LoggingSyslogAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingSyslogAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingSyslogAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingSyslogAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LoggingSyslogAdditional) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LoggingSyslogAdditional) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUseTls

`func (o *LoggingSyslogAdditional) GetUseTls() LoggingUseTlsString`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *LoggingSyslogAdditional) GetUseTlsOk() (*LoggingUseTlsString, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *LoggingSyslogAdditional) SetUseTls(v LoggingUseTlsString)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *LoggingSyslogAdditional) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


