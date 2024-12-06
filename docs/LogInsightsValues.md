# LogInsightsValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheHitRatio** | Pointer to **float32** | The cache hit ratio for the URL specified in the dimension. | [optional] 
**Region** | Pointer to **string** | The client&#39;s country subdivision code as defined by ISO 3166-2. | [optional] 
**RegionChr** | Pointer to **float32** | The cache hit ratio for the region. | [optional] 
**RegionErrorRate** | Pointer to **float32** | The error rate for the region. | [optional] 
**URL** | Pointer to **string** | The HTTP request path. | [optional] 
**RatePerStatus** | Pointer to **float32** | The URL accounts for this percentage of the status code in this dimension. | [optional] 
**RatePerURL** | Pointer to **float32** | The rate at which the reason in this dimension occurs among responses to this URL with a 503 status code. | [optional] 
**Var503RatePerURL** | Pointer to **float32** | The rate at which 503 status codes are returned for this URL. | [optional] 
**BrowserVersion** | Pointer to **string** | The version of the client&#39;s browser. | [optional] 
**Rate** | Pointer to **float32** | The percentage of requests matching the value in the current dimension. | [optional] 
**AverageBandwidthBytes** | Pointer to **float32** | The average bandwidth in bytes for responses to requests to the URL in the current dimension. | [optional] 
**BandwidthPercentage** | Pointer to **float32** | The total bandwidth percentage for all responses to requests to the URL in the current dimension. | [optional] 
**AverageResponseTime** | Pointer to **float32** | The average time in seconds to respond to requests to the URL in the current dimension. | [optional] 
**P95ResponseTime** | Pointer to **float32** | The P95 time in seconds to respond to requests to the URL in the current dimension. | [optional] 
**ResponseTimePercentage** | Pointer to **float32** | The total percentage of time to respond to all requests to the URL in the current dimension. | [optional] 
**MissRate** | Pointer to **float32** | The miss rate for requests to the URL in the current dimension. | [optional] 
**RequestPercentage** | Pointer to **float32** | The percentage of all requests made to the URL in the current dimension. | [optional] 

## Methods

### NewLogInsightsValues

`func NewLogInsightsValues() *LogInsightsValues`

NewLogInsightsValues instantiates a new LogInsightsValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInsightsValuesWithDefaults

`func NewLogInsightsValuesWithDefaults() *LogInsightsValues`

NewLogInsightsValuesWithDefaults instantiates a new LogInsightsValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheHitRatio

`func (o *LogInsightsValues) GetCacheHitRatio() float32`

GetCacheHitRatio returns the CacheHitRatio field if non-nil, zero value otherwise.

### GetCacheHitRatioOk

`func (o *LogInsightsValues) GetCacheHitRatioOk() (*float32, bool)`

GetCacheHitRatioOk returns a tuple with the CacheHitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHitRatio

`func (o *LogInsightsValues) SetCacheHitRatio(v float32)`

SetCacheHitRatio sets CacheHitRatio field to given value.

### HasCacheHitRatio

`func (o *LogInsightsValues) HasCacheHitRatio() bool`

HasCacheHitRatio returns a boolean if a field has been set.

### GetRegion

`func (o *LogInsightsValues) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogInsightsValues) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogInsightsValues) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogInsightsValues) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionChr

`func (o *LogInsightsValues) GetRegionChr() float32`

GetRegionChr returns the RegionChr field if non-nil, zero value otherwise.

### GetRegionChrOk

`func (o *LogInsightsValues) GetRegionChrOk() (*float32, bool)`

GetRegionChrOk returns a tuple with the RegionChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionChr

`func (o *LogInsightsValues) SetRegionChr(v float32)`

SetRegionChr sets RegionChr field to given value.

### HasRegionChr

`func (o *LogInsightsValues) HasRegionChr() bool`

HasRegionChr returns a boolean if a field has been set.

### GetRegionErrorRate

`func (o *LogInsightsValues) GetRegionErrorRate() float32`

GetRegionErrorRate returns the RegionErrorRate field if non-nil, zero value otherwise.

### GetRegionErrorRateOk

`func (o *LogInsightsValues) GetRegionErrorRateOk() (*float32, bool)`

GetRegionErrorRateOk returns a tuple with the RegionErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionErrorRate

`func (o *LogInsightsValues) SetRegionErrorRate(v float32)`

SetRegionErrorRate sets RegionErrorRate field to given value.

### HasRegionErrorRate

`func (o *LogInsightsValues) HasRegionErrorRate() bool`

HasRegionErrorRate returns a boolean if a field has been set.

### GetURL

`func (o *LogInsightsValues) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LogInsightsValues) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LogInsightsValues) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LogInsightsValues) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetRatePerStatus

