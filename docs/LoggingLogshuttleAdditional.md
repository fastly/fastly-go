# LoggingLogshuttleAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** | The data authentication token associated with this endpoint. | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. | [optional] 

## Methods

### NewLoggingLogshuttleAdditional

`func NewLoggingLogshuttleAdditional() *LoggingLogshuttleAdditional`

NewLoggingLogshuttleAdditional instantiates a new LoggingLogshuttleAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogshuttleAdditionalWithDefaults

`func NewLoggingLogshuttleAdditionalWithDefaults() *LoggingLogshuttleAdditional`

NewLoggingLogshuttleAdditionalWithDefaults instantiates a new LoggingLogshuttleAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LoggingLogshuttleAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingLogshuttleAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingLogshuttleAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingLogshuttleAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LoggingLogshuttleAdditional) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LoggingLogshuttleAdditional) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetURL

`func (o *LoggingLogshuttleAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingLogshuttleAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingLogshuttleAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingLogshuttleAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
