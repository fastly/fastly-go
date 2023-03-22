# LoggingLogshuttleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** | The data authentication token associated with this endpoint. | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. | [optional] 

## Methods

### NewLoggingLogshuttleAllOf

`func NewLoggingLogshuttleAllOf() *LoggingLogshuttleAllOf`

NewLoggingLogshuttleAllOf instantiates a new LoggingLogshuttleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogshuttleAllOfWithDefaults

`func NewLoggingLogshuttleAllOfWithDefaults() *LoggingLogshuttleAllOf`

NewLoggingLogshuttleAllOfWithDefaults instantiates a new LoggingLogshuttleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LoggingLogshuttleAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingLogshuttleAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingLogshuttleAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingLogshuttleAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LoggingLogshuttleAllOf) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LoggingLogshuttleAllOf) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetURL

`func (o *LoggingLogshuttleAllOf) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingLogshuttleAllOf) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingLogshuttleAllOf) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingLogshuttleAllOf) HasURL() bool`

HasURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