`func (o *LogInsightsValues) GetRatePerStatus() float32`

GetRatePerStatus returns the RatePerStatus field if non-nil, zero value otherwise.

### GetRatePerStatusOk

`func (o *LogInsightsValues) GetRatePerStatusOk() (*float32, bool)`

GetRatePerStatusOk returns a tuple with the RatePerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerStatus

`func (o *LogInsightsValues) SetRatePerStatus(v float32)`

SetRatePerStatus sets RatePerStatus field to given value.

### HasRatePerStatus

`func (o *LogInsightsValues) HasRatePerStatus() bool`

HasRatePerStatus returns a boolean if a field has been set.

### GetRatePerURL

`func (o *LogInsightsValues) GetRatePerURL() float32`

GetRatePerURL returns the RatePerURL field if non-nil, zero value otherwise.

### GetRatePerURLOk

`func (o *LogInsightsValues) GetRatePerURLOk() (*float32, bool)`

GetRatePerURLOk returns a tuple with the RatePerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerURL

`func (o *LogInsightsValues) SetRatePerURL(v float32)`

SetRatePerURL sets RatePerURL field to given value.

### HasRatePerURL

`func (o *LogInsightsValues) HasRatePerURL() bool`

HasRatePerURL returns a boolean if a field has been set.

### GetVar503RatePerURL

`func (o *LogInsightsValues) GetVar503RatePerURL() float32`

GetVar503RatePerURL returns the Var503RatePerURL field if non-nil, zero value otherwise.

### GetVar503RatePerURLOk

`func (o *LogInsightsValues) GetVar503RatePerURLOk() (*float32, bool)`

GetVar503RatePerURLOk returns a tuple with the Var503RatePerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar503RatePerURL

`func (o *LogInsightsValues) SetVar503RatePerURL(v float32)`

SetVar503RatePerURL sets Var503RatePerURL field to given value.

### HasVar503RatePerURL

`func (o *LogInsightsValues) HasVar503RatePerURL() bool`

HasVar503RatePerURL returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *LogInsightsValues) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *LogInsightsValues) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *LogInsightsValues) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *LogInsightsValues) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetRate

