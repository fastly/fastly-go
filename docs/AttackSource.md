# AttackSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | Country code of the attack source | 
**CountryName** | **string** | Name of the country | 
**RequestCount** | **int32** | Number of requests from this country | 
**TotalCount** | **int32** | Total number of attacks considered | 

## Methods

### NewAttackSource

`func NewAttackSource(countryCode string, countryName string, requestCount int32, totalCount int32, ) *AttackSource`

NewAttackSource instantiates a new AttackSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackSourceWithDefaults

`func NewAttackSourceWithDefaults() *AttackSource`

NewAttackSourceWithDefaults instantiates a new AttackSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *AttackSource) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AttackSource) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AttackSource) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCountryName

`func (o *AttackSource) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *AttackSource) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *AttackSource) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetRequestCount

`func (o *AttackSource) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *AttackSource) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *AttackSource) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetTotalCount

`func (o *AttackSource) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AttackSource) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AttackSource) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
