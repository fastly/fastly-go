# LoggingSplunkAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL to post logs to. | [optional] 
**Token** | Pointer to **string** | A Splunk token for use in posting logs over HTTP to your collector. | [optional] 
**UseTls** | Pointer to [**LoggingUseTlsString**](LoggingUseTlsString.md) |  | [optional] [default to LOGGINGUSETLSSTRING_no_tls]

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

### GetUrl

`func (o *LoggingSplunkAdditional) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingSplunkAdditional) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingSplunkAdditional) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingSplunkAdditional) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

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

### GetUseTls

`func (o *LoggingSplunkAdditional) GetUseTls() LoggingUseTlsString`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *LoggingSplunkAdditional) GetUseTlsOk() (*LoggingUseTlsString, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *LoggingSplunkAdditional) SetUseTls(v LoggingUseTlsString)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *LoggingSplunkAdditional) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


