# ValuesStatusCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The HTTP request path. | [optional] 
**RatePerStatus** | Pointer to **float32** | The URL accounts for this percentage of the status code in this dimension. | [optional] 
**RatePerUrl** | Pointer to **float32** | The rate at which the status code in this dimension occurs for this URL. | [optional] 

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

### GetUrl

`func (o *ValuesStatusCodes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ValuesStatusCodes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ValuesStatusCodes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ValuesStatusCodes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

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

### GetRatePerUrl

`func (o *ValuesStatusCodes) GetRatePerUrl() float32`

GetRatePerUrl returns the RatePerUrl field if non-nil, zero value otherwise.

### GetRatePerUrlOk

`func (o *ValuesStatusCodes) GetRatePerUrlOk() (*float32, bool)`

GetRatePerUrlOk returns a tuple with the RatePerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUrl

`func (o *ValuesStatusCodes) SetRatePerUrl(v float32)`

SetRatePerUrl sets RatePerUrl field to given value.

### HasRatePerUrl

`func (o *ValuesStatusCodes) HasRatePerUrl() bool`

HasRatePerUrl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


