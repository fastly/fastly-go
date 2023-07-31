# LoggingSplunkAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | The URL to post logs to. | [optional] 
**Token** | Pointer to **string** | A Splunk token for use in posting logs over HTTP to your collector. | [optional] 
**UseTLS** | Pointer to [**LoggingUseTLS**](LoggingUseTLS.md) |  | [optional] [default to LOGGINGUSETLS_no_tls]

## Methods

### NewLoggingSplunkAdditional

`func NewLoggingSplunkAdditional() *LoggingSplunkAdditional`

NewLoggingSplunkAdditional instantiates a new LoggingSplunkAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSplunkAdditionalWithDefaults

`func NewLoggingSplunkAdditionalWithDefaults() *LoggingSplunkAdditional`

NewLoggingSplunkAdditionalWithDefaults instantiates a new LoggingSplunkAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *LoggingSplunkAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingSplunkAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingSplunkAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingSplunkAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetToken

`func (o *LoggingSplunkAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingSplunkAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingSplunkAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingSplunkAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUseTLS

`func (o *LoggingSplunkAdditional) GetUseTLS() LoggingUseTLS`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingSplunkAdditional) GetUseTLSOk() (*LoggingUseTLS, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingSplunkAdditional) SetUseTLS(v LoggingUseTLS)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingSplunkAdditional) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
