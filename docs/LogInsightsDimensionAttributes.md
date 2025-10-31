# LogInsightsDimensionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | Pointer to **float32** | The rate at which the value in the current dimension occurs. | [optional] 
**CountryChr** | Pointer to **float32** | The cache hit ratio for the country. | [optional] 
**CountryErrorRate** | Pointer to **float32** | The error rate for the country. | [optional] 
**CountryRequestRate** | Pointer to **float32** | This country&#39;s percentage of the total requests. | [optional] 

## Methods

### NewLogInsightsDimensionAttributes

`func NewLogInsightsDimensionAttributes() *LogInsightsDimensionAttributes`

NewLogInsightsDimensionAttributes instantiates a new LogInsightsDimensionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInsightsDimensionAttributesWithDefaults

`func NewLogInsightsDimensionAttributesWithDefaults() *LogInsightsDimensionAttributes`

NewLogInsightsDimensionAttributesWithDefaults instantiates a new LogInsightsDimensionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *LogInsightsDimensionAttributes) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *LogInsightsDimensionAttributes) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *LogInsightsDimensionAttributes) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *LogInsightsDimensionAttributes) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCountryChr

`func (o *LogInsightsDimensionAttributes) GetCountryChr() float32`

GetCountryChr returns the CountryChr field if non-nil, zero value otherwise.

### GetCountryChrOk

`func (o *LogInsightsDimensionAttributes) GetCountryChrOk() (*float32, bool)`

GetCountryChrOk returns a tuple with the CountryChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryChr

`func (o *LogInsightsDimensionAttributes) SetCountryChr(v float32)`

SetCountryChr sets CountryChr field to given value.

### HasCountryChr

`func (o *LogInsightsDimensionAttributes) HasCountryChr() bool`

HasCountryChr returns a boolean if a field has been set.

### GetCountryErrorRate

`func (o *LogInsightsDimensionAttributes) GetCountryErrorRate() float32`

GetCountryErrorRate returns the CountryErrorRate field if non-nil, zero value otherwise.

### GetCountryErrorRateOk

`func (o *LogInsightsDimensionAttributes) GetCountryErrorRateOk() (*float32, bool)`

GetCountryErrorRateOk returns a tuple with the CountryErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryErrorRate

`func (o *LogInsightsDimensionAttributes) SetCountryErrorRate(v float32)`

SetCountryErrorRate sets CountryErrorRate field to given value.

### HasCountryErrorRate

`func (o *LogInsightsDimensionAttributes) HasCountryErrorRate() bool`

HasCountryErrorRate returns a boolean if a field has been set.

### GetCountryRequestRate

`func (o *LogInsightsDimensionAttributes) GetCountryRequestRate() float32`

GetCountryRequestRate returns the CountryRequestRate field if non-nil, zero value otherwise.

### GetCountryRequestRateOk

`func (o *LogInsightsDimensionAttributes) GetCountryRequestRateOk() (*float32, bool)`

GetCountryRequestRateOk returns a tuple with the CountryRequestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRequestRate

`func (o *LogInsightsDimensionAttributes) SetCountryRequestRate(v float32)`

SetCountryRequestRate sets CountryRequestRate field to given value.

### HasCountryRequestRate

`func (o *LogInsightsDimensionAttributes) HasCountryRequestRate() bool`

HasCountryRequestRate returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


