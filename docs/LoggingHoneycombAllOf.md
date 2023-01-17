# LoggingHoneycombAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest. | [optional] 
**Dataset** | Pointer to **string** | The Honeycomb Dataset you want to log to. | [optional] 
**Token** | Pointer to **string** | The Write Key from the Account page of your Honeycomb account. | [optional] 

## Methods

### NewLoggingHoneycombAllOf

`func NewLoggingHoneycombAllOf() *LoggingHoneycombAllOf`

NewLoggingHoneycombAllOf instantiates a new LoggingHoneycombAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHoneycombAllOfWithDefaults

`func NewLoggingHoneycombAllOfWithDefaults() *LoggingHoneycombAllOf`

NewLoggingHoneycombAllOfWithDefaults instantiates a new LoggingHoneycombAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *LoggingHoneycombAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHoneycombAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHoneycombAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHoneycombAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDataset

`func (o *LoggingHoneycombAllOf) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *LoggingHoneycombAllOf) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *LoggingHoneycombAllOf) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *LoggingHoneycombAllOf) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetToken

`func (o *LoggingHoneycombAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingHoneycombAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingHoneycombAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingHoneycombAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
