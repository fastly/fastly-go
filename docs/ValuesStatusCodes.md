# ValuesStatusCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | The HTTP request path. | [optional] 
**RatePerStatus** | Pointer to **float32** | The URL accounts for this percentage of the status code in this dimension. | [optional] 
**RatePerURL** | Pointer to **float32** | The rate at which the status code in this dimension occurs for this URL. | [optional] 

## Methods

### NewValuesStatusCodes

`func NewValuesStatusCodes() *ValuesStatusCodes`

NewValuesStatusCodes instantiates a new ValuesStatusCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesStatusCodesWithDefaults

`func NewValuesStatusCodesWithDefaults() *ValuesStatusCodes`

NewValuesStatusCodesWithDefaults instantiates a new ValuesStatusCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *ValuesStatusCodes) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *ValuesStatusCodes) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *ValuesStatusCodes) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *ValuesStatusCodes) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetRatePerStatus

`func (o *ValuesStatusCodes) GetRatePerStatus() float32`

GetRatePerStatus returns the RatePerStatus field if non-nil, zero value otherwise.

### GetRatePerStatusOk

`func (o *ValuesStatusCodes) GetRatePerStatusOk() (*float32, bool)`

GetRatePerStatusOk returns a tuple with the RatePerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerStatus

`func (o *ValuesStatusCodes) SetRatePerStatus(v float32)`

SetRatePerStatus sets RatePerStatus field to given value.

### HasRatePerStatus

`func (o *ValuesStatusCodes) HasRatePerStatus() bool`

HasRatePerStatus returns a boolean if a field has been set.

### GetRatePerURL

`func (o *ValuesStatusCodes) GetRatePerURL() float32`

GetRatePerURL returns the RatePerURL field if non-nil, zero value otherwise.

### GetRatePerURLOk

`func (o *ValuesStatusCodes) GetRatePerURLOk() (*float32, bool)`

GetRatePerURLOk returns a tuple with the RatePerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerURL

`func (o *ValuesStatusCodes) SetRatePerURL(v float32)`

SetRatePerURL sets RatePerURL field to given value.

### HasRatePerURL

`func (o *ValuesStatusCodes) HasRatePerURL() bool`

HasRatePerURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
