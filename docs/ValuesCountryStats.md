# ValuesCountryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The client&#39;s country subdivision code as defined by ISO 3166-2. | [optional] 
**RegionChr** | Pointer to **float32** | The cache hit ratio for the region. | [optional] 
**RegionErrorRate** | Pointer to **float32** | The error rate for the region. | [optional] 

## Methods

### NewValuesCountryStats

`func NewValuesCountryStats() *ValuesCountryStats`

NewValuesCountryStats instantiates a new ValuesCountryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesCountryStatsWithDefaults

`func NewValuesCountryStatsWithDefaults() *ValuesCountryStats`

NewValuesCountryStatsWithDefaults instantiates a new ValuesCountryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *ValuesCountryStats) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ValuesCountryStats) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ValuesCountryStats) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ValuesCountryStats) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionChr

`func (o *ValuesCountryStats) GetRegionChr() float32`

GetRegionChr returns the RegionChr field if non-nil, zero value otherwise.

### GetRegionChrOk

`func (o *ValuesCountryStats) GetRegionChrOk() (*float32, bool)`

GetRegionChrOk returns a tuple with the RegionChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionChr

`func (o *ValuesCountryStats) SetRegionChr(v float32)`

SetRegionChr sets RegionChr field to given value.

### HasRegionChr

`func (o *ValuesCountryStats) HasRegionChr() bool`

HasRegionChr returns a boolean if a field has been set.

### GetRegionErrorRate

`func (o *ValuesCountryStats) GetRegionErrorRate() float32`

GetRegionErrorRate returns the RegionErrorRate field if non-nil, zero value otherwise.

### GetRegionErrorRateOk

`func (o *ValuesCountryStats) GetRegionErrorRateOk() (*float32, bool)`

GetRegionErrorRateOk returns a tuple with the RegionErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionErrorRate

`func (o *ValuesCountryStats) SetRegionErrorRate(v float32)`

SetRegionErrorRate sets RegionErrorRate field to given value.

### HasRegionErrorRate

`func (o *ValuesCountryStats) HasRegionErrorRate() bool`

HasRegionErrorRate returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


