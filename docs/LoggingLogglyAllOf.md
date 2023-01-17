# LoggingLogglyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The token to use for authentication ([https://www.loggly.com/docs/customer-token-authentication-token/](https://www.loggly.com/docs/customer-token-authentication-token/)). | [optional] 

## Methods

### NewLoggingLogglyAllOf

`func NewLoggingLogglyAllOf() *LoggingLogglyAllOf`

NewLoggingLogglyAllOf instantiates a new LoggingLogglyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogglyAllOfWithDefaults

`func NewLoggingLogglyAllOfWithDefaults() *LoggingLogglyAllOf`

NewLoggingLogglyAllOfWithDefaults instantiates a new LoggingLogglyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LoggingLogglyAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingLogglyAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingLogglyAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingLogglyAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