`func (o *LogInsightsValues) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *LogInsightsValues) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *LogInsightsValues) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *LogInsightsValues) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetAverageBandwidthBytes

`func (o *LogInsightsValues) GetAverageBandwidthBytes() float32`

GetAverageBandwidthBytes returns the AverageBandwidthBytes field if non-nil, zero value otherwise.

### GetAverageBandwidthBytesOk

`func (o *LogInsightsValues) GetAverageBandwidthBytesOk() (*float32, bool)`

GetAverageBandwidthBytesOk returns a tuple with the AverageBandwidthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBandwidthBytes

`func (o *LogInsightsValues) SetAverageBandwidthBytes(v float32)`

SetAverageBandwidthBytes sets AverageBandwidthBytes field to given value.

### HasAverageBandwidthBytes

`func (o *LogInsightsValues) HasAverageBandwidthBytes() bool`

HasAverageBandwidthBytes returns a boolean if a field has been set.

### GetBandwidthPercentage

`func (o *LogInsightsValues) GetBandwidthPercentage() float32`

GetBandwidthPercentage returns the BandwidthPercentage field if non-nil, zero value otherwise.

### GetBandwidthPercentageOk

`func (o *LogInsightsValues) GetBandwidthPercentageOk() (*float32, bool)`

GetBandwidthPercentageOk returns a tuple with the BandwidthPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPercentage

`func (o *LogInsightsValues) SetBandwidthPercentage(v float32)`

SetBandwidthPercentage sets BandwidthPercentage field to given value.

### HasBandwidthPercentage

`func (o *LogInsightsValues) HasBandwidthPercentage() bool`

HasBandwidthPercentage returns a boolean if a field has been set.

### GetAverageResponseTime

`func (o *LogInsightsValues) GetAverageResponseTime() float32`

GetAverageResponseTime returns the AverageResponseTime field if non-nil, zero value otherwise.

### GetAverageResponseTimeOk

`func (o *LogInsightsValues) GetAverageResponseTimeOk() (*float32, bool)`

GetAverageResponseTimeOk returns a tuple with the AverageResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageResponseTime

`func (o *LogInsightsValues) SetAverageResponseTime(v float32)`

SetAverageResponseTime sets AverageResponseTime field to given value.

### HasAverageResponseTime

`func (o *LogInsightsValues) HasAverageResponseTime() bool`

HasAverageResponseTime returns a boolean if a field has been set.

### GetP95ResponseTime

`func (o *LogInsightsValues) GetP95ResponseTime() float32`

GetP95ResponseTime returns the P95ResponseTime field if non-nil, zero value otherwise.

### GetP95ResponseTimeOk

`func (o *LogInsightsValues) GetP95ResponseTimeOk() (*float32, bool)`

GetP95ResponseTimeOk returns a tuple with the P95ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95ResponseTime

`func (o *LogInsightsValues) SetP95ResponseTime(v float32)`

SetP95ResponseTime sets P95ResponseTime field to given value.

### HasP95ResponseTime

`func (o *LogInsightsValues) HasP95ResponseTime() bool`

HasP95ResponseTime returns a boolean if a field has been set.

### GetResponseTimePercentage

`func (o *LogInsightsValues) GetResponseTimePercentage() float32`

GetResponseTimePercentage returns the ResponseTimePercentage field if non-nil, zero value otherwise.

### GetResponseTimePercentageOk

`func (o *LogInsightsValues) GetResponseTimePercentageOk() (*float32, bool)`

GetResponseTimePercentageOk returns a tuple with the ResponseTimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimePercentage

`func (o *LogInsightsValues) SetResponseTimePercentage(v float32)`

SetResponseTimePercentage sets ResponseTimePercentage field to given value.

### HasResponseTimePercentage

`func (o *LogInsightsValues) HasResponseTimePercentage() bool`

HasResponseTimePercentage returns a boolean if a field has been set.

### GetMissRate

`func (o *LogInsightsValues) GetMissRate() float32`

GetMissRate returns the MissRate field if non-nil, zero value otherwise.

### GetMissRateOk

`func (o *LogInsightsValues) GetMissRateOk() (*float32, bool)`

GetMissRateOk returns a tuple with the MissRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissRate

`func (o *LogInsightsValues) SetMissRate(v float32)`

SetMissRate sets MissRate field to given value.

### HasMissRate

`func (o *LogInsightsValues) HasMissRate() bool`

HasMissRate returns a boolean if a field has been set.

### GetRequestPercentage

`func (o *LogInsightsValues) GetRequestPercentage() float32`

GetRequestPercentage returns the RequestPercentage field if non-nil, zero value otherwise.

### GetRequestPercentageOk

`func (o *LogInsightsValues) GetRequestPercentageOk() (*float32, bool)`

GetRequestPercentageOk returns a tuple with the RequestPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPercentage

`func (o *LogInsightsValues) SetRequestPercentage(v float32)`

SetRequestPercentage sets RequestPercentage field to given value.

### HasRequestPercentage

`func (o *LogInsightsValues) HasRequestPercentage() bool`

HasRequestPercentage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
