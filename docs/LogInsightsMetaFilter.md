# LogInsightsMetaFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceID** | Pointer to **string** | Specifies the ID of the service for which data should be returned. | [optional] 
**Start** | Pointer to **string** | Start time for the query as supplied in the request. | [optional] 
**End** | Pointer to **string** | End time for the query as supplied in the request. | [optional] 
**DomainExactMatch** | Pointer to **bool** | Value of the `domain_exact_match` filter as supplied in the request. | [optional] 
**Limit** | Pointer to **int32** | Number of records per page. | [optional] [default to 20]

## Methods

### NewLogInsightsMetaFilter

`func NewLogInsightsMetaFilter() *LogInsightsMetaFilter`

NewLogInsightsMetaFilter instantiates a new LogInsightsMetaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInsightsMetaFilterWithDefaults

`func NewLogInsightsMetaFilterWithDefaults() *LogInsightsMetaFilter`

NewLogInsightsMetaFilterWithDefaults instantiates a new LogInsightsMetaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceID

`func (o *LogInsightsMetaFilter) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LogInsightsMetaFilter) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LogInsightsMetaFilter) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LogInsightsMetaFilter) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetStart

`func (o *LogInsightsMetaFilter) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LogInsightsMetaFilter) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LogInsightsMetaFilter) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *LogInsightsMetaFilter) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *LogInsightsMetaFilter) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LogInsightsMetaFilter) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LogInsightsMetaFilter) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LogInsightsMetaFilter) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDomainExactMatch

`func (o *LogInsightsMetaFilter) GetDomainExactMatch() bool`

GetDomainExactMatch returns the DomainExactMatch field if non-nil, zero value otherwise.

### GetDomainExactMatchOk

`func (o *LogInsightsMetaFilter) GetDomainExactMatchOk() (*bool, bool)`

GetDomainExactMatchOk returns a tuple with the DomainExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainExactMatch

`func (o *LogInsightsMetaFilter) SetDomainExactMatch(v bool)`

SetDomainExactMatch sets DomainExactMatch field to given value.

### HasDomainExactMatch

`func (o *LogInsightsMetaFilter) HasDomainExactMatch() bool`

HasDomainExactMatch returns a boolean if a field has been set.

### GetLimit

`func (o *LogInsightsMetaFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogInsightsMetaFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogInsightsMetaFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogInsightsMetaFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
