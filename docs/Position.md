# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **string** | The id of the rule that this rule should be placed before. | [optional] 
**After** | Pointer to **string** | The id of the rule that this rule should be placed after. | [optional] 

## Methods

### NewPosition

`func NewPosition() *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *Position) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Position) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Position) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Position) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetAfter

`func (o *Position) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Position) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Position) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Position) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


