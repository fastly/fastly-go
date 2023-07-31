# LoggingHerokuAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)). | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. | [optional] 

## Methods

### NewLoggingHerokuAdditional

`func NewLoggingHerokuAdditional() *LoggingHerokuAdditional`

NewLoggingHerokuAdditional instantiates a new LoggingHerokuAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHerokuAdditionalWithDefaults

`func NewLoggingHerokuAdditionalWithDefaults() *LoggingHerokuAdditional`

NewLoggingHerokuAdditionalWithDefaults instantiates a new LoggingHerokuAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LoggingHerokuAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingHerokuAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingHerokuAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingHerokuAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetURL

`func (o *LoggingHerokuAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingHerokuAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingHerokuAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingHerokuAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
