# LoggingHoneycombAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Honeycomb can ingest. | [optional] 
**Dataset** | Pointer to **string** | The Honeycomb Dataset you want to log to. | [optional] 
**Token** | Pointer to **string** | The Write Key from the Account page of your Honeycomb account. | [optional] 

## Methods

### NewLoggingHoneycombAdditional

`func NewLoggingHoneycombAdditional() *LoggingHoneycombAdditional`

NewLoggingHoneycombAdditional instantiates a new LoggingHoneycombAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHoneycombAdditionalWithDefaults

`func NewLoggingHoneycombAdditionalWithDefaults() *LoggingHoneycombAdditional`

NewLoggingHoneycombAdditionalWithDefaults instantiates a new LoggingHoneycombAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *LoggingHoneycombAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHoneycombAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHoneycombAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHoneycombAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDataset

`func (o *LoggingHoneycombAdditional) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *LoggingHoneycombAdditional) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *LoggingHoneycombAdditional) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *LoggingHoneycombAdditional) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetToken

`func (o *LoggingHoneycombAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingHoneycombAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingHoneycombAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingHoneycombAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


