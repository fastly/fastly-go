# LoggingSplunkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | The URL to post logs to. | [optional] 
**Token** | Pointer to **string** | A Splunk token for use in posting logs over HTTP to your collector. | [optional] 
**UseTLS** | Pointer to [**LoggingUseTLS**](LoggingUseTLS.md) |  | [optional] [default to LOGGINGUSETLS_no_tls]

## Methods

### NewLoggingSplunkAllOf

`func NewLoggingSplunkAllOf() *LoggingSplunkAllOf`

NewLoggingSplunkAllOf instantiates a new LoggingSplunkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSplunkAllOfWithDefaults

`func NewLoggingSplunkAllOfWithDefaults() *LoggingSplunkAllOf`

NewLoggingSplunkAllOfWithDefaults instantiates a new LoggingSplunkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *LoggingSplunkAllOf) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingSplunkAllOf) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingSplunkAllOf) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingSplunkAllOf) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetToken

`func (o *LoggingSplunkAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingSplunkAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingSplunkAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingSplunkAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUseTLS

`func (o *LoggingSplunkAllOf) GetUseTLS() LoggingUseTLS`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingSplunkAllOf) GetUseTLSOk() (*LoggingUseTLS, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingSplunkAllOf) SetUseTLS(v LoggingUseTLS)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingSplunkAllOf) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
