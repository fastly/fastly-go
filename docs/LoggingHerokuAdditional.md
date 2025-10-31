# LoggingHerokuAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)). | [optional] 
**Url** | Pointer to **string** | The URL to stream logs to. | [optional] 

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

### GetUrl

`func (o *LoggingHerokuAdditional) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingHerokuAdditional) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingHerokuAdditional) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingHerokuAdditional) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


