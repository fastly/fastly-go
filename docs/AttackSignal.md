# AttackSignal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | Name of the attack signal tag | 
**TagCount** | **int32** | Count of requests with this attack signal | 
**TotalCount** | **int32** | Total number of attacks considered | 

## Methods

### NewAttackSignal

`func NewAttackSignal(tagName string, tagCount int32, totalCount int32, ) *AttackSignal`

NewAttackSignal instantiates a new AttackSignal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackSignalWithDefaults

`func NewAttackSignalWithDefaults() *AttackSignal`

NewAttackSignalWithDefaults instantiates a new AttackSignal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *AttackSignal) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *AttackSignal) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *AttackSignal) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTagCount

`func (o *AttackSignal) GetTagCount() int32`

GetTagCount returns the TagCount field if non-nil, zero value otherwise.

### GetTagCountOk

`func (o *AttackSignal) GetTagCountOk() (*int32, bool)`

GetTagCountOk returns a tuple with the TagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCount

`func (o *AttackSignal) SetTagCount(v int32)`

SetTagCount sets TagCount field to given value.


### GetTotalCount

`func (o *AttackSignal) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AttackSignal) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AttackSignal) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


